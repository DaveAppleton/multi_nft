const fs = require("fs");
const {network, ethers} = require("hardhat");

async function main() {
    if ((network.name != "kovan") && (network.name != "mainnet")) {
        console.warn("This needs to be on Kovan or Mainnet");
        process.exit(1)
    }
    const [owner] = await ethers.getSigners();
    console.log(network.name);
    console.log("owner",await owner.getAddress());
    await deployStuff();
}

async function deployStuff() {
    const [deployer] = await ethers.getSigners();
    //const LOCKING = await ethers.getContractFactory("locking")
    const SALE2   = await ethers.getContractFactory("mintyMultiSale")
    const M1155   = await ethers.getContractFactory("mintyMultiTokens")

    const locking = "0x259a2E0BfC2Da23Bd4Cd56c6b1B1569b6C571adb"

    let wallet = "0x31EFd75bc0b5fbafc6015Bd50590f4fDab6a3F22"
    let weth = "0xA6FA4fB5f76172d178d61B04b0ecd319C5d1C0aa"
    let Virgo = "0x39d07f321cAF5b0668459DB5Bcf039A462A9273d"
    let NFTDaddy = "0xa1EFF0887B93e18bB0f59334a0CB57148BC2086f"
    let Amir1 = "0xFE029208b267daBF5077bB3E3E7B1cc9916e9943"
    let Amir2 = "0xc355b9A5dE199F811542D8C5D3E309789FfBEf86"
    let Seri1 = "0xa454515041892eB78132293ABd5763a730412F65"
    let Neumann = '0x31a09F5A107A5214A62ec15aB3C97Dd1E4D3382b'
    let Neuman = '0x6e93Deb7FDa0E5A4CA8A3566785F0c7f4D0A2cb3'
    let LaylaHifi = '0xcad809e60e7B93fF7343E25F891A4959F0a2BBC7'

    let retroAuth = false
    if (retroAuth) {
        console.log("retroAuth")
        addresses = [
            "0xC1D2c492BD2EEb8451d4d08015a13e1cA224BAf8", // (Owner)
            "0x3393cECf6795b01fe40dd8d60592d89aAe4e8b1c", // (Virgo)
            "0xe9aCbfa0d526fbA72490e242D700B388459d918C", // (NFTDaddy)
            "0x370EF21A012Baa0C9017d93B4C958E6693A6e33C", // (Amir 1)
            "0xEEfF94Ac50d3526c559A4CaCAf4b6cC42852D586", // (Amir 2)
            "0x12A26F74c544363800334fE86902e9667184C56a" // (Seri 1)
        ]
        for (j = 0; j < addresses.length; j++) {
            m1155 = await M1155.attach(addresses[j])
        }
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
        os.exit(0)
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