# List
* [获取部署的智能合约地址、abi和code](#获取部署的智能合约地址、abi和code)
* [测试用例智能合约调用另外一个智能合约](#测试用例智能合约调用另外一个智能合约)

# Summary
## 获取部署的智能合约地址、abi和code

```js
var SimpleConstract = artifacts.require("./SimpleConstract.sol");

module.exports = function(deployer) {
  deployer.deploy(SimpleConstract)
    // Console log the address:
    .then(() => console.log(SimpleConstract.address))

    // Retrieve the contract instance and get the address from that:
    .then(() => SimpleConstract.deployed())
    .then(_instance => console.log(_instance.address));
};
```

## 测试用例智能合约调用另外一个智能合约
