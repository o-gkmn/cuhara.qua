package dto

type ClaimDTO struct {
	ID          int64
	Name        string
	Description string
}

type CreateClaimRequest struct {
	Name        string
	Description *string
}

type CreateClaimResponse struct {
	ID int64
}

type UpdateClaimRequest struct {
	ID          int64
	Name        *string
	Description *string
}

type UpdateClaimResponse struct {
	ID int64
}

type DeleteClaimRequest struct {
	ID int64
}

type DeleteClaimResponse struct {
	ID int64
}
