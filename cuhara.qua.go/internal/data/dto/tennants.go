package dto

type TenantDTO struct {
	ID   int64
	Name string
}

type CreateTenantRequest struct {
	Name string
}

type CreateTenantResponse struct {
	ID int64
}

type UpdateTenantRequest struct {
	ID   int64
	Name *string
}

type UpdateTenantResponse struct {
	ID int64
}

type DeleteTenantRequest struct {
	ID int64
}

type DeleteTenantResponse struct {
	ID int64
}
