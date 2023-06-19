package employee

import (
	"attendance/app/dao"
	"attendance/pkg/entity"
	"attendance/pkg/utils"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// CreateEmployee func for create a Employee.
// @Description Create a new Employee.
// @Summary create a new Employee
// @Tags Employee
// @Accept json
// @Produce json
// @Param Emloyee body entity.EmployeeReq true "create Employee"
// @Success 201 {array} entity.BaseResponse
// @Failure 400 {array} entity.BaseResponse
// @Failure 404 {array} entity.BaseResponse
// @Security ApiKeyAuth
// @Router /v1/employee [post]
func CreateEmployee(c *fiber.Ctx) error {

	functionName := "CreateEmployee"

	// Create new AdminUser struct
	empReq := &entity.EmployeeReq{}

	validate := utils.NewValidator()

	// Check, if received JSON data is valid.
	if err := c.BodyParser(empReq); err != nil {
		// Return status 400 and error message.
		baseResponse := &entity.BaseResponse{
			StatusCode: c.Response().StatusCode(),
			Message:    "Failed Body Parser",
			Errors:     utils.GetError("10001", functionName, err.Error()),
		}
		return c.JSON(baseResponse)
	}

	// Check, validation.
	if err := validate.Struct(empReq); err != nil {
		baseResponse := &entity.BaseResponse{
			StatusCode: c.Response().StatusCode(),
			Message:    "Failed in Validation",
			Errors:     utils.ValidatorErrors(err),
		}
		return c.JSON(baseResponse)
	}

	// Check email is exist
	_, rows, _ := dao.GetAdminUserByEmail(empReq.AdminEmail)
	if rows > 0 {
		// Return status 500 and error message.
		baseResponse := &entity.BaseResponse{
			StatusCode: c.Response().StatusCode(),
			Message:    "Email Already Exists",
			Errors:     utils.GetError("10003", functionName, ""),
		}
		return c.Status(fiber.StatusBadRequest).JSON(baseResponse)
	}

	pass := "12345"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	if err != nil {
		baseResponse := &entity.BaseResponse{
			StatusCode: fiber.StatusBadRequest,
			Message:    utils.ValidatorErrors(err),
			Errors:     utils.GetError("10003", functionName, err.Error()),
		}
		return c.Status(fiber.StatusBadRequest).JSON(baseResponse)
	}

	empNoGen := utils.Generate6Digit()
	empReq.AdminPassword = string(hashedPassword)
	empReq.EmpNo = &empNoGen
	// status := adminUser.UcUsrStatus
	// adminUser.UcUsrStatus = strings

	// Create AdminUser by given model.
	_, _, err = dao.CreateEmp(empReq)
	if err != nil {
		// Return status 500 and error message.
		baseResponse := &entity.BaseResponse{
			StatusCode: fiber.StatusBadRequest,
			Message:    utils.ValidatorErrors(err),
			Errors:     utils.GetError("10001", functionName, err.Error()),
		}
		return c.Status(fiber.StatusBadRequest).JSON(baseResponse)
	}

	// Return status 201 CREATED.
	baseResponse := &entity.BaseResponse{
		StatusCode: fiber.StatusCreated,
		Message:    "Success create Employee",
		Errors:     false,
	}
	return c.JSON(baseResponse)
}
