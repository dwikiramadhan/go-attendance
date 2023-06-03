package attendance

import (
	"attendance/app/dao"
	model "attendance/app/models"
	"attendance/pkg/entity"
	"attendance/pkg/utils"
	"encoding/base64"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

// Attendance func for Attendance.
// @Description for Create Attendance
// @Summary Attendance
// @Tags Attendance
// @Accept json
// @Produce json
// @Param Attendance body entity.AttendanceReq true "form add attendance"
// @Success 200 {array} entity.BaseResponse
// @Failure 400 {array} entity.BaseResponse
// @Failure 404 {array} entity.BaseResponse
// @Router /v1/attendance [post]
func CreateAttendance(c *fiber.Ctx) (err error) {
	functionName := "CreateAttendance"

	// Create new AttendanceReq struct
	input := &entity.AttendanceReq{}

	validate := utils.NewValidator()

	// Check, if received JSON data is valid.
	if err := c.BodyParser(input); err != nil {
		fmt.Println("errrrrrr: ", err)
		// Return status 400 and error message.
		baseResponse := &entity.BaseResponse{
			StatusCode: c.Response().StatusCode(),
			Message:    "Failed Body Parser",
			Errors:     utils.GetError("10001", functionName, err.Error()),
		}
		return c.JSON(baseResponse)
	}

	// Check, validation.
	if err := validate.Struct(input); err != nil {
		baseResponse := &entity.BaseResponse{
			StatusCode: c.Response().StatusCode(),
			Message:    "Failed in Validation",
			Errors:     utils.ValidatorErrors(err),
		}
		return c.JSON(baseResponse)
	}

	// Decode the base64 string
	data, err := base64.StdEncoding.DecodeString(input.Selfie)
	if err != nil {
		log.Fatal(err)
	}

	now := time.Now()
	var late_count int32
	var status int32
	var addmission_time_limit string = "08:00"
	var time_out_limit string = "17:00"

	hour := 8
	minute := 0
	second := 0
	// Create a specific time
	setTime := time.Date(now.Year(), now.Month(), now.Day(), hour, minute, second, 0, now.Location())

	// Compare the specific time with the current time
	if setTime.After(now) {
		late_count = 0
		fmt.Println("setTime is bigger than the current time.")
	} else if setTime.Before(now) {
		late_count = 1
		fmt.Println("setTime is smaller than the current time.")
	} else {
		late_count = 0
		fmt.Println("setTime is the same as the current time.")
	}

	foundedToday, totalRows, err := dao.GetAttendanceTodayByEmpNo(input.EmpNo)
	if err != nil {
		// Return status 404 and book not found error.
		baseResponse := &entity.BaseResponse{
			StatusCode: fiber.StatusBadRequest,
			Message:    "Failed GetAttendanceByEmpNo",
			Errors:     utils.GetError("10001", functionName, err.Error()),
		}
		return c.JSON(baseResponse)
	}

	if totalRows < 2 {
		if totalRows == 0 {
			status = 1
		} else {
			status = 2
		}
		// Create AdminUserRoles by given model.
		_, _, err = dao.CreateAttendance(&model.Attendance{
			EmpNo:              input.EmpNo,
			Latitude:           &input.Latitude,
			Longitude:          &input.Longitude,
			Selfie:             &data,
			CheckIn:            &now,
			CreatedAt:          &now,
			LateCount:          &late_count,
			AdmissionTimeLimit: &addmission_time_limit,
			TimeOutLimit:       &time_out_limit,
			Status:             &status,
		})

		if err != nil {
			// Return status 500 and error message.
			baseResponse := &entity.BaseResponse{
				StatusCode: fiber.StatusBadRequest,
				// Message:    utils.ValidatorErrors(err),
				Errors: utils.GetError("10004", functionName, err.Error()),
			}
			return c.Status(fiber.StatusBadRequest).JSON(baseResponse)
		}
	} else {
		for i := 0; i < len(foundedToday); i++ {
			if i == 1 {
				status = 2
				late_count = 0
				// Create AdminUserRoles by given model.
				_, _, err = dao.UpdateAttendance(foundedToday[i].ID, &model.Attendance{
					EmpNo:              input.EmpNo,
					Latitude:           &input.Latitude,
					Longitude:          &input.Longitude,
					Selfie:             &data,
					CheckIn:            &now,
					CreatedAt:          &now,
					LateCount:          &late_count,
					AdmissionTimeLimit: &addmission_time_limit,
					TimeOutLimit:       &time_out_limit,
					Status:             &status,
				})

				if err != nil {
					// Return status 500 and error message.
					baseResponse := &entity.BaseResponse{
						StatusCode: fiber.StatusBadRequest,
						// Message:    utils.ValidatorErrors(err),
						Errors: utils.GetError("10004", functionName, err.Error()),
					}
					return c.Status(fiber.StatusBadRequest).JSON(baseResponse)
				}
			}
		}
	}

	// Return status 201 CREATED.
	baseResponse := &entity.BaseResponse{
		StatusCode: fiber.StatusCreated,
		Message:    "Success",
		Errors:     false,
	}
	return c.JSON(baseResponse)
}
