package dtos

type ModuleDTO struct {
	ID      string      `json:"id"`
	Title   string      `json:"title"`
	Order   int         `json:"order"`
	Lessons []LessonDTO `json:"lessons"`
}

type ModuleInsertDTO struct {
	Title   string            `json:"title" binding:"required"`
	Order   int               `json:"order" binding:"required,min=0"`
	Lessons []LessonInsertDTO `json:"lessons" binding:"dive"`
}
