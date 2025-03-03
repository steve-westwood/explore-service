package main

import (
	"os"

	"github.com/steve-westwood/explore-service/src/persistence"
	"github.com/steve-westwood/explore-service/src/server"
)

func main() {
	db := persistence.NewDB(os.Getenv("DATABASE_URL"))
	recipientService := persistence.NewRecipientService(db)
	grpcServer := server.NewServer(recipientService)
	server.StartServer(grpcServer)
}
