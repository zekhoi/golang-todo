package dto

type ActivityCreateRequest struct {
	Title string `json:"title" binding:"required" validate:"required"`
	Email string `json:"email" binding:"required" validate:"required"`
}

type ActivityUpdateRequest struct {
	Title string `json:"title" binding:"required" validate:"required"`
}
