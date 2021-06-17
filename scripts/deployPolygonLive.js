const fs = require("fs");
const {network, ethers} = require("hardhat");

// remember to add ERC20Drain functions

async function main() {
    if (network.name != "matic")  {
        console.warn("This needs to be on MATIC");
        process.exit(1)
    }
    const [owner] = await ethers.getSigners();
    console.log(network.name);
    console.log("owner",await owner.getAddress());
    // deployer address : 0xA45cb6B905F14b4B38bf76d79445304b5C2F355f
    await deployStuff();
}

async function deployStuff() {
    const [deployer] = await ethers.getSigners();
    const MINTY   = await ethers.getContractFactory("contracts/flat/pminty.sol:pMinty")
    const M721    = await ethers.getContractFactory("contracts/flat/pMintyUnique.sol:pMintyUnique")
    const SALE    = await ethers.getContractFactory("contracts/flat/psale.sol:pMintysale")
    const LOCKING = await ethers.getContractFactory("contracts/locking/locking.sol:locking")
    const SALE2   = await ethers.getContractFactory("contracts/flat/pMintyMultiSale.sol:pMintyMultiSale")
    //const M1155   = await ethers.getContractFactory("mintyMultiTokens")
    
    let weth = "0x7ceB23fD6bC0adD59E62ac25578270cFf1b9f619"     // Main Net 
    let wallet = "0x31EFd75bc0b5fbafc6015Bd50590f4fDab6a3F22"
    let minty = "0x474Ba20088174612427cf8440ac5712e98652AD2"
    let locking = "0x756fe78b65A400F07b6fcA2F92E0482f6DcBF25B"
    let sale = "0x37a8De5647D5a71e2508991A033D3E46765d92af"
    let mintyunique = "0x1332690FC3AB1C5434E5C252C5BA71Ac7A72F964"

    let s1 = false
    if (s1) {
        let minty = await MINTY.deploy()
        console.log("MINTY",minty.address, minty.deployTransaction.hash)
        process.exit(0)
    } 

    let dl = false
    if (dl) {
        let locking = await LOCKING.deploy(minty, ethers.utils.parseEther("100.0"))
        console.log("locking",locking.address, locking.deployTransaction.hash)
        process.exit(0)
    }

    let s0 = false
    if (s0) {
        let sale = await SALE.deploy(weth,wallet,900,100,1025)
        console.log("SALE",sale.address, sale.deployTransaction.hash)
        process.exit(0)
    }


    let s2 = false
    if (s2) {
        let sale2 = await SALE2.deploy(wallet,weth,1025)
        console.log("MULTI SALE",sale2.address, sale2.deployTransaction.hash)
        process.exit(0)
    }

    let s3 = false
    if (s3) {
        let m721 = await M721.deploy(sale)
        console.log("MINTY UNIQUE",m721.address, m721.deployTransaction.hash)
    }

    let s4 = true
    if (s4) {
        saleContract = await SALE.attach(sale)
        tx = await saleContract.setMintyUnique(mintyunique)
        console.log(tx.hash)
        process.exit(0)
    }

    let smu = false
    if (smu) {
        sale = await SALE.attach("0xE01017da9bd2DD291A05E7b3d24A1655dc9ca892")
        tx = await sale.setMintyUnique("0xd9dB3cCEe640EEEC9321Ddc7b4D8404665f71503")
        console.log(tx.hash)
        process.exit(0)
    }

    let minters = [ ] // "0x39d07f321cAF5b0668459DB5Bcf039A462A9273d" ] //"0xa454515041892eB78132293ABd5763a730412F65"]
    if (minters.length > 0) {
        sale = await SALE.attach("0xE01017da9bd2DD291A05E7b3d24A1655dc9ca892")
        for (j = 0; j < minters.length; j++){
            tx = await sale.setMinter(minters[j],true)
            console.log(tx.hash)
        }
        process.exit(0)
    }

   



    let now = false
    if (now) {
        if (network.name == mumbai) {
            let wethC = await DWETH.deploy(wallet)
            console.log("WETH", wethC.address, wethC.deployTransaction.hash)
            let weth = wethC.address

        } else {
            let weth = "HUH"
            let minty = await MINTY.deploy(0)
            console.log("MINTY",minty.address, minty.deployTransaction.hash)
 
        }
    

        let m721 = await M721.deploy(sale.address)
        console.log("MINTY UNIQUE",m721.address, m721.deployTransaction.hash)

        let sale2 = await SALE2.deploy(wallet,weth,1025)
        console.log("MULTI SALE",sale2.address, sale2.deployTransaction.hash)

        process.exit(0)
    }
    let m1155
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
    let deployLoop = false
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