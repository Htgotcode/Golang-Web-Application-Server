package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Htgotcode/Golang-Web-Application-Server/api/models"
	"github.com/cockroachdb/cockroach-go/v2/crdb/crdbgorm"
	"github.com/google/uuid"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var cardIDs []uuid.UUID

func addCards(db *gorm.DB, numRows int, transferCard int) error {
	log.Printf("Creating %d new cards...", numRows)
	for i := 0; i < numRows; i++ {
		newID := uuid.New()
		//newBalance := rand.Intn(10000) + transferAmount
		if err := db.Create(&models.Card{}).Error; err != nil {
			return err
		}
		cardIDs = append(cardIDs, newID)
	}
	log.Println("Cards created.")
	return nil
}

func transferCard(db *gorm.DB, fromID uuid.UUID, toID uuid.UUID, card string) error {
	log.Printf("Transferring %d from user %s to user %s...", card, fromID, toID)
	var fromSeller models.Card
	var toBuyer models.Card

	db.First(&fromSeller, fromID)
	db.First(&toBuyer, toID)

	log.Println("Cards transferred.")
	return nil
}

func printCardPurchases(db *gorm.DB) {
	var cards []models.Card
	db.Find(&cardIDs)
	fmt.Printf("Balance at '%s':\n", time.Now())
	for _, cards := range cards {
		fmt.Printf("%s %d\n", cards.ID, cards.CardID)
	}
}

func removeCards(db *gorm.DB, accountIDs []uuid.UUID) error {
	log.Println("Deleting accounts created...")
	err := db.Where("id IN ?", cardIDs).Delete(models.Card{}).Error
	if err != nil {
		return err
	}
	log.Println("Accounts deleted.")
	return nil
}

func TestDBSetup() {
	// Read in connection string

	fmt.Println("Connecting to database...")
	connstring := os.Getenv("cockroachDB_URL")

	// Connect to the "card" database
	db, err := gorm.Open(postgres.Open(connstring), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: "bank.",
		},
	})
	if err != nil {
		log.Fatal("error configuring the database: ", err)
	}

	db.AutoMigrate(&models.Card{})

	printCardPurchases(db)

	//fromID := cardIDs[0]
	//toID := cardIDs[0:][rand.Intn(len(cardIDs))]

	printCardPurchases(db)

	if err := crdbgorm.ExecuteTx(context.Background(), db, nil,
		func(tx *gorm.DB) error {
			return removeCards(db, cardIDs)
		},
	); err != nil {
		fmt.Println(err)
	}
}
