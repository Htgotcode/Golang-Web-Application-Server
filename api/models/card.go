package models

import (
	"time"
)

type Card struct {
	ID           string    `bson:"id"`
	Brand        string    `json:"brand" validate:"required"`
	Set          string    `json:"set" validate:"required"`
	Rarity       string    `json:"rarity" validate:"required"`
	SellingPrice uint      `json:"selling_price" validate:"required"`
	UploadedAt   time.Time `json:"uploaded_at"`
	CardID       string    `json:"card_id"`
}

type Set struct {
	ID           string `bson:"id"`
	Name         string `bson:"name"`
	Series       string `bson:"series"`
	PrintedTotal uint   `bson:"printedTotal"`
	Total        uint   `bson:"total"`
	//Legalities 		hash		`bson:"legalities"`
	PtcgoCode   string `bson:"ptcgoCode"`
	ReleaseDate string `bson:"releaseDate"`
	UpdatedAt   string `bson:"updatedAt"`
	//Images 			hash		`bson:"images"`
}

type PokeCard struct {
	ID           string            `bson:"id"`
	Name         string            `bson:"name"`
	Supertype    string            `bson:"supertype"`
	Subtypes     []string          `bson:"subtypes"`
	Level        string            `bson:"level"`
	HP           uint              `bson:"hp"`
	Types        []string          `bson:"types"`
	EvolvesFrom  string            `bson:"evolvesFrom"`
	EvolvesTo    []string          `bson:"evolvesTo"`
	Rules        []string          `bson:"rules"`
	AncientTrait map[string]string `bson:"ancientTrait"`
	//Abilities    list(hash)			`bson:"abilities"`
	//Attacks      list(hash)			`bson:"attacks"`
	//Weaknesses   list(hash)			`bson:"weaknesses"`
	//Resistances  list(hash)			`bson:"resistances"`
	RetreatCost          []string `bson:"retreatCost"`
	ConvertedRetreatCost uint     `bson:"convertedRetreatCost"`
	//Set 		 hash				`bson:"set"`
	Number                 string `bson:"number"`
	Artist                 string `bson:"artist"`
	Rarity                 string `bson:"rarity"`
	FlavorText             string `bson:"flavorText"`
	NationalPokedexNumbers []uint `bson:"nationalPokedexNumbers"`
	//Legalities   hash				`bson:"legalities"`
	//Images       hash				`bson:"images"`
	//Tcgplayer    hash				`bson:"tcgplayer"`
	//Cardmarket   hash				`bson:"cardmarket"`
}
