pragma solidity ^0.7.5;
//SPDX-License-Identifier: UNLICENSED

import "@openzeppelin/contracts/token/ERC20/IERC20.sol";


    struct Offer {
        address  creator;
        string   itemHash;
        uint256  price;
        bool     available;
        bool     minted;
    } 

interface IMintyToken {
    function mint(address buyer, address artist,uint256 tokenId, string memory ipfsHash) external;

    function safeTransferFrom(address from, address to, uint256 tokenId, bytes calldata data) external;

    function artist(uint256 tokenId) external view returns (address);

    function tokenExists(uint256 tokenId) external view returns (bool);

    function ownerOf(uint256 tokenId) external view  returns (address);

    function isApprovedForAll(address owner, address operator) external view returns (bool) ;

    function getApproved(uint256 tokenId) external view returns (address);
}

contract pMintysale {

    address                     public owner = msg.sender;
    IMintyToken                 public token;
    IERC20                      public weth;

    uint256                     public nextToken;
    mapping(uint256 => Offer)   public items; 

    mapping(address => bool)    public isMinter;

    bool                               entered;
    uint                        public ownerPerMille;
    uint                        public creatorPerMille;
    uint                        public divisor;
    address                     public minty;

    mapping(uint => mapping(address => uint256)) public bids;

    event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);
    event WalletTransferred(address indexed previousWallet, address indexed newWallet);

    event SharesUpdated(uint256 ownerShare, uint256 creatorShare, uint256 divisor);
    event NewOffer(uint256 tokenId, address owner, uint256 price, string hash);
    event ResaleOffer(uint256 tokenId, address owner, uint256 price);
    event SaleResubmitted(uint256 tokenId, uint256 price);
    event OfferAccepted(address buyer, uint256 tokenId, uint256 price);
    event SaleRetracted(uint256 tokenId);

    event BidReceived(address  bidder, uint256 tokenId, uint256 bid);
    event BidIncreased(address bidder, uint256 tokenId, uint256 previous_bid, uint256 this_bid);


    event Payment(address wallet,address creator, address _owner);

    modifier onlyOwner() {
        require(msg.sender == owner, "unauthorised");
        _;
    }

    constructor(IERC20 _weth, address wallet, uint256 _ownerPerMille, uint256 _creatorPerMille, uint _divisor) {
        require(_creatorPerMille + _ownerPerMille == 1000,"sum(_creatorPerMille + _ownerPerMille) must equal 1000");
        require(_divisor >= 1000,"divisor is less than 1000");
        weth = _weth;
        minty = wallet;
        ownerPerMille = _ownerPerMille;
        creatorPerMille = _creatorPerMille;
        divisor = _divisor;
        emit SharesUpdated(_ownerPerMille, _creatorPerMille, _divisor);
    }

    function setMintyUnique(IMintyToken m) external onlyOwner {
        token = m;
    }

    function setMinter(address minter, bool status) external onlyOwner {
        isMinter[minter] = status;
    }

    function updateShares(uint256 _ownerPerMille, uint256 _creatorPerMille, uint _divisor) external onlyOwner {
        require(_creatorPerMille + _ownerPerMille == 1000,"sum(_creatorPerMille + _ownerPerMille) must equal 1000");
        require(divisor >= 1000,"divisor is less than 1000");
        ownerPerMille = _ownerPerMille;
        creatorPerMille = _creatorPerMille;
        divisor = _divisor;
        emit SharesUpdated(_ownerPerMille, _creatorPerMille, _divisor);
    }

    function offerNew(uint256 tokenId, string memory ipfsString, uint256 price) external {
        require(!token.tokenExists(tokenId),"Invalid token ID");
        require(isMinter[msg.sender],"You are not allowed to mint tokens here");
        Offer memory offer = items[tokenId];
        require((offer.creator == address(0) ),"Attempt to modify an existing offer");
        items[tokenId] = Offer(msg.sender, ipfsString,price, true,false);
        emit NewOffer(tokenId, msg.sender, price, ipfsString);
    }

    function offerResale(uint256 tokenId, uint256 price) external {
        require(token.tokenExists(tokenId),"Token does not exist!");
        require(msg.sender == token.ownerOf(tokenId),"You do not own this token");
        require( (token.isApprovedForAll(msg.sender, address(this))) || (token.getApproved(tokenId) == address(this)),"You have not approved this contract to sell your tokens");
        items[tokenId] = Offer(msg.sender, "",price, true,true);
        emit ResaleOffer(tokenId, msg.sender, price);
    }

    // only needed if we are replacing an old contract
    function offerSpecial(uint256 tokenId, address creator, string memory ipfsString, uint256 price) external onlyOwner {
        require(!token.tokenExists(tokenId),"Invalid token ID");
        items[tokenId] = Offer(creator, ipfsString,price, true,false);
        emit NewOffer(tokenId, creator, price, ipfsString);
    }

    function retractOffer(uint256 tokenId) external {
        Offer memory offer = items[tokenId];
        address _owner = offer.creator;
        if (token.tokenExists(tokenId)) {
            _owner = token.ownerOf(tokenId);
        }
        require(_owner == msg.sender,"Unauthorised");
        offer.available = false;
        items[tokenId] = offer;
        emit SaleRetracted(tokenId);
    }

    function reSubmitOffer(uint256 tokenId, uint256 price) external {
        Offer memory offer = items[tokenId];
        address _owner = offer.creator;
        if (token.tokenExists(tokenId)) {
            _owner = token.ownerOf(tokenId);
        }
        require(_owner == msg.sender,"Unauthorised");
        offer.available = true;
        offer.price = price;
        items[tokenId] = offer;
        emit SaleResubmitted(tokenId, price);
    }

    function acceptOffer(uint tokenId) external  {
        uint256 price = items[tokenId].price;
        require(weth.transferFrom(msg.sender, address(this), price),"Cannot transfer funds");
        accept(tokenId, price);
    }

    function accept(uint tokenId,uint value) internal {
        require(!entered,"No reentrancy please");
        entered = true;
        bytes memory data;
        Offer memory offer = items[tokenId];
        address _owner = offer.creator;
        require(offer.available,"Item not available");
        require(value >= offer.price, "Price not met");
        if (offer.minted) {
            address _realOwner = token.ownerOf(tokenId);
            require(_realOwner == _owner,"Item not owned by offerer");
            token.safeTransferFrom(_owner,msg.sender,tokenId,data);
        } else {
            token.mint(msg.sender,offer.creator,tokenId,offer.itemHash);
            offer.minted = true;
        }
        offer.available = false;
        items[tokenId] = offer;
        address creator = token.artist(tokenId); 
        emit Payment(minty,creator,_owner);
        splitFee(creator, _owner, value);
        entered = false;
        emit OfferAccepted(msg.sender, tokenId, value);
    }

    function splitFee(address  creator, address  _owner, uint value) internal {
        
        uint creatorPart = value * creatorPerMille / divisor;
        uint ownerPart   = value * ownerPerMille / divisor;
        uint mintyPart   = value - (creatorPart + ownerPart);
        if (creator == _owner) {
            require(weth.transfer(creator,ownerPart+creatorPart),"Cannot transfer funds");
        } else {
            require(weth.transfer(creator,creatorPart),"Cannot transfer funds");
            require(weth.transfer(_owner, ownerPart),"Cannot transfer funds");
        }
        require(weth.transfer(minty,mintyPart),"Cannot transfer funds");
    }

    function makeBid(uint256 tokenId, uint256 topup) external  {
        require(weth.transferFrom(msg.sender, address(this), topup),"Cannot transfer funds");
        Offer memory offer = items[tokenId];
        require(offer.available,"Item not available");
        uint myBid = topup + bids[tokenId][msg.sender];
        if (myBid > offer.price) {
            bids[tokenId][msg.sender] = 0;
            accept(tokenId, myBid);
            return;
        }
        bids[tokenId][msg.sender] = myBid;
        if (myBid == topup) {
            emit BidReceived(msg.sender, tokenId, myBid);
        } else {
            emit BidIncreased(msg.sender, tokenId, myBid-topup, topup);
        }
    }

    function acceptBid(uint256 tokenId, address bidder) external {
        bytes memory data;
        require(!entered,"No reentrancy please");
        entered = true;

        Offer memory offer = items[tokenId];
        address _owner = offer.creator;
        address _realOwner = token.ownerOf(tokenId);
        require(offer.available,"Item not available");
        require(_realOwner == _owner,"Item not owned by offerer");
        require(msg.sender == _owner,"Not your item to sell");
        uint256 bid = bids[tokenId][bidder];
        bids[tokenId][bidder] = 0;

        if (offer.minted) {
            token.safeTransferFrom(_owner,bidder,tokenId,data);
        } else {
            token.mint(bidder,offer.creator,tokenId,offer.itemHash);
            offer.minted = true;
        }
        offer.available = false;
        items[tokenId] = offer;
        emit Payment(minty,offer.creator,_owner);
        splitFee(offer.creator, _owner,bid);
        entered = false;
        emit OfferAccepted(msg.sender, tokenId, bid);
    }

    function withdrawBid(uint256 tokenId) external {
        require(!entered,"No reentrancy please");
        entered = true;
        uint256 bid = bids[tokenId][msg.sender];
        require(bid > 0,"nothing to withdraw");
        bids[tokenId][msg.sender] = 0;
        require(weth.transfer(msg.sender,bid),"Cannot transfer funds");
        entered = false;
    }

    // ------ UTILS

    function available(uint256 tokenId) external view returns (bool) {
        Offer memory offer = items[tokenId];
        if (!offer.available) return false;
        if (!offer.minted) return true;
        if (token.ownerOf(tokenId) != items[tokenId].creator) return false;
        if (token.isApprovedForAll(offer.creator, address(this))) return true;
        return (token.getApproved(tokenId) == address(this));
    }

    function transferOwnership(address newOwner) public onlyOwner {
        require(newOwner != address(0),"Do not set to address zero");
        emit OwnershipTransferred(owner, newOwner);
        owner = newOwner;
    }

    function changeWallet(address newWallet) public onlyOwner {
        require(newWallet != address(0),"Do not set to address zero");
        emit WalletTransferred(minty, newWallet);
        minty = newWallet;
    }

}