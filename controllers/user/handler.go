package user

import (
	"gamePicker/business/user"
	"gamePicker/controllers"
	"gamePicker/controllers/user/request"
	"gamePicker/controllers/user/response"
	"net/http"
	"time"

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
	userLogin.Email = c.FormValue("email")
	userLogin.Password = c.FormValue("password")
	// c.Bind(&userLogin)

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

// Username        string `json:"username"`
// Password        string `json:"password"`
// Email           string `json:"email"`
// Name            string `json:"name"`
// SteamProfile_id string `json:"steamProfile_id"`
func (handler UserController) CreateUserController(c echo.Context) error {
	userInsert := user.Domain{}
	userInsert.Username = c.FormValue("username")
	userInsert.Password = c.FormValue("password")
	userInsert.Email = c.FormValue("email")
	userInsert.Name = c.FormValue("name")
	userInsert.SteamProfile_id = c.FormValue("steamProfile_id")
	userInsert.Created_at = time.Now()
	userInsert.Updated_at = time.Now()
	// fmt.Print(userInsert.SteamProfile_id)

	ctx := c.Request().Context()

	user, err := handler.UserUseCase.CreateUserController(ctx, userInsert)
	if err != nil {
		return controllers.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controllers.NewSuccessResponse(c, response.CreateFromDomain(user))
}
