package main

/* 交易结构
- nonce: 账户已发现交易数
- to: 接收方地址(部署合约时未空)
- value: 转账金额(ETH)
- data: 调用合约时的编码数据(函数选择器 + 参数)
- gasLimit: 愿意为本次执行支付的最大Gas数量
- gasPrice / maxFeePerGas / maxPriorityFeePerGas: Gas单价相关字段
*/
// ETH 交易的生命周期
// 1, 创建与签名(本地)
// 构造交易对象: 填充 nonce to value data gasLimit gasPrice等字段
// 使用私钥签名: 适用发送方账户的私钥对交易进行ECDSA签名
// 生成交易哈希: 签名后的交易会生成唯一的交易哈希(txHash), 用于后续查询和追踪

// 2, 广播到网络(Pending)
// 发送到节点: 通过RPC调用, 将签名后的交易发送到以太坊节点
// 进入交易池(Mempool): 节点验证交易有效性(签名,nonce等)后, 将交易放入本地交易池
// 网络传播: 节点将交易广播给其他对等节点, 交易在网络中传播
// 状态: Pending: 此时交易尚未被打包进区块, 可以通过 eth_getTransactionByHash 查询到交易, 但返回null

// 3, 打包进区块(Mined)
// 验证者选择: 验证者(Validator)从交易池中选择交易打包进新区块
// 执行交易: 在区块中执行交易, 更新账户状态(余额, nonce等)
// 声场交易回执: 执行完成后生成交易回执(Receipt), 包含执行结果, Gas消耗, 事件日志等
// 状态: 已打包: 此时可以通过 eth_getTransactionReceipt获取交易回执, Status字段为1(成功)

// 4, 确认(Confirmed)
// 区块确认: 交易所在区块被后续区块引用 每增加一个后续区块, 确认数+1
// 最终性: 在PoS机制下, 经过一定数量的确认后, 交易达到最终性(Finality), 几乎不可能被回滚
// 状态: 已确认: 交易成功为链上历史的一部分, 可以查询任意历史块的状态

import (
	"context"
	"flag"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := ethclient.DialContext(ctx, "http://localhost:8545")
	if err != nil {
		log.Fatal("连接以太坊失败", err)
	}
	defer client.Close()

	chainId, err := client.ChainID(ctx)
	if err != nil {
		log.Fatalf("failed to get chain id: %v", err)
	}

	header, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		log.Fatalf("failed to get latest block header: %v", err)
	}

	fmt.Println("=== Ethereum Node Info ===")
	fmt.Printf("Chain id : %s\n", chainId.String())
	fmt.Printf("Latest Block: %d\n", header.Number.Uint64())
	fmt.Printf("Block Hash : %s\n", header.Hash().Hex())
	fmt.Printf("Block Time :%s\n", time.Unix(int64(header.Time), 0).Format(time.RFC3339))

	// 获取最新区块 指定区块 以及批量查询区块范围的信息
	// flag 用来解析 命令行参数  flag.Typename(名字 string, 默认值 T, 帮助说明 string) *T
	// 定义 flag (返回的是指针)
	blockNumberFlag := flag.Uint64("number", 0, "block number to query (0 means skip)")
	// rangeStartFlag := flag.Uint64("range-start", 0, "start block number for range query")
	// rangeEndFlag := flag.Uint64("range-end", 0, "end block number for range query")
	// rateLimitFlag := flag.Int("rate-limit", 200, "rate limit in milliseconds between requests")

	flag.Parse() // 必须调用, 否则拿到的永远是默认值

	/*
		go run xxx.go -number 123456
		go run xxx.go -number=123456
		go run xxx.go --number=123456
	*/

	latestBlock, err := client.BlockByNumber(ctx, nil)
	if err != nil {
		log.Fatalf("failed to get latest block: %v", err)
	}
	printBlockInfo("Latest Block", latestBlock)

	// 通过输出参数查询指定区块
	if *blockNumberFlag > 0 {
		num := big.NewInt(0).SetUint64(*blockNumberFlag)
		block, err := fetchBlockWithRetry(ctx, client, num, 3)
		if err != nil {
			log.Fatalf("failed to get block %d: %v", *blockNumberFlag)
		}

		printBlockInfo(fmt.Sprintf("Block %d", *blockNumberFlag), block)
	}

}

func fetchBlockWithRetry(ctx context.Context, client *ethclient.Client, blockNumber *big.Int, maxRetries int) (*types.Block, error) {
	var lastErr error
	for i := 0; i < maxRetries; i++ {
		// 每次重试适用新的超时上下文, 避免上下文被取消
		reqCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
		block, err := client.BlockByNumber(reqCtx, blockNumber)
		cancel()

		if err == nil {
			return block, nil
		}

		lastErr = err
		if i < maxRetries-1 {
			backoff := time.Duration(i+1) * 500 * time.Millisecond
			log.Printf("[WARN] failed to fetch block %s, retry %d/%d after %v/v%", 
			blockNumber.String(), i+1, maxRetries, backoff, err)
			time.Sleep(backoff)
		}
	}
	return nil, fmt.Errorf("failed fater %d retries: %w", maxRetries, lastErr)
}


// printBlockInfo 打印详细的区块信息
func printBlockInfo(title string, block *types.Block) {
	fmt.Println("======================================")
	fmt.Println(title)
	fmt.Println("======================================")
	fmt.Printf("Block: %+v\n", block)

	// 基本信息
	fmt.Printf("Number       : %d\n", block.Number().Uint64())
	fmt.Printf("Hash         : %s\n", block.Hash().Hex())
	fmt.Printf("Parent Hash  : %s\n", block.ParentHash().Hex())

	// 时间信息
	blockTime := time.Unix(int64(block.Time()), 0)
	fmt.Printf("Time         : %s\n", blockTime.Format(time.RFC3339))
	fmt.Printf("Time (Local) : %s\n", blockTime.Local().Format("2006-01-02 15:04:05 MST"))

	// Gas 信息
	gasUsed := block.GasUsed()
	gasLimit := block.GasLimit()
	gasUsagePercent := float64(gasUsed) / float64(gasLimit) * 100
	fmt.Printf("Gas Used     : %d (%.2f%%)\n", gasUsed, gasUsagePercent)
	fmt.Printf("Gas Limit    : %d\n", gasLimit)

	// 交易信息
	txCount := len(block.Transactions())
	fmt.Printf("Tx Count     : %d\n", txCount)

	// 区块根信息（Merkle 树根）
	fmt.Printf("State Root   : %s\n", block.Root().Hex())
	fmt.Printf("Tx Root      : %s\n", block.TxHash().Hex())
	fmt.Printf("Receipt Root : %s\n", block.ReceiptHash().Hex())

	// 区块大小估算（简化版，实际大小还包括其他字段）
	if txCount > 0 {
		fmt.Printf("\nFirst Tx Hash: %s\n", block.Transactions()[0].Hash().Hex())
		if txCount > 1 {
			fmt.Printf("Last Tx Hash : %s\n", block.Transactions()[txCount-1].Hash().Hex())
		}
	}

	// 难度信息（PoW 相关，PoS 后基本固定）
	fmt.Printf("Difficulty   : %s\n", block.Difficulty().String())

	// 区块奖励相关信息
	coinbase := block.Coinbase()
	if coinbase != (common.Address{}) {
		fmt.Printf("Coinbase     : %s\n", coinbase.Hex())
	}

	fmt.Println("======================================")
	fmt.Println()
}

