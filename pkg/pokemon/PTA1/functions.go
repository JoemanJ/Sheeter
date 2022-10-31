package PTA1

import (
	"fmt"
	"os"

	actions "Joe/sheet-hole/pkg/general"

	"github.com/mtslzr/pokeapi-go"
)

const ABILITYDATA string = "./data/testeAbilities.txt"
const MOVEDATA string = "./data/testeMoves.txt"
const SPECIESDATA string = "./data/speciesData"

func RegisterAbility(name string, activation string, description string) (*PokemonAbility, error) {
	abf, err := os.OpenFile(ABILITYDATA, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0444)
	if err != nil {
		abf.Close()
		return nil, err
	}

	defer abf.Chmod(0444)
	defer abf.Close()

	ability := &PokemonAbility{
		IsHighAbility: false,
		Name:          name,
		Activation:    activation,
		Description:   description,
	}

	text := fmt.Sprintf("%s\n%s\n%s\n\n", name, activation, description)

	if text == "" {
		return nil, fmt.Errorf("Ability %s could not be written\n", name)
	}

	_, err = abf.WriteString(text)
	if err != nil {
		return nil, err
	}

	return ability, nil
}

func RegisterMove(name string, Type string, aptitude string, descriptors []string, accDiff int, dice *actions.DiceSet, reach string, frequency string, contests string, effect string) (*PokemonMove, error) {
	mvf, err := os.OpenFile(MOVEDATA, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		mvf.Close()
		return nil, err
	}

	defer mvf.Close()

	move := &PokemonMove{
		Name:        name,
		Type:        Type,
		Aptitude:    aptitude,
		Descriptors: descriptors,
		AccDiff:     accDiff,
		Damage:      dice,
		Reach:       reach,
		Frequency:   frequency,
		Contests:    contests,
		Effect:      effect,
	}

	text := fmt.Sprintf("%s\n%s\n%s\n", name, Type, aptitude)

	for _, str := range descriptors {
		text = string(fmt.Appendf([]byte(text), "%s ", str))
	}

	text = string(fmt.Appendf([]byte(text), "\n%d\n%d %d %d\n%s\n%s\n%s\n%s\n\n", accDiff, dice.N, dice.X, dice.Mod, reach, frequency, contests, effect))

	if text == "" {
		return nil, fmt.Errorf("Move  %s could not be written\n", name)
	}

	_, err = mvf.WriteString(text)
	if err != nil {
		return nil, err
	}

	return move, nil
}

func RegisterSpecies(name string, diet string, abilities []*PokemonAbility) (*PokemonSpecies, error) {
	spcf, err := os.OpenFile(MOVEDATA, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		spcf.Close()
		return nil, err
	}

	defer spcf.Close()

	pokeData, err := pokeapi.Pokemon(name)
	if err != nil {
		panic(err)
	}

	fmt.Println(pokeData)

	return &PokemonSpecies{}, nil
}
