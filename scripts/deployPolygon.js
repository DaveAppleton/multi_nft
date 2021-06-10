const fs = require("fs");
const {network, ethers} = require("hardhat");

// remember to add ERC20Drain functions

async function main() {
    if (network.name != "mumbai")  {
        console.warn("This needs to be on MUMBAI");
        process.exit(1)
    }
    const [owner] = await ethers.getSigners();
    console.log(network.name);
    console.log("owner",await owner.getAddress());
    await deployStuff();
}

async function deployStuff() {
    const [deployer] = await ethers.getSigners();
    const MINTY   = await ethers.getContractFactory("pMinty")
    const M721    = await ethers.getContractFactory("pMintyUnique")
    const SALE    = await ethers.getContractFactory("pMintysale")
    const LOCKING = await ethers.getContractFactory("contracts/locking/locking.sol:locking")
    const SALE2   = await ethers.getContractFactory("pMintyMultiSale")
    //const M1155   = await ethers.getContractFactory("mintyMultiTokens")
    const DWETH   = await ethers.getContractFactory("dummyWeth")

    let wallet = "0x31EFd75bc0b5fbafc6015Bd50590f4fDab6a3F22"

    let dl = true
    if (dl) {
        let minty = "0x91509AD882E530f0934CA0901e3098fB6e1e3Af1"
        let locking = await LOCKING.deploy(minty, ethers.utils.parseEther("100.0"))
        console.log("locking",locking.address, locking.deployTransaction.hash)
        process.exit(0)
    }

    let now = false
    if (now) {
        if (network.name == mumbai) {
            let wethC = await DWETH.deploy(wallet)
            console.log("WETH", wethC.address, wethC.deployTransaction.hash)
            let weth = wethC.address

            let minty = await MINTY.deploy(wallet)
            console.log("MINTY",minty.address, minty.deployTransaction.hash)
        } else {
            let weth = "HUH"
            let minty = await MINTY.deploy(0)
            console.log("MINTY",minty.address, minty.deployTransaction.hash)
 
        }
        let sale = await SALE.deploy(weth,wallet,900,100,1025)
        console.log("SALE",sale.address, sale.deployTransaction.hash)
    

        let m721 = await M721.deploy(sale.address)
        console.log("MINTY UNIQUE",m721.address, m721.deployTransaction.hash)

        let sale2 = await SALE2.deploy(wallet,weth,1025)
        console.log("MULTI SALE",sale2.address, sale2.deployTransaction.hash)

        process.exit(0)
    }
    let newToken = false
    if (newToken) {
        m1155 = await M1155.deploy(Seri1,[locking],"You need to be a MINTY patron or TEST PROJECT patron")
        await m1155.deployed();
        console.log("M1155",m1155.address,"deployed at",m1155.deployTransaction.hash)
    } else {
        tokenAddress = "0xC1D2c492BD2EEb8451d4d08015a13e1cA224BAf8"
        m1155 = M1155.attach(tokenAddress)
    }
    let newSale = false
    if (newSale) {
        sale2 = await SALE2.deploy(deployer.address, 10, 10)
        await sale2.deployed()
        console.log("sale", sale2.address,"deployed at",sale2.deployTransaction.hash)
    } else {
        multiSale = "0xFfba9038a0d6bB400c341e82349097983d09998d"
        sale2 = SALE2.attach(multiSale)
    }
    console.log("sale at ",sale2.address)
    let deployLoop = true
    if (deployLoop) {
        addrs = [Neuman] // [Neumann,LaylaHifi] //[NFTDaddy,Amir1,Amir2]
        for (j = 0; j < addrs.length; j++) {
            m1155 = await M1155.deploy(addrs[j],multiSale,[locking],"You need to be a MINTY patron or TEST PROJECT patron")
            await m1155.deployed();
            console.log("M1155 for ",addrs[j],":",m1155.address,"deployed at",m1155.deployTransaction.hash)
            tx = await m1155.setAuth(sale2.address,true);
            console.log("setAuth",tx.hash)
            receipt = await tx.wait(); 
        }                      
        process.exit(0)
    }
    tx = await m1155.connect(deployer).setAuth(sale2.address,true);
    console.log("setAuth",tx.hash)
    receipt = await tx.wait();


}   


main()
    .then(()=>process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });