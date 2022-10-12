package dto

type CreatePostDTO struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	ImageURL    string
	SectionID   uint `json:"section_id"`
}

type EditPostDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageURL    string
	Body        string `json:"editorjs_body"`
}

type CreateCourseDTO struct {
	Title     string `json:"title"`
	DescShort string `json:"desc_short"`
	Cost      int    `json:"cost"`
}

type EditCourseDTO struct {
	Title     string `json:"title"`
	DescBody  string `json:"desc_body"`
	DescShort string `json:"desc_short"`
	Cost      int    `json:"cost"`
}

type CreateSectionDTO struct {
	CourseID    int    `json:"course_id"`
	Title       string `json:"title"`
	Description string `json:"desc"`
}

type EditSectionDTO struct {
	Title       string `json:"title"`
	Description string `json:"desc"`
}

type ChangeOrderSectionDTO struct {
	SectionID uint `json:"section_id"`
	Order     int  `json:"order"`
}
