package main

import "fmt"

type Rarity struct {
	id   int
	name string
	rate int // 万分率
}

func main() {
	fmt.Printf("[Information]\n")
	fmt.Printf("Rarity & Rate:\n")
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
