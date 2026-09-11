package requests

type DailyVisitsQuery struct {
	StartDate string `form:"start_date" validate:"required,datetime=2006-01-02"`
	EndDate   string `form:"end_date" validate:"required,datetime=2006-01-02"`
	Page      int    `form:"page"`
	Limit     int    `form:"limit"`
}
