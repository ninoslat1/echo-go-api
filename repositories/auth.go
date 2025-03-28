package repositories

import (
	"echo-api/configs"
	dbmodels "echo-api/models/db"
	loginmodels "echo-api/models/login"
	"echo-api/utils"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

func LoginHandler(c echo.Context, log *logrus.Logger) error {
	var localLog = logrus.New()
	db, err := configs.RunDatabase(log, "bromousr")

	if err != nil {
		localLog.Fatal("ERROR:", err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"ERROR": err.Error()})
	}

	loginReq := new(loginmodels.LoginRequest)
	if err := c.Bind(loginReq); err != nil {
		log.Error("Error binding request: ", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"ERROR": "Invalid request"})
	}

	if loginReq.UserName == "" || loginReq.Password == "" {
		log.Infof("User in %s want to access without credentials", c.RealIP())
		return c.JSON(http.StatusBadRequest, map[string]string{"ERROR": "Credentials are required"})
	}

	encodePassword := utils.EncodeToBase64Password(loginReq.Password)

	var user dbmodels.User
	if err := db.Table("myuser").Where("UserName = ? AND Password LIKE ?", loginReq.UserName, encodePassword[:len(encodePassword)-2]+"%").First(&user).Error; err != nil {
		log.Info("User in %s want to access with wrong credentials", c.RealIP())
		return c.JSON(http.StatusUnauthorized, map[string]string{"ERROR": err.Error()})
	}

	cookie, err := utils.SetSecureCookies(c)
	if err != nil {
		log.Error("User %s failed to create a session cookie", user.UserCode)
		localLog.Error("User %s failed to create a session cookie. %s", user.UserCode, err)
		return c.JSON(http.StatusInternalServerError, map[string]string{"ERROR": "Failed to generate session"})
	}

	response := loginmodels.LoginResponse{
		Message: "Welcome " + user.UserCode,
		Cookie:  cookie.Value,
	}

	return c.JSON(http.StatusOK, response)
}
