package param

type TaskPage struct {
	BasePage
	Status int
}

type TaskCreate struct {
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	PeriodType int8   `json:"period_type"`
	Type       int8   `json:"type"`
}
