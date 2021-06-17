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
    let musician
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

        M1155 = await ethers.getContractFactory("contracts/flat/pMintyMultiToken.sol:pMintyMultiToken")

        WETH = await ethers.getContractFactory("WETH9")
        weth = await WETH.deploy()

        console.log("deploy sale")
        SALE = await ethers.getContractFactory("contracts/flat/pMintyMultiSale.sol:pMintyMultiSale")
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
        locking2 = await LOCKING.deploy(minty.address,lock100)
        console.log("locking1",locking1.address)
        console.log("locking2",locking2.address)

    })

    it("deploy m1155", async function(){
        creator = "0x39d07f321cAF5b0668459DB5Bcf039A462A9273d" // Virgo
        patron = "0x08ef5C04D2F93d9858F0d87697f22cBeCD990076" // Libra
        musician = "0x11318D5cAB7f211c0386E803b454bBA8F482e739" // capricorn
        creator1 = { beneficiary: creator, share: 500 }
        patron1 = { beneficiary: patron, share: 500 }

        creator1R = { beneficiary: creator, share: 750 }
        patron1R = { beneficiary: patron, share: 250 }


        console.log("deploy token")
        m1155 = await M1155.deploy(
            addr1.address,
            sale.address,
            [locking1.address,locking2.address],
            "You need to be a MINTY patron or TEST PROJECT patron",
            100,
            [creator1,patron1],
            [creator1R,patron1R],
            "C+P"
            )
        console.log("ERC1155 ",m1155.address)
        

        await expect(minty.transfer(addr1.address,lock200)).to.emit(minty,'Transfer')
        await expect(minty.transfer(addr2.address,lock200)).to.emit(minty,'Transfer')
        await expect(minty.transfer(addr3.address,lock200)).to.emit(minty,'Transfer')


        await m1155.connect(addr1).setApprovalForAll(sale.address,true);
        console.log("IS APP 4 ALL",await m1155.isApprovedForAll(addr1.address,sale.address))
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
        await sale.connect(addr1).offerNew(m1155.address,10,"1212",200,ethers.utils.parseEther("0.1025"),0)
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
        buyerBefore = await weth.balanceOf(addr2.address)
        sellerBefore = await weth.balanceOf(addr1.address)
        patronBefore = await weth.balanceOf(patron)
        creatorBefore = await weth.balanceOf(creator)
        mintyBefore = await weth.balanceOf(alice)


        await expect( sale.connect(addr2).acceptOffer(m1155.address,10,0,5)).to.emit(m1155, 'TransferSingle')
        expect(await m1155.balanceOf(addr1.address,10)).to.equal(195);
        expect(await m1155.balanceOf(addr2.address,10)).to.equal(5);

        buyerAfter = await weth.balanceOf(addr2.address)
        sellerAfter = await weth.balanceOf(addr1.address)
        patronAfter = await weth.balanceOf(patron)
        creatorAfter = await weth.balanceOf(creator)
        mintyAfter = await weth.balanceOf(alice)

        console.log("buyer spends ",ethers.utils.formatEther(buyerBefore.sub(buyerAfter)))
        console.log("seller gains ",ethers.utils.formatEther(sellerAfter.sub(sellerBefore)))
        console.log("patron ",ethers.utils.formatEther(patronAfter.sub(patronBefore)))
        console.log("creator ",ethers.utils.formatEther(creatorAfter.sub(creatorBefore)))
        console.log("minty ",ethers.utils.formatEther(mintyAfter.sub(mintyBefore)))
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

        sellerBefore = await weth.balanceOf(addr2.address)
        buyerBefore = await weth.balanceOf(addr3.address)
        patronBefore = await weth.balanceOf(patron)
        creatorBefore = await weth.balanceOf(creator)
        mintyBefore = await weth.balanceOf(alice)


        await expect( sale.connect(addr3).acceptOffer(m1155.address,10,1,1)).to.emit(m1155, 'TransferSingle')
        expect(await m1155.balanceOf(addr2.address,10)).to.equal(4);
        expect(await m1155.balanceOf(addr3.address,10)).to.equal(1);
        expect(await sale.available(m1155.address,10,1)).to.equal(2);

        sellerAfter = await weth.balanceOf(addr2.address)
        buyerAfter = await weth.balanceOf(addr3.address)
        patronAfter = await weth.balanceOf(patron)
        creatorAfter = await weth.balanceOf(creator)
        mintyAfter = await weth.balanceOf(alice)

        console.log("buyer spend ",ethers.utils.formatEther(buyerBefore.sub(buyerAfter)))
        console.log("seller gains ",ethers.utils.formatEther(sellerAfter.sub(sellerBefore)))
        console.log("patron ",ethers.utils.formatEther(patronAfter.sub(patronBefore)))
        console.log("creator ",ethers.utils.formatEther(creatorAfter.sub(creatorBefore)))
        console.log("minty ",ethers.utils.formatEther(mintyAfter.sub(mintyBefore)))

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

    it("Have another pool", async function() {
        creator2 = { beneficiary: creator, share: 500 }
        musician2 = { beneficiary: musician, share: 300 }
        patron2 = { beneficiary: patron, share: 200 }

        creator2R = { beneficiary: creator, share: 700 }
        musician2R = { beneficiary: musician, share: 200 }
        patron2R = { beneficiary: patron, share: 100 }

        await m1155.connect(addr1).addPools([creator2,musician2,patron2],[creator2R,musician2R,patron2R],"C+P+M")

    })

    it("offer an NFT for sale at 0.2 + 0.005 ", async function() { 
        //await m1155.connect(addr1).setAuth(sale.address,true);
        await sale.connect(addr1).offerNew(m1155.address,11,"hash",180,ethers.utils.parseEther("0.205"),1)
        expect(await m1155.balanceOf(addr1.address,11)).to.equal(180);
    })


    it("No 2 buys five (5 x 0.2 = 1.0)", async function() {
        sellerBefore = await weth.balanceOf(addr1.address)
        buyerBefore = await weth.balanceOf(addr2.address)
        patronBefore = await weth.balanceOf(patron)
        creatorBefore = await weth.balanceOf(creator)
        musicianBefore = await weth.balanceOf(musician)
        mintyBefore = await weth.balanceOf(alice)
        await expect( sale.connect(addr2).acceptOffer(m1155.address,11,0,5)).to.emit(m1155, 'TransferSingle')
        expect(await m1155.balanceOf(addr1.address,11)).to.equal(175);
        expect(await m1155.balanceOf(addr2.address,11)).to.equal(5);
        sellerAfter = await weth.balanceOf(addr1.address)
        buyerAfter = await weth.balanceOf(addr2.address)
        patronAfter = await weth.balanceOf(patron)
        creatorAfter = await weth.balanceOf(creator)
        musicianAfter = await weth.balanceOf(musician)
        mintyAfter = await weth.balanceOf(alice)

        console.log("check balances")
        console.log("buyer spend ",ethers.utils.formatEther(buyerBefore.sub(buyerAfter)))
        console.log("seller gains ",ethers.utils.formatEther(sellerAfter.sub(sellerBefore)))
        console.log("patron ",ethers.utils.formatEther(patronAfter.sub(patronBefore)))
        console.log("creator ",ethers.utils.formatEther(creatorAfter.sub(creatorBefore)))
        console.log("musician ",ethers.utils.formatEther(musicianAfter.sub(musicianBefore)))
        console.log("minty ",ethers.utils.formatEther(mintyAfter.sub(mintyBefore)))
        //expect(await weth.balanceOf(addr2.address)).to.equal(ethers.utils.parseEther("99.5"))
    })

    it("No 2 sells and 3 buys five (5 x 0.2 = 1.0) AGAIN", async function() {
        sellerBefore = await weth.balanceOf(addr2.address)
        buyerBefore = await weth.balanceOf(addr3.address)
        patronBefore = await weth.balanceOf(patron)
        creatorBefore = await weth.balanceOf(creator)
        musicianBefore = await weth.balanceOf(musician)
        mintyBefore = await weth.balanceOf(alice)

        await m1155.connect(addr2).setApprovalForAll(sale.address,true);

        await expect(sale.connect(addr2).offerResale(m1155.address,11,5,ethers.utils.parseEther("0.205"))).to.emit(sale,'ResaleOffer');

        await expect( sale.connect(addr3).acceptOffer(m1155.address,11,1,5)).to.emit(m1155, 'TransferSingle')
        expect(await m1155.balanceOf(addr2.address,11)).to.equal(0);
        expect(await m1155.balanceOf(addr3.address,11)).to.equal(5);
        sellerAfter = await weth.balanceOf(addr2.address)
        buyerAfter = await weth.balanceOf(addr3.address)
        patronAfter = await weth.balanceOf(patron)
        creatorAfter = await weth.balanceOf(creator)
        musicianAfter = await weth.balanceOf(musician)
        mintyAfter = await weth.balanceOf(alice)

        console.log("check balances")
        console.log("buyer spend ",ethers.utils.formatEther(buyerBefore.sub(buyerAfter)))
        console.log("seller gains ",ethers.utils.formatEther(sellerAfter.sub(sellerBefore)))
        console.log("patron ",ethers.utils.formatEther(patronAfter.sub(patronBefore)))
        console.log("creator ",ethers.utils.formatEther(creatorAfter.sub(creatorBefore)))
        console.log("musician ",ethers.utils.formatEther(musicianAfter.sub(musicianBefore)))
        console.log("minty ",ethers.utils.formatEther(mintyAfter.sub(mintyBefore)))
        //expect(await weth.balanceOf(addr2.address)).to.equal(ethers.utils.parseEther("99.5"))
    })


    async function bid(signer, amount, residual) {
        
    }
})