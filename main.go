package Socially

import (
	"database/sql"
	"fmt"
	"github.com/hako/branca"
	"github.com/iamstymsingh/Social-Network/internal/service"
	"github.com/joho/godotenv"
	"log"
	"os"
)

const (
	databaseURL = "postgresql://root@127.0.0.1:26257/defaultdb?sslmode=disable"
	port        = 3000
)

func main() {
	fmt.Println("Socially - social network, simplified.")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := sql.Open("pgx", databaseURL)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
		return
	}

	defer db.Close()
	if err = db.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
		return
	}

	brancaSecretKey := os.Getenv("BRANCA_SECRET_KEY")
	codec := branca.NewBranca(brancaSecretKey)

	// Initialise Service
	_ := service.New(db, codec)
}
