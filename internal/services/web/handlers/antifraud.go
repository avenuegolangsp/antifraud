package handlers

import (
	"fmt"
	"net/http"

	"github.com/avenuegolangsp/antifraud/internal/services/web/dto/request"
	"github.com/avenuegolangsp/antifraud/internal/services/web/handlers/rules"
	"github.com/emicklei/go-restful/v3"
	"github.com/go-playground/validator/v10"
)

type AntifraudHandler struct {
	RuleManager rules.IRuleManager
	validator   *validator.Validate
}

func NewAntifraudHandler() *AntifraudHandler {
	return &AntifraudHandler{
		RuleManager: rules.NewRuleManager(),
		validator:   validator.New(),
	}
}

func (h *AntifraudHandler) AnalyzeTransaction(req *restful.Request, resp *restful.Response) {
	var bodyReq request.Analyze

	if err := req.ReadEntity(&bodyReq); err != nil {
		resp.WriteHeaderAndEntity(http.StatusBadRequest, map[string]string{
			"message": "invalid.json",
		})
		return
	}

	if err := h.validator.Struct(bodyReq); err != nil {
		var errors []string
		for _, e := range err.(validator.ValidationErrors) {
			errors = append(errors, fmt.Sprintf("field '%s' failed to validate '%s'", e.Field(), e.Tag()))
		}
		resp.WriteHeaderAndEntity(http.StatusBadRequest, map[string]interface{}{
			"validation_errors": errors,
		})
		return
	}

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
	_, _ = resp.Write([]byte("OK - HealthCheck"))
}

func (h *AntifraudHandler) GetStats(req *restful.Request, resp *restful.Response) {
	_, _ = resp.Write([]byte("OK - GetStats"))
}
