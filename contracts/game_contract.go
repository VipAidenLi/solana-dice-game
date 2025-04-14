package contracts

import (
	"crypto/sha256"
	"fmt"
	"github.com/blocto/solana-go-sdk/common"
	"github.com/dedis/kyber/v3/group/edwards25519"
	"github.com/dedis/kyber/v3/vrf/ecvrf"
	"go.dedis.ch/kyber/v3"
)

// Game 表示游戏状态
type Game struct {
	Player       common.PublicKey
	House        common.PublicKey
	BetAmount    uint64
	TargetNumber int
	VRFPrivKey   kyber.Scalar
	VRFPubKey    kyber.Point
}

// NewGame 创建一个新的游戏实例
func NewGame(player, house common.PublicKey, betAmount uint64, targetNumber int) (*Game, error) {
	suite := edwards25519.NewBlakeSHA256Ed25519()
	privKey := suite.Scalar().Pick(suite.RandomStream())
	pubKey := suite.Point().Mul(privKey, nil)

	return &Game{
		Player:       player,
		House:        house,
		BetAmount:    betAmount,
		TargetNumber: targetNumber,
		VRFPrivKey:   privKey,
		VRFPubKey:    pubKey,
	}, nil
}

// RollDice 执行掷骰子操作并返回结果
func (g *Game) RollDice() (bool, []byte, error) {
	// 生成随机种子
	seed := fmt.Sprintf("%s%s%d%d", g.Player.String(), g.House.String(), g.BetAmount, g.TargetNumber)
	seedHash := sha256.Sum256([]byte(seed))

	// 使用 VRF 生成随机数
	vrf := ecvrf.NewEd25519Sha512()
	proof, err := vrf.Prove(g.VRFPrivKey, seedHash[:])
	if err != nil {
		return false, nil, err
	}

	// 验证 VRF 证明
	ok, err := vrf.Verify(g.VRFPubKey, seedHash[:], proof)
	if err != nil || !ok {
		return false, nil, err
	}

	// 从 VRF 输出中提取随机数
	output, err := vrf.ProofToHash(proof)
	if err != nil {
		return false, nil, err
	}

	// 将随机数映射到 1 - 6 的范围
	result := int(output[0]%6) + 1

	// 判断玩家是否获胜
	win := result == g.TargetNumber

	return win, proof, nil
}
