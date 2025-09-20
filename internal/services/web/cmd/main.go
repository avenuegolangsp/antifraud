package main

import (
	"log"
	"net/http"

	"github.com/avenuegolangsp/antifraud/internal/services/web/handlers"
	restful "github.com/emicklei/go-restful/v3"
	_ "github.com/lib/pq"
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
