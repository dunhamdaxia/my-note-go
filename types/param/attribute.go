package param

type AttributeCreate struct {
	ID       int64  `json:"id"`
	ParentID int64  `json:"parent_id"`
	UserID   int64  `json:"user_id"`
	Name     string `json:"name"`
	Sort     int    `json:"sort"`
}

type AttributesPage struct {
	BasePage
	ParentId int `json:"parent_id"`
}

type AttributeInfo struct {
	ID int64 `json:"id"`
}
