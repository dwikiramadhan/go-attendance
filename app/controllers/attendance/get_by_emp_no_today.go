package attendance

import (
	"attendance/app/dao"
	"attendance/pkg/entity"
	"attendance/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

// Attendance func for Attendance.
// @Description for Get Attendance Get Today by Emp No
// @Summary Attendance
// @Tags Attendance
// @Accept json
// @Produce json
// @Param emp_no path string true "emp_no"
// @Success 200 {array} entity.BaseResponse
// @Failure 400 {array} entity.BaseResponse
// @Failure 404 {array} entity.BaseResponse
// @Router /v1/attendance/get-today/{emp_no} [get]
func GetAttendanceByEmpNoToday(c *fiber.Ctx) (err error) {
	functionName := "GetAttendanceByEmpNo"

	emp_no := c.Params("emp_no")

	// Checking, if book with given ID is exists.
	foundedUser, _, err := dao.GetAttendanceTodayByEmpNo(emp_no)
	if err != nil {
		// Return status 404 and book not found error.
		baseResponse := &entity.BaseResponse{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Failed GetAttendanceByEmpNo",
			Errors:     utils.GetError("10001", functionName, err.Error()),
		}
		return c.JSON(baseResponse)
	}

	var attClassified entity.AttendanceClassified
	for i := 0; i < len(foundedUser); i++ {
		if i == 0 {
			attClassified.ComeIn = foundedUser[i].CheckIn
			attClassified.AdmissionTimeLimit = *foundedUser[i].AdmissionTimeLimit
			attClassified.LateCount = *foundedUser[i].LateCount
			attClassified.Selfie = foundedUser[i].Selfie
		}

		if i > 0 && i == len(foundedUser)-1 {
			attClassified.ComeOut = foundedUser[i].CheckIn
		}
	}

	// Return status 200 OK.
	baseResponse := &entity.BaseResponse{
		StatusCode: fiber.StatusOK,
		Message:    "Success GetAttendanceTodayByEmpNo",
		Data:       attClassified,
		Errors:     false,
	}
	return c.JSON(baseResponse)

}
