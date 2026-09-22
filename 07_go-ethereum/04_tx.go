package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	toAddrHex := "0xd4d249A88D6cB574c273e0F3a0b2eC6A937eAe6F"
	amountEth := 99

	sendTransaction(toAddrHex, float64(amountEth))
}

func sendTransaction(toAddrHex string, amountEth float64) {
	rpcURL := "http://localhost:8545"
	senderPrivateKey := "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	// senderPrivateKey := "0xb528c71913fd8135b7d824d01a0949023fe2ff883914a998adcb0349f8eac7a2"

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := ethclient.DialContext(ctx, rpcURL)
	if err != nil {
		log.Fatalf("failed to connect to Ethereum node: %v", err)
	}

	// 解析私钥
	privKey, err := crypto.HexToECDSA(trim0x(senderPrivateKey))
	if err != nil {
		log.Fatalf("invalid private key %v", err)
	}

	// 获取发送方地址
	publicKey := privKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddr := crypto.PubkeyToAddress(*publicKeyECDSA)
	toAddr := common.HexToAddress(toAddrHex)

	// 获取链Id
	chanID, err := client.ChainID(ctx)
	if err != nil {
		log.Fatalf("failed to get chain id: %v", err)
	}

	// 获取nonce
	nonce, err := client.PendingNonceAt(ctx, fromAddr)
	if err != nil {
		log.Fatalf("failed to get nonce: %v", err)
	}

	// 获取gas价格(使用 EIP-1559动态费用)
	gasTipCap, err := client.SuggestGasTipCap(ctx)
	if err != nil {
		log.Fatalf("failed to get gas tip cap: %v", err)
	}

	// 获取 base fee, 计算 fee cap
	header, err := client.HeaderByNumber(ctx, nil)
	if err != nil {
		log.Fatalf("failed to get header: %v", err)
	}

	baseFee := header.BaseFee
	if baseFee == nil {
		// 如果比支持EIP-1559, 适用传统 gas price
		gasPrice, err := client.SuggestGasPrice(ctx)
		if err != nil {
			log.Fatalf("failed to get gas price: %v", err)
		}
		baseFee = gasPrice
	}

	// fee cap = base fee * 2 + tim cap(简单策略)
	gasFeeCap := new(big.Int).Add(new(big.Int).Mul(baseFee, big.NewInt(2)), gasTipCap)

	// 估算 gas limit(普通装张固定21000)
	gasLimit := uint64(21000)

	// 转换 ETH 金额为 wei
	amountWei := new(big.Float).Mul(big.NewFloat(amountEth), big.NewFloat(1e18))

	valueWei, _ := amountWei.Int(nil)

	// 检查余额是否足够
	balance, err := client.BalanceAt(ctx, fromAddr, nil)
	if err != nil {
		log.Fatalf("failed to get balance: %v", err)
	}

	// value + gasFeeCap * gasLimit
	totalCost := new(big.Int).Add(valueWei, new(big.Int).Mul(gasFeeCap, big.NewInt(int64(gasLimit))))

	if balance.Cmp(totalCost) < 0 {
		log.Fatalf("insuffeicient balance: bave %s wei, need %s wei",
			balance.String(), totalCost.String(),
		)
	}

	// 构造交易(EIP-1559 动态交易费用)
	txData := &types.DynamicFeeTx{
		ChainID:   chanID,
		Nonce:     nonce,
		GasTipCap: gasTipCap,
		GasFeeCap: gasFeeCap,
		Gas:       gasLimit,
		To:        &toAddr,
		Value:     valueWei,
		Data:      nil,
	}

	tx := types.NewTx(txData)

	// 签名交易
	signer := types.NewLondonSigner(chanID)
	signedTx, err := types.SignTx(tx, signer, privKey)
	if err != nil {
		log.Fatalf("failed to sign stransaction: %v", err)
	}

	// 发送交易
	if err := client.SendTransaction(ctx, signedTx); err != nil {
		log.Fatalf("failed to send transaction: %v", err)
	}

	// 输出交易信息
	fmt.Println("=== Transaction Sent ===")
	fmt.Printf("From       : %s\n", fromAddr.Hex())
	fmt.Printf("To         : %s\n", toAddr.Hex())
	fmt.Printf("Value      : %s ETH (%s Wei)\n", fmt.Sprintf("%.6f", amountEth), valueWei.String())
	fmt.Printf("Gas Limit  : %d\n", gasLimit)
	fmt.Printf("Gas Tip Cap: %s Wei\n", gasTipCap.String())
	fmt.Printf("Gas Fee Cap: %s Wei\n", gasFeeCap.String())
	fmt.Printf("Nonce      : %d\n", nonce)
	fmt.Printf("Tx Hash    : %s\n", signedTx.Hash().Hex())
	fmt.Println("\nTransaction is pending. Use --tx flag to query status:")
	fmt.Printf("  go run main.go --tx %s\n", signedTx.Hash().Hex())

}

// trim0x 移除十六进制字符串前缀 "0x"
func trim0x(s string) string {
	if len(s) >= 2 && s[:2] == "0x" {
		return s[2:]
	}
	return s
}
