pragma solidity ^0.7.5;
//SPDX-License-Identifier: UNLICENSED

import "hardhat/console.sol";
/*

    mapping(address => mapping(uint => []offer))

    1) offerNew -> token, tokenID, quantity, price
    2) offerResale -> token, tokenId, quantity, price






*/

interface IMintyToken {
    function mint(address buyer, address artist,uint256 tokenId, string memory ipfsHash) external;

    function safeTransferFrom(address from, address to, uint256 tokenId, bytes calldata data) external;

    function creator(uint256 tokenId) external view returns (address);

    function tokenExists(uint256 tokenId) external view returns (bool);

    function ownerOf(uint256 tokenId) external view  returns (address);

    function owner() external view returns (address);
}

interface IMintyMultiToken {

    function isApprovedForAll(address account, address operator) external view returns (bool);

    function mint(uint256 tokenId, uint256 quantity, string memory ipfsHash) external; // mints tokens to contract owner

    function minted(uint256 id) external view returns (bool) ;

    function safeTransferFrom(address from, address to, uint256 tokenId, uint256 amount, bytes calldata data) external;

    function creator(uint256 tokenId) external view returns (address);

    function balanceOf(address account, uint256 id) external view returns (uint256); 

    function owner() external view returns (address);

    function uri(uint256 id) external view returns (string memory);
}

    struct Offer721 {
        address      creator;
        string       itemHash;
        uint256      price;
        bool         available;
        bool         minted;
    }

    struct Offer1155 {
        address           creator;
        uint256           quantity;
        string            itemHash;
        uint256           unitPrice;
    }

contract mintyMultiSale {

    //          token                         tokenid     offer
    mapping(IMintyMultiToken      => mapping(uint256 => Offer1155[]))   public items; 

    mapping(IMintyMultiToken      => address) public multiTokenOwners;

    address                        public owner = msg.sender;

    bool                                  entered;
    uint                           public ownerPerMille;
    uint                           public creatorPerMille;
    address payable                public minty;

    //mapping(IMintyMultiToken => mapping(uint => mapping(address => uint256))) public bids;

    event SharesUpdated(uint256 ownerShare, uint256 creatorShare);
    event NewOffer(IMintyMultiToken token, uint256 tokenId, address owner, uint256 quantity, uint256 price, string hash);
    event OfferWithdrawn(IMintyMultiToken token, uint256 tokenId);

    event ResaleOffer(IMintyMultiToken token,uint256 tokenId, uint256 quantity, address owner, uint256 price, uint256 position);
    event SaleRepriced(IMintyMultiToken token, uint256 tokenId, uint256 pos, uint256 price, address owner);
    event OfferAccepted(address buyer, IMintyMultiToken token, uint256 tokenId, uint256 pos, uint256 quantity, uint256 price);
    event SaleRetracted(IMintyMultiToken token, uint256 tokenId,uint256 pos, address owner);

    event BidReceived(address  bidder, IMintyMultiToken token, uint256 tokenId, uint256 bid);
    event BidIncreased(address bidder, IMintyMultiToken token, uint256 tokenId, uint256 previous_bid, uint256 this_bid);
    event BidWithdrawn(IMintyMultiToken token, uint256 tokenId);


    event Payment(address wallet,address creator, address _owner,uint256 amount);

    constructor(address payable wallet, uint256 opm, uint256 cpm) {
        minty = wallet;
        ownerPerMille = opm;
        creatorPerMille = cpm;
        emit SharesUpdated(opm, cpm);
    }

    function updateShares(uint256 opm, uint256 cpm) external {
        require(msg.sender == owner, "unauthorised");
        ownerPerMille = opm;
        creatorPerMille = cpm;
        emit SharesUpdated(opm, cpm);
    }

    function offerNew(IMintyMultiToken token, uint256 tokenId, string memory ipfsString, uint256 quantity, uint256 price) external {
        require(token.isApprovedForAll(msg.sender,address(this)),"You have not approved this contract to sell your tokens");
        require(!token.minted(tokenId),"Token ID already minted");
        require(token.owner() == msg.sender,"Unauthorised");
        require(items[token][tokenId].length == 0,"Unable to offer new");
        token.mint(tokenId,quantity,ipfsString);
        items[token][tokenId].push(Offer1155(msg.sender, quantity, ipfsString,price));
        emit NewOffer(token, tokenId, msg.sender, quantity, price, ipfsString);
        console.log("length at",address(token),tokenId,items[token][tokenId].length);
        return;        
    }

    //   SistineToken--->132-->[(initial Offer)]

 
    function offerResale(IMintyMultiToken token, uint256 tokenId, uint256 quantity, uint256 price) external {
        require((quantity <= token.balanceOf(msg.sender,tokenId) && (quantity > 0)),"You do not own enough of this token");
        require(token.isApprovedForAll(msg.sender,address(this)),"You have not approved this contract to sell your tokens");
        uint pos = items[token][tokenId].length;
        items[token][tokenId].push(Offer1155(msg.sender, quantity, token.uri(tokenId),price));
        emit ResaleOffer(token,tokenId, quantity,msg.sender, price, pos);
    }

   //   SistineToken--->132-->[(initial Offer),(resale offer)]

   //  SistineToken--->132-->[( Artist 5pcs )]
   //  SistineToken--->132-->[( Artist 3pcs )( Dave 1 pcs)]

   //                art_id


    function retractRemainingOffer(IMintyMultiToken token,uint256 tokenId, uint256 pos) external {
        require(pos < items[token][tokenId].length,"invalid offer position");
        Offer1155 memory offer = items[token][tokenId][pos];
        require(msg.sender == offer.creator, "not your offer");
        offer.quantity = 0;
        items[token][tokenId][pos] = offer;
        emit SaleRetracted(token,tokenId, pos, msg.sender);
    }

    function reSubmitOffer(IMintyMultiToken token, uint256 tokenId, uint256 pos, uint256 price) external {
        require(pos < items[token][tokenId].length,"invalid offer position");
        Offer1155 memory offer = items[token][tokenId][pos];
        require(msg.sender == offer.creator, "not your offer");
        offer.unitPrice = price;
        items[token][tokenId][pos] = offer;
        emit SaleRepriced(token,tokenId, pos, price, msg.sender);
    }

    function acceptOffer(IMintyMultiToken token,uint tokenId, uint256 pos, uint256 quantity) external payable {
        accept(token,tokenId,pos, quantity, msg.value);
    }

    function accept(IMintyMultiToken token,uint tokenId, uint256 pos, uint256 quantity, uint value) internal {
        require(!entered,"No reentrancy please");
        entered = true;
        bytes memory data;
        console.log(address(token),tokenId, pos, items[token][tokenId].length);
        Offer1155 memory offer = items[token][tokenId][pos];
        address _owner = offer.creator;
        require(offer.quantity >= quantity,"not enough items available");
        require(value >= mul(offer.unitPrice,quantity), "Price not met");
        require(token.balanceOf(_owner,tokenId) >= quantity,"not enough items owned by offerer");
        token.safeTransferFrom(_owner,msg.sender,tokenId,quantity, data);

        offer.quantity -= quantity;
        items[token][tokenId][pos] = offer;
        emit Payment(minty,offer.creator,_owner,value);
        splitFee(payable(offer.creator), payable(_owner), value);
        entered = false;
        emit OfferAccepted(msg.sender, token, tokenId, pos, quantity, value);
    }

    //-------- UTILITY ------

    function numberOfBids(IMintyMultiToken token,uint tokenId) external view returns (uint) {
        return items[token][tokenId].length;
    }

    function splitFee(address payable creator, address payable _owner, uint value) internal {
        uint creatorPart = value * creatorPerMille / 1000;
        uint ownerPart   = value - creatorPart;
        if (creator == _owner) {
            creator.transfer(ownerPart+creatorPart);
        } else {
            creator.transfer(creatorPart);
            _owner.transfer(ownerPart);
        }
        minty.transfer(value - (creatorPart + ownerPart));
    }

    function mul(uint256 a, uint256 b) internal pure returns (uint256 c) {
        if (a == 0) {
        return 0;
        }
        c = a * b;
        assert(c / a == b);
        return c;
    }


}