package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

//
//func ZapLogger(lg *zap.Logger) gin.HandlerFunc {
//	return func(c *gin.Context) {
//
//		start := time.Now()
//		path := c.Request.URL.Path
//		query := c.Request.URL.RawQuery
//		post := ""
//
//		if c.Request.Method == "POST" {
//			// 把request的内容读取出来
//			bodyBytes, _ := io.ReadAll(c.Request.Body)
//			err := c.Request.Body.Close()
//			if err != nil {
//				return
//			}
//			// 把刚刚读出来的再写进去
//			c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
//			switch c.ContentType() {
//			case "application/json":
//				var result map[string]interface{}
//				d := jsoniter.NewDecoder(bytes.NewReader(bodyBytes))
//				d.UseNumber()
//				if err := d.Decode(&result); err == nil {
//					bt, _ := jsoniter.Marshal(result)
//					post = string(bt)
//				}
//			default:
//				post = string(bodyBytes)
//			}
//		}
//
//		c.Next()
//
//		cost := time.Since(start)
//		lg.Info(path,
//			zap.Int("status", c.Writer.Status()),
//			zap.String("method", c.Request.Method),
//			zap.String("path", path),
//			zap.String("query", query),
//			zap.String("post", post),
//			zap.String("ip", c.ClientIP()),
//			zap.String("user-agent", c.Request.UserAgent()),
//			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
//			zap.Duration("cost", cost),
//		)
//	}
//}

func ZapLogger(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery
		c.Next()

		cost := time.Since(start)
		logger.Info(path,
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("cost", cost),
		)
	}
}
