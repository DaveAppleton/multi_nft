// Sources flattened with hardhat v2.3.0 https://hardhat.org

// File contracts/sale2M/pMintyMultiSale.sol

pragma solidity ^0.7.5;
pragma abicoder v2;
//SPDX-License-Identifier: MIT


/*

    1. Allow Multiple Mints at once
    2. Restrict level 0 sales to people who satisfy the criteria set out in the 1155

*/

struct PoolEntry {
    address  beneficiary;
    uint256  share;
} 

interface IMintyMultiToken {

    function isApprovedForAll(address account, address operator) external view returns (bool);

    function mint(uint256 tokenId, uint256 quantity, string memory ipfsHash, uint256 poolId) external; // mints tokens to contract owner

    function mintBatch(uint [] memory tokenIds, uint [] memory quantities, string[] memory hashes, uint256 poolId) external;

    function minted(uint256 id) external view returns (bool) ;

    function safeTransferFrom(address from, address to, uint256 tokenId, uint256 amount, bytes calldata data) external;

    function creator(uint256 tokenId) external view returns (address);

    function balanceOf(address account, uint256 id) external view returns (uint256); 

    function owner() external view returns (address);

    function uri(uint256 id) external view returns (string memory);

    function validateBuyer(address buyer) external; // should revert is not valid

    function getRoyalties(uint saleNumber, uint256 tokenId) external view returns (PoolEntry[] memory);

    function royaltyPerMille() external view returns (uint256);
}

abstract contract oldContract  {
    mapping(IMintyMultiToken      => mapping(uint256 => Offer1155[]))   public items; 
    function numberOfOffers(IMintyMultiToken token,uint tokenId) external view virtual returns (uint);
}

struct Offer1155 {
    address           creator;
    uint256           quantity;
    string            itemHash;
    uint256           unitPrice;
}

interface IERC20 {
    function transferFrom(address owner, address receiver, uint256 amount) external returns (bool);
}

contract pMintyMultiSale {

    //          token                         tokenid     offer
    mapping(IMintyMultiToken      => mapping(uint256 => Offer1155[]))   public items; 

    mapping(IMintyMultiToken      => address) public multiTokenOwners;

    address                        public owner = msg.sender;
    IERC20                         public weth;

    bool                                  entered;
    uint                           public divisor;
    address                        public minty;

    bool                           public paused;


    //mapping(IMintyMultiToken => mapping(uint => mapping(address => uint256))) public bids;

    event OwnershipTransferred(address indexed previousOwner, address indexed newOwner);
    event WalletTransferred(address indexed previousWallet, address indexed newWallet);

    event FeeUpdated(uint256 divisor);
    event NewOffer(IMintyMultiToken token, uint256 tokenId, address owner, uint256 quantity, uint256 price, string hash, uint256 poolId);
    event OfferWithdrawn(IMintyMultiToken token, uint256 tokenId);

    event ResaleOffer(IMintyMultiToken token,uint256 tokenId, uint256 quantity, address owner, uint256 price, uint256 position);
    event SaleRepriced(IMintyMultiToken token, uint256 tokenId, uint256 pos, uint256 price, address owner);
    event OfferAccepted(address buyer, IMintyMultiToken token, uint256 tokenId, uint256 pos, uint256 quantity, uint256 price);
    event SaleRetracted(IMintyMultiToken token, uint256 tokenId,uint256 pos, address owner);

    event BidReceived(address  bidder, IMintyMultiToken token, uint256 tokenId, uint256 bid);
    event BidIncreased(address bidder, IMintyMultiToken token, uint256 tokenId, uint256 previous_bid, uint256 this_bid);
    event BidWithdrawn(IMintyMultiToken token, uint256 tokenId);


    event Payment(address wallet,address creator, address _owner,uint256 amount);

    event Paused(bool isPaused);

    modifier onlyOwner() {
        require(msg.sender == owner, "unauthorised");
        _;
    }

    modifier notPaused() {
        require(!paused,"operation not permitted when sale is paused");
        _;
    }


    constructor(address wallet, IERC20 _weth, uint256 _divisor) {
        require(_divisor >= 1000, "Divisor must be >= 1000");
        minty = wallet;
        weth = _weth;
        divisor = _divisor;
        emit FeeUpdated(_divisor);
    }

    function updateShares(uint256 _divisor) external onlyOwner {
        require(_divisor >= 1000, "Divisor must be >= 1000");
        divisor = _divisor;
        emit FeeUpdated(_divisor);
    }

    function offerNew(IMintyMultiToken token, uint256 tokenId, string memory ipfsString, uint256 quantity, uint256 price, uint256 poolId) external notPaused {
        require(token.isApprovedForAll(msg.sender,address(this)),"You have not approved this contract to sell your tokens");
        require(!token.minted(tokenId),"Token ID already minted");
        require(token.owner() == msg.sender,"Unauthorised");
        require(items[token][tokenId].length == 0,"Unable to offer new");
        token.mint(tokenId,quantity,ipfsString, poolId);
        items[token][tokenId].push(Offer1155(msg.sender, quantity, ipfsString,price));
        emit NewOffer(token, tokenId, msg.sender, quantity, price, ipfsString, poolId);
        return;        
    }

    function offerNewBatch(IMintyMultiToken token, uint256[] memory tokenIds, string[] memory ipfsStrings, uint256[] memory quantities, uint256[] memory prices, uint256 poolId) external notPaused {
        require(token.isApprovedForAll(msg.sender,address(this)),"You have not approved this contract to sell your tokens");
        require(token.owner() == msg.sender,"Unauthorised");
        for (uint j = 0; j < tokenIds.length; j++) {
            require(!token.minted(tokenIds[j]),"Token ID already minted");
            require(items[token][tokenIds[j]].length == 0,"Unable to offer new");
            items[token][tokenIds[j]].push(Offer1155(msg.sender, quantities[j], ipfsStrings[j],prices[j]));
            emit NewOffer(token, tokenIds[j], msg.sender, quantities[j], prices[j], ipfsStrings[j],poolId);
        }
        token.mintBatch(tokenIds,quantities,ipfsStrings, poolId);
        return;        
    }

    //   SistineToken--->132-->[(initial Offer)]

 
    function offerResale(IMintyMultiToken token, uint256 tokenId, uint256 quantity, uint256 price) external notPaused {
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

    function reSubmitOffer(IMintyMultiToken token, uint256 tokenId, uint256 pos, uint256 price) external notPaused {
        require(pos < items[token][tokenId].length,"invalid offer position");
        Offer1155 memory offer = items[token][tokenId][pos];
        require(msg.sender == offer.creator, "not your offer");
        offer.unitPrice = price;
        items[token][tokenId][pos] = offer;
        emit SaleRepriced(token,tokenId, pos, price, msg.sender);
    }

    function acceptOffer(IMintyMultiToken token,uint tokenId, uint256 pos, uint256 quantity) external notPaused {
        require(!entered,"No reentrancy please");
        if (pos == 0) {
            token.validateBuyer(msg.sender);
        }
        entered = true;
        bytes memory data;
        Offer1155 memory offer = items[token][tokenId][pos];
        address _owner = offer.creator;
        require(offer.quantity >= quantity,"not enough items available");
        uint value = mul(offer.unitPrice,quantity);
        require(token.balanceOf(_owner,tokenId) >= quantity,"not enough items owned by offerer");
        token.safeTransferFrom(_owner,msg.sender,tokenId,quantity, data);

        offer.quantity -= quantity;
        items[token][tokenId][pos] = offer;
        emit Payment(minty,offer.creator,_owner,value);
        splitFee(token,tokenId, pos, _owner, value);
        entered = false;
        emit OfferAccepted(msg.sender, token, tokenId, pos, quantity, value);
    }

    //-------- UTILITY ------

    function numberOfOffers(IMintyMultiToken token,uint tokenId) external view returns (uint) {
        return items[token][tokenId].length;
    }

    function available(IMintyMultiToken token,uint tokenId, uint offerId) external view returns (uint) {
        require(offerId < items[token][tokenId].length,"OfferID not valid");
        Offer1155 memory offer = items[token][tokenId][offerId];
        if (!token.isApprovedForAll(offer.creator,address(this))) return 0;
        uint256 onOffer = offer.quantity;
        uint256 owned   = token.balanceOf(offer.creator,tokenId);
        return min(onOffer,owned);
    }

    function price(IMintyMultiToken token,uint tokenId, uint offerId) external view returns (uint) {
        require(offerId < items[token][tokenId].length,"OfferID not valid");
        Offer1155 memory offer = items[token][tokenId][offerId];
        if (!token.isApprovedForAll(offer.creator,address(this))) return 0;
        return offer.unitPrice;
    }


    function splitFee(IMintyMultiToken token, uint256 tokenId, uint256 position, address _seller, uint value) internal {
        
        uint royaltyPerMille = token.royaltyPerMille();
        uint royaltyPart = mul(value , royaltyPerMille) / divisor;

        uint sellerPart  = mul(value , (1000 - royaltyPerMille)) / divisor;
        require(weth.transferFrom(msg.sender,_seller,sellerPart),"cannot transfer funds");

        uint sent = sellerPart;
        PoolEntry[] memory royalties = token.getRoyalties(position, tokenId);
        for (uint j = 0; j < royalties.length; j++) {
            uint amount = mul(royaltyPart , royalties[j].share) / 1000;
            require(weth.transferFrom(msg.sender,royalties[j].beneficiary,amount),"cannot transfer funds");
            sent += amount;
        }

        uint mintyPart = value - sent;
        require(weth.transferFrom(msg.sender,minty,mintyPart),"cannot transfer funds");
    }

    function mul(uint256 a, uint256 b) internal pure returns (uint256 c) {
        if (a == 0) {
            return 0;
        }
        c = a * b;
        assert(c / a == b);
        return c;
    }

    function min(uint a, uint b) internal pure returns (uint) {
        if (a < b) return a;
        return b;
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

    // don't over engineer this - just the offers
    function transferItems(oldContract old, IMintyMultiToken token, uint256 tokenId) external onlyOwner {
        require(1 == old.numberOfOffers(token,tokenId),"Only when there is a single offer");
        // address           creator;
        // uint256           quantity;
        // string            itemHash;
        // uint256           unitPrice;
        Offer1155   memory   theOffer;
        (theOffer.creator,theOffer.quantity,theOffer.itemHash,theOffer.unitPrice) = old.items(token,tokenId,0);

        items[token][tokenId].push(theOffer);
    }

    function PauseSale(bool putOnHold) external onlyOwner {
        paused = putOnHold;
        emit Paused(putOnHold);
    }

}
