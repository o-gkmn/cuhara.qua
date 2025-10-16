package dto

type TenantDTO struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type CreateTenantRequest struct {
	Name string `json:"name"`
}

type CreateTenantResponse struct {
	ID int64 `json:"id"`
}

type UpdateTenantRequest struct {
	ID   int64   `json:"id"`
	Name *string `json:"name"`
}

type UpdateTenantResponse struct {
	ID int64 `json:"id"`
}

type DeleteTenantRequest struct {
	ID int64 `json:"id"`
}

type DeleteTenantResponse struct {
	ID int64 `json:"id"`
}
