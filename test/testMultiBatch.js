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

        M1155 = await ethers.getContractFactory("mintyMultiToken")
 
        m1155 = await M1155.deploy(addr1.address,[],"")
        console.log("ERC1155 ",m1155.address)
        console.log

        SALE = await ethers.getContractFactory("mintyMultiSale")
        sale = await SALE.deploy(alice, 10, 10)

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

    it("offer an NFT for sale", async function() { 
        await m1155.connect(addr1).setAuth(sale.address,true);
        await sale.connect(addr1).offerNew(m1155.address,10,"1212",200,ethers.utils.parseEther("0.1"))
        expect(await m1155.balanceOf(addr1.address,10)).to.equal(200);
    })

    it("offer a range of nfts for sale", async function(){
        await m1155.connect(addr1).setBase(">")
        let tokenIds = [20,21,22,23]
        let hashes = ["2 1","2 2", "2 3","2 4"]
        let quantities = [201,202,203,204]
        let prices = [ethers.utils.parseEther("0.1"),ethers.utils.parseEther("0.2"),ethers.utils.parseEther("0.3"),ethers.utils.parseEther("0.4")]
        await sale.connect(addr1).offerNewBatch(m1155.address,tokenIds,hashes,quantities,prices)
        for (j = 0; j < tokenIds.length; j++){
            expect(await m1155.balanceOf(addr1.address,tokenIds[j])).to.equal(quantities[j])
            expect(await m1155.uri(tokenIds[j])).to.equal(">"+hashes[j])
            expect(await sale.numberOfOffers(m1155.address,tokenIds[j])).to.equal(1)
        }
    })

    it("offer a shit load of nfts for sale", async function(){
        let priceBase = [ethers.utils.parseEther("0.1"),ethers.utils.parseEther("0.2"),ethers.utils.parseEther("0.3"),ethers.utils.parseEther("0.4")]
        let tokenIds = []
        let hashes = []
        let quantities = []
        let prices = []
        for (j = 100; j < 150; j++) {
            tokenIds.push(j)
            hashes.push((j-100)/10 + "" + j%10)
            quantities.push(j % 5 + 1)
            prices.push(priceBase[j % 4])
        }
        await sale.connect(addr1).offerNewBatch(m1155.address,tokenIds,hashes,quantities,prices)
        for (j = 0; j < tokenIds.length; j++){
            expect(await m1155.balanceOf(addr1.address,tokenIds[j])).to.equal(quantities[j])
            expect(await sale.numberOfOffers(m1155.address,tokenIds[j])).to.equal(1)
            expect(await sale.price(m1155.address,tokenIds[j],0)).to.equal(priceBase[j % 4])
            expect(await sale.available(m1155.address,tokenIds[j],0)).to.equal(quantities[j])
        }
    })

    it("buy one of dem mass mintys",async function(){
        expect(await sale.price(m1155.address,124,0)).to.equal(ethers.utils.parseEther("0.1"))
        let ovx = { value: ethers.utils.parseEther("0.3")}
        await expect(sale.connect(addr2).acceptOffer(m1155.address,124,0,3,ovx)).to.emit(m1155, 'TransferSingle')
    })
   
 
    it("No 2 buys five", async function() {
        let override = {
            value: ethers.utils.parseEther("0.5")
        };
        await expect( sale.connect(addr2).acceptOffer(m1155.address,10,0,5,override)).to.emit(m1155, 'TransferSingle')
        expect(await m1155.balanceOf(addr1.address,10)).to.equal(195);
        expect(await m1155.balanceOf(addr2.address,10)).to.equal(5);
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
        expect(await m1155.balanceOf(addr2.address,10)).to.equal(4);
        expect(await m1155.balanceOf(addr3.address,10)).to.equal(1);
        expect(await sale.available(m1155.address,10,1)).to.equal(2);
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