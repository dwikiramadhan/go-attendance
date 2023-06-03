package utils

import "attendance/pkg/entity"

// GetError get Error entity
func GetError(code string, source string, custDetail string) (error *entity.Error) {
	var err entity.Error

	switch code {
	case "201001":
		err = entity.Error{ErrCode: code, Source: "", Status: "Error", Title: "Incorrect Format", Detail: "Request format does not meet the request criteria "}
	case "10001":
		err = entity.Error{ErrCode: "10001", Source: "", Status: "Error", Title: "Incorrect Format", Detail: "Request format does not meet the request criteria "}
	case "10002":
		err = entity.Error{ErrCode: "10002", Source: "", Status: "Error", Title: "Datas", Detail: "Record Not Found"}
	case "10003":
		err = entity.Error{ErrCode: "10003", Source: "", Status: "Error", Title: "Error Validation", Detail: "Email already exist"}
	case "10004":
		err = entity.Error{ErrCode: "10004", Source: "", Status: "Error", Title: "Error Query Insert", Detail: ""}
	case "10005":
		err = entity.Error{ErrCode: "10005", Source: "", Status: "Error", Title: "Error Query Update", Detail: ""}
	case "10006":
		err = entity.Error{ErrCode: "10006", Source: "", Status: "Error", Title: "Error Query Delete", Detail: ""}
	case "10007":
		err = entity.Error{ErrCode: "10006", Source: "", Status: "Error", Title: "Authentication Problem", Detail: ""}
	}

	e := err
	e.Source = source

	if len(custDetail) > 0 {
		e.Detail = custDetail
	}

	return &e
}
