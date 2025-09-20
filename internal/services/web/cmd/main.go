
package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/avenuegolangsp/antifraud/internal/services/web/handlers"
	restful "github.com/emicklei/go-restful/v3"
	_ "github.com/lib/pq"
)

func main() {
	// ...existing code...
	db, err := sql.Open("postgres", "postgres://avenue:password@localhost:5432/antifraud?sslmode=disable")
	if err != nil {
		log.Fatalf("failed to connect to db: %v", err)
	}
	repo := handlers.NewTransactionRepository(db)
	antifraudHandler := &handlers.AntifraudHandler{Repo: repo}

	ws := new(restful.WebService)
	ws.Consumes("application/json").Produces("application/json")

	cors := restful.CrossOriginResourceSharing{
		AllowedMethods: []string{"GET", "POST", "OPTIONS", "DELETE", "PUT", "PATCH"},
		CookiesAllowed: false,
	}
	ws.Filter(cors.Filter)
	ws.Route(ws.GET("/").To((&handlers.StaticHandler{}).Render200Ok))
	ws.Route(ws.POST("/analyze").To(antifraudHandler.AnalyzeTransaction))
	ws.Route(ws.GET("/alerts").To(antifraudHandler.ListAlerts))
	ws.Route(ws.GET("/risk/{transactionId}").To(antifraudHandler.GetRisk))
	ws.Route(ws.POST("/rules").To(antifraudHandler.SetRules))
	ws.Route(ws.GET("/stats").To(antifraudHandler.GetStats))
	ws.Route(ws.GET("/health").To(antifraudHandler.HealthCheck))
	ws.Route(ws.GET("/clients").To(antifraudHandler.ListClientsWithTransactions))

	restful.DefaultContainer.Router(restful.CurlyRouter{})
	restful.Add(ws)

	log.Println("Servidor iniciado em :8080")
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	ws := handlers.NewInternalWebRestfulContainer()

	restful.DefaultContainer.Router(restful.CurlyRouter{})
	restful.Add(ws.GetWS())

	dsn := "host=localhost user=avenue password=password dbname=antifraud port=5432 sslmode=disable"
	_, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	http.ListenAndServe(":8080", nil)
}
