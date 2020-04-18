package main

import (
	crand "crypto/rand"
	"fmt"
	"github.com/seehuhn/mt19937"
	"math"
	"math/big"
	"math/rand"
	"strconv"
)

type Item struct {
	id   int
	name string
}

type ProbabilityCalculator struct {
	denominator int64
	rarities    []Rarity
}

type Rarity struct {
	id    int
	name  string
	rate  int64 // 万分率
	items []Item
}

func main() {
	fmt.Printf("[Information]\n")
	rarities := GetRarities()
	calculator := GetProbabilityCalculator(rarities)
	fmt.Printf("Rarity & Rate(%%):\n")
	for _, r := range rarities {
		fmt.Printf("%s=%s ", r.name, calculator.GetRate(r))
	}
	fmt.Printf("\n")

	fmt.Printf("Rarity & Item:\n")
	for _, r := range rarities {
		fmt.Printf("%s=[ ", r.name)
		for _, i := range r.items {
			fmt.Printf("%s ", i.name)
		}
		fmt.Printf("]\n")
	}
	fmt.Printf("\n")

	fmt.Printf("[Execute]\n")
	fmt.Printf("Draw...\n")
	rarityResult := calculator.DrawRarity()
	fmt.Printf("\n")

	fmt.Printf("[Result]\n")
	fmt.Printf("Rarity:%s\n", rarityResult.name)
	fmt.Printf("Item:---\n")
}

func GetRand() *rand.Rand {
	seed, _ := crand.Int(crand.Reader, big.NewInt(math.MaxInt64))
	rng := rand.New(mt19937.New())
	rng.Seed(seed.Int64())
	return rng
}

func GetRarities() []Rarity {
	items := []Item{
		{1, "きれいな石ころ"},
		{2, "イイカンジの枝"},
		{3, "ドライバー"},
		{4, "ネジ"},
		{5, "泥水"},
		{6, "真水"},
		{7, "付箋"},
		{8, "定規"},
		{9, "土"},
		{10, "草"},
		{11, "イケてるTシャツ"},
		{12, "ビール缶6本セット"},
		{13, "チョコレートアソート"},
		{14, "スタバカード"},
		{15, "医療用マスク詰め合わせ"},
		{16, "ちょっといいぬいぐるみ"},
		{17, "1000円分の商品券"},
		{18, "金の延べ棒"},
		{19, "ダイヤの指輪"},
		{20, "ディズニーペアチケット"},
	}
	return []Rarity{
		{1, "R", 6000, items[0:10]},
		{2, "SR", 3500, items[10:17]},
		{3, "UR", 500, items[17:20]},
	}
}

func GetProbabilityCalculator(rarities []Rarity) ProbabilityCalculator {
	var denominator int64
	for _, r := range rarities {
		denominator += r.rate
	}
	return ProbabilityCalculator{denominator, rarities}
}

func (calculator ProbabilityCalculator) GetRate(rarity Rarity) string {
	return strconv.FormatFloat(float64(rarity.rate)/float64(calculator.denominator)*100.0, 'f', 4, 64)
}

func (calculator ProbabilityCalculator) DrawRarity() Rarity {
	rng := GetRand()
	dart := rng.Int63n(calculator.denominator)

	var result Rarity
	var board int64
	for _, r := range calculator.rarities {
		board += r.rate
		fmt.Printf("Debug:%s %d %d %d\n", r.name, r.rate, board, dart)
		if dart < board {
			result = r
			break
		}
	}

	return result
}

func (rarity Rarity) DrawItem() Item {
	rng := GetRand()
	dart := rng.Int63n(int64(len(rarity.items)))
	return rarity.items[dart]
}
