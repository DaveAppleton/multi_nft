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
    
    let MINTY
    let minty

    let LOCKING 
    let locking
 
    let owner
    let addr1
    let addr2
    let addr3
    let addrs

    let lock100
    let lock200
    let lock300
    // let addr2
    // let addr3

    

    it("beforeAll",async function() {
        console.log("start");
        [owner,addr1, addr2,addr3, ...addrs] = await ethers.getSigners();
        console.log("OWNER",owner.address)
        console.log("ADDR1",addr1.address)

        lock100 = ethers.utils.parseEther("100.0")
        lock200 = ethers.utils.parseEther("200.0")
        lock300 = ethers.utils.parseEther("300.0")
       
        alice = "0x31EFd75bc0b5fbafc6015Bd50590f4fDab6a3F22"

        MINTY = await ethers.getContractFactory("Minty")
 
        minty = await MINTY.deploy(owner.address)
        console.log("MINTY ",minty.address)
        console.log

        LOCKING = await ethers.getContractFactory("locking")
        locking = await LOCKING.deploy(minty.address,lock100)

        await expect(minty.transfer(addr1.address,lock100)).to.emit(minty,'Transfer')
        await expect(minty.transfer(addr2.address,lock200)).to.emit(minty,'Transfer')
        await expect(minty.transfer(addr3.address,lock300)).to.emit(minty,'Transfer')
    })

    it("Lock via transferAndCall", async function() {
        await expect(minty.connect(addr1).transferAndCall(locking.address,lock100,"0x")).to.emit(locking,'LockedBaseToken').withArgs(addr1.address, lock100);
        ld1 = await locking.LockData(addr1.address)
        console.log( ethers.utils.formatEther(ld1.amount))
    })

    it("Lock via transfer", async function() {
        await expect(minty.connect(addr2).approve(locking.address,lock300)).to.emit(minty,'Approval')
        await expect(locking.connect(addr2).lock()).to.emit(locking,'LockedBaseToken').withArgs(addr2.address, lock100);
        ld2 = await locking.LockData(addr2.address)
        console.log( ethers.utils.formatEther(ld2.amount))
    })


    it("Re Lock via transferAndCall", async function() { 
        await expect(minty.connect(addr2).transferAndCall(locking.address,lock100,"0x")).to.be.revertedWith("You already have locked tokens")
    })

    it("Re Lock via transfer", async function() { 
        await expect(locking.connect(addr2).lock()).to.be.revertedWith("You already have locked tokens")
    })

    it("No 1 unlocks", async function() {
        await expect(locking.connect(addr1).unlock()).to.emit(locking,"UnlockRequested").withArgs(addr1.address);
    })

    it("No 1 unlocks AGAIN", async function() {
        await expect(locking.connect(addr1).unlock()).to.be.revertedWith("release already requested but not ready")
    })

    it("No 1 unlocks AGAIN after a week", async function() {
        ethers.provider.send("evm_increaseTime", [60*60*24*7])   // 1 week
        console.log("addr1 before ",ethers.utils.formatEther(await minty.balanceOf(addr1.address)))
        await expect(locking.connect(addr1).unlock()).to.emit(locking,"WithdrawBaseTokens").withArgs(addr1.address, lock100);
        console.log("addr1 after ",ethers.utils.formatEther(await minty.balanceOf(addr1.address)))
    })


    /*
    it("No 2 buys it", async function() {
        let override = {
            value: ethers.utils.parseEther("0.1")
        };
        await expect( sale.connect(addr2).acceptOffer(10,override)).to.emit(minty, 'Transfer')
        expect(await minty.balanceOf(addr1.address)).to.equal(0);
        expect(await minty.balanceOf(addr2.address)).to.equal(1);
    })

    it("No 2 cannot resell (approve)", async function() { 
        await expect(sale.connect(addr2).offerResale(10,ethers.utils.parseEther("0.5"))).to.be.revertedWith("You have not approved this contract to sell your tokens");
    })

    it("No 2 cannot resell (doesn't exist)", async function() { 
        await expect(sale.connect(addr2).offerResale(11,ethers.utils.parseEther("0.5"))).to.be.revertedWith("Token does not exist!");
    })

    it("No 2 cannot resell (ain't yours)", async function() { 
        let override = {
            value: ethers.utils.parseEther("0.1")
        };
        await sale.connect(addr1).offerNew(11,"1212",ethers.utils.parseEther("0.1"))
        await expect( sale.connect(addr3).acceptOffer(11,override)).to.emit(minty, 'Transfer')
        await expect(sale.connect(addr2).offerResale(11,ethers.utils.parseEther("0.5"))).to.be.revertedWith("You do not own this token");
    })


    it("No 2 can resell after approve", async function() { 
        await minty.connect(addr2).setApprovalForAll(sale.address,true);
        await expect(sale.connect(addr2).offerResale(10,ethers.utils.parseEther("0.5"))).to.emit(sale,'ResaleOffer');
    })

    it("No 3 buys it", async function() {
        let override = {
            value: ethers.utils.parseEther("0.5")
        };
        expect(await sale.available(10)).to.equal(true);
        await expect( sale.connect(addr3).acceptOffer(10,override)).to.emit(minty, 'Transfer')
        expect(await minty.ownerOf(10)).to.equal(addr3.address);
        expect(await sale.available(10)).to.equal(false);
    })

    it("No 3 puts on sale then changes price", async function() {
        await minty.connect(addr3).setApprovalForAll(sale.address,true);
        await expect(sale.connect(addr3).offerResale(10,ethers.utils.parseEther("0.5"))).to.emit(sale,'ResaleOffer');
    })

    it("wrong party tries to resubmit offer", async function() {
        await expect(sale.connect(addr2).reSubmitOffer(10,ethers.utils.parseEther("1.5"))).to.be.revertedWith("Unauthorised");
    })

    it("correct party tries to resubmit offer", async function() {
        let override = {
            value: ethers.utils.parseEther("0.5")
        };
        await expect(sale.connect(addr3).reSubmitOffer(10,ethers.utils.parseEther("1.5"))).to.emit(sale,'SaleResubmitted')
        await expect( sale.connect(addr1).acceptOffer(10,override)).to.be.revertedWith("Price not met");
        override.value = ethers.utils.parseEther("1.5")
        await expect( sale.connect(addr1).acceptOffer(10,override)).to.emit(minty, 'Transfer')
    })

    it("Check when No 3 does something naughty", async function() {
        let override = {
            value: ethers.utils.parseEther("0.5")
        };
        expect(await minty.ownerOf(11)).to.equal(addr3.address);
        await expect(sale.connect(addr3).offerResale(11,ethers.utils.parseEther("0.5"))).to.emit(sale,'ResaleOffer');
        expect(await sale.available(11)).to.equal(true);
        await (expect(minty.connect(addr3).transferFrom(addr3.address,addr1.address,11))).to.emit(minty,'Transfer')
        expect(await minty.ownerOf(11)).to.equal(addr1.address);
        expect(await sale.available(11)).to.equal(false);  // sale should check balance
        await expect( sale.connect(addr1).acceptOffer(11,override)).to.be.revertedWith("Item not owned by offerer");
    })

    async function bid(signer, amount, residual) {
        
    }
    */
})