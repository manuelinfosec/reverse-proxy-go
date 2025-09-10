package p

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strconv"
)

var shouldEnableHttps = getBooleanConfig("enableHttps")
var shouldVerifyApiKey = getBooleanConfig("verifyApiKey")

func init() {
}

func Proxy(responseWriter http.ResponseWriter, request *http.Request) {
	target, _ := url.Parse("https://api.predic8.de:443/shop/categories/")
	proxy := httputil.NewSingleHostReverseProxy(target)

	if shouldEnableHttps {
		enableHttps(request, target)
	}
	if shouldVerifyApiKey {

	}

	proxy.ServeHTTP(responseWriter, request)
}

func getBooleanConfig(configParameterName string) bool {
	result, _ := strconv.ParseBool(os.Getenv(configParameterName))
	return result
}

func enableHttps(request *http.Request, target *url.URL) {
	request.URL.Host = target.Host
	request.URL.Scheme = target.Scheme
	request.Header.Set("X-Forwarded-Host", request.Header.Get("Host"))
	request.Host = target.Host
}

func Main() {
	Proxy(nil, nil)
}
