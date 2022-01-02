package model

import (
	"fmt"
	"net/http"

	"andromeda.ottopay.id/ottopoint/ottopoint-shared/model"
	"github.com/gofiber/fiber/v2"
)

// Response base response from fiber
func Response(c *fiber.Ctx, res *model.BaseResp) error {
	status := http.StatusOK
	if res.Code != 0 {
		status = res.Code
	}
	return c.Status(status).JSON(res)
}

func ResponseErr(c *fiber.Ctx, status int, msg string) error {
	return c.Status(status).JSON(
		fiber.Map{
			"responseCode": fmt.Sprintf("%d", status),
			"responseDesc": msg,
			"data":         nil,
		})
}
