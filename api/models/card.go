package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Card struct {
	ID           primitive.ObjectID `bson:"_id"`
	Name         string             `bson:"name"`
	Description  string             `bson:"description"`
	Brand        string             `json:"brand" validate:"required"`
	SetName      string             `json:"set_name" validate:"required"`
	Rarity       string             `json:"rarity" validate:"required"`
	SellingPrice uint               `json:"selling_price" validate:"required"`
	UploadedAt   time.Time          `json:"uploaded_at"`
	CardID       string             `json:"card_id"`
}

type Set struct {
	ID           primitive.ObjectID `bson:"_id"`
	Name         string             `bson:"name"`
	Series       string             `bson:"series"`
	PrintedTotal uint               `bson:"printedTotal"`
	Total        uint               `bson:"total"`

	Legalities struct {
		Standard  string `bson:"standard"`
		Expanded  string `bson:"expanded"`
		Unlimited string `bson:"unlimited"`
	} `bson:"legalities"`

	PtcgoCode   string `bson:"ptcgoCode"`
	ReleaseDate string `bson:"releaseDate"`
	UpdatedAt   string `bson:"updatedAt"`

	Images struct {
		Small string `bson:"small"`
		Large string `bson:"large"`
	} `bson:"images"`
}

type PokeCard struct {
	ID          primitive.ObjectID `bson:"_id"`
	Name        string             `bson:"name"`
	Supertype   string             `bson:"supertype"`
	Subtypes    []string           `bson:"subtypes"`
	Level       string             `bson:"level"`
	HP          uint               `bson:"hp"`
	Types       []string           `bson:"types"`
	EvolvesFrom string             `bson:"evolvesFrom"`
	EvolvesTo   []string           `bson:"evolvesTo"`
	Rules       []string           `bson:"rules"`

	AncientTrait struct {
		Name string `bson:"name"`
		Text string `bson:"text"`
	} `bson:"ancientTrait"`

	Abilities []struct {
		Name string `bson:"name"`
		Text string `bson:"text"`
		Type string `bson:"type"`
	} `bson:"abilities"`

	Attacks []struct {
		Cost                []string `bson:"cost"`
		Name                string   `bson:"name"`
		Text                string   `bson:"text"`
		Damage              string   `bson:"damage"`
		ConvertedEnergyCost int      `bson:"convertedEnergyCost"`
	} `bson:"attacks"`

	Weaknesses []struct {
		Type  string `bson:"type"`
		Value string `bson:"value"`
	} `bson:"weaknesses"`

	Resistances []struct {
		Type  string `bson:"type"`
		Value string `bson:"value"`
	} `bson:"resistances"`

	RetreatCost          []string `bson:"retreatCost"`
	ConvertedRetreatCost uint     `bson:"convertedRetreatCost"`

	Set struct {
		ID           string `bson:"id"`
		Name         string `bson:"name"`
		Series       string `bson:"series"`
		PrintedTotal uint   `bson:"printedTotal"`
		Total        uint   `bson:"total"`

		Legalities struct {
			Standard  string `bson:"standard"`
			Expanded  string `bson:"expanded"`
			Unlimited string `bson:"unlimited"`
		} `bson:"legalities"`

		PtcgoCode   string `bson:"ptcgoCode"`
		ReleaseDate string `bson:"releaseDate"`
		UpdatedAt   string `bson:"updatedAt"`

		Images struct {
			Small string `bson:"small"`
			Large string `bson:"large"`
		} `bson:"images"`
	} `bson:"set"`

	Number                 string `bson:"number"`
	Artist                 string `bson:"artist"`
	Rarity                 string `bson:"rarity"`
	FlavorText             string `bson:"flavorText"`
	NationalPokedexNumbers []uint `bson:"nationalPokedexNumbers"`

	Legalities struct {
		Standard  string `bson:"standard"`
		Expanded  string `bson:"expanded"`
		Unlimited string `bson:"unlimited"`
	} `bson:"legalities"`

	Images struct {
		Small string `bson:"small"`
		Large string `bson:"large"`
	} `bson:"images"`

	Tcgplayer struct {
		URL       string `bson:"url"`
		UpdatedAt string `bson:"updatedAt"`

		Prices struct {
			AverageSellPrice float64 `bson:"averageSellPrice"`
			LowPrice         float64 `bson:"lowPrice"`
			TrendPrice       float64 `bson:"trendPrice"`
			GermanProLow     float64 `bson:"germanProLow"`
			SuggestedPrice   float64 `bson:"suggestedPrice"`
			ReverseHoloSell  float64 `bson:"reverseHoloSell"`
			ReverseHoloLow   float64 `bson:"reverseHoloLow"`
			ReverseHoloTrend float64 `bson:"reverseHoloTrend"`
			LowPriceExPlus   float64 `bson:"lowPriceExPlus"`
			Avg1             float64 `bson:"avg1"`
			Avg7             float64 `bson:"avg7"`
			Avg30            float64 `bson:"avg30"`
			ReverseHoloAvg1  float64 `bson:"reverseHoloAvg1"`
			ReverseHoloAvg7  float64 `bson:"reverseHoloAvg7"`
			ReverseHoloAvg30 float64 `bson:"reverseHoloAvg30"`
		} `bson:"prices"`
	} `bson:"cardmarket"`
}

type ScryfallBulkData struct {
	ID               primitive.ObjectID `bson:"id"`
	URI              string             `bson:"uri"`
	Type             string             `bson:"type"`
	Name             string             `bson:"name"`
	Description      string             `bson:"description"`
	Download_uri     string             `bson:"download_uri"`
	Updated_at       time.Time          `bson:"updated_at"`
	Compressed_size  int                `bson:"compressed_size"`
	Content_type     string             `bson:"content_type"`
	Content_encoding string             `bson:"content_encoding"`
}
