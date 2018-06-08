var Reentrancy = artifacts.require("Reentrancy");
var ReentrancyAttack = artifacts.require("ReentrancyAttack");

module.exports = function(deployer) {

    deployer.deploy(Reentrancy).then(() => 
        deployer.deploy(ReentrancyAttack, Reentrancy.address, {from: "0xc5fdf4076b8f3a5357c5e395ab970b5b54098fef"}) 
    );
};
