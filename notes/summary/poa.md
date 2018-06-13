## POA (proof-of-authority)

允许授权节点(signers)对区块进行签名挖矿, 通过控制节点难度系数进行延迟出块， 来确保难度系数最低的优先出块并进行广播。

- 难度系数计算通过满足 (block_number % signers) == (index of signers) 条件的signer标记为diffInTurn, 其余的signer标记为diffNoTurn；以此来设置延迟时长； 
    - [calc Difficulty 源码](https://github.com/ethereum/go-ethereum/blob/master/consensus/clique/clique.go#L669)
    - [diffInTurn 源码](https://github.com/ethereum/go-ethereum/blob/master/consensus/clique/snapshot.go#L304)


- 共识机制保证同一个signer只能签名 连续 (SIGNER_COUNT / 2) + 1) blocks 中的一个block;
    - [保证算法 源码](https://github.com/ethereum/go-ethereum/blob/master/consensus/clique/clique.go#L619)


简言之： 对signers进行轮询出块；

### 特点
- PoA是依靠预设好的授权节点(signers)，负责产生block；
- 可以由已授权的signer选举(投票超过50%)加入新的signer；
- 即使存在恶意signer, 他最多只能攻击连续块(数量是 (SIGNER_COUNT / 2) + 1) 中的1个, 期间可以由其他signer投票踢出该恶意signer；
- 可指定产生block的时间；


### POA 工作流程
1. 在创世块中指定一组初始授权的signers, 所有地址 保存在创世块Extra字段中；

2. 启动挖矿后, 该组signers开始对生成的block 进行签名并广播；

3. 签名结果 保存在区块头的Extra字段中；

4. Extra中更新当前高度已授权的所有signers的地址 ,因为有新加入或踢出的signer；

5. 每一高度都有一个signer处于IN-TURN状态, 其他signer处于OUT-OF-TURN状态, IN-TURN的signer签名的block会 立即广播 , OUT-OF-TURN的signer签名的block会 延时 一点随机时间后再广播, 保证IN-TURN的签名block有更高的优先级上链；

6. 如果需要加入一个新的signer, signer通过API接口发起一个proposal, 该proposal通过复用区块头 Coinbase(新signer地址)和Nonce("0xffffffffffffffff") 字段广播给其他节点. 所有已授权的signers对该新的signer进行"加入"投票, 如果赞成票超过signers总数的50%, 表示同意加入；

7. 如果需要踢出一个旧的signer, 所有已授权的signers对该旧的signer进行"踢出"投票, 如果赞成票超过signers总数的50%, 表示同意踢出；


## Resources
- [Clique PoA protocol & Rinkeby PoA testnet #EIP-225](https://github.com/ethereum/EIPs/issues/225)  (中译 [链接](https://github.com/ZtesoftCS/go-ethereum-code-analysis/blob/master/%E4%BB%A5%E5%A4%AA%E5%9D%8A%E6%B5%8B%E8%AF%95%E7%BD%91%E7%BB%9CClique_PoA%E4%BB%8B%E7%BB%8D.md))

- [简书： 以太坊PoA共识引擎算法介绍](https://www.jianshu.com/p/9025a523ab0f)

- [Medium： 使用 go-ethereum 1.6 Clique PoA consensus 建立 Private chain](https://medium.com/taipei-ethereum-meetup/%E4%BD%BF%E7%94%A8-go-ethereum-1-6-clique-poa-consensus-%E5%BB%BA%E7%AB%8B-private-chain-1-4d359f28feff)