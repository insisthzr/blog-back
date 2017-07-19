package api

import (
	"github.com/Sirupsen/logrus"
	"github.com/labstack/echo"

	"github.com/insisthzr/blog-back/busniess"
	"github.com/insisthzr/blog-back/util"
)

func Login(c echo.Context) error {
	userIn := &busniess.UserIn{}
	err := c.Bind(userIn)
	if err != nil {
		return err
	}
	in := &busniess.LoginIn{UserIn: *userIn}
	out, err := busniess.Login(in)
	logrus.WithFields(logrus.Fields{"out": out, "error": err}).Infoln("Login")
	if err != nil {
		return err
	}
	token := util.NewJWTToken(out.ID)
	return c.JSON(200, JSON{"user": out, "token": token})
}

func Signup(c echo.Context) error {
	userIn := &busniess.UserIn{}
	err := c.Bind(userIn)
	if err != nil {
		return err
	}
	in := &busniess.SignupIn{UserIn: *userIn}
	out, err := busniess.Signup(in)
	logrus.WithFields(logrus.Fields{"out": out, "error": err}).Infoln("Signup")
	if err != nil {
		return err
	}
	return c.JSON(200, JSON{"user": out})
}