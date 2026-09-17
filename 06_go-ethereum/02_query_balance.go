package main

import (
	"context"
	"log"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	// 连接以太坊节点
	client, err := ethclient.Dial("http://localhost:8545")
	if err != nil {
		log.Fatal("连接节点失败", err)
	}

	// 要查询的地址
	address := common.HexToAddress("0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65")

	// 查询余额, nil表示最新区块
	balance, err := client.BalanceAt(context.Background(), address, nil)
	if err != nil {
		log.Fatal("查询余额失败", err)
	}

	// 余额单位是 Wei (1 ETH = 10^18 Wei)
	fmt.Printf("余额 (Wei): %s\n", balance.String())

	// 转换为 ETH
	fBalance := new(big.Float).SetInt(balance)
	// big.NewFloat(1e18) 创建一个值为 10^18 的大浮点数
    // Quo = quotient（商），执行 fBalance / 1e18，结果仍是 *big.Float
	ethValue := new(big.Float).Quo(fBalance, big.NewFloat(1e18))
    fmt.Printf("余额 (ETH): %s\n", ethValue.Text('f', 6))

	/*
	go math包常用API介绍 -> 00hello_world/13_math_demo.go
	*/
}