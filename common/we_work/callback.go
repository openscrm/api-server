package we_work

import (
	"openscrm/conf"
	work_wx "openscrm/pkg/easywework"
)

var Callback *work_wx.CallBackHandler

func SetupWXCallback() {
	var err error
	Callback, err = work_wx.NewCBHandler(conf.Settings.WeWork.CallbackToken, conf.Settings.WeWork.CallbackAesKey)
	if err != nil {
		panic("init callback handler failed")
	}
}
