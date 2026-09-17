`go-ethereum` (通常称为 Geth) 是用 Go 语言实现的以太坊协议官方客户端。它既是一个完整的以太坊节点软件，也提供了一套丰富的 Go 语言库，允许开发者在应用程序中直接与以太坊区块链交互。

# 核心概念

- Geth 节点: `go-ethereum` 本身是一个命令行工具, 可以启动一个以太坊节点, 同步主网、测试网或私链的数据
- JSON-RPC: Geth 节点通过 JSON-RPC 接口对外提供服务, go 程序本质上是作为一个客户端, 通过 HTTP、WebSocket 或 IPC 链接到 Geth 节点, 然后发送 RPC 请求
- ethclient: 这是 `go-ethereum` 库中最重要的包之一. 它封装了底层 JSON-RPC 调用, 可以用 Go 方法(如 `BlockByNumber`, `BalanceAt`)来与区块链交互, 无需手动拼接 JSON 请求

