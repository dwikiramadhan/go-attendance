package authentication

import (
	"attendance/app/dao"
	"attendance/pkg/entity"
	"attendance/pkg/utils"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

// Authentication func for Login.
// @Description for Login
// @Summary Login
// @Tags Authentication
// @Accept json
// @Produce json
// @Param Authentication body entity.LoginReq true "form login"
// @Success 200 {array} entity.BaseResponse
// @Failure 400 {array} entity.BaseResponse
// @Failure 404 {array} entity.BaseResponse
// @Router /v1/login [post]
func Login(c *fiber.Ctx) (err error) {
	functionName := "Login"

	formLogin := &entity.LoginReq{}

	validate := utils.NewValidator()

	// Check, if received JSON data is valid.
	if err := c.BodyParser(formLogin); err != nil {
		// Return status 400 and error message.
		baseResponse := &entity.BaseResponse{
			Message: "Form JSON data validation checking",
			Errors:  utils.GetError("10001", functionName, err.Error()),
		}
		return c.JSON(baseResponse)
	}

	// Check, validation.
	if err := validate.Struct(formLogin); err != nil {
		baseResponse := &entity.BaseResponse{
			StatusCode: c.Response().StatusCode(),
			Message:    "Failed in Validation",
			Errors:     utils.ValidatorErrors(err),
		}
		return c.JSON(baseResponse)
	}

	// Check email is exist
	foundedUser, _, err := dao.GetAdminUserByEmail(formLogin.Email)
	if err != nil {
		// Return status 500 and error message.
		baseResponse := &entity.BaseResponse{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Email not found",
			Errors:     utils.GetError("10002", functionName, ""),
		}
		return c.Status(fiber.StatusBadRequest).JSON(baseResponse)
	}

	fmt.Println("foundedUser.AdminPassword: ", foundedUser.AdminPassword)
	fmt.Println("formLogin.Password: ", formLogin.Password)
	err = bcrypt.CompareHashAndPassword([]byte(foundedUser.AdminPassword), []byte(formLogin.Password))
	if err != nil {
		baseResponse := &entity.BaseResponse{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Email/password is invalid",
			Errors:     utils.GetError("10007", functionName, err.Error()),
		}
		return c.Status(fiber.StatusBadRequest).JSON(baseResponse)
	}

	// Generate a new pair of access and refresh tokens.
	tokens, err := utils.GenerateNewTokens(fmt.Sprint(foundedUser.AdminID), *foundedUser.EmpNo)
	if err != nil {
		// Return status 500 and token generation error.
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	// Return status 200 OK.
	baseResponse := &entity.BaseResponse{
		StatusCode: fiber.StatusOK,
		Message:    "Login success",
		Data:       tokens,
		Errors:     false,
	}
	return c.JSON(baseResponse)
}
