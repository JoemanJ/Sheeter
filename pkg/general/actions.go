package actions

import "math/rand"

type DiceSet struct {
	X   int
	N   int
	Mod int
}

// Returns the result of a pseudo-random "XdN + mod" dice roll
func Roll(X int, N int, mod int) int {
	total := 0
	for i := 0; i < X; i++ {
		total += rand.Intn(N) + 1
	}

	return total + mod
}

// Rolls a preset roll object (type DiceSet)
func RollSet(roll *DiceSet) int {
	total := 0
	for i := 0; i < roll.X; i++ {
		total += rand.Intn(roll.N) + 1
	}

	return total + roll.Mod
}

func CreateDiceSet(X int, N int, mod int) *DiceSet {
	return &DiceSet{X: X, N: N, Mod: mod}
}

func Capped(value, minCap, maxCap int) int {
	if value < minCap {
		return minCap
	}

	if value > maxCap {
		return maxCap
	}

	return value
}
