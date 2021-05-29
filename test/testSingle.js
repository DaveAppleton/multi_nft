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
    
    let S721
    let s721

    let SALE 
    let sale
 
    let owner
    let addr1
    let addr2
    let addr3
    let addrs
    // let addr2
    // let addr3

    

    it("beforeAll",async function() {
        console.log("start");
        [owner,addr1, addr2,addr3, ...addrs] = await ethers.getSigners();
        console.log("OWNER",owner.address)
        console.log("ADDR1",addr1.address)
       
        alice = "0x31EFd75bc0b5fbafc6015Bd50590f4fDab6a3F22"

        M721 = await ethers.getContractFactory("MintyUnique")
 
        m721 = await M721.deploy()
        console.log("ERC721 ",m721.address)
        console.log

        SALE = await ethers.getContractFactory("mintysale")
        sale = await SALE.deploy(m721.address,alice, 10, 10)

        m721.connect(addr1).setApprovalForAll(sale.address,true);
    })

    it("set contract URI", async function() {
        let str = "beans on toast"
        await m721.setContractURI(str)
        expect(await m721.contractURI()).to.equal(str)
    })

    it("offer an NFT for sale", async function() { 
        await m721.setAuth(sale.address,true);
        await sale.connect(addr1).offerNew(10,"1212",ethers.utils.parseEther("0.1"))
        expect(await sale.available(10)).to.equal(true);
    })

 
    it("No 2 buys it", async function() {
        let override = {
            value: ethers.utils.parseEther("0.1")
        };
        await expect( sale.connect(addr2).acceptOffer(10,override)).to.emit(m721, 'Transfer')
        expect(await m721.balanceOf(addr1.address)).to.equal(0);
        expect(await m721.balanceOf(addr2.address)).to.equal(1);
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
        await expect( sale.connect(addr3).acceptOffer(11,override)).to.emit(m721, 'Transfer')
        await expect(sale.connect(addr2).offerResale(11,ethers.utils.parseEther("0.5"))).to.be.revertedWith("You do not own this token");
    })


    it("No 2 can resell after approve", async function() { 
        await m721.connect(addr2).setApprovalForAll(sale.address,true);
        await expect(sale.connect(addr2).offerResale(10,ethers.utils.parseEther("0.5"))).to.emit(sale,'ResaleOffer');
    })

    it("No 3 buys it", async function() {
        let override = {
            value: ethers.utils.parseEther("0.5")
        };
        expect(await sale.available(10)).to.equal(true);
        await expect( sale.connect(addr3).acceptOffer(10,override)).to.emit(m721, 'Transfer')
        expect(await m721.ownerOf(10)).to.equal(addr3.address);
        expect(await sale.available(10)).to.equal(false);
    })

    it("No 3 puts on sale then changes price", async function() {
        await m721.connect(addr3).setApprovalForAll(sale.address,true);
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
        await expect( sale.connect(addr1).acceptOffer(10,override)).to.emit(m721, 'Transfer')
    })

    it("Check when No 3 does something naughty", async function() {
        let override = {
            value: ethers.utils.parseEther("0.5")
        };
        expect(await m721.ownerOf(11)).to.equal(addr3.address);
        await expect(sale.connect(addr3).offerResale(11,ethers.utils.parseEther("0.5"))).to.emit(sale,'ResaleOffer');
        expect(await sale.available(11)).to.equal(true);
        await (expect(m721.connect(addr3).transferFrom(addr3.address,addr1.address,11))).to.emit(m721,'Transfer')
        expect(await m721.ownerOf(11)).to.equal(addr1.address);
        expect(await sale.available(11)).to.equal(false);  // sale should check balance
        await expect( sale.connect(addr1).acceptOffer(11,override)).to.be.revertedWith("Item not owned by offerer");
    })

    async function bid(signer, amount, residual) {
        
    }
})