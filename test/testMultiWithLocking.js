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
    let locking1
    let locking2

    let M1155
    let m1155

    let SALE 
    let sale

    let lock100
    let lock200
 
    let owner
    let addr1
    let addr2
    let addr3
    let addr0
    let addrs
    // let addr2
    // let addr3

    

    it("beforeAll",async function() {
        console.log("start");
        [owner,addr1, addr2,addr3,addr0, ...addrs] = await ethers.getSigners();
        console.log("OWNER",owner.address)
        console.log("ADDR0",addr0.address)
        console.log("ADDR1",addr1.address)
        console.log("ADDR2",addr2.address)
        console.log("ADDR3",addr3.address)
       
        lock100 = ethers.utils.parseEther("100.0")
        lock200 = ethers.utils.parseEther("200.0")

        alice = "0x31EFd75bc0b5fbafc6015Bd50590f4fDab6a3F22"

        MINTY = await ethers.getContractFactory("Minty")
        minty = await MINTY.deploy(owner.address)
        console.log("MINTY ",minty.address)

        LOCKING = await ethers.getContractFactory("locking")
        locking1 = await LOCKING.deploy(minty.address,lock100)
        locking2 = await LOCKING.deploy(minty.address,lock100)
        console.log("locking1",locking1.address)
        console.log("locking2",locking2.address)

        M1155 = await ethers.getContractFactory("mintyMultiToken")
 
        m1155 = await M1155.deploy(owner.address,[locking1.address,locking2.address],"You need to be a MINTY patron or TEST PROJECT patron")
        console.log("ERC1155 ",m1155.address)
        console.log

        SALE = await ethers.getContractFactory("mintyMultiSale")
        sale = await SALE.deploy(owner.address, 10, 10)

        await m1155.connect(owner).setApprovalForAll(sale.address,true);

        await expect(minty.transfer(addr0.address,lock200)).to.emit(minty,'Transfer')
        await expect(minty.transfer(addr1.address,lock200)).to.emit(minty,'Transfer')
        await expect(minty.transfer(addr2.address,lock200)).to.emit(minty,'Transfer')
        await expect(minty.transfer(addr3.address,lock200)).to.emit(minty,'Transfer')

    })


    it("lock 1 2 and 3", async function() { 
        await expect(minty.connect(addr1).transferAndCall(locking1.address,lock100,"0x")).to.emit(locking1,'LockedBaseToken').withArgs(addr1.address, lock100);
        await expect(minty.connect(addr2).transferAndCall(locking2.address,lock100,"0x")).to.emit(locking2,'LockedBaseToken').withArgs(addr2.address, lock100);
        await expect(minty.connect(addr3).transferAndCall(locking1.address,lock100,"0x")).to.emit(locking1,'LockedBaseToken').withArgs(addr3.address, lock100);
        await expect(minty.connect(addr3).transferAndCall(locking2.address,lock100,"0x")).to.emit(locking2,'LockedBaseToken').withArgs(addr3.address, lock100);
    })

    it("check lock status", async function() {
        expect(await locking1.isLocked(addr0.address)).to.equal(false)
        expect(await locking2.isLocked(addr0.address)).to.equal(false)
        expect(await locking1.isLocked(addr1.address)).to.equal(true)
        expect(await locking2.isLocked(addr1.address)).to.equal(false)
        expect(await locking1.isLocked(addr2.address)).to.equal(false)
        expect(await locking2.isLocked(addr2.address)).to.equal(true)
        expect(await locking1.isLocked(addr3.address)).to.equal(true)
        expect(await locking2.isLocked(addr3.address)).to.equal(true)
    })

    it("check validate", async function() {
        await expect(m1155.validateBuyer(addr0.address)).to.be.revertedWith("You need to be a MINTY patron or TEST PROJECT patron");
        await m1155.validateBuyer(addr1.address)
        await m1155.validateBuyer(addr2.address)
        await m1155.validateBuyer(addr3.address)
    })

    it("offer an NFT for sale", async function() { 
        await m1155.setAuth(sale.address,true);
        await sale.offerNew(m1155.address,10,"1212",200,ethers.utils.parseEther("0.1"))
        expect(await m1155.balanceOf(owner.address,10)).to.equal(200);
    })

 
    it("No 0 cannot buy 5", async function() {
        let override = {
            value: ethers.utils.parseEther("0.5")
        };
        await expect( sale.connect(addr0).acceptOffer(m1155.address,10,0,5,override)).to.be.revertedWith("You need to be a MINTY patron or TEST PROJECT patron");
    })

    it("No 1 can buy 5", async function() {
        let override = {
            value: ethers.utils.parseEther("0.5")
        };
        await expect( sale.connect(addr1).acceptOffer(m1155.address,10,0,5,override)).to.emit(m1155, 'TransferSingle')
    })

    it("No 2 can buy 5", async function() {
        let override = {
            value: ethers.utils.parseEther("0.5")
        };
        await expect( sale.connect(addr2).acceptOffer(m1155.address,10,0,5,override)).to.emit(m1155, 'TransferSingle')
    })

    it("No 3 can buy 5", async function() {
        let override = {
            value: ethers.utils.parseEther("0.5")
        };
        await expect( sale.connect(addr3).acceptOffer(m1155.address,10,0,5,override)).to.emit(m1155, 'TransferSingle')
    })




    it("No 2 cannot resell (approve)", async function() { 
        await expect(sale.connect(addr2).offerResale(m1155.address,10,3,ethers.utils.parseEther("0.5"))).to.be.revertedWith("You have not approved this contract to sell your tokens");
    })

    it("No 2 cannot resell (quantity)", async function() { 
        await expect(sale.connect(addr2).offerResale(m1155.address,10,13,ethers.utils.parseEther("0.5"))).to.be.revertedWith("You do not own enough of this token");
    })

    it("No 2 can resell after approve", async function() { 
        await m1155.connect(addr2).setApprovalForAll(sale.address,true);
        await expect(sale.connect(addr2).offerResale(m1155.address,10,3,ethers.utils.parseEther("0.5"))).to.emit(sale,'ResaleOffer');
    })

    it("No 3 buys one", async function() {
        let override = {
            value: ethers.utils.parseEther("0.5")
        };
        await expect( sale.connect(addr3).acceptOffer(m1155.address,10,1,1,override)).to.emit(m1155, 'TransferSingle')
    })

    it("No 3 fails to buy three", async function() {
        let override = {
            value: ethers.utils.parseEther("1.5")
        };
        await expect( sale.connect(addr3).acceptOffer(m1155.address,10,1,3,override)).to.be.revertedWith("not enough items available");
    })

    it("Check when No 2 does something naughty", async function() {
        expect(await m1155.balanceOf(addr2.address,10)).to.equal(4);
        expect(await sale.available(m1155.address,10,1)).to.equal(2);
        await (expect(m1155.connect(addr2).safeTransferFrom(addr2.address,addr1.address,10,3,"0x"))).to.emit(m1155,'TransferSingle')
        expect(await m1155.balanceOf(addr2.address,10)).to.equal(1);
        expect(await sale.available(m1155.address,10,1)).to.equal(1);  // sale should check balance
        await (m1155.connect(addr2).setApprovalForAll(sale.address,false));
        expect(await sale.available(m1155.address,10,1)).to.equal(0);  // sale checks that approval is in place
    })

    async function bid(signer, amount, residual) {
        
    }
})