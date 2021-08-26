package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
)

type AccessLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w AccessLogWriter) Write(p []byte) (int, error) {
	if n, err := w.body.Write(p); err != nil {
		return n, err
	}
	return w.ResponseWriter.Write(p)
}

func AccessLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		//bodyWriter := &AccessLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		//c.Writer = bodyWriter

		//beginTime := time.Now().Seconds()
		//c.Next()
		//endTime := time.Now().Seconds()

		//fields := logger.Fields{
		//	"app":  c.Request.PostForm.Encode(),
		//	"app": bodyWriter.body.String(),
		//}
		//s := "access log: method: %s, status_code: %d, " +
		//	"begin_time: %d, end_time: %d"
		//global.Logger.WithFields(fields).Infof(c, s,
		//	c.Request.Method,
		//	bodyWriter.OutFlowStatus(),
		//	beginTime,
		//	endTime,
		//)
	}
}
