package handlers

import (
	"github.com/emicklei/go-restful/v3"
)

type RestfulWebServer struct {
	ws *restful.WebService
}

type StaticHandler struct {
	//PubFolder string
}

func (s *StaticHandler) Render200Ok(req *restful.Request, resp *restful.Response) {
	_, _ = resp.Write([]byte("OK"))
}

func NewInternalWebRestfulContainer() *RestfulWebServer {
	ws := new(restful.WebService)
	ws.Consumes("application/json").Produces("application/json")

	cors := restful.CrossOriginResourceSharing{
		AllowedMethods: []string{"GET", "POST", "OPTIONS", "DELETE", "PUT", "PATCH"},
		CookiesAllowed: false,
	}

	antifraudHandler := NewAntifraudHandler()

	ws.Filter(cors.Filter)
	ws.Route(ws.GET("/").To((&StaticHandler{}).Render200Ok))
	ws.Route(ws.POST("/analyze").To(antifraudHandler.AnalyzeTransaction))
	ws.Route(ws.GET("/alerts").To(antifraudHandler.ListAlerts))
	ws.Route(ws.GET("/risk/{transactionId}").To(antifraudHandler.GetRisk))
	ws.Route(ws.GET("/patterns/{userId}").To(antifraudHandler.GetPatterns))
	ws.Route(ws.POST("/rules").To(antifraudHandler.SetRules))
	ws.Route(ws.GET("/stats").To(antifraudHandler.GetStats))
	ws.Route(ws.GET("/health").To(antifraudHandler.HealthCheck))

	return &RestfulWebServer{ws: ws}
}

func (s *RestfulWebServer) GetWS() *restful.WebService {
	return s.ws
}

