package utils

import (
	"auth-services/model"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-resty/resty/v2"
	"net/http"
	"strconv"
	"strings"
	"sync"
)

var (
	PlnUtilsContentTypeKey        = http.CanonicalHeaderKey("Content-Type")
	PlnUtilsContentDispositionKey = http.CanonicalHeaderKey("Content-Disposition")
	PlnUtilsContentLengthKey      = http.CanonicalHeaderKey("Content-Length")
)

func PlnMapQuerystring(ctx *gin.Context) map[string]string {
	// Transformation from map[string][]string to map[string]string:
	reqParams := map[string]string{}
	if len(ctx.Request.URL.Query()) > 0 {
		for k, v := range ctx.Request.URL.Query() {
			reqParams[k] = v[0]
		}
	}

	return reqParams
}

func PlnReturnResponse(ctx *gin.Context, returnData interface{}, resp *resty.Response, err error) {
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &model.Response{
			Success:     false,
			Status:      http.StatusInternalServerError,
			ErrorCode:   "4402",
			RespMessage: "Failed to process",
			RespData:    gin.H{"error": err.Error()},
		})
		return
	}

	if resp.StatusCode() != 200 {
		var returnDataErr gin.H
		_ = json.Unmarshal([]byte(resp.String()), &returnDataErr)
		ctx.JSON(resp.StatusCode(), returnDataErr)
		return
	}

	ctx.JSON(http.StatusOK, returnData)
}

func PlnReturnStreamResponse(ctx *gin.Context, returnData interface{}, resp *resty.Response, err error) {
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &model.Response{
			Success:     false,
			Status:      http.StatusInternalServerError,
			ErrorCode:   "4402",
			RespMessage: "Failed to process",
			RespData:    gin.H{"error": err.Error()},
		})
		return
	}

	if resp.StatusCode() != 200 {
		ctx.JSON(resp.StatusCode(), &model.Response{
			Success:     false,
			Status:      resp.StatusCode(),
			ErrorCode:   strconv.Itoa(resp.StatusCode()),
			RespMessage: resp.Status(),
			RespData:    gin.H{"error": resp.Status()},
		})
		ctx.Status(resp.StatusCode())
		return
	}

	contentType := resp.Header().Get(PlnUtilsContentTypeKey)
	if contentType == "application/octet-stream" {
		ctx.Header(PlnUtilsContentDispositionKey, resp.Header().Get(PlnUtilsContentDispositionKey))
		ctx.Header(PlnUtilsContentLengthKey, resp.Header().Get(PlnUtilsContentLengthKey))
		ctx.Header(PlnUtilsContentTypeKey, resp.Header().Get(PlnUtilsContentTypeKey))
		ctx.Status(resp.StatusCode())
		ctx.Data(http.StatusOK, contentType, resp.Body())
		return
	} else {
		ctx.JSON(http.StatusOK, returnData)
	}
}

func PlnReturnAnyResponse(ctx *gin.Context, resp *resty.Response, err error) {
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, &model.Response{
			Success:     false,
			Status:      http.StatusInternalServerError,
			ErrorCode:   "4402",
			RespMessage: "Failed to process",
			RespData:    gin.H{"error": err.Error()},
		})
		return
	}

	if resp.StatusCode() != 200 {
		ctx.JSON(resp.StatusCode(), &model.Response{
			Success:     false,
			Status:      resp.StatusCode(),
			ErrorCode:   strconv.Itoa(resp.StatusCode()),
			RespMessage: resp.Status(),
			RespData:    gin.H{"error": resp.Status()},
		})
		ctx.Status(resp.StatusCode())
		return
	}

	if resp.Header().Get(PlnUtilsContentDispositionKey) != "" {
		ctx.Header(PlnUtilsContentDispositionKey, resp.Header().Get(PlnUtilsContentDispositionKey))
	}
	if resp.Header().Get(PlnUtilsContentLengthKey) != "" {
		ctx.Header(PlnUtilsContentLengthKey, resp.Header().Get(PlnUtilsContentLengthKey))
	}
	if resp.Header().Get(PlnUtilsContentTypeKey) != "" {
		ctx.Header(PlnUtilsContentTypeKey, resp.Header().Get(PlnUtilsContentTypeKey))
	}

	ctx.Status(resp.StatusCode())
	ctx.Data(http.StatusOK, resp.Header().Get(PlnUtilsContentTypeKey), resp.Body())
	return
}

func PlnPrintOutResponse(resp *resty.Response, err error) {
	if resp != nil {
		fmt.Println("\nRequest Info: ")
		fmt.Println(fmt.Sprintf("  URL        : %v %v", resp.Request.Method, resp.Request.URL))

		//Response Header
		r := resp.Request
		// Loop over header names
		for name, values := range r.Header {
			// Loop over all values for the name.
			for _, value := range values {
				fmt.Println(fmt.Sprintf("  HEADER     : %v  %v", name, value))
			}
		}

		//Response body
		if !strings.HasPrefix(resp.Request.Header.Get(PlnUtilsContentTypeKey), "multipart/form-data") {
			if resp.Request.Method == resty.MethodPost {
				body := resp.Request.Body
				if b, ok := body.([]byte); ok {
					fmt.Println(fmt.Sprintf("  BODY        : %v", string(b)))
				} else {
					fmt.Println(fmt.Sprintf("  BODY        : %v", body))
				}
			}
		}

		fmt.Println("Response Info:")
		fmt.Println("  Error      :", err)
		fmt.Println("  Status Code:", resp.StatusCode())
		fmt.Println("  Status     :", resp.Status())
		fmt.Println("  Proto      :", resp.Proto())
		fmt.Println("  Time       :", resp.Time())
		fmt.Println("  Received At:", resp.ReceivedAt())
		fmt.Println("  ContentType:", resp.Header().Get(PlnUtilsContentTypeKey))
		if resp.Header().Get(PlnUtilsContentTypeKey) == "application/octet-stream" {
			fmt.Println("  ContentLength:", resp.Header().Get(PlnUtilsContentLengthKey))
			fmt.Println("  ContentDisposition:", resp.Header().Get(PlnUtilsContentDispositionKey))
		} else {
			fmt.Println("  Body       :\n", resp)
		}
		fmt.Println()
	}
}

func PlnAdditionalHeaderFromRequest(ctx *gin.Context) map[string]string {
	additionalHeaders := map[string]string{}
	if ctx.Request.Header.Get(PlnUtilsContentTypeKey) != "" {
		additionalHeaders[PlnUtilsContentTypeKey] = ctx.Request.Header.Get(PlnUtilsContentTypeKey)
	}
	if ctx.Request.Header.Get(PlnUtilsContentLengthKey) != "" {
		additionalHeaders[PlnUtilsContentLengthKey] = ctx.Request.Header.Get(PlnUtilsContentLengthKey)
	}
	if ctx.Request.Header.Get(PlnUtilsContentDispositionKey) != "" {
		additionalHeaders[PlnUtilsContentDispositionKey] = ctx.Request.Header.Get(PlnUtilsContentDispositionKey)
	}

	return additionalHeaders
}

func PlnSetHeader(ctx *gin.Context, additionalHeaders ...map[string]string) map[string]string {
	headers := map[string]string{}

	//Get from context
	if ctx != nil {
		userPln := &model.DetailUserPln{}
		var mu sync.Mutex
		mu.Lock()
		if u, exist := ctx.Get("user-pln"); exist && u != nil {
			if v, ok := u.(*model.DetailUserPln); ok {
				userPln = v
			} else if v, ok := u.(model.DetailUserPln); ok {
				userPln = &v
			}
		}
		mu.Unlock()

		if userPln != nil && userPln.CorporateId != "" {
			headers["corporate-id"] = userPln.CorporateId
			headers["user-id"] = userPln.UserId
			headers["bnidirect-api-key"] = userPln.ApiKey
		}
	}

	//Set additional headers
	if len(additionalHeaders) >= 1 {
		for i, h := range additionalHeaders[0] {
			headers[i] = h
		}
	}

	return headers
}

func Round(x, unit float64) float64 {
	var rounded float64
	if x > 0 {
		rounded = float64(int64(x/unit+0.5)) * unit
	} else {
		rounded = float64(int64(x/unit-0.5)) * unit
	}
	formatted, err := strconv.ParseFloat(fmt.Sprintf("%.2f", rounded), 64)
	if err != nil {
		return rounded
	}
	return formatted
}
