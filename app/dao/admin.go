package dao

import (
	model "attendance/app/models"
	"attendance/pkg/entity"
)

// GetAdminUserByEmail method for getting admin user by Email.
func GetAdminUserByEmail(email string) (result *model.Admin, rows int64, err error) {
	resultOrm := DB.Model(&model.Admin{})
	resultOrm = resultOrm.Where(&model.Admin{AdminEmail: email})
	resultOrm.Count(&rows)
	if err := resultOrm.Debug().First(&result).Error; err != nil {
		// Return empty object and error.
		err = ErrNotFound
		return nil, -1, err
	}
	// Return query result.
	return result, rows, nil
}

// GetAdminUserByEmpNo method for getting admin user by Email.
func GetAdminUserByEmpNo(emp_no string) (result *model.Admin, err error) {
	resultOrm := DB.Model(&model.Admin{})
	resultOrm = resultOrm.Where(&model.Admin{EmpNo: &emp_no})
	if err := resultOrm.Debug().First(&result).Error; err != nil {
		// Return empty object and error.
		err = ErrNotFound
		return nil, err
	}
	// Return query result.
	return result, nil
}

// GetEmployees method for getting employees.
func GetEmployees(page, pagesize int, order string) (result []model.Admin, totalRows int64, err error) {
	resultOrm := DB.Model(&model.Admin{})

	if order != "" {
		resultOrm = resultOrm.Order(order)
	}

	resultOrm.Count(&totalRows)

	if page > 0 {
		offset := (page - 1) * pagesize
		resultOrm = resultOrm.Offset(offset).Limit(pagesize)
	} else {
		resultOrm = resultOrm.Limit(pagesize)
	}

	if err := resultOrm.Debug().Find(&result).Error; err != nil {
		// Return empty object and error.
		err = ErrNotFound
		return nil, -1, err
	}
	// Return query result.
	return result, totalRows, nil
}

// CreateAdminUser method for creating AdminUser by given AdminUser object.
func CreateEmp(record *entity.EmployeeReq) (result *entity.EmployeeReq, RowAffected int64, err error) {

	// record.UcAdmUserUID = uuid.NewString()

	db := DB.Debug().Create(record)

	if err != nil {
		// Return only error.
		return nil, -1, ErrInsertFailed
	}

	// This query returns nothing.
	return record, db.RowsAffected, nil

}
