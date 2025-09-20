package main

import (
	"log"
	"net/http"

	"github.com/avenuegolangsp/antifraud/internal/services/web/handlers"
	"github.com/avenuegolangsp/antifraud/internal/services/web/repository"
	restful "github.com/emicklei/go-restful/v3"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	ws := handlers.NewInternalWebRestfulContainer()

	restful.DefaultContainer.Router(restful.CurlyRouter{})
	restful.Add(ws.GetWS())

	dsn := "host=localhost user=postgres password=123 dbname=testdb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Migrar tabelas
	db.AutoMigrate(&repository.Transaction{}, &repository.TransactionAnalysis{}, &repository.TransactionAlert{})

	http.ListenAndServe(":8080", nil)
}
