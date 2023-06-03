package entity

// type BaseResponse struct {
// 	Data DataStructure `json:"data"`
// }

type BaseResponseError struct {
	StatusCode int         `json:"status_code"`
	Message    string      `json:"msg"`
	MessageId  string      `json:"msg_id"`
	Errors     interface{} `json:"errors"`
}

//SetError error object setter helper
func SetError(errCode string, status string, source string, title string, detail string) Error {

	var e Error

	e.ErrCode = errCode
	e.Status = status
	e.Source = source
	e.Title = title
	e.Detail = detail

	return e
}
