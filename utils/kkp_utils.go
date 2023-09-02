package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	KkpUtilsContentTypeKey        = http.CanonicalHeaderKey("Content-Type")
	KkpUtilsContentDispositionKey = http.CanonicalHeaderKey("Content-Disposition")
	KkpUtilsContentLengthKey      = http.CanonicalHeaderKey("Content-Length")
)

func KkpMapQuerystring(ctx *gin.Context) map[string]string {
	// Transformation from map[string][]string to map[string]string:
	reqParams := map[string]string{}
	if len(ctx.Request.URL.Query()) > 0 {
		for k, v := range ctx.Request.URL.Query() {
			reqParams[k] = v[0]
		}
	}

	return reqParams
}

func KkpAdditionalHeaderFromRequest(ctx *gin.Context) map[string]string {
	additionalHeaders := map[string]string{}
	if ctx.Request.Header.Get(KkpUtilsContentTypeKey) != "" {
		additionalHeaders[KkpUtilsContentTypeKey] = ctx.Request.Header.Get(KkpUtilsContentTypeKey)
	}
	if ctx.Request.Header.Get(KkpUtilsContentLengthKey) != "" {
		additionalHeaders[KkpUtilsContentLengthKey] = ctx.Request.Header.Get(KkpUtilsContentLengthKey)
	}
	if ctx.Request.Header.Get(KkpUtilsContentDispositionKey) != "" {
		additionalHeaders[KkpUtilsContentDispositionKey] = ctx.Request.Header.Get(KkpUtilsContentDispositionKey)
	}

	return additionalHeaders
}

func KkpSetHeader(ctx *gin.Context, additionalHeaders ...map[string]string) map[string]string {
	headers := map[string]string{}

	//Set additional headers
	if len(additionalHeaders) >= 1 {
		for i, h := range additionalHeaders[0] {
			headers[i] = h
		}
	}

	return headers
}
