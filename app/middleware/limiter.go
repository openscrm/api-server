package middleware

//func RateLimiter(l limiter.LimiterIface) gin.HandlerFunc {
//	return func(c *gin.Context) {
//key := l.Key(c)
//if bucket, ok := l.GetBucket(key); ok {
//	count := bucket.TakeAvailable(1)
//	if count == 0 {
//		app := app.NewResponse(c)
//		app.ResponseError(errcode.TooManyRequests)
//		c.Abort()
//		return
//	}
//}
//
//c.Next()
//}
//}
