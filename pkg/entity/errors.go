package entity

type ErrorList struct {
	Errors []*Error `json:"errors"`
}

type Error struct {
	ErrCode string `json:"errorcode"`
	Status  string `json:"status"`
	Source  string `json:"source"`
	Title   string `json:"title"`
	Detail  string `json:"detail"`
}
