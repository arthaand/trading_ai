package middleware

import (
	"encoding/json"
	"net/http"
	"trading-ai/domain/model"

	shared "andromeda.ottopay.id/ottopoint/ottopoint-shared"
	"andromeda.ottopay.id/ottopoint/ottopoint-shared/constant"
	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp"
)

func SignatureAuth(c *fiber.Ctx) error {
	authBaseUrl := shared.Config.AuthService.Host
	if shared.Config.AuthService.Port != "" {
		authBaseUrl += ":" + shared.Config.AuthService.Port
	}

	// init fasthttp client because Fiber does not wrap fasthttp
	client := fasthttp.Client{
		NoDefaultUserAgentHeader: true,
		DisablePathNormalizing:   true,
	}
	// init request object from fasthttp
	req := c.Request()
	req.SetRequestURI(authBaseUrl + "/auth/v2/signature")

	// init response object from fasthttp
	var (
		resBody map[string]json.RawMessage
		resCode string
	)
	res := c.Response()
	if err := client.Do(req, res); err != nil {
		return model.ResponseErr(c, http.StatusUnprocessableEntity, err.Error())
	}
	// get responseCode
	json.Unmarshal(res.Body(), &resBody)
	json.Unmarshal(resBody["response_code"], &resCode)

	if resCode != constant.RC_SUCCESS {
		return model.ResponseErr(c, http.StatusUnauthorized, "invalid signature")
	}
	return c.Next()
}
