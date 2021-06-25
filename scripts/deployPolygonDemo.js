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
    const M1155   = await ethers.getContractFactory("contracts/flat/pMintyMultiToken.sol:pMintyMultiToken")
    //const M1155   = await ethers.getContractFactory("mintyMultiTokens")
    const weth = "0xf738b83Fa52A7Ab570918Afe61b78b8E2DC6F4EF"
    const wallet = "0x31EFd75bc0b5fbafc6015Bd50590f4fDab6a3F22"
    const sale = '0x57b5eD0DcA25fcf5A939807dD06D0e6216b1FB85' // "0x383a9962A1f55eF866297A7606427639a34B16B7"
    const mintyUnique = "0xb78cBDe4c5036E597555b5dc0885EB2827b1C5Ae" // "0x13891510d68D7a9358D090E4393B10724C649c3A"

    let s1 = false
    if (s1) {
        let minty = await MINTY.deploy()
        console.log("MINTY",minty.address, minty.deployTransaction.hash)
        process.exit(0)
    } 

    let samourai = true  // uses MUD acount
    if (samourai) {
        let theOwner = "0x0EA0281db722F5642457B5E785dda24eA75f6149"
        hash = "Qmd3w3c8X69yJuAjGyiXUi7XcFQLLycWsFjPWToThjUADj"
        const samToken = await M1155.attach("0x668Ee0f5bfA0A0497F8F8315212888d5a64779Eb")
        console.log("before",await samToken.balanceOf(theOwner,62))
        // tx = await samToken.mint(62,40,hash,0)
        // console.log(tx.hash)

        multisale = await SALE2.attach("0x6F2a567826554D2526703a6BD265C42E53FdF4C2")

        oldItem = await multisale.items(samToken.address,62,0) 

        console.log(ethers.utils.formatEther(oldItem.unitPrice))

       

        tx = await multisale.retractRemainingOffer(samToken.address,62,0)
        console.log("retract ",tx.hash)
        tx = await multisale.offerResale(samToken.address,62,45,oldItem.unitPrice)
        console.log("re offer ",tx.hash)
        
        console.log("after",await samToken.balanceOf(theOwner,62))
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
        process.exit(0)
    }

    let s4 = false
    if (s4) {
        saleContract = await SALE.attach(sale)
        tx = await saleContract.setMintyUnique(mintyunique)
        console.log(tx.hash)
        process.exit(0)
    }

    let smu = false
    if (smu) {
        saleC = await SALE.attach(sale)
        tx = await saleC.setMintyUnique(mintyUnique)
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