package handlers

import "github.com/emicklei/go-restful/v3"

type AntifraudHandler struct {
}

func (h *AntifraudHandler) AnalyzeTransaction(req *restful.Request, resp *restful.Response) {
	_, _ = resp.Write([]byte("OK - AnalyzeTransaction"))
}

func (h *AntifraudHandler) ListAlerts(req *restful.Request, resp *restful.Response) {
	_, _ = resp.Write([]byte("OK - ListAlerts"))
}

func (h *AntifraudHandler) GetRisk(req *restful.Request, resp *restful.Response) {
	_, _ = resp.Write([]byte("OK - GetRisk"))
}

func (h *AntifraudHandler) GetPatterns(req *restful.Request, resp *restful.Response) {
	_, _ = resp.Write([]byte("OK - GetPatterns"))
}

func (h *AntifraudHandler) SetRules(req *restful.Request, resp *restful.Response) {
	_, _ = resp.Write([]byte("OK - SetRules"))
}

func (h *AntifraudHandler) HealthCheck(req *restful.Request, resp *restful.Response) {
	_, _ = resp.Write([]byte("OK - HealthCheck"))
}

func (h *AntifraudHandler) GetStats(req *restful.Request, resp *restful.Response) {
	_, _ = resp.Write([]byte("OK - GetStats"))
}