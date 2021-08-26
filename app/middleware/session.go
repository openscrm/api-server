package middleware

import (
	"github.com/gin-gonic/gin"
	"openscrm/app/constants"
	"openscrm/app/models"
	"openscrm/common/app"
)

func RequireStaffAdminLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		handler := app.NewHandler(c)
		item := models.Staff{}
		err = handler.GetSessionVal(constants.StaffAdminSessionName, constants.StaffInfo, &item)
		if err != nil {
			handler.ResponseError(err)
			c.Abort()
			return
		}
		c.Next()
	}
}

func RequireStaffLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		handler := app.NewHandler(c)
		item := models.Staff{}
		err = handler.GetSessionVal(constants.StaffSessionName, constants.StaffInfo, &item)
		if err != nil {
			handler.ResponseError(err)
			c.Abort()
			return
		}
		c.Next()
	}
}

func RequireCustomerLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		handler := app.NewHandler(c)
		item := models.Customer{}
		err = handler.GetSessionVal(constants.CustomerSessionName, constants.CustomerInfo, &item)
		if err != nil {
			handler.ResponseError(err)
			c.Abort()
			return
		}
		c.Next()
	}
}
