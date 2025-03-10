package dtos

type ResourceDTO struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Type  string `json:"type"`
	URL   string `json:"url"`
}

type ResourceInsertDTO struct {
	Title string `json:"title" binding:"required"`
	Type  string `json:"type" binding:"required,oneof=PDF SLIDES LINK CODE OTHER"`
	URL   string `json:"url" binding:"required,url"`
}
