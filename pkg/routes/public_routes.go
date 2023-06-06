package routes

import (
	"attendance/app/controllers/attendance"
	"attendance/app/controllers/authentication"
	"attendance/app/controllers/employee"

	"github.com/gofiber/fiber/v2"
)

// PublicRoutes func for describe group of public routes.
func PublicRoutes(a *fiber.App) {
	// Create routes group.
	route := a.Group("/api/v1")

	//AUTH
	route.Post("/login", authentication.Login) // auth, return Access & Refresh tokens
	route.Post("/attendance", attendance.CreateAttendance)
	route.Get("/attendances", attendance.GetAttendances)
	route.Get("/attendance/get-today/:emp_no", attendance.GetAttendanceByEmpNoToday)
	route.Get("/attendance/:emp_no", attendance.GetAttendanceByEmpNo)
	route.Get("/employee/:emp_no", employee.GetEmpByEmpNo)
}
