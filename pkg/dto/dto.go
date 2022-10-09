package dto

type CreatePostDTO struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	ImageURL    string
	IsBlogPost  bool `json:"is_blog_post"`
}

type EditPostDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	ImageURL    string
	Body        string `json:"editorjs_body"`
	IsBlogPost  bool   `json:"is_blog_post"`
}

type CreateCourseDTO struct {
	Title    string `json:"title"`
	DescBody string `json:"desc_body"`
	Cost     int    `json:"cost"`
}

type EditCourseDTO struct {
	Title    string `json:"title"`
	DescBody string `json:"desc_body"`
	Cost     int    `json:"cost"`
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

type CreateSubsectionDTO struct {
	SectionID   uint   `json:"section_id"`
	Title       string `json:"title"`
	Description string `json:"desc"`
}

type EditSubsectionDTO struct {
	SectionID   uint   `json:"section_id"`
	Title       string `json:"title"`
	Description string `json:"desc"`
	PostID      uint   `json:"post_id"`
}

type ChangeOrderSectionDTO struct {
	SectionID uint `json:"section_id"`
	Order     int  `json:"order"`
}

type ChangeOrderSubsectionDTO struct {
	SubsectionID uint `json:"subsection_id"`
	Order        int  `json:"order"`
}

type AddPostSubsectionDTO struct {
	CreatePostDTO
	SubsectionID uint `json:"subsection_id"`
}
