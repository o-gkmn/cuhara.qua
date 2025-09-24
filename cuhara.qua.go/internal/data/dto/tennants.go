package dto

type TennantDTO struct {
	ID   int64
	Name string
}

type CreateTennantRequest struct {
	Name string
}

type CreateTennantResponse struct {
	ID int64
}

type UpdateTennantRequest struct {
	ID int64
	Name string
}

type UpdateTennantResponse struct {
	ID int64
}

type DeleteTennantRequest struct {
	ID int64
}

type DeleteTennantResponse struct {
	ID int64
}