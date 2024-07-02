package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	infuraURL := flag.String("infura", "", "Infura URL")
	privateKey := flag.String("private", "", "Private key")
	publicKey := flag.String("public", "", "Public key")
	abiPath := flag.String("abi", "", "Path to the ABI file")
	binPath := flag.String("bin", "", "Path to the BIN file")
	flag.Parse()

	if *infuraURL == "" || *privateKey == "" || *publicKey == "" || *abiPath == "" || *binPath == "" {
		log.Fatalf("All flags (infura, private, public, abi, bin) are required")
	}

	client, err := ethclient.Dial(*infuraURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	privKey, err := crypto.HexToECDSA(*privateKey)
	if err != nil {
		log.Fatalf("Invalid private key: %v", err)
	}

	publicAddress := common.HexToAddress(*publicKey)

	nonce, err := client.PendingNonceAt(context.Background(), publicAddress)
	if err != nil {
		log.Fatalf("Failed to get nonce: %v", err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Failed to get gas price: %v", err)
	}

	contractABI, err := os.ReadFile(*abiPath)
	if err != nil {
		log.Fatalf("Failed to read contract ABI: %v", err)
	}

	contractBin, err := os.ReadFile(*binPath)
	if err != nil {
		log.Fatalf("Failed to read contract bytecode: %v", err)
	}

	parsedABI, err := abi.JSON(bytes.NewReader(contractABI))
	if err != nil {
		log.Fatalf("Failed to parse ABI: %v", err)
	}

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatalf("Failed to get chain ID: %v", err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privKey, chainID)
	if err != nil {
		log.Fatalf("Failed to create auth: %v", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice
	constructorArgs := []interface{}{publicAddress}

	packed, err := parsedABI.Pack("", constructorArgs...)
	if err != nil {
		log.Fatalf("Failed to pack constructor arguments: %v", err)
	}

	data := append(common.FromHex(string(contractBin)), packed...)
	tx := types.NewContractCreation(nonce, big.NewInt(0), auth.GasLimit, auth.GasPrice, data)

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privKey)
	if err != nil {
		log.Fatalf("Failed to sign transaction: %v", err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatalf("Failed to send transaction: %v", err)
	}

	fmt.Printf("Contract deployed! Tx Hash: %s\n", signedTx.Hash().Hex())
}
