package main

import (
	"context"
	"fmt"
	"log"

	"github.com/blocto/solana-go-sdk/client"
	"github.com/blocto/solana-go-sdk/rpc"
	"github.com/blocto/solana-go-sdk/types"
)

func main() {
	// 初始化玩家和庄家账户
	playerAccount, err := types.AccountFromBase58("your-player-private-key-base58")
	if err != nil {
		log.Fatalf("Failed to create player account: %v", err)
	}
	houseAccount, err := types.AccountFromBase58("your-house-private-key-base58")
	if err != nil {
		log.Fatalf("Failed to create house account: %v", err)
	}

	// 玩家下注
	betAmount := uint64(100)
	targetNumber := 3

	// 创建游戏实例
	game, err := contracts.NewGame(playerAccount.PublicKey, houseAccount.PublicKey, betAmount, targetNumber)
	if err != nil {
		log.Fatalf("Failed to create game: %v", err)
	}

	// 玩家掷骰子
	win, proof, err := game.RollDice()
	if err != nil {
		log.Fatalf("Failed to roll dice: %v", err)
	}

	// 打印结果
	fmt.Printf("Player bet %d on number %d\n", betAmount, targetNumber)
	if win {
		fmt.Println("Player wins!")
	} else {
		fmt.Println("Player loses!")
	}
	fmt.Printf("VRF proof: %x\n", proof)

	// 发送交易到 Solana 链
	c := client.NewClient(rpc.LocalnetRPCEndpoint)
	res, err := c.GetLatestBlockhash(context.Background())
	if err != nil {
		log.Fatalf("Get recent block hash error: %v", err)
	}

	// 这里可以根据输赢情况添加相应的转账指令，示例中仅为占位
	var instructions []types.Instruction
	if win {
		// 玩家获胜，庄家给玩家转账
		// instructions = append(instructions, ...)
	} else {
		// 玩家失败，玩家给庄家转账
		// instructions = append(instructions, ...)
	}

	tx, err := types.NewTransaction(types.NewTransactionParam{
		Message: types.NewMessage(types.NewMessageParam{
			FeePayer:        playerAccount.PublicKey,
			RecentBlockhash: res.Blockhash,
			Instructions:    instructions,
		}),
		Signers: []types.Account{playerAccount},
	})
	if err != nil {
		log.Fatalf("Generate tx error: %v", err)
	}

	txhash, err := c.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatalf("Send tx error: %v", err)
	}

	fmt.Println("Transaction hash:", txhash)
}
