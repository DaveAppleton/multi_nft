pragma solidity ^0.7.0;
//SPDX-License-Identifier: UNLICENSED

import "@openzeppelin/contracts/token/ERC20/ERC20.sol";

abstract contract ERC677Receiver {
  function onTokenTransfer(address _sender, uint _value, bytes memory _data) public virtual;
}

abstract contract ERC677{
  function transferAndCall(address to, uint value, bytes memory data) public virtual returns (bool success);

  event ERC677Transfer(address indexed from, address indexed to, uint value, bytes data);
}

contract dummyWeth is ERC20, ERC677 {
    using SafeMath for uint256;

    uint256 constant public MAX_SUPPLY = 1000000 ether;
    bool             public minting = true;

    mapping(address => bool) public operators;

    constructor(address mintysig) ERC20("Wrapped Ether","WETH") {
        _mint(mintysig, 100000000 ether);
    }


    function transferAndCall(address _to, uint _value, bytes memory _data)
        public override
        returns (bool success)
    {
        super.transfer(_to, _value);
        emit ERC677Transfer(msg.sender, _to, _value, _data);
        if (isContract(_to)) {
            contractFallback(_to, _value, _data);
        }
        return true;
    }

  function contractFallback(address _to, uint _value, bytes memory _data)  private
  {
    ERC677Receiver receiver = ERC677Receiver(_to);
    receiver.onTokenTransfer(msg.sender, _value, _data);
  }

  function isContract(address _addr)
    private view
    returns (bool hasCode)
  {
    uint length;
    assembly { length := extcodesize(_addr) }
    return length > 0;
  }
}