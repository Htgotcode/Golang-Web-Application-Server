package api

import (
	"log"

	tcg "github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg"
	"github.com/PokemonTCG/pokemon-tcg-sdk-go-v2/pkg/request"
)

func GetACard(name, types string) {
	c := tcg.NewClient("pokeAPIKey")

	cards, err := c.GetCards(
		request.Query("name:"+name, "types:"+types),
		request.OrderBy("+name"),
		request.PageSize(3),
		request.Page(2),
	)
	if err != nil {
		log.Fatal(err)
	}

	for _, card := range cards {
		log.Printf("%s: %s\n", card.Name, card.Set.Name)
	}
}

func SearchCards(cardID string) {
	c := tcg.NewClient("pokeAPIKey")

	card, err := c.GetCardByID(cardID)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%s: %s\n", card.Name, card.Set.Name)
}

func GetASet() {
	c := tcg.NewClient("pokeAPIKey")

	sets, err := c.GetSets(
		request.Query("legalities.standard:legal"),
		request.PageSize(5),
		request.Page(2),
	)
	if err != nil {
		log.Fatal(err)
	}

	for _, set := range sets {
		log.Printf("%s: %s\n", set.ID, set.Name)
	}
}

func SearchSets(setID string) {
	c := tcg.NewClient("pokeAPIKey")

	set, err := c.GetSetByID(setID)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%s: %s\n", set.ID, set.Name)
}

func GetTypes() {

}

func GetSubtypes() {

}

func GetSupertypes() {

}

func GetRarities() {

}
