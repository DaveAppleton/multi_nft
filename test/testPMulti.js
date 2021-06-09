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

    let WETH 
    let weth 

    let MINTY 
    let minty 

    let LOCKING 
    let locking1 
    let locking2 
    let lock100
    let lock200

    let creator
    
    let patron
    

 
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

        M1155 = await ethers.getContractFactory("pMintyMultiToken")

        WETH = await ethers.getContractFactory("WETH9")
        weth = await WETH.deploy()

        console.log("deploy sale")
        SALE = await ethers.getContractFactory("pMintyMultiSale")
        sale = await SALE.deploy(alice, weth.address, 1025)
    })

    it ("deploy locking", async function(){
        MINTY = await ethers.getContractFactory("Minty")
        minty = await MINTY.deploy(owner.address)
        console.log("MINTY ",minty.address)

        lock100 = ethers.utils.parseEther("100.0")
        lock200 = ethers.utils.parseEther("200.0")

        LOCKING = await ethers.getContractFactory("contracts/locking/locking.sol:locking")
        locking1 = await LOCKING.deploy(minty.address,lock100)
        locking2 = await LOethers.utils.parseEther("100.0")CKING.deploy(minty.address,lock100)
        console.log("locking1",locking1.address)
        console.log("locking2",locking2.address)

    })

    it("deploy m1155", async function(){
        creator = "0x39d07f321cAF5b0668459DB5Bcf039A462A9273d" // Virgo
        patron = "0x08ef5C04D2F93d9858F0d87697f22cBeCD990076" // Libra
        creator1 = { beneficiary: creator, share: 500 }
        creatorN = { beneficiary: creator, share: 750 }
        patron1 = { beneficiary: patron, share: 500 }
        patronN = { beneficiary: patron, share: 250 }
        console.log("deploy token")
        m1155 = await M1155.deploy(
            addr1.address,
            sale.address,
            [locking1.address,locking2.address],
            "You need to be a MINTY patron or TEST PROJECT patron",
            100,
            [creator1,patron1],
            [creatorN,patronN]
            )
        console.log("ERC1155 ",m1155.address)
        console.log

        await expect(minty.transfer(addr1.address,lock200)).to.emit(minty,'Transfer')
        await expect(minty.transfer(addr2.address,lock200)).to.emit(minty,'Transfer')
        await expect(minty.transfer(addr3.address,lock200)).to.emit(minty,'Transfer')


        m1155.connect(addr1).setApprovalForAll(sale.address,true);
    })

    it("check owner", async function() { 
        expect(await m1155.owner()).to.equal(addr1.address);
    })

    it("set contract URI", async function() {
        let str = "beans on toast"
        await m1155.connect(addr1).setContractURI(str)
        expect(await m1155.contractURI()).to.equal(str)
    })

    it("offer an NFT for sale at 0.1 + 0.025 ", async function() { 
        await m1155.connect(addr1).setAuth(sale.address,true);
        await sale.connect(addr1).offerNew(m1155.address,10,"1212",200,ethers.utils.parseEther("0.1025"))
        expect(await m1155.balanceOf(addr1.address,10)).to.equal(200);
    })

    it("No 2 buys 100 WETH", async function() {
        await weth.connect(addr2).deposit({ value: lock100})
        expect(await weth.balanceOf(addr2.address)).to.equal(ethers.utils.parseEther("100.0"))
    })

    it("No 2 allows sale to use 1 WETH", async function() {
        await weth.connect(addr2).approve(sale.address,ethers.utils.parseEther("100.0"))
        console.log("allowance:",ethers.utils.formatEther(await weth.allowance(addr2.address,sale.address)))
    })

    it("No 2 fails to buy five before locking", async function() {
        await expect( sale.connect(addr2).acceptOffer(m1155.address,10,0,5)).to.be.revertedWith('You need to be a MINTY patron or TEST PROJECT patron')
    })

    it("lock 1 2 and 3", async function() { 
        await expect(minty.connect(addr1).transferAndCall(locking1.address,lock100,"0x")).to.emit(locking1,'LockedBaseToken').withArgs(addr1.address, lock100);
        await expect(minty.connect(addr2).transferAndCall(locking2.address,lock100,"0x")).to.emit(locking2,'LockedBaseToken').withArgs(addr2.address, lock100);
        await expect(minty.connect(addr3).transferAndCall(locking1.address,lock100,"0x")).to.emit(locking1,'LockedBaseToken').withArgs(addr3.address, lock100);
        await expect(minty.connect(addr3).transferAndCall(locking2.address,lock100,"0x")).to.emit(locking2,'LockedBaseToken').withArgs(addr3.address, lock100);
    })

    it("No 2 buys five (5 x 0.1 = 0.5)", async function() {
        console.log("buyer before ",ethers.utils.formatEther(await weth.balanceOf(addr2.address)))
        await expect( sale.connect(addr2).acceptOffer(m1155.address,10,0,5)).to.emit(m1155, 'TransferSingle')
        expect(await m1155.balanceOf(addr1.address,10)).to.equal(195);
        expect(await m1155.balanceOf(addr2.address,10)).to.equal(5);
        console.log("buyer ",ethers.utils.formatEther(await weth.balanceOf(addr2.address)))
        console.log("seller ",ethers.utils.formatEther(await weth.balanceOf(addr1.address)))
        console.log("patron ",ethers.utils.formatEther(await weth.balanceOf(patron)))
        console.log("creator ",ethers.utils.formatEther(await weth.balanceOf(creator)))
        console.log("minty ",ethers.utils.formatEther(await weth.balanceOf(alice)))
        //expect(await weth.balanceOf(addr2.address)).to.equal(ethers.utils.parseEther("99.5"))
    })


    it("No 2 cannot resell (approve)", async function() { 
        await expect(sale.connect(addr2).offerResale(m1155.address,10,3,ethers.utils.parseEther("0.5"))).to.be.revertedWith("You have not approved this contract to sell your tokens");
    })

    it("No 2 cannot resell (quantity)", async function() { 
        await expect(sale.connect(addr2).offerResale(m1155.address,10,13,ethers.utils.parseEther("0.5"))).to.be.revertedWith("You do not own enough of this token");
    })

    it("No 2 can resell after approve", async function() { 
        await m1155.connect(addr2).setApprovalForAll(sale.address,true);
        await expect(sale.connect(addr2).offerResale(m1155.address,10,3,ethers.utils.parseEther("0.5125"))).to.emit(sale,'ResaleOffer');
    })

    it("No 3 buys one", async function() {
        await weth.connect(addr3).deposit({ value: lock100})
        await weth.connect(addr3).approve(sale.address,ethers.utils.parseEther("100.0"))
        console.log("allowance:",ethers.utils.formatEther(await weth.allowance(addr3.address,sale.address)))
        await expect( sale.connect(addr3).acceptOffer(m1155.address,10,1,1)).to.emit(m1155, 'TransferSingle')
        expect(await m1155.balanceOf(addr2.address,10)).to.equal(4);
        expect(await m1155.balanceOf(addr3.address,10)).to.equal(1);
        expect(await sale.available(m1155.address,10,1)).to.equal(2);
    })

    it("No 3 fails to buy three", async function() {
        await expect( sale.connect(addr3).acceptOffer(m1155.address,10,1,3)).to.be.revertedWith("not enough items available");
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