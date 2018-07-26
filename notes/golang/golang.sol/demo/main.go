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
	//	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.sol/demo/token"
)

const gKeystore = `{"address":"9b4eabea5d69a3c434c40f84f65282f6b4d9b232","crypto":{"cipher":"aes-128-ctr","ciphertext":"0c1a562d3a28682f28a02de89927adbacd99168e9efa48fe3ff0a85df70febac","cipherparams":{"iv":"6cdadf4f3f38af7a4aee1843198a9c00"},"kdf":"scrypt","kdfparams":{"dklen":32,"n":262144,"p":1,"r":8,"salt":"20b525e4dbfac089dd9c5c65fb873a9e530d42f47610647fe31b23f6e348f58e"},"mac":"1651c2597ccedb675f372ad49f0ad30fb5a6c604ee5bda7283e35ed3f9da3ba7"},"id":"6d3f052d-bdf7-411f-bc99-86b7e73fb6a3","version":3}`
const gKdfKeystore = `{"id": "99bc6e3f-494d-435e-acdf-8de9ca575fea", "crypto": {"cipher": "aes-128-ctr", "kdf": "pbkdf2", "kdfparams": {"prf": "hmac-sha256", "dklen": 32, "salt": "eec7d57b21d816db9cc77928b28bb9bb", "c": 1000000}, "cipherparams": {"iv": "b9235c9c5a1ded08c5a950f72b30e8e1"}, "ciphertext": "ceb1e49cbbd7322c513dc5e541899a880724f5997191b5e8b53eeee8c11c1c16", "mac": "1f6ca2bcc386ac1d4befd64ab302310fda013bbaa967582c9861f4a1ea6abb64"}, "version": 3, "address": "9b4eabea5d69a3c434c40f84f65282f6b4d9b232"}`
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
	log.Printf("decrypt keystore private key: %x", crypto.FromECDSA(key.PrivateKey))
	return bind.NewKeyedTransactor(key.PrivateKey)
}

func testDeployContract() {
	auth := GetAuth(gKeystore, gPassphrase)
	pEt := NewEthToken("ws://192.168.4.136:8641")
	contract := pEt.DeployToken(auth, 100000000)
	b, _ := contract.BalanceOf(&bind.CallOpts{}, auth.From)
	log.Printf("Owner: %x, balaceOf: %s", auth.From, b.String())
}

func main() {
	testDeployContract()
}
