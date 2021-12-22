package v1

import (
	"github.com/labstack/echo/v4"
	"golang_hexagonal_architecture/adapters/api/v1/user"
	"net/http"
)

type Router struct {
	UserController user.Controller
	Echo           *echo.Echo
}

func RegisterRouter(param Router) {
	param.Echo.GET("/", func(c echo.Context) error {
		return c.HTML(http.StatusOK, "Welcome")
	})
	param.Echo.GET("/health", func(c echo.Context) error {
		return c.NoContent(200)
	})
	prefixV1 := "api/v1/"
	userV1 := param.Echo.Group(prefixV1 + "users")
	userV1.POST("", param.UserController.AddNew)
	userV1.GET("/:userId", param.UserController.GetById)
	userV1.PUT("/:userId", param.UserController.EditById)
	userV1.DELETE("/:userId", param.UserController.Delete)

}
