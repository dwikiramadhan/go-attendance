package entity

import "time"

type AttendanceReq struct {
	EmpNo     string `json:"emp_no" validate:"required" example:"00015"`
	Latitude  string `json:"latitude" validate:"required" example:"-6.242304"`
	Longitude string `json:"longitude" validate:"required" example:"106.8957696"`
	Selfie    string `json:"selfie"`
}

type AttendanceClassified struct {
	ComeIn             *time.Time `json:"come_in"`
	ComeOut            *time.Time `json:"come_out"`
	AdmissionTimeLimit string     `json:"admission_time_limit"`
	LateCount          int32      `json:"late_count"`
	Selfie             []byte     `json:"selfie"`
}

type AttendanceList struct {
	ID                 int64      `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	EmpNo              string     `gorm:"column:emp_no;not null" json:"emp_no"`
	CheckIn            *time.Time `gorm:"column:check_in" json:"check_in"`
	CreatedAt          *time.Time `gorm:"column:created_at" json:"created_at"`
	Description        *string    `gorm:"column:description" json:"description"`
	Status             *int32     `gorm:"column:status" json:"status"`
	LateCount          *int32     `gorm:"column:late_count" json:"late_count"`
	AdmissionTimeLimit *string    `gorm:"column:admission_time_limit" json:"admission_time_limit"`
	TimeOutLimit       *string    `gorm:"column:time_out_limit" json:"time_out_limit"`
	Latitude           *string    `gorm:"column:latitude" json:"latitude"`
	Longitude          *string    `gorm:"column:longitude" json:"longitude"`
	Selfie             []byte     `gorm:"column:selfie" json:"selfie"`
}
