const fs = require("fs");
const {network, ethers} = require("hardhat");

async function main() {
    if (network.name != "goerli") {
        console.warn("This needs to be on Goerli");
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
    const MINTY   = await ethers.getContractFactory("Minty")

 
    let wallet = "0x31EFd75bc0b5fbafc6015Bd50590f4fDab6a3F22"
 
    minty = await MINTY.deploy(wallet)

    console.log("Minty on Goerli",minty.address,"deployed at",minty.deployTransaction.hash)
   

}   


main()
    .then(()=>process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });