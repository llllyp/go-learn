package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	txHashHex := "0x0f1446ceb26a8504efd149cf63d6290747bd1f32efb17579aa6ada7c08b4f25a"
	queryTransaction(txHashHex)
}

// 查询交易
func queryTransaction(txHashHex string) {
	rpcURL := "http://localhost:8545"
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := ethclient.DialContext(ctx, rpcURL)
	if err != nil {
		log.Fatalf("failed to connect Ethereum node: %v", err)
	}
	defer client.Close()

	txHash := common.HexToHash(txHashHex)

	tx, isPending , err  := client.TransactionByHash(ctx, txHash)
	if err != nil {
		log.Fatalf("failed to get transaction: %v", err)
	}

	fmt.Println("=== Transaction ===")
	printTxBasicInfo(tx, isPending)

	// 回执可能尚不可用(Pending交易)
	receipt, err := client.TransactionReceipt(ctx, txHash)
	if err != nil {
		log.Printf("failed to get receipt(maybe pending): %v", err)
		return
	}

	fmt.Println("=== Receipt ===")
	printReceiptInfo(receipt)

}

// 打印交易信息, 打包发送交易时就会有 transaction 信息
func printTxBasicInfo(tx *types.Transaction, isPending bool) {
	fmt.Printf("Hash        : %s\n", tx.Hash().Hex())
	fmt.Printf("Nonce       : %d\n", tx.Nonce())
	fmt.Printf("Gas         : %d\n", tx.Gas())
	fmt.Printf("Gas Price   : %s\n", tx.GasPrice().String())
	fmt.Printf("To          : %v\n", tx.To())
	fmt.Printf("Value (Wei) : %s\n", tx.Value().String())
	fmt.Printf("Data Len    : %d bytes\n", len(tx.Data()))
	fmt.Printf("Pending     : %v\n", isPending)
}

// 打印成功交易的回执信息 交易成功后才有
func printReceiptInfo(r *types.Receipt) {
	fmt.Printf("Status      : %d\n", r.Status)
	fmt.Printf("BlockNumber : %d\n", r.BlockNumber.Uint64())
	fmt.Printf("BlockHash   : %s\n", r.BlockHash.Hex())
	fmt.Printf("TxIndex     : %d\n", r.TransactionIndex)
	fmt.Printf("Gas Used    : %d\n", r.GasUsed)
	fmt.Printf("Logs        : %d\n", len(r.Logs))
	if len(r.Logs) > 0 {
		fmt.Printf("First Log Address : %s\n", r.Logs[0].Address.Hex())
	}
}
