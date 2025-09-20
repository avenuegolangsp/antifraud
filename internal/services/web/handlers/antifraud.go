package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/emicklei/go-restful/v3"
)

type AntifraudHandler struct {
	Repo TransactionRepository
}

func (h *AntifraudHandler) AnalyzeTransaction(req *restful.Request, resp *restful.Response) {
	var tx Transaction
	if err := json.NewDecoder(req.Request.Body).Decode(&tx); err != nil {
		resp.WriteErrorString(http.StatusBadRequest, "invalid payload")
		return
	}
	// Salva a transação normalmente
	err := h.Repo.Insert(context.Background(), &tx)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, "db error: "+err.Error())
		return
	}

	// Busca dados do cliente e suas transações
	var clienteDados struct {
		User       *User
		Transacoes []Transaction
	}
	clienteDados.User, _ = h.Repo.GetUserByID(context.Background(), tx.UserID)
	clienteDados.Transacoes, _ = h.Repo.GetTransactionsByUserID(context.Background(), tx.UserID)
	// Variável clienteDados pode ser usada em outras camadas

	resp.WriteHeaderAndEntity(http.StatusCreated, map[string]string{"status": "created"})
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

// ListClientsWithTransactions handles GET /clients requests
func (h *AntifraudHandler) ListClientsWithTransactions(req *restful.Request, resp *restful.Response) {
	ctx := context.Background()
	users, err := h.Repo.GetAllUsers(ctx)
	if err != nil {
		resp.WriteErrorString(http.StatusInternalServerError, "erro ao buscar usuários: "+err.Error())
		return
	}

	type UserWithTransactions struct {
		*User
		Transactions []Transaction `json:"transactions"`
	}
	var result []UserWithTransactions
	for _, user := range users {
		txs, err := h.Repo.GetTransactionsByUserID(ctx, user.ID)
		if err != nil {
			resp.WriteErrorString(http.StatusInternalServerError, "erro ao buscar transações do usuário: "+user.ID+": "+err.Error())
			return
		}
		result = append(result, UserWithTransactions{
			User:         user,
			Transactions: txs,
		})
	}
	resp.WriteHeaderAndEntity(http.StatusOK, result)
}
