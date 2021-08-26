package util

import "github.com/gogf/gf/crypto/gmd5"

func Password(raw string, salt string) string {
	return gmd5.MustEncryptString(gmd5.MustEncryptString(salt) + salt + gmd5.MustEncryptString(raw))
}
