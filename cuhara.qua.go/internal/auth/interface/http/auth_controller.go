package http

import (
	"encoding/json"
	"net/http"

	authqry "cuhara.qua.go/internal/auth/application/query"
	"cuhara.qua.go/internal/common/cqrs"
)

type AuthController struct {
	queryBus *cqrs.QueryBus
}

func NewAuthController(qb *cqrs.QueryBus) *AuthController {
	return &AuthController{queryBus: qb}
}

// Login godoc
// @Summary Login user with email
// @Tags auth
// @Accept json
// @Produce json
// @Param auth body authqry.LoginRequest true " "
// @Success 200 {object} authqry.LoginResponse
// @Router /auth/login [post]
func (ac *AuthController) Login(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var req authqry.LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json body", http.StatusBadRequest)
	}

	qry := authqry.NewLoginQuery(req)
	res, err := cqrs.ExecuteQuery[authqry.LoginResponse](ac.queryBus, r.Context(), qry)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(res)
}
