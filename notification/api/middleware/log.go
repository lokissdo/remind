package middleware

import (
	"bytes"
	"io/ioutil"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

// Log to file
func LogRequest() gin.HandlerFunc {

	return func(c *gin.Context) {
		//Start time
		startTime := time.Now()
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw
		//Process request
		bodyReq, _ := ioutil.ReadAll(c.Request.Body)
		c.Request.Body = ioutil.NopCloser(bytes.NewReader(bodyReq))
		c.Next()
		errormsg, ok := c.Get("error")
		if !ok {
			logrus.WithFields(logrus.Fields{
				"latency_time": time.Since(startTime),
				"resq": gin.H{
					"client_ip":  c.ClientIP(),
					"req_method": c.Request.Method,
					"req_uri":    c.Request.RequestURI,
					// "req_header": c.Request.Header,
					// "req_body": string(bodyReq),
				},
				"resp": gin.H{
					// "body":        blw.body.String(),
					"status_code": c.Writer.Status(),
				},
			}).Info("request access")
		} else {
			logrus.WithFields(logrus.Fields{
				"latency_time": time.Since(startTime),
				"resq": gin.H{
					"client_ip":  c.ClientIP(),
					"req_method": c.Request.Method,
					"req_uri":    c.Request.RequestURI,
					// "req_header": c.Request.Header,
					// "req_body":   string(bodyReq),
				},
				"resp": gin.H{
					// "body":        blw.body.String(),
					"status_code": c.Writer.Status(),
				},
			}).Error(errormsg)
		}
	}
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}
