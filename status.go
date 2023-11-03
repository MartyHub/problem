package problem

import "net/http"

// Status returns the standard HTTP phrase for the given HTTP status code.
func Status(code int) string { //nolint:cyclop,funlen,gocyclo
	switch code {
	// 4xx
	case http.StatusBadRequest:
		return "Bad Request"
	case http.StatusUnauthorized:
		return "Unauthorized"
	case http.StatusPaymentRequired:
		return "Payment Required"
	case http.StatusForbidden:
		return "Forbidden"
	case http.StatusNotFound:
		return "Not Found"
	case http.StatusMethodNotAllowed:
		return "Method Not Allowed"
	case http.StatusNotAcceptable:
		return "Not Acceptable"
	case http.StatusProxyAuthRequired:
		return "Proxy Authentication Required"
	case http.StatusRequestTimeout:
		return "Request Timeout"
	case http.StatusConflict:
		return "Conflict"
	case http.StatusGone:
		return "Gone"
	case http.StatusLengthRequired:
		return "Length Required"
	case http.StatusPreconditionFailed:
		return "Precondition Failed"
	case http.StatusRequestEntityTooLarge:
		return "Request Entity Too Large"
	case http.StatusRequestURITooLong:
		return "Request-URI Too Long"
	case http.StatusUnsupportedMediaType:
		return "Unsupported Media Type"
	case http.StatusRequestedRangeNotSatisfiable:
		return "Requested Range Not Satisfiable"
	case http.StatusExpectationFailed:
		return "Expectation Failed"
	case http.StatusTeapot:
		return "I'm a teapot"
	case http.StatusMisdirectedRequest:
		return "Misdirected Request"
	case http.StatusUnprocessableEntity:
		return "Unprocessable Entity"
	case http.StatusLocked:
		return "Locked"
	case http.StatusFailedDependency:
		return "Failed Dependency"
	case http.StatusTooEarly:
		return "Too Early"
	case http.StatusUpgradeRequired:
		return "Upgrade Required"
	case http.StatusPreconditionRequired:
		return "Precondition Required"
	case http.StatusTooManyRequests:
		return "Too Many Requests"
	case http.StatusRequestHeaderFieldsTooLarge:
		return "Request Header Fields Too Large"
	case http.StatusUnavailableForLegalReasons:
		return "Unavailable For Legal Reasons"

	// 5xx
	case http.StatusInternalServerError:
		return "Internal Server Error"
	case http.StatusNotImplemented:
		return "Not Implemented"
	case http.StatusBadGateway:
		return "Bad Gateway"
	case http.StatusServiceUnavailable:
		return "Service Unavailable"
	case http.StatusGatewayTimeout:
		return "Gateway Timeout"
	case http.StatusHTTPVersionNotSupported:
		return "HTTP Version Not Supported"
	case http.StatusVariantAlsoNegotiates:
		return "Variant Also Negotiates"
	case http.StatusInsufficientStorage:
		return "Insufficient Storage"
	case http.StatusLoopDetected:
		return "Loop Detected"
	case http.StatusNotExtended:
		return "Not Extended"
	case http.StatusNetworkAuthenticationRequired:
		return "Network Authentication Required"
	}

	return ""
}
