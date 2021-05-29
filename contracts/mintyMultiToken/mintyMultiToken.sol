pragma solidity ^0.7.5;

import "@openzeppelin/contracts/token/ERC1155/ERC1155.sol";

contract mintyMultiToken is ERC1155 {

    mapping (uint => bool) public minted;
    mapping(uint => string) _tokenURIs;
    string                  base;
    address                 public owner;
    mapping (address => bool) auth;


    modifier onlyAuth() {
        require (auth[msg.sender] || (msg.sender == owner),"unauthorised");
        _;
    }

    constructor(address _owner) ERC1155("") {
        owner = _owner;
        base = "https://minty.mypinata.cloud/ipfs/";
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

    function setAuth(address operator, bool state) external onlyAuth {
        auth[operator] = state;

    }
}