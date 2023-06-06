package dao

import (
	model "attendance/app/models"
	"attendance/pkg/entity"
)

// CreateAttendance method for creating book by given Book object.
func CreateAttendance(record *model.Attendance) (result *model.Attendance, RowAffected int64, err error) {
	db := DB.Debug().Save(record)

	if err = db.Error; err != nil {
		// Return only error.
		return nil, -1, ErrInsertFailed
	}

	// This query returns nothing.
	return record, db.RowsAffected, nil
}

// UpdateAttendance method for creating book by given Book object.
func UpdateAttendance(id int64, updated *model.Attendance) (result *model.Attendance, RowsAffected int64, err error) {
	result = &model.Attendance{}

	db := DB.First(result, model.Attendance{ID: id})
	if err = db.Error; err != nil {
		return nil, -1, ErrNotFound
	}

	if err = Copy(result, updated); err != nil {
		return nil, -1, ErrUpdateFailed
	}

	db = db.Save(result)
	if err = db.Error; err != nil {
		return nil, -1, ErrUpdateFailed
	}

	return result, db.RowsAffected, nil
}

// GetAdminUserByEmpNo method for getting admin user by Email.
func GetAttendanceTodayByEmpNo(emp_no string) (result []entity.AttendanceList, totalRows int64, err error) {
	resultOrm := DB.Model(&model.Attendance{})
	resultOrm = resultOrm.Where(&model.Attendance{EmpNo: emp_no})
	resultOrm = resultOrm.Where("DATE(check_in) = CURDATE()")
	resultOrm = resultOrm.Order("check_in asc")
	resultOrm.Count(&totalRows)
	if err := resultOrm.Debug().Find(&result).Error; err != nil {
		// Return empty object and error.
		err = ErrNotFound
		return nil, -1, err
	}
	// Return query result.
	return result, totalRows, nil
}

// GetAttendanceByEmpNo method for getting admin user by Email.
func GetAttendanceByEmpNo(page, pagesize int, order, emp_no string) (result []entity.AttendanceList, totalRows int64, err error) {
	resultOrm := DB.Model(&model.Attendance{})

	if order != "" {
		resultOrm = resultOrm.Order(order)
	}

	resultOrm = resultOrm.Where(&model.Attendance{EmpNo: emp_no})
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

// GetAttendances method for getting admin user by Email.
func GetAttendances(page, pagesize int, order string) (result []entity.AttendanceList, totalRows int64, err error) {
	resultOrm := DB.Model(&model.Attendance{})

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
