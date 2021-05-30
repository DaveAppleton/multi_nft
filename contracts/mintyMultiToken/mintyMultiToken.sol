pragma solidity ^0.7.5;
pragma abicoder v2;

import "@openzeppelin/contracts/token/ERC1155/ERC1155.sol";
import "../locking/locking.sol";
import "hardhat/console.sol";

contract mintyMultiToken is ERC1155 {

    mapping (uint => bool) public minted;
    mapping(uint => string)       _tokenURIs;
    string                        base;
    address                public owner;
    mapping (address => bool)     auth;
    string                 public contractURI;
    locking []             public locs;
    string                 public lockError;


    modifier onlyAuth() {
        require (auth[msg.sender] || (msg.sender == owner),"unauthorised");
        _;
    }

    constructor(address _owner, locking[] memory _locs, string memory _lockError ) ERC1155("") {
        owner = _owner;
        locs = _locs;
        lockError = _lockError;
        base = "https://minty.mypinata.cloud/ipfs/";
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

    function setAuth(address operator, bool state) external onlyAuth {
        auth[operator] = state;
    }

    function setBase(string memory _base) external onlyAuth {
        base = _base;
    }

    function validateBuyer(address buyer) external {
        for (uint j = 0; j < locs.length; j++) {
            bool res = locs[j].isLocked(buyer);
            console.log("trying ",address(locs[j]),buyer,res);
            if (res) return;
        }
        require(locs.length == 0,lockError);
    } 
}