package requests

type ProjectsQuery struct {
	Page  int  `form:"page"`
	Limit int  `form:"limit"`
	Cache bool `form:"cache"`
}
