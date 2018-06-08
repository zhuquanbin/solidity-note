var SafeMath = artifacts.require("./utils/SafeMath.sol");
var ABCToken = artifacts.require("./ABCToken.sol");

module.exports = function(deployer) {
  deployer.deploy(SafeMath);
  deployer.link(SafeMath, ABCToken);
  deployer.deploy(ABCToken, "AbcCoin", "ABC", 10000, 8);
};
