package dto

type ClaimDTO struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CreateClaimRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type CreateClaimResponse struct {
	ID int64 `json:"id"`
}

type UpdateClaimRequest struct {
	ID          int64   `json:"id"`
	Name        *string `json:"name"`
	Description *string `json:"description"`
}

type UpdateClaimResponse struct {
	ID int64 `json:"id"`
}

type DeleteClaimRequest struct {
	ID int64 `json:"id"`
}

type DeleteClaimResponse struct {
	ID int64 `json:"id"`
}
