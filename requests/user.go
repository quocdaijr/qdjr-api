package requests

type SearchUserRequest struct {
	KeyWord string `form:"keyword" json:"keyword"`
	Page    int    `form:"page,default=1" json:"page"`
	PerPage int    `form:"per_page,default=1" json:"per_page"`
	Sort    string `form:"sort" json:"sort"`
}
