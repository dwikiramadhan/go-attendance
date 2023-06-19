package entity

const TableNameAdmin = "admin"

type EmployeeReq struct {
	AdminName     string  `gorm:"column:admin_name;not null" json:"admin_name"`
	AdminEmail    string  `gorm:"column:admin_email;not null" json:"admin_email"`
	AdminPassword string  `gorm:"column:admin_password;not null" json:"admin_password"`
	EmpNo         *string `gorm:"column:emp_no" json:"emp_no"`
	Role          *string `gorm:"column:role" json:"role"`
}

// TableName admin's table name
func (*EmployeeReq) TableName() string {
	return TableNameAdmin
}
