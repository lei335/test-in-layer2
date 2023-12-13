package main

import (
	"bytes"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/zl/test-in-layer2/go/store"

	com "github.com/memoio/contractsv2/common"
)

var (
	eth   string
	hexSk string

	// AdminAddr admin address
	AdminAddr = common.HexToAddress("0x9c570A8e0783Adf20b982172999Eccad59a17518")
)

// 验证多次质押、分润、取款后，账户在质押池中的余额（质押额和分润之和）以及totalPledge是否与质押池的余额相等
func main() {
	inputeth := flag.String("eth", "", "eth api Address;")
	sk := flag.String("sk", "", "signature for sending transaction")

	flag.Parse()
	eth = *inputeth
	hexSk = *sk

	fmt.Println()

	// get client
	client, err := ethclient.DialContext(context.Background(), eth)
	if err != nil {
		log.Fatal(err)
	}

	// make auth to send transaction
	chainId, err := client.ChainID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	txAuth, err := com.MakeAuth(chainId, hexSk)
	if err != nil {
		log.Fatal(err)
	}
	txAuth.GasPrice = nil

	// deploy Store
	storeAddr, tx, storeIns, err := store.DeployStore(txAuth, client)
	if err != nil {
		log.Fatal("deploy Store err:", err)
	}
	fmt.Println("store addr: ", storeAddr.Hex())
	err = com.CheckTx(eth, tx.Hash(), "deploy Store")
	if err != nil {
		log.Fatal(err)
	}

	// get tx info
	receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	receiptJson, err := receipt.MarshalJSON()
	if err != nil {
		log.Fatal(err)
	}
	// format input
	var str bytes.Buffer
	json.Indent(&str, receiptJson, "", " ")
	fmt.Println(str.String())

	// test, send 10 transactions
	fmt.Println()
	for i := 0; i < 10; i++ {
		tx, err = storeIns.Add(txAuth, big.NewInt(1))
		if err != nil {
			fmt.Println(i, err)
		}
		time.Sleep(time.Second)
	}
	err = com.CheckTx(eth, tx.Hash(), "store add")
	if err != nil {
		log.Fatal(err)
	}

	// get tx info
	receipt, err = client.TransactionReceipt(context.Background(), tx.Hash())
	if err != nil {
		log.Fatal(err)
	}
	receiptJson, err = receipt.MarshalJSON()
	if err != nil {
		log.Fatal(err)
	}
	// format input
	str.Reset()
	json.Indent(&str, receiptJson, "", " ")
	fmt.Println(str.String())

	// get count
	fmt.Println()
	time.Sleep(2*time.Second)
	count, err := storeIns.GetCount(nil, AdminAddr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("count: ", count)
}
