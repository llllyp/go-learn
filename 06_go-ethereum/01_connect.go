package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/ethclient"
)

/*
连接到节点并获取最新区块号
*/
func main() {
	// 连接到以太坊节点, 使用功能 RPC 端点示例
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal("连接节点失败", err)
	}

	// 获取最新区块号
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal("获取最新区块号失败", err)
	}

	fmt.Printf("最新区块号: %d\n", header.Number)

	// 根据区块号获取完整区块
	blockNumber := big.NewInt(header.Number.Int64() - 1) // 获取上一个区块
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal("获取区块失败", err)
	}

	fmt.Printf("区块哈希: %s\n", block.Hash().Hex())

}
