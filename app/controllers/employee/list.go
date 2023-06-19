package employee

import (
	"attendance/app/dao"
	"attendance/pkg/entity"
	"attendance/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

// Employee func for Employee.
// @Description for Get Employees
// @Summary Employee
// @Tags Employee
// @Accept json
// @Produce json
// @Param       page     query    int    false "page requested (defaults to 0)"
// @Param       pagesize query    int    false "number of records in a page  (defaults to 10)"
// @Param       order    query    string false "sort order by column. ex: check_in asc. field name with space then asc/desc"
// @Success 200 {array} entity.BaseResponse
// @Failure 400 {array} entity.BaseResponse
// @Failure 404 {array} entity.BaseResponse
// @Router /v1/employees [get]
func GetEmployees(c *fiber.Ctx) (err error) {
	functionName := "GetEmployees"

	page, err := utils.ReadInt(c.Request(), "page", 0)
	if err != nil || page < 0 {
		baseResponse := &entity.BaseResponse{
			Message: "Failed",
			Data:    page,
			Errors:  utils.GetError("10001", functionName, err.Error()),
		}
		return c.JSON(baseResponse)
	}

	pagesize, err := utils.ReadInt(c.Request(), "pagesize", 10)
	if err != nil || pagesize <= 0 {
		baseResponse := &entity.BaseResponse{
			Message: "Failed",
			Data:    page,
			Errors:  utils.GetError("10001", functionName, err.Error()),
		}
		return c.JSON(baseResponse)
	}

	order := c.Query("order")

	// Checking, if book with given ID is exists.
	data, totalRows, err := dao.GetEmployees(int(page), int(pagesize), order)
	if err != nil {
		// Return status 404 and book not found error.
		baseResponse := &entity.BaseResponse{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Failed GetAttendances",
			Errors:     utils.GetError("10001", functionName, err.Error()),
		}
		return c.JSON(baseResponse)
	}

	results := entity.PagedResults{
		Page:         page,
		PageSize:     pagesize,
		Data:         data,
		TotalRecords: int(totalRows),
	}

	// Return status 200 OK.
	baseResponse := &entity.BaseResponse{
		StatusCode: fiber.StatusOK,
		Message:    "Success GetEmployees",
		Data:       results,
		Errors:     false,
	}
	return c.JSON(baseResponse)

}
