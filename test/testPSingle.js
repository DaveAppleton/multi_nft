//const { ZERO_ADDRESS, ROLE, Data } = require('./helpers/common');
const { BigNumber } = require("@ethersproject/bignumber");
const { expect } = require("chai");
const { ethers, network } = require("hardhat");
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

    let WETH 
    let weth

    let lock100

    let alice
 
    let owner
    let addr1
    let addr2
    let addr3
    let addrs
    // let addr2
    // let addr3

    

    it("beforeAll",async function() {
        if (network.name != "hardhat") {
            console.log("PLEASE USE --network hardhat")
            process.exit(0)
        }
        console.log("start");
        [owner,addr1, addr2,addr3, ...addrs] = await ethers.getSigners();
        console.log("OWNER",owner.address)
        console.log("ADDR1",addr1.address)

        console.log("deploy WETH")
        WETH = await ethers.getContractFactory("WETH9")
        weth = await WETH.deploy()

        lock100 = ethers.utils.parseEther("100.0")
       
        alice = "0x31EFd75bc0b5fbafc6015Bd50590f4fDab6a3F22"

        //SALE = await ethers.getContractFactory("pMintysale")

        console.log("Start...")

         M721    = await ethers.getContractFactory("contracts/flat/pMintyUnique.sol:pMintyUnique")
         SALE    = await ethers.getContractFactory("contracts/flat/psale.sol:pMintysale")
    
        console.log("deploy sale")
        sale = await SALE.deploy(weth.address,alice,900,100,1025)

        console.log("deploy M721")
        m721 = await M721.deploy(sale.address)
        console.log("ERC721 ",m721.address)
        console.log


        await(sale.setMintyUnique(m721.address));
    })

    it("set contract URI", async function() {
        let str = "beans on toast"
        await m721.setContractURI(str)
        expect(await m721.contractURI()).to.equal(str)
    })

    it("cannot offer an NFT for sale", async function() { 
        //await m721.setAuth(sale.address,true);
        await expect(sale.connect(addr1).offerNew(10,"1212",ethers.utils.parseEther("0.1025"))).to.be.revertedWith('You are not allowed to mint tokens here')
        expect(await sale.available(10)).to.equal(false);
    })

    it("offer an NFT for sale after being approved", async function() { 
        await sale.setMinter(addr1.address,true);
        await expect(sale.connect(addr1).offerNew(10,"1212",ethers.utils.parseEther("0.1025"))).to.emit(sale, 'NewOffer')
        expect(await sale.available(10)).to.equal(true);
    })

    it("No 2 & No 3 buy 100 WETH", async function() {
        await weth.connect(addr1).deposit({ value: lock100})
        expect(await weth.balanceOf(addr1.address)).to.equal(ethers.utils.parseEther("100.0"))
        await weth.connect(addr2).deposit({ value: lock100})
        expect(await weth.balanceOf(addr2.address)).to.equal(ethers.utils.parseEther("100.0"))
        await weth.connect(addr3).deposit({ value: lock100})
        expect(await weth.balanceOf(addr3.address)).to.equal(ethers.utils.parseEther("100.0"))
    })


    
    it("No 2 buys it", async function() {
        await weth.connect(addr2).approve(sale.address,ethers.utils.parseEther("100.0"))
        await expect( sale.connect(addr2).acceptOffer(10)).to.emit(m721, 'Transfer')
        expect(await m721.balanceOf(addr1.address)).to.equal(0);
        expect(await m721.balanceOf(addr2.address)).to.equal(1);
        console.log("seller : ", ethers.utils.formatEther(await weth.balanceOf(addr1.address)))
        console.log("buyer : ",  ethers.utils.formatEther(await weth.balanceOf(addr2.address)))
        console.log("minty : ",  ethers.utils.formatEther(await weth.balanceOf(alice)))
    })

    it("No 2 cannot resell (approve)", async function() { 
        await expect(sale.connect(addr2).offerResale(10,ethers.utils.parseEther("0.5"))).to.be.revertedWith("You have not approved this contract to sell your tokens");
    })

    it("No 2 cannot resell (doesn't exist)", async function() { 
        await expect(sale.connect(addr2).offerResale(11,ethers.utils.parseEther("0.5"))).to.be.revertedWith("Token does not exist!");
    })

    it("No 2 cannot resell (ain't yours)", async function() { 
        await weth.connect(addr3).approve(sale.address,ethers.utils.parseEther("100.0"))
        await sale.connect(addr1).offerNew(11,"1212",ethers.utils.parseEther("0.1"))
        await expect( sale.connect(addr3).acceptOffer(11)).to.emit(m721, 'Transfer')
        console.log("try to offer something we don't have")
        await expect(sale.connect(addr2).offerResale(11,ethers.utils.parseEther("0.5"))).to.be.revertedWith("You do not own this token");
    })


    it("No 2 can resell after approve", async function() { 
        await m721.connect(addr2).setApprovalForAll(sale.address,true);
        await expect(sale.connect(addr2).offerResale(10,ethers.utils.parseEther("0.5"))).to.emit(sale,'ResaleOffer');
    })

    it("No 3 buys it", async function() {
        expect(await sale.available(10)).to.equal(true);
        await expect( sale.connect(addr3).acceptOffer(10)).to.emit(m721, 'Transfer')
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
        await weth.connect(addr1).approve(sale.address,ethers.utils.parseEther("1.4"))
        await expect(sale.connect(addr3).reSubmitOffer(10,ethers.utils.parseEther("1.5"))).to.emit(sale,'SaleResubmitted')
        await expect( sale.connect(addr1).acceptOffer(10)).to.be.reverted;
        await weth.connect(addr1).approve(sale.address,ethers.utils.parseEther("1.5"))
        await expect( sale.connect(addr1).acceptOffer(10)).to.emit(m721, 'Transfer')
    })

    it("Check when No 3 does something naughty", async function() {
        await weth.connect(addr1).approve(sale.address,ethers.utils.parseEther("1.0"))
        expect(await m721.ownerOf(11)).to.equal(addr3.address);
        await expect(sale.connect(addr3).offerResale(11,ethers.utils.parseEther("1.025"))).to.emit(sale,'ResaleOffer');
        expect(await sale.available(11)).to.equal(true);
        await (expect(m721.connect(addr3).transferFrom(addr3.address,addr1.address,11))).to.emit(m721,'Transfer')
        expect(await m721.ownerOf(11)).to.equal(addr1.address);
        expect(await sale.available(11)).to.equal(false);  // sale should check balance
        await expect( sale.connect(addr1).acceptOffer(11)).to.be.reverted;
    })

    it("Transfer it back", async function() {
        await (expect(m721.connect(addr1).transferFrom(addr1.address,addr3.address,11))).to.emit(m721,'Transfer')
        expect(await m721.ownerOf(11)).to.equal(addr3.address);
    })

    it("number 2 buys it", async function() {
        expect(await m721.artist(11)).to.equal(addr1.address)
        let bal1b4 = await weth.balanceOf(addr1.address)
        let bal2b4 = await weth.balanceOf(addr2.address)
        let bal3b4 = await weth.balanceOf(addr3.address)
        let balAb4 = await weth.balanceOf(alice)
        await expect( sale.connect(addr2).acceptOffer(11)).to.emit(m721, 'Transfer')
        let bal1x = await weth.balanceOf(addr1.address)
        let bal2x = await weth.balanceOf(addr2.address)
        let bal3x = await weth.balanceOf(addr3.address)
        let balAx = await weth.balanceOf(alice)

        console.log("addr1 ^",ethers.utils.formatEther(bal1x.sub(bal1b4)))
        console.log("addr2 v",ethers.utils.formatEther(bal2b4.sub(bal2x)))
        console.log("addr3 ^",ethers.utils.formatEther(bal3x.sub(bal3b4)))
        console.log("alice ^",ethers.utils.formatEther(balAx.sub(balAb4)))

    })

    async function bid(signer, amount, residual) {
        
    }
})