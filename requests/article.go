package requests

type SearchArticleRequest struct {
	KeyWord string `form:"keyword" json:"keyword"`
	Page    int    `form:"page,default=1" json:"page"`
	PerPage int    `form:"per_page,default=1" json:"per_page"`
	Sort    string `form:"sort" json:"sort"`
}

type CreateArticleRequest struct {
	Title       string `json:"title" binding:"required"`
	Slug        string `json:"slug" binding:"required"`
	Description string `json:"description" binding:"omitempty,max=255"`
	Content     string `json:"content" binding:"omitempty"`
	Status      uint8  `json:"status" binding:"required"`
	Author      string `json:"author" binding:"required,max=255"`
	PublishedAt string `json:"published_at" binding:"required,datetime=2006-01-02 15:04:05"`
}

type UpdateArticleRequest struct {
	Title       string `json:"title" binding:"omitempty"`
	Slug        string `json:"slug" binding:"omitempty"`
	Description string `json:"description" binding:"omitempty,max=255"`
	Content     string `json:"content" binding:"omitempty"`
	Status      uint8  `json:"status" binding:"omitempty"`
	Author      string `json:"author" binding:"omitempty,max=255"`
	PublishedAt string `json:"published_at" binding:"omitempty,datetime=2006-01-02 15:04:05"`
}
