# 1 Solidity

## 1.1 Knowledge

- [基础学习](./notes/summary/knowledge.md#basics)
- [进阶学习](https://github.com/androlo/solidity-workshop)
<!-- https://github.com/androlo/solidity-workshop/blob/master/tutorials/2016-03-11-advanced-solidity-III.md#calldata -->

## 1.2 Personal Summary
- [solidity函数和变量](./notes/summary/knowledge.md#solidity-functions-and-variables)
- [修饰符 public 和 external 的区别](./notes/summary/knowledge.md#different-between-public-and-external)
- [修饰符 prue 和 view 的区别](./notes/summary/knowledge.md#different-between-prue-and-view)
- [Different between require and assert](./notes/summary/knowledge.md#different-between-require-and-assert)
- [send ether from contract to another contract](./notes/summary/knowledge.md#send-ether-from-contract-to-another-contract)
- [modifer function with parameters](./notes/summary/knowledge.md#modifer-function-with-parameters)


# 2 开发环境及框架
* [以太坊私链POW或POA部署](./notes/summary/geth.md)

## 2.1 Reminx
网页版开发工具
- 在线地址：http://remix.ethereum.org
- Github： https://github.com/ethereum/remixd

优点：简单直接<BR>
缺点：工程化程度不足， 如：不能编写测试用例

### 2.1.1 summary
- [环境安装](./nodes/remix/install.md)
- [使用说明](./nodes/remix/guide.md)
- [Reminx + MetaMask 部署智能合约](./nodes/remix/deploy_example.md)

## 2.2 Truffle + VScode
本地项目开发框架 https://github.com/trufflesuite/truffle

优点：功能丰富， 支持Reminx的所有功能；<BR>
缺点：命令行操作， 略复杂；

### 2.2.1 Sumary
- [环境安装](./notes/truffle/install.md)

- [获取部署的智能合约地址、abi和code](./notes/truffle/summary.md#获取部署的智能合约地址、abi和code)

- [测试用例智能合约调用另外一个智能合约](./notes/truffle/summary.md#测试用例智能合约调用另外一个智能合约)

- [使用truffle框架管理编写一个ERC20 token、并编写测试用例和部署](./notes/truffle/erc20-token/)


# 3 智能合约调用
## 3.1 Golang Api
- [智能合约API生成](./notes/golang/api.md)
- [部署]() 
- [交易]()
- [Event监听]()

## 3.2 Python Api
- [web3.py api](https://github.com/ethereum/web3.py)
- [部署]()
- [交易]()


## 3.3 Nodejs Api
- [web3.js api](https://github.com/ethereum/web3.js)

# 4 智能合约安全

## 4.1 漏洞
- [可重入攻击](./notes/attacks/example01/)
- [短地址攻击](./notes/attacks/1.md)
- [整数溢出](./notes/attacks/2.md)

## 4.2 安全工具
- [Formal Verification of Ethereum Smart Contracts](https://securify.ch/)
- [Security Tools](https://consensys.github.io/smart-contract-best-practices/security_tools/)

## 4.3 推荐实践
- [openzeppelin-solidity](https://github.com/OpenZeppelin/openzeppelin-solidity)
- [Ethereum Smart Contract Security Best Practices](https://consensys.github.io/smart-contract-best-practices/)


# 5 Resources
- [Soldity Document](https://solidity.readthedocs.io/en/v0.4.24/)
- [Awesome-solidity](https://github.com/bkrem/awesome-solidity)
