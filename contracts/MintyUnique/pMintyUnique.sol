pragma solidity ^0.7.0;
//SPDX-License-Identifier: UNLICENSED

import "@openzeppelin/contracts/token/ERC721/ERC721.sol";
import "@openzeppelin/contracts/access/Ownable.sol";


contract pMintyUnique is ERC721, Ownable {

    mapping(address=>bool)          public operators;
    uint256                         public nextToken = 1;
    mapping (uint256 => address)    public artist;
    string                          _contractURI;

    modifier onlyOperator () {
        require(operators[msg.sender],"onlyOperator: unauthorised");
        _;
    }

    event OperatorSet(address operator, bool enabled);

    constructor(address saleContract) ERC721("Minty Unique NFT","MNTNFT") {
        _setBaseURI("https://minty.mypinata.cloud/ipfs/");
        operators[msg.sender] = true;
        operators[saleContract] = true;
    }

    function contractURI() external view returns (string memory) {
        return _contractURI;
    }

    function mint(address _owner, address _artist, uint tokenId, string memory _tokenURI) external onlyOperator {
        require(!_exists(tokenId),"Token already exists");
        artist[tokenId] =_artist;
        _safeMint(_owner, tokenId);
        _setTokenURI(tokenId, _tokenURI);
    }

    // emergency use : to be used if we replace our IPFS gateway
    function setBaseURI(string memory _uri) external onlyOwner {
        _setBaseURI(_uri);
    }

    function setContractURI(string memory _uri) external onlyOwner {
        _contractURI =  _uri;
    }

    function tokenExists(uint256 tokenId) external view returns (bool) {
        return _exists(tokenId);
    }



    function setAuth(address _op, bool status) external onlyOwner {
        operators[_op] = status;
        emit OperatorSet(_op,status);
    }

}