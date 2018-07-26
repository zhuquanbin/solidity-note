# Golang Call Smart Contract

## Generate smart contract api

- 安装 go-ethereum
    ```bash
    # golang & solc install  略

    # 安装 ethereum 到 GOPATH 目录下    
    go get github.com/ethereum/go-ethereum
    # 编译 geth 
    go install github.com/ethereum/go-ethereum/cmd/geth
    # 编译 abigen
    go install github.com/ethereum/go-ethereum/cmd/abigen
    ```

- 准备测试智能合约 `BoreyToken.sol`
    ```sol
    pragma solidity ^0.4.23;

    contract BoreyToken {
        event Transfer(address indexed _from, address indexed _to, uint256 _value);

        address  public owner;
        mapping (address => uint)  public balanceOf;

        constructor(uint256 supply) public {
            owner = msg.sender;
            balanceOf[msg.sender] = supply;
        }

        function transfer(address _to, uint256 _value) public returns (bool) {
            require(_to != address(0));
            require(_value <= balanceOf[msg.sender]);

            balanceOf[msg.sender] = balanceOf[msg.sender] - _value;
            balanceOf[_to] = balanceOf[_to] + _value;
            emit Transfer(msg.sender, _to, _value);
            return true;
        }   
    }
    ```

- abigen

    ```bash
    # abigen help
    prod@ubuntu:~/solidity/go$ abigen -h
    Usage of abigen:
    -abi string
            Path to the Ethereum contract ABI json to bind
    -bin string
            Path to the Ethereum contract bytecode (generate deploy method)
    -exc string
            Comma separated types to exclude from binding
    -lang string
            Destination language for the bindings (go, java, objc) (default "go")
    -out string
            Output file for the generated binding (default = stdout)
    -pkg string
            Package name to generate the binding into
    -sol string
            Path to the Ethereum contract Solidity source to build and bind
    -solc string
            Solidity compiler to use if source builds are requested (default "solc")
    -type string
            Struct name for the binding (default = package name)

    # abigen 生成 boreytoken.go
    prod@ubuntu:~/solidity/go$ abigen --sol BoreyToken.sol --pkg token --out boreytoken.go
    prod@ubuntu:~/solidity/go$ 
    prod@ubuntu:~/solidity/go$ ls
    boreytoken.go  BoreyToken.sol

    ```
- 智能合约的API 文件 [boreytoken.go](./golang.sol/demo/token/boreytoken.go)


## Deploy contract with golang code

- 运行 geth 命令
    ```bash
    # 1、geth 初始化： 略， 见geth私链部署

    # 2、启动websock 为后续 event 监听提供准备
    geth --ws --wsaddr 0.0.0.0 --wsport 8641  --wsorigins "*" --rpc --rpcport "8541" --rpcaddr "0.0.0.0" --datadir node1/data --port "30301" --rpccorsdomain "*" --rpcapi "personal,db,eth,net,web3,admin,txpool,miner" --networkid 1024 --nodiscover
    ```

- 部署代码
    ```golang
    package main

    import (
        "context"
        "log"
        "math/big"
        "time"

        "github.com/ethereum/go-ethereum"
        "github.com/ethereum/go-ethereum/accounts/abi/bind"
        "github.com/ethereum/go-ethereum/accounts/keystore"
        "github.com/ethereum/go-ethereum/crypto"
        "github.com/ethereum/go-ethereum/ethclient"
        "golang.sol/demo/token"
    )

    // Keystore & Password， 钱包中有一定余额的ETH用于部署合约的GAS消耗
    const gKeystore = `{"address":"9b4eabea5d69a3c434c40f84f65282f6b4d9b232","crypto":{"cipher":"aes-128-ctr","ciphertext":"0c1a562d3a28682f28a02de89927adbacd99168e9efa48fe3ff0a85df70febac","cipherparams":{"iv":"6cdadf4f3f38af7a4aee1843198a9c00"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"20b525e4dbfac089dd9c5c65fb873a9e530d42f47610647fe31b23f6e348f58e"},"mac":"1651c2597ccedb675f372ad49f0ad30fb5a6c604ee5bda7283e35ed3f9da3ba7"},"id":"6d3f052d-bdf7-411f-bc99-86b7e73fb6a3","version":3}`

    const gPassphrase = "123456"

    type EthToken struct {
        client *ethclient.Client // ethclient client instance
    }

    func NewEthToken(url string) *EthToken {
        ec, err := ethclient.Dial(url)
        if err != nil {
            log.Fatalf("Failed to instantiate a ethereum client: %v", err)
        }
        return &EthToken{client: ec}
    }

    func (t *EthToken) DeployToken(_tops *bind.TransactOpts, supply int64) *token.BoreyToken {
        _, tx, contract, err := token.DeployBoreyToken(_tops, t.client, big.NewInt(supply))
        
        if err != nil {
            log.Fatalf("Failed to deploy new token contract: %v", err)
        }

        log.Printf("Transaction waiting to be mined: 0x%x", tx.Hash())

        // 等待交易确认后记录在区块上， 获取该笔交易的信息
        for {
            r, err := t.client.TransactionReceipt(context.Background(), tx.Hash())
            if err == ethereum.NotFound {
                time.Sleep(1 * time.Second)
                continue
            }

            if err != nil {
                log.Fatalf("Failed to deploy new token contract: %v", err)
            }

            if r != nil {
                log.Printf("Transaction had been mined !")
                log.Printf("Deploy token contract successfully, contract address is : %s", r.ContractAddress.String())
                break
            }
        }

        return contract
    }

    // TransactOpts is the collection of authorization data required to create a valid Ethereum transaction.
    func GetAuth(_keystore string, _passphrase string) *bind.TransactOpts {
        key, err := keystore.DecryptKey([]byte(_keystore), _passphrase)
        if err != nil {
            log.Fatalf("Failed to decrypt key: %v", err)
        }
        // 对keystore采取对称加密解析出私钥
        log.Printf("decrypt keystore private key: %x", crypto.FromECDSA(key.PrivateKey))
        return bind.NewKeyedTransactor(key.PrivateKey)
    }

    func testDeployContract() {

        auth := GetAuth(gKeystore, gPassphrase)
        // 部署合约
        pEt := NewEthToken("ws://192.168.4.136:8641")
        contract := pEt.DeployToken(auth, 100000000)
        // 获取balance
        b, _ := contract.BalanceOf(&bind.CallOpts{}, auth.From)
        log.Printf("Owner: %x, balaceOf: %s", auth.From, b.String())
    }

    func main() {
        testDeployContract()
    }

    ```

- Application log
    ```bash
    c:/go/bin/go.exe run main.go [D:/golang/gopath/src/golang.sol/demo]
    2018/07/26 18:30:50 decrypt keystore private key: 0465160528ae598e77ac573be21e61a706117d0217b410510a61cbf90d66a2a5
    2018/07/26 18:30:50 Transaction waiting to be mined: 0xc7e01febb8610c31cdf1592d2d3fb958cc6ca5cc26769c08c0b37afd748559b8
    2018/07/26 18:31:14 Transaction had been mined !
    2018/07/26 18:31:14 Deploy token contract successfully, contract address is : 0xB48e9b2993977777A0cC06D535FBB8A642029251
    2018/07/26 18:31:14 Owner: 9b4eabea5d69a3c434c40f84f65282f6b4d9b232, balaceOf: 100000000
    成功: 进程退出代码 0.
    ```

- Geth log
    ```
    INFO [07-26|18:30:39] Commit new mining work                   number=3440 txs=0 uncles=0 elapsed=107.261µs
    INFO [07-26|18:30:50] Submitted contract creation              fullhash=0xc7e01febb8610c31cdf1592d2d3fb958cc6ca5cc26769c08c0b37afd748559b8 contract=0xB48e9b2993977777A0cC06D535FBB8A642029251
    INFO [07-26|18:31:03] Successfully sealed new block            number=3440 hash=cf758e…0b1d57
    INFO [07-26|18:31:03] 🔗 block reached canonical chain          number=3435 hash=61f6ef…8fc3ff
    INFO [07-26|18:31:03] 🔨 mined potential block                  number=3440 hash=cf758e…0b1d57
    INFO [07-26|18:31:03] Commit new mining work                   number=3441 txs=1 uncles=0 elapsed=371.219µs
    INFO [07-26|18:31:13] Successfully sealed new block            number=3441 hash=74482b…201edc
    INFO [07-26|18:31:13] 🔗 block reached canonical chain          number=3436 hash=fd39c5…1481a1
    INFO [07-26|18:31:13] 🔨 mined potential block                  number=3441 hash=74482b…201edc
    INFO [07-26|18:31:13] Commit new mining work                   number=3442 txs=0 uncles=0 elapsed=164.888µs

    ```