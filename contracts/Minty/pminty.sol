pragma solidity ^0.6.6;
//SPDX-License-Identifier: UNLICENSED

import "@maticnetwork/pos-portal/contracts/child/ChildToken/ChildERC20.sol";
import "@openzeppelin/contracts/math/SafeMath.sol";

abstract contract ERC677Receiver {
  function onTokenTransfer(address _sender, uint _value, bytes memory _data) public virtual;
}

abstract contract ERC677{
  function transferAndCall(address to, uint value, bytes memory data) public virtual returns (bool success);

  event ERC677Transfer(address indexed from, address indexed to, uint value, bytes data);
}

//contract Minty is ERC20, ERC677 {
contract pMinty is ChildERC20, ERC677 {
    using SafeMath for uint256;
    //
    //address   constant childChainManager = 0xb5505a6d998549090530911180f38aC5130101c6 ;
    //
    address   constant childChainManager = 0xA6FA4fB5f76172d178d61B04b0ecd319C5d1C0aa ;

    uint256 constant public MAX_SUPPLY = 100000000 ether;
    
    constructor() ChildERC20("Minty Art","MINTY", 18,childChainManager) public  {
        
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