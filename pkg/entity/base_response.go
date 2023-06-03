package entity

type BaseResponse struct {
	StatusCode int         `json:"status_code"`
	Message    interface{} `json:"msg"`
	MessageId  string      `json:"msg_id"`
	Data       interface{} `json:"data"`
	Errors     interface{} `json:"errors"`
}

// PagedResults results for pages GetAll results.
type PagedResults struct {
	Page         int64       `json:"page"`
	PageSize     int64       `json:"page_size"`
	Data         interface{} `json:"data"`
	TotalRecords int         `json:"total_records"`
}

type StatusResponse struct {
	Status string `json:"status"`
}
