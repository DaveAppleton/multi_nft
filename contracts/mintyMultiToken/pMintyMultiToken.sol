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
    address                public deployer = msg.sender;

    uint256                         public royaltyPerMille;
    uint256                         public numPools;
    mapping(uint => PoolEntry [])   public initialPools;
    mapping(uint => PoolEntry [])   public resalePools;
    mapping(uint => string)         public poolNames;
    mapping(uint => uint)           public poolByTokenId;

    event OperatorSet(address operator, bool enabled);
    event PoolAdded(uint256 poolId, string poolName);

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
        PoolEntry [] memory _resaleEntries,
        string       memory _poolName
    ) ERC1155("") {
        owner = _owner;
        auth[saleContract] = true;
        emit OperatorSet(saleContract, true);
        locs = _locs;
        lockError = _lockError;
        base = "https://minty.mypinata.cloud/ipfs/";
        royaltyPerMille = _royaltyPerMille;
        uint256 iPerMille = 0;
        _addPools(_initialEntries, _resaleEntries, _poolName);
    }

    function addPools(
        PoolEntry [] memory _initialEntries,
        PoolEntry [] memory _resaleEntries,
        string       memory _poolName
    ) public  {
        require ( msg.sender == deployer ||  auth[msg.sender] || (msg.sender == owner) , "Unauthorised");
        _addPools(_initialEntries, _resaleEntries,_poolName);
    }

    function _addPools(
        PoolEntry [] memory _initialEntries,
        PoolEntry [] memory _resaleEntries,
        string       memory _poolName
    ) internal {
        uint thisPool = numPools++;
        uint256 iPerMille = 0;
        for (uint j = 0; j < _initialEntries.length; j++) {
            PoolEntry memory p = _initialEntries[j];
            initialPools[thisPool].push(p);
            iPerMille += p.share;
        }
        require(iPerMille == 1000,"Initial shares do not add up");
        uint256 rPerMille = 0;
        for (uint j = 0; j < _resaleEntries.length; j++) {
            PoolEntry memory p = _resaleEntries[j];
            resalePools[thisPool].push(p);
            rPerMille += p.share;
        }
        require(rPerMille == 1000,"Resale shares do not add up");
        poolNames[thisPool] = _poolName;
        emit PoolAdded(thisPool,_poolName);
    }

    function setContractURI(string memory _uri) external onlyAuth {
        contractURI = _uri;
    } 
   

    function uri(uint256 tokenId) external view override returns (string memory) {
        string memory _tokenURI = _tokenURIs[tokenId];
        return string(abi.encodePacked(base, _tokenURI));
    }

    function mint(uint tokenId, uint quantity, string memory hash, uint poolId) external onlyAuth {
        require(poolId < numPools,"Invalid Pool Number");
        bytes memory data;
        _tokenURIs[tokenId] = hash;
        minted[tokenId] = true;
        _mint(owner, tokenId, quantity, data);
        poolByTokenId[tokenId] = poolId;
    }

    function mintBatch(uint [] memory tokenIds, uint [] memory quantities, string[] memory hashes, uint256 poolId) external onlyAuth {
        bytes memory data;
        require(tokenIds.length == hashes.length,"array lengths do not match");
        require(poolId < numPools,"Invalid Pool Number");
        for (uint j = 0; j < tokenIds.length; j++) {
            _tokenURIs[tokenIds[j]] = hashes[j];
            minted[tokenIds[j]] = true;
            poolByTokenId[tokenIds[j]] = poolId;
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

    function getRoyalties(uint saleNumber, uint tokenId) external view returns (PoolEntry[] memory) {
        uint thisPool = poolByTokenId[tokenId];
        if (saleNumber == 0) return initialPools[thisPool];
        return resalePools[thisPool];
    } 


}