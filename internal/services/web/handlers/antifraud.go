package handlers

import (
	"github.com/avenuegolangsp/antifraud/internal/services/web/handlers/rules"
	"github.com/emicklei/go-restful/v3"
)

type AntifraudHandler struct {
	RuleManager rules.IRuleManager
}

func NewAntifraudHandler() *AntifraudHandler {
	return &AntifraudHandler{
		RuleManager: rules.NewRuleManager(),
	}
}

func (h *AntifraudHandler) AnalyzeTransaction(req *restful.Request, resp *restful.Response) {

	h.RuleManager.AnalyzeTransaction(rules.AnalyzeRequest{})

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
	if _, err := resp.Write([]byte("OK - HealthCheck")); err != nil {
		resp.WriteHeaderAndEntity(500, map[string]string{"error": "Health check failed"})
	} else {
		resp.WriteHeader(200)
	}
}

func (h *AntifraudHandler) GetStats(req *restful.Request, resp *restful.Response) {
	_, _ = resp.Write([]byte("OK - GetStats"))
}
