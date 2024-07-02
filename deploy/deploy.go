package main

import (
	"bytes"
	"context"
	"crypto/ecdsa"
	"encoding/hex"
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
	keygen := flag.Bool("keygen", false, "Generate and log Ethereum keys")
	infuraURL := flag.String("infura", "", "Infura URL")
	privateKey := flag.String("private", "", "Private key")
	publicKey := flag.String("public", "", "Public key")
	abiPath := flag.String("abi", "", "Path to the ABI file")
	binPath := flag.String("bin", "", "Path to the BIN file")
	flag.Parse()

	if *keygen {
		generateKeys()
	} else {
		deployContract(*infuraURL, *privateKey, *publicKey, *abiPath, *binPath)
	}
}

func generateKeys() {
	// Generate a new private key
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatalf("Failed to generate private key: %v", err)
	}

	// Convert the private key to bytes and log it out
	privateKeyBytes := crypto.FromECDSA(privateKey)
	fmt.Printf("Private Key: %s\n", hex.EncodeToString(privateKeyBytes))

	// Derive the public key from the private key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatalf("Failed to cast public key to ECDSA")
	}

	// Convert the public key to bytes and log it out
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	fmt.Printf("Public Key: %s\n", hex.EncodeToString(publicKeyBytes))

	// Generate the Ethereum address from the public key
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Printf("Ethereum Address: %s\n", address)
}

func deployContract(infuraURL, privateKey, publicKey, abiPath, binPath string) {
	if infuraURL == "" || privateKey == "" || publicKey == "" || abiPath == "" || binPath == "" {
		log.Fatalf("All flags (infura, private, public, abi, bin) are required")
	}

	// Connect to the Ethereum client
	client, err := ethclient.Dial(infuraURL)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	// Load the private key
	privKey, err := crypto.HexToECDSA(privateKey)
	if err != nil {
		log.Fatalf("Invalid private key: %v", err)
	}

	// Get the public address
	publicAddress := common.HexToAddress(publicKey)

	// Get the nonce
	nonce, err := client.PendingNonceAt(context.Background(), publicAddress)
	if err != nil {
		log.Fatalf("Failed to get nonce: %v", err)
	}

	// Get the gas price
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatalf("Failed to get gas price: %v", err)
	}

	// Load the contract ABI and bytecode
	contractABI, err := os.ReadFile(abiPath)
	if err != nil {
		log.Fatalf("Failed to read contract ABI: %v", err)
	}

	contractBin, err := os.ReadFile(binPath)
	if err != nil {
		log.Fatalf("Failed to read contract bytecode: %v", err)
	}

	// Parse the ABI
	parsedABI, err := abi.JSON(bytes.NewReader(contractABI))
	if err != nil {
		log.Fatalf("Failed to parse ABI: %v", err)
	}

	// Set the chain ID (e.g., 1 for mainnet, 3 for ropsten, etc.)
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatalf("Failed to get chain ID: %v", err)
	}

	// Create the auth object
	auth, err := bind.NewKeyedTransactorWithChainID(privKey, chainID)
	if err != nil {
		log.Fatalf("Failed to create auth: %v", err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = gasPrice

	// Deploy the contract
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
