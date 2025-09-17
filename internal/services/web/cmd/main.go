package main

import (
	"github.com/avenuegolangsp/antifraud/internal/services/web/handlers"
	"net/http"
	restful "github.com/emicklei/go-restful/v3"
)

func main() {
	ws := handlers.NewInternalWebRestfulContainer()
	
	restful.DefaultContainer.Router(restful.CurlyRouter{})
	restful.Add(ws.GetWS())

	http.ListenAndServe(":8080", nil)	
}