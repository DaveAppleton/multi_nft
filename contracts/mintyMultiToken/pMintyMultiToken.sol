pragma solidity ^0.7.5;
pragma abicoder v2;

import "@openzeppelin/contracts/token/ERC1155/ERC1155.sol";
import "../locking/locking.sol";

    struct PoolEntry {
        address  beneficiary;
        uint256  share;
    } 

contract pMintyMultiToken is ERC1155 {

    mapping (uint => bool) public minted;
    mapping(uint => string)       _tokenURIs;
    string                        base;
    address                public owner;
    mapping (address => bool)     auth;
    string                 public contractURI;
    locking []             public locs;
    string                 public lockError;

    uint256                public royaltyPerMille;
    PoolEntry []           public initialPool;
    PoolEntry []           public resalePool;

    event OperatorSet(address operator, bool enabled);

    modifier onlyAuth() {
        require (auth[msg.sender] || (msg.sender == owner),"unauthorised");
        _;
    }

    constructor(
        address _owner, 
        address saleContract, 
        locking[] memory _locs, 
        string memory _lockError, 
        uint256             _royaltyPerMille,
        PoolEntry [] memory _initialEntries,
        PoolEntry [] memory _resaleEntries
    ) ERC1155("") {
        owner = _owner;
        auth[saleContract] = true;
        emit OperatorSet(saleContract, true);
        locs = _locs;
        lockError = _lockError;
        base = "https://minty.mypinata.cloud/ipfs/";
        royaltyPerMille = _royaltyPerMille;
        uint256 iPerMille = 0;
        for (uint j = 0; j < _initialEntries.length; j++) {
            PoolEntry memory p = _initialEntries[j];
            initialPool.push(p);
            iPerMille += p.share;
        }
        require(iPerMille == 1000,"Initial shares do not add up");
        uint256 rPerMille = 0;
        for (uint j = 0; j < _resaleEntries.length; j++) {
            PoolEntry memory p = _resaleEntries[j];
            resalePool.push(p);
            rPerMille += p.share;
        }
        require(rPerMille == 1000,"Resale shares do not add up");
    }


    
    function setContractURI(string memory _uri) external onlyAuth {
        contractURI = _uri;
    } 
   

    function uri(uint256 tokenId) external view override returns (string memory) {
        string memory _tokenURI = _tokenURIs[tokenId];
        return string(abi.encodePacked(base, _tokenURI));
    }

    function mint(uint tokenId, uint quantity, string memory hash) external onlyAuth {
        bytes memory data;
        _tokenURIs[tokenId] = hash;
        minted[tokenId] = true;
        _mint(owner, tokenId, quantity, data);
    }

    function mintBatch(uint [] memory tokenIds, uint [] memory quantities, string[] memory hashes) external onlyAuth {
        bytes memory data;
        require(tokenIds.length == hashes.length,"array lengths do not match");
        for (uint j = 0; j < tokenIds.length; j++) {
            _tokenURIs[tokenIds[j]] = hashes[j];
            minted[tokenIds[j]] = true;
        }
        _mintBatch(owner, tokenIds, quantities, data);
    }

    function setAuth(address operator, bool enabled) external onlyAuth {
        auth[operator] = enabled;
        emit OperatorSet(operator, enabled);
    }

    function setBase(string memory _base) external onlyAuth {
        base = _base;
    }

    function validateBuyer(address buyer) external {
        for (uint j = 0; j < locs.length; j++) {
            bool res = locs[j].isLocked(buyer);
            if (res) return;
        }
        require(locs.length == 0,lockError);
    }

    function getRoyalties(uint saleNumber) external view returns (PoolEntry[] memory) {
        if (saleNumber == 0) return initialPool;
        return resalePool;
    } 


}