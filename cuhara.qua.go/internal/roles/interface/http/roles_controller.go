package http

import (
	"encoding/json"
	"net/http"

	"cuhara.qua.go/internal/common/cqrs"
	rolecmd "cuhara.qua.go/internal/roles/application/command"
)

type RoleController struct {
	commandBus *cqrs.CommandBus
}

func NewRoleController(cb *cqrs.CommandBus) *RoleController {
	return &RoleController{commandBus: cb}
}

// CreateRole godoc
// @Summary Create a new role
// @Tags roles
// @Accept json
// @Produce json
// @Param role body rolecmd.CreateRoleRequest true " "
// @Success 201 {object} rolecmd.CreateRoleResponse
// @Router /roles [post]
func (rc *RoleController) CreateRole(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	var req rolecmd.CreateRoleRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid json body", http.StatusBadRequest)
	}

	cmd := rolecmd.NewCreateRoleCommand(req)
	res, err := cqrs.Execute[rolecmd.CreateRoleResponse](rc.commandBus, r.Context(), cmd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(res)
}
