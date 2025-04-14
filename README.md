# Solana 可证明公平的 Dice 游戏 Demo

这是一个基于 Solana 链的可证明公平的 Dice 游戏简易 Demo，使用 Go 语言结合 `blocto/solana-go-sdk` 和 `kyber` 库开发。

## 项目结构
- `contracts/`: 包含游戏合约逻辑代码
    - `game_contract.go`: 实现游戏的初始化、掷骰子等逻辑
- `client/`: 包含客户端代码
    - `main.go`: 模拟玩家与游戏交互，验证游戏结果并发送交易到 Solana 链
    - `go.mod`: Go 模块依赖文件

## 操作步骤

### 1. 克隆项目
```sh
git clone <项目仓库地址>
cd solana-dice-game