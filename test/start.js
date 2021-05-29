//const { ZERO_ADDRESS, ROLE, Data } = require('./helpers/common');
const { BigNumber } = require("@ethersproject/bignumber");
const { expect } = require("chai");
const { ethers } = require("hardhat");
//const { time } = require('@openzeppelin/test-helpers');


    // evm_increaseTime
    // // suppose the current block has a timestamp of 01:00 PM
    // await network.provider.send("evm_increaseTime", [3600])
    // await network.provider.send("evm_mine") // this one will have 02:00 PM as its timestamp

    // evm_setNextBlockTimestamp
    // await network.provider.send("evm_setNextBlockTimestamp", [1625097600])
    // await network.provider.send("evm_mine") // this one will have 2021-07-01 12:00 AM as its timestamp, no matter what the previous block has

describe("Token contract", function () {
    
    let M1155
    let m1155

    let SALE 
    let sale
 
    let owner
    let addr1
    let addr2
    let addrs
    // let addr2
    // let addr3

    

    beforeEach (async function() {
        console.log("start");
        [owner,addr1, addr2, ...addrs] = await ethers.getSigners();
        console.log("OWNER",owner.address)
        console.log("ADDR1",addr1.address)
       
        alice = "0x31EFd75bc0b5fbafc6015Bd50590f4fDab6a3F22"

        M1155 = await ethers.getContractFactory("mintyMultiToken")
 
        m1155 = await M1155.deploy(addr1.address)
        console.log("ERC1155 ",m1155.address)
        console.log

        SALE = await ethers.getContractFactory("mintyMultiSale")
        sale = await SALE.deploy(alice, 10, 10)
    })

    it("check owner", async function() { 
        expect(await m1155.owner()).to.equal(addr1.address);
    })

    it("bid after it starts", async function() { 
       await sale.connect(addr1).offerNew(m1155.address,10,"1212",200,100)

    })

 
    it("16 bids", async function() {

    })

    it("bid after it ends", async function() { 
        
    })


    async function bid(signer, amount, residual) {
        
    }
})