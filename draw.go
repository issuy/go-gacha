package main

import (
	"fmt"
	"strconv"
)

type ProbabilityCalculator struct {
	denominator int
	rarities    []Rarity
}

type Rarity struct {
	id   int
	name string
	rate int // 万分率
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
	fmt.Printf("[Execute]\n")
	fmt.Printf("Draw\n")
	fmt.Printf("[Result]\n")
	fmt.Printf("Rarity:---\n")
	fmt.Printf("Item:---\n")
}

func GetRarities() []Rarity {
	return []Rarity{
		{1, "R", 6000},
		{2, "SR", 3500},
		{3, "UR", 500},
	}
}

func GetProbabilityCalculator(rarities []Rarity) ProbabilityCalculator {
	denominator := 0
	for _, r := range rarities {
		denominator += r.rate
	}
	return ProbabilityCalculator{denominator, rarities}
}

func (calculator ProbabilityCalculator) GetRate(rarity Rarity) string {
	return strconv.FormatFloat(float64(rarity.rate)/float64(calculator.denominator)*100.0, 'f', 4, 64)
}
