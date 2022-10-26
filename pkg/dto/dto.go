package dto

type CreatePostDTO struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description"`
	SectionID   uint   `json:"section_id" binding:"required"`
}

type EditPostDTO struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Markdown    string `json:"markdown"`
	HTML        string `json:"html"`
	Publish     bool   `json:"publish"`
}

type CreateCourseDTO struct {
	Title     string `json:"title" binding:"required"`
	DescShort string `json:"desc_short"`
	Cost      int    `json:"cost"`
}

type EditCourseDTO struct {
	Title      string     `json:"title"`
	Markdown   string     `json:"markdown"`
	HTML       string     `json:"html"`
	DescShort  string     `json:"desc_short"`
	Cost       int        `json:"cost"`
	Publish    bool       `json:"publish"`
	Categories []Category `json:"categories"`
}

type Category struct {
	Label string `json:"label"`
	Value string `json:"value"`
	IsNew bool   `json:"__isNew__"`
}

type CategoryDTO struct {
	Name string `json:"name"`
}

type CreateSectionDTO struct {
	CourseID    int    `json:"course_id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"desc"`
}

type EditSectionDTO struct {
	Title        string `json:"title"`
	Description  string `json:"desc"`
	ExpectedTime int    `json:"expected_time"`
	Publish      bool   `json:"publish"`
}

type ChangeOrderSectionDTO struct {
	SectionID uint `json:"section_id" binding:"required"`
	Order     *int `json:"order" binding:"required"` // since order can be  zero we need to have a pointer
}

type ChangeOrderPostDTO struct {
	PostID uint `json:"post_id" binding:"required"`
	Order  *int `json:"order" binding:"required"` //order=0 will result in required validation failed, no desirable thus pointer
}
