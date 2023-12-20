package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/metadata"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

// TracerWrapper tracer 中间件
func TracerWrapper(c *gin.Context) {
	md := make(map[string]string)
	spanCtx, _ := opentracing.GlobalTracer().Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(c.Request.Header))
	sp := opentracing.GlobalTracer().StartSpan(c.Request.URL.Path, opentracing.ChildOf(spanCtx))
	defer sp.Finish()

	if err := opentracing.GlobalTracer().Inject(sp.Context(),
		opentracing.TextMap,
		opentracing.TextMapCarrier(md)); err != nil {
		fmt.Println(err)
	}

	ctx := context.TODO()
	ctx = opentracing.ContextWithSpan(ctx, sp)
	ctx = metadata.NewContext(ctx, md)
	c.Set("contextTracerKey", ctx)

	c.Next()

	statusCode := c.Writer.Status()
	ext.HTTPStatusCode.Set(sp, uint16(statusCode))
	ext.HTTPMethod.Set(sp, c.Request.Method)
	ext.HTTPUrl.Set(sp, c.Request.URL.EscapedPath())
	if statusCode >= http.StatusInternalServerError {
		ext.Error.Set(sp, true)
	}
}

// ContextWithSpan 返回context
func ContextWithSpan(c *gin.Context) (ctx context.Context, ok bool) {
	v, exist := c.Get("contextTracerKey")
	if !exist {
		ok = false
		return
	}

	ctx, ok = v.(context.Context)
	return
}
