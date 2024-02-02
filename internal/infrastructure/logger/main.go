package logger

import (
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func isProdEnv() bool {
	return os.Getenv("ENV") == "prod"
}

func init() {
	zapConf := zap.Must(zap.NewDevelopment())
	if isProdEnv() {
		zapConf = zap.Must(zap.NewProduction())
	}
	zap.ReplaceGlobals(zapConf)
}

func ZapMiddleware(c echo.Context, next echo.HandlerFunc) error {
	req := c.Request()
	res := c.Response()
	start := time.Now()
	stop := time.Now()

	userAgent := req.UserAgent()
	bytesIn := req.Header.Get(echo.HeaderContentLength)
	rid := res.Header().Get(echo.HeaderXRequestID)
	if req.RequestURI == "/health" {
		return nil
	}

	zap.S().With(
		"trace", rid,
		"host", req.Host,
		"uri", req.RequestURI,
		"method", req.Method,
		"path", req.URL.Path,
		"status", res.Status,
		"user_agent", userAgent,
		"latency_human", stop.Sub(start).String(),
		"bytes_in", bytesIn,
	).Info("[incoming_request]")

	if err := next(c); err != nil {
		c.Error(err)
	}

	return nil
}

func BodyDumpHandler(c echo.Context, _, resBody []byte) {
	res := c.Response()
	rid := res.Header().Get(echo.HeaderXRequestID)

	zap.S().With("trace", rid, "body", string(resBody)).Info("[response_body]")
}

func handlerFunc(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		return ZapMiddleware(c, next)
	}
}

// LoggerHook is a function to process middleware.
func LoggerHook() echo.MiddlewareFunc {
	return handlerFunc
}
