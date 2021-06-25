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
    hash = "0xef838204bc223bb6acf0129fffde1eca36a908d55b766d3f66c66e186bf7f5f3"
    provider = await network.provider
    tx = await provider.getTransaction(hash)
    console.log(tx)
}   


main()
    .then(()=>process.exit(0))
    .catch((error) => {
        console.error(error);
        process.exit(1);
    });