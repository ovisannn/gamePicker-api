package user

import (
	"gamePicker/business/user"
	"gamePicker/controllers"
	"gamePicker/controllers/user/request"
	"gamePicker/controllers/user/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	UserUseCase user.UseCase
}

func NewUserController(userUseCase user.UseCase) *UserController {
	return &UserController{
		UserUseCase: userUseCase,
	}
}

func (handler UserController) Login(c echo.Context) error {
	userLogin := request.UserLogin{}
	// userLogin.Email = c.FormValue("email")
	// userLogin.Password = c.FormValue("password")
	c.Bind(&userLogin)

	ctx := c.Request().Context()
	user, err := handler.UserUseCase.Login(ctx, userLogin.Email, userLogin.Password)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, response.FromDomain(user))
}

func (handler UserController) GetUserController(c echo.Context) error {

	ctx := c.Request().Context()
	user, err := handler.UserUseCase.GetUserController(ctx)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, user)
}
