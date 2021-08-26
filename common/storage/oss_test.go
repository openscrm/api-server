package storage

import (
	"net/http"
	"openscrm/conf"
	"testing"
)

func TestOSSStorage_SignURL(t *testing.T) {
	err := conf.SetupSetting()
	if err != nil {
		t.Failed()
	}
	Setup(conf.Settings.Storage)
	signedURL, err := FileStorage.SignURL("wwacfbf964143dc303/quick_reply/admin/4566456.jpg", http.MethodPut, 3600*24*365)
	if err != nil {
		t.Failed()
	}
	println(signedURL)

}
