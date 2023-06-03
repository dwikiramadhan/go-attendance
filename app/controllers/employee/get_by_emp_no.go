package employee

import (
	"attendance/app/dao"
	"attendance/pkg/entity"
	"attendance/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

// GetEmpByEmpNo func gets Employee by emp_no.
// @Description Get Employee by emp_no.
// @Summary get Employee by emp_no.
// @Tags Employee
// @Accept json
// @Produce json
// @Param emp_no path string true "emp_no"
// @Success 200 {array} entity.BaseResponse
// @Failure 400 {array} entity.BaseResponse
// @Failure 404 {array} entity.BaseResponse
// @Router /v1/employee/{emp_no} [get]
func GetEmpByEmpNo(c *fiber.Ctx) (err error) {
	functionName := "GetEmpByEmpNo"

	emp_no := c.Params("emp_no")

	// Checking, if book with given ID is exists.
	foundedUser, err := dao.GetAdminUserByEmpNo(emp_no)
	if err != nil {
		// Return status 404 and book not found error.
		baseResponse := &entity.BaseResponse{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Failed get Emp by Emp No",
			Errors:     utils.GetError("10001", functionName, err.Error()),
		}
		return c.JSON(baseResponse)
	}

	// Return status 200 OK.
	baseResponse := &entity.BaseResponse{
		StatusCode: fiber.StatusOK,
		Message:    "Success get data Emp by Emp No",
		Data:       foundedUser,
		Errors:     false,
	}
	return c.JSON(baseResponse)
}
