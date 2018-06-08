## ERC20介绍
&emsp; &emsp;[ERC-20 Token Standard](https://github.com/ethereum/EIPs/blob/master/EIPS/eip-20.md)

## 初始化
注意： 在windows下使用 `truffle.cmd`
```
# 1. 创建 erc20-token 目录
prod@ubuntu:~/solidity$ mkdir erc20-token

# 2. 初始化目录
prod@ubuntu:~/solidity/erc20-token$ truffle init
Downloading...
Unpacking...
Setting up...
Unbox successful. Sweet!

Commands:

  Compile:        truffle compile
  Migrate:        truffle migrate
  Test contracts: truffle test
# 3. 查看目录树
prod@ubuntu:~/solidity/erc20-token$ tree
.
├── contracts
│   └── Migrations.sol
├── migrations
│   └── 1_initial_migration.js
├── test
├── truffle-config.js
└── truffle.js

3 directories, 4 files
```

## 智能合约创建
```
# 4. 创建 ABCToken contract
prod@ubuntu:~/solidity/erc20-token$ truffle create contract ABCToken
# 5. 创建 ABCToken migration
prod@ubuntu:~/solidity/erc20-token$ truffle create migration AbcToken
# 6. 进行重命名migration， 按照顺序 1，2, 3 ... n
prod@ubuntu:~/solidity/erc20-token$ mv migrations/1528368256_abc_token.js migrations/2_abc_token.js 
# 7. 创建测试用例
prod@ubuntu:~/solidity/erc20-token$ truffle create test AbcToken

# 查看创建的文件
prod@ubuntu:~/solidity/erc20-token$ tree
.
├── contracts
│   ├── ABCToken.sol
│   └── Migrations.sol
├── migrations
│   ├── 1_initial_migration.js
│   └── 2_abc_token.js
├── test
│   └── abc_token.js
├── truffle-config.js
└── truffle.js

```

## 智能合约编写
[ABCToken.sol](./contracts/ABCToken.sol)

## 编译

```
# 8. 编译
prod@ubuntu:~/solidity/erc20-token$ truffle compile
Compiling ./contracts/ABCToken.sol...
Compiling ./contracts/Migrations.sol...
Compiling ./contracts/utils/ERC20Token.sol...
Compiling ./contracts/utils/SafeMath.sol...
Writing artifacts to ./build/contracts

```

## 本地开发测试环境部署
- 编辑migrations文件
```js
var SafeMath = artifacts.require("./utils/SafeMath.sol");
var ABCToken = artifacts.require("./ABCToken.sol");

module.exports = function(deployer) {
  deployer.deploy(SafeMath);
  deployer.link(SafeMath, ABCToken);
  deployer.deploy(ABCToken, "AbcCoin", "ABC", 10000, 8);
};
```

- 进行开发环境

```bash
# 进入develop环境
prod@ubuntu:~/solidity/erc20-token$ truffle develop
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

# 部署合约， 采用migrate  -f 指定migration文件
truffle(develop)> migrate -f 2
Using network 'develop'.

Running migration: 2_abc_token.js
  Deploying SafeMath...
  ... 0x751687331e5f86a1afe097d7038815b1406e763c58c2d4c89d028dd580e5f603
  SafeMath: 0x8cdaf0cd259887258bc13a92c0a6da92698644c0
  Deploying ABCToken...
  ... 0x7e3c6add244aa2786d554fb1f16ddd197398b4c62bf088712e43709fe65547c7
  ABCToken: 0xf12b5dd4ead5f743c6baa640b0216200e89b60da
Saving artifacts...

# 查看合约地址
truffle(develop)> ABCToken.deployed().then(instance => instance.address);
'0xf12b5dd4ead5f743c6baa640b0216200e89b60da'
# 查看合约name
truffle(develop)> ABCToken.deployed().then(instance => instance.name());
'AbcCoin'
# 查看合约symbol
truffle(develop)> ABCToken.deployed().then(instance => instance.symbol());
'ABC'
# 查看发行量
truffle(develop)> ABCToken.deployed().then(instance => instance.totalSupply());
BigNumber { s: 1, e: 12, c: [ 1000000000000 ] }

```

## 测试用例
编写test

```javascript
const ABCToken = artifacts.require("ABCToken");

contract('ABCToken', function(accounts) {
    let owner   = accounts[0];
    let wallet1 = accounts[1];
    let wallet2 = accounts[2];
    let TestInstance = null;
    
    beforeEach('setup contract for each test', async() => {
        TestInstance = await ABCToken.new("AbcCoin","ABC", 10000, 8);
    })

    it('1) 检测初始化参数', async() => {
        assert.equal(await TestInstance.name(),"AbcCoin");
        assert.equal(await TestInstance.symbol(), "ABC");
        assert.equal(await TestInstance.totalSupply(), 10 ** 12);
        assert.equal(await TestInstance.owner(), owner);
    })

    it('2) 钱包转账测试', async() => {
        let balance1 = await TestInstance.balanceOf.call(owner)
        assert.equal(balance1.toNumber(), 10 ** 12);

        await TestInstance.transfer(wallet1, 1.2 * 10**8);
        
        let balance2 = await TestInstance.balanceOf(owner);
        assert.equal(balance2.toNumber(), 10 ** 12 - 1.2 * 10**8);

        let balance3 = await TestInstance.balanceOf(wallet1);
        assert.equal(balance3.toNumber(), 1.2 * 10**8);
    })

    it('3) 代币铸造权限检测', async() => {

        try{
            await TestInstance.mint(wallet1, 0.3 * 10**8, {from: wallet2});
        }catch(error){
            const revertFound = error.message.search('revert') >= 0;
            assert(revertFound, `Expected "revert", got ${error} instead`);
        }
        let balance1 = await TestInstance.balanceOf(wallet1);
        assert.equal(balance1.toNumber(), 0);
    })

    it('4) 代币铸造测试', async() => {

        try{
            await TestInstance.mint(wallet1, 0.3 * 10**8);
        }catch(error){
            const revertFound = error.message.search('revert') >= 0;
            assert(revertFound, `Expected "revert", got ${error} instead`);
        }
        let balance1 = await TestInstance.balanceOf(wallet1);
        assert.equal(balance1.toNumber(), 3 * 10**7);

        assert.equal(await TestInstance.totalSupply(), 10 ** 12 + 0.3 * 10 ** 8);
    })

    it('5) 代币销毁权限检测', async() => {
        await TestInstance.transfer(wallet1, 5 * 10**8);
        try{
            await TestInstance.burn(wallet1, 2 * 10**8, {from: wallet2});
        }catch(error){
            const revertFound = error.message.search('revert') >= 0;
            assert(revertFound, `Expected "revert", got ${error} instead`);
        }
        let balance1 = await TestInstance.balanceOf(wallet1);
        assert.equal(balance1.toNumber(), 5 * 10 ** 8);
    })

    it('6) 代币销毁测试', async() => {
        await TestInstance.transfer(wallet1, 5 * 10**8);
        try{
            await TestInstance.burn(wallet1, 2 * 10**8);
        }catch(error){
            const revertFound = error.message.search('revert') >= 0;
            assert(revertFound, `Expected "revert", got ${error} instead`);
        }
        let balance1 = await TestInstance.balanceOf(wallet1);
        assert.equal(balance1.toNumber(), 3 * 10**8);

        assert.equal(await TestInstance.totalSupply(), 10 ** 12 - 2 * 10 ** 8);
    })
});

```


```
# 9. 运行测试用例
prod@ubuntu:~/solidity/erc20-token$ truffle test
Using network 'test'.



  Contract: ABCToken
    ✓ 1) 检测初始化参数 (75ms)
    ✓ 2) 钱包转账测试 (68ms)
    ✓ 3) 代币铸造权限检测
    ✓ 4) 代币铸造测试 (49ms)
    ✓ 5) 代币销毁权限检测 (54ms)
    ✓ 6) 代币销毁测试 (72ms)


  6 passing (689ms)

prod@ubuntu:~/solidity/erc20-token$ 

```

## 部署生成环境
### 节点环境
- 运行geth
```bash
geth --identity "chain" --ethash.dagsinmem 0 --rpc --rpcport "8541" --rpcaddr "0.0.0.0" --datadir node1/data --port "30301" --rpccorsdomain "*" --rpcapi "personal,db,eth,net,web3,admin,txpool,miner" --networkid 1024 --nodiscover
```

- geth 进入 console

```bash
prod@ubuntu:~/ethereum/nodes$ geth attach node1/data/geth.ipc 
Welcome to the Geth JavaScript console!

instance: Geth/chain/v1.8.9-stable-ff9b1461/linux-amd64/go1.10.1
coinbase: 0x9b4eabea5d69a3c434c40f84f65282f6b4d9b232
at block: 369 (Wed, 30 May 2018 16:39:14 CST)
 datadir: /home/prod/ethereum/nodes/node1/data
 modules: admin:1.0 debug:1.0 eth:1.0 miner:1.0 net:1.0 personal:1.0 rpc:1.0 txpool:1.0 web3:1.0

# 开始挖矿记账
> miner.start()

# 获取部署合约的 coinbase 钱包地址
> eth.coinbase
"0x9b4eabea5d69a3c434c40f84f65282f6b4d9b232"

```


### 项目配置
- 编辑truffle.js
```js
module.exports = {
    // See <http://truffleframework.com/docs/advanced/configuration>
    // to customize your Truffle configuration!
    solc: {
      optimizer: {
        enabled: true,
      }
    },
    networks: {
      prodwork: {
        host: "127.0.0.1",
        port: 8541,
        network_id: "*",
        address: "0x9b4eabea5d69a3c434c40f84f65282f6b4d9b232",
        gasPrice: 18000000000, // eth.gasPrice
        gas: 0x47b760, // 取 genesis.json 中的 GasLimit ，保证此次transaction gas 
      }
    }
};
```

### 部署
- 在geth console 端 unlock coinbase 钱包地址
```
# 解锁 1000 s
> personal.unlockAccount(eth.coinbase, null, 1000)
Unlock account 0x9b4eabea5d69a3c434c40f84f65282f6b4d9b232
Passphrase: 
true
> 
```
- 切换到项目根目录
```
# migrate 到 prodwork 网络
prod@ubuntu:~/solidity/erc20-token$ truffle migrate -f 2 --network prodwork
Using network 'prodwork'.

Running migration: 2_abc_token.js
  Deploying SafeMath...
  ... 0xdae2d4c89c2e9524a9e4f7326fb12acd7d3888e1b32894a844a39959cfdca5c8
  SafeMath: 0x5e8309efcaff694bd1f43a8651a2a34378eae48a
  Deploying ABCToken...
  ... 0xe3960f2ec083b0cb5d0987af4b28477894352577d166a39bacaca879a1976e57
  ABCToken: 0x580a0c67940008a8f6773a3a9c991eec9712b6d2
Saving artifacts...

# 查看部署地址
prod@ubuntu:~/solidity/erc20-token$ truffle console  --network prodwork
truffle(prodwork)> ABCToken.deployed().then(instance => instance.address);
'0x580a0c67940008a8f6773a3a9c991eec9712b6d2'
truffle(prodwork)> 

```

- python web3测试
pip3 install eth_utils web3 ipython

```python
prod@ubuntu:~/solidity/erc20-token$ ipython3 
Python 3.5.2 (default, Nov 23 2017, 16:37:01) 
Type 'copyright', 'credits' or 'license' for more information
IPython 6.4.0 -- An enhanced Interactive Python. Type '?' for help.

In [1]: from web3 import Web3, HTTPProvider

In [2]: from eth_utils import to_checksum_address

In [3]: import json

In [4]: address = to_checksum_address("0x580a0c67940008a8f6773a3a9c991eec9712b6d2")

In [5]: abi = json.loads(open("./build/contracts/ABCToken.json").read())["abi"]

In [6]: w3 = Web3(HTTPProvider("http://127.0.0.1:8541"))

In [7]: contract = w3.eth.contract(abi=abi, address=address)

In [8]: contract.functions.totalSupply().call()
Out[8]: 1000000000000

In [9]: contract.functions.name().call()
Out[9]: 'AbcCoin'

In [10]: contract.functions.symbol().call()
Out[10]: 'ABC'

```
- 更多交易代码<BR>
    通过Keystore进行发起交易参考 [web3py_contract_transaction](https://github.com/zhuquanbin/ethereum-bip44/blob/master/eth_bip44/__init__.py#L175)
