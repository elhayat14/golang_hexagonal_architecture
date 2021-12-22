package user

import (
	"github.com/devfeel/mapper"
	"github.com/labstack/echo/v4"
	"golang_hexagonal_architecture/adapters/api/utils/response"
	"golang_hexagonal_architecture/adapters/api/utils/validator"
	userCore "golang_hexagonal_architecture/core/user"
	portContractUser "golang_hexagonal_architecture/ports/contract/user"
	"net/http"
)

type Controller struct {
	userService userCore.IService
}

func NewController(userService userCore.IService) *Controller {
	return &Controller{
		userService: userService,
	}
}
func (controller *Controller) AddNew(c echo.Context) error {
	bodyRequest := new(portContractUser.CreateUserRequest)
	if err := c.Bind(bodyRequest); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewResponse(nil, response.MetaFromError(err)))
	}
	if err := c.Validate(bodyRequest); err != nil {
		data := validator.BuildErrorBodyRequestValidator(err)
		return c.JSON(http.StatusBadRequest, response.NewResponse(data, response.DefaultMetaError))
	}
	var userObj portContractUser.Object
	mapper.AutoMapper(bodyRequest, &userObj)
	newUser, err := controller.userService.DoCreateUser(userObj)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.NewResponse(nil, response.MetaFromError(err)))
	}
	var defaultResponse portContractUser.DefaultResponse
	mapper.AutoMapper(newUser, &defaultResponse)

	return c.JSON(http.StatusCreated, response.NewResponse(defaultResponse, response.ConstantMeta["created"]))
}
func (controller *Controller) EditById(c echo.Context) error {
	userId := c.Param("userId")
	bodyRequest := new(portContractUser.UpdateUserRequest)
	if err := c.Bind(bodyRequest); err != nil {
		return c.JSON(http.StatusBadRequest, response.NewResponse(nil, response.MetaFromError(err)))
	}
	if err := c.Validate(bodyRequest); err != nil {
		data := validator.BuildErrorBodyRequestValidator(err)
		return c.JSON(http.StatusBadRequest, response.NewResponse(data, response.DefaultMetaError))
	}
	var userObj portContractUser.Object
	mapper.AutoMapper(bodyRequest, &userObj)
	updatedUser, err := controller.userService.DoEditUser(userId, userObj)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.NewResponse(nil, response.MetaFromError(err)))
	}
	var defaultResponse portContractUser.DefaultResponse
	mapper.AutoMapper(updatedUser, &defaultResponse)

	return c.JSON(http.StatusOK, response.NewResponse(defaultResponse, response.ConstantMeta["update"]))
}
func (controller *Controller) GetById(c echo.Context) error {
	userId := c.Param("userId")
	user, err := controller.userService.DoGetUserById(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, response.NewResponse(nil, response.MetaFromError(err)))
	}
	var defaultResponse portContractUser.DefaultResponse
	mapper.AutoMapper(user, &defaultResponse)
	return c.JSON(http.StatusOK, response.NewResponse(defaultResponse, response.DefaultMeta))
}
func (controller *Controller) Delete(c echo.Context) error {
	userId := c.Param("userId")
	_, err := controller.userService.DoDeleteUser(userId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, response.NewResponse(map[string]string{"id": userId}, response.DefaultMeta))
}
