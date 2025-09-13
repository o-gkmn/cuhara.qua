package http

import (
	"encoding/json"
	"net/http"

	"cuhara.qua.go/internal/common/cqrs"
	tennantcmd "cuhara.qua.go/internal/tennants/application/command"
)

type TennantsController struct {
	CommandBus *cqrs.CommandBus
}

func NewTennantsController(cb *cqrs.CommandBus) *TennantsController {
	return &TennantsController{CommandBus: cb}
}

// CreateTennant godoc
// @Summary Create a new Tennant
// @Tags Tennant
// @Accept json
// @Produce json
// @Param tennant body tennantcmd.CreateTennantRequest true " "
// @Success 201 {object} tennantcmd.CreateTennantResponse
// @Router /tennants [post]
func (tc *TennantsController) Create(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var req tennantcmd.CreateTennantRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json body", http.StatusBadRequest)
	}

	cmd := tennantcmd.NewCreateTennantCommand(req)
	res, err := cqrs.Execute[tennantcmd.CreateTennantResponse](tc.CommandBus, r.Context(), cmd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(res)
}
