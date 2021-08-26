package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/thoas/go-funk"
	"openscrm/app/constants"
	"openscrm/app/models"
	"openscrm/common/app"
	"openscrm/common/ecode"
)

func Guard(biz constants.BizIdentity, operation constants.Operation) gin.HandlerFunc {
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

		// todo 改为从缓存取，提高性能
		staff, err := item.Get(item.ExtID, item.ExtCorpID, false)
		if err != nil {
			handler.ResponseError(err)
			c.Abort()
			return
		}

		if staff.Status != constants.UserStatusActivated {
			handler.ResponseError(ecode.ForbiddenError)
			c.Abort()
			return
		}

		role, err := (&models.Role{}).CachedGet(staff.RoleID)
		if err != nil {
			handler.ResponseError(err)
			c.Abort()
			return
		}

		identity := fmt.Sprintf("%s_%s", biz, operation)

		if !funk.ContainsString(role.PermissionIDs, identity) {
			handler.ResponseBadRequestError(ecode.ForbiddenError)
			c.Abort()
			return
		}

		c.Next()
	}
}
