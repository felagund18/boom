package boom

// 500 Internal Server Error
func InternalServerError(text string, data ...interface{}) string {
	return string(marshalJson(500, "Internal Server Error", text, data...))
}

// 501 Not Implemented
func NotImplemented(text string, data ...interface{}) string {
	return string(marshalJson(501, "Not Implemented", text, data...))
}

// 502 Bad Gateway
func BadGateway(text string, data ...interface{}) string {
	return string(marshalJson(502, "Bad Gateway", text, data...))
}

// 503 Service Unavailable
func ServiceUnavailable(text string, data ...interface{}) string {
	return string(marshalJson(503, "Service Unavailable", text, data...))
}

// 504 Gateway Timeout
func GatewayTimeout(text string, data ...interface{}) string {
	return string(marshalJson(504, "Gateway Timeout", text, data...))
}

// 505 HTTP Version Not Supported
func HTTPVersionNotSupported(text string, data ...interface{}) string {
	return string(marshalJson(505, "HTTP Version Not Supported", text, data...))
}

// 506 Variant Also Negotiates (Experimental)
func VariantAlsoNegotiates(text string, data ...interface{}) string {
	return string(marshalJson(506, "Variant Also Negotiates (Experimental)", text, data...))
}

// 507 Insufficient Storage (WebDAV)
func InsufficientStorage(text string, data ...interface{}) string {
	return string(marshalJson(507, "Insufficient Storage (WebDAV)", text, data...))
}

// 508 Loop Detected (WebDAV)
func LoopDetected(text string, data ...interface{}) string {
	return string(marshalJson(508, "Loop Detected (WebDAV)", text, data...))
}

// 509 Bandwidth Limit Exceeded (Apache)
func BandwidthLimitExceeded(text string, data ...interface{}) string {
	return string(marshalJson(509, "Bandwidth Limit Exceeded (Apache)", text, data...))
}

// 510 Not Extended
func NotExtended(text string, data ...interface{}) string {
	return string(marshalJson(510, "Not Extended", text, data...))
}

// 511 Network Authentication Required
func NetworkAuthenticationRequired(text string, data ...interface{}) string {
	return string(marshalJson(511, "Network Authentication Required", text, data...))
}

// 598 Network read timeout error
func NetworkReadTimeout(text string, data ...interface{}) string {
	return string(marshalJson(598, "Network read timeout error", text, data...))
}

// 599 Network connect timeout error
func NetworkConnectTimeout(text string, data ...interface{}) string {
	return string(marshalJson(599, "Network connect timeout error", text, data...))
}
