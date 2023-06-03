package dao

import model "attendance/app/models"

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
