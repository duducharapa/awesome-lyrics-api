package dtos

type CreateLyricDTO struct {
	Name   string `json:"name" binding:"required"`
	Author string `json:"author"`
}
