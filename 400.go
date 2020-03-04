package boom

// 400 Bad Request
func BadRequest(text string, data ...interface{}) string {
	return string(marshalJson(400, "Bad Request", text, data...))
}

// 401 Unauthorized
func Unauthorized(text string, data ...interface{}) string {
	return string(marshalJson(401, "Unauthorized", text, data...))
}

// 402 Payment Required
func PaymentRequired(text string, data ...interface{}) string {
	return string(marshalJson(402, "Payment Required", text, data...))
}

// 403 Forbidden
func Forbidden(text string, data ...interface{}) string {
	return string(marshalJson(403, "Forbidden", text, data...))
}

// 404 Not Found
func NotFound(text string, data ...interface{}) string {
	return string(marshalJson(404, "Not Found", text, data...))
}

// 405 Method Not Allowed
func MethodNotAllowed(text string, data ...interface{}) string {
	return string(marshalJson(405, "Method Not Allowed", text, data...))
}

// 406 Not Acceptable
func NotAcceptable(text string, data ...interface{}) string {
	return string(marshalJson(406, "Not Acceptable", text, data...))
}

// 407 Proxy Authentication Required
func ProxyAuthenticationRequired(text string, data ...interface{}) string {
	return string(marshalJson(407, "Proxy Authentication Required", text, data...))
}

// 408 Request Timeout
func RequestTimeout(text string, data ...interface{}) string {
	return string(marshalJson(408, "Request Timeout", text, data...))
}

// 409 Conflict
func Conflict(text string, data ...interface{}) string {
	return string(marshalJson(409, "Conflict", text, data...))
}

// 410 Gone
func Gone(text string, data ...interface{}) string {
	return string(marshalJson(410, "Gone", text, data...))
}

// 411 Length Required
func LengthRequired(text string, data ...interface{}) string {
	return string(marshalJson(411, "Length Required", text, data...))
}

// 412 Precondition Failed
func PreconditionFailed(text string, data ...interface{}) string {
	return string(marshalJson(412, "Precondition Failed", text, data...))
}

// 413 Request Entity Too Large
func RequestEntityTooLarge(text string, data ...interface{}) string {
	return string(marshalJson(413, "Request Entity Too Large", text, data...))
}

// 414 Request-URI Too Long
func RequestURITooLong(text string, data ...interface{}) string {
	return string(marshalJson(414, "Request-URI Too Long", text, data...))
}

// 415 Unsupported Media Type
func UnsupportedMediaType(text string, data ...interface{}) string {
	return string(marshalJson(415, "Unsupported Media Type", text, data...))
}

// 416 Requested Range Not Satisfiable
func RequestedRangeNotSatisfiable(text string, data ...interface{}) string {
	return string(marshalJson(416, "Requested Range Not Satisfiable", text, data...))
}

// 417 Expectation Failed
func ExpectationFailed(text string, data ...interface{}) string {
	return string(marshalJson(417, "Expectation Failed", text, data...))
}

// 418 I'm a teapot (RFC 2324)
func ImATeapot(text string, data ...interface{}) string {
	return string(marshalJson(418, "I'm a teapot (RFC 2324)", text, data...))
}

// 420 Enhance Your Calm (Twitter)
func EnhanceYourCalm(text string, data ...interface{}) string {
	return string(marshalJson(420, "Enhance Your Calm (Twitter)", text, data...))
}

// 422 Unprocessable Entity
func UnprocessableEntity(text string, data ...interface{}) string {
	return string(marshalJson(422, "Unprocessable Entity (WebDAV)", text, data...))
}

// 423 Locked (WebDAV)
func Locked(text string, data ...interface{}) string {
	return string(marshalJson(423, "Locked (WebDAV)", text, data...))
}

// 424 Failed Dependency (WebDAV)
func FailedDependency(text string, data ...interface{}) string {
	return string(marshalJson(424, "Failed Dependency (WebDAV)", text, data...))
}

// 425 Reserved for WebDAV
func ReservedForWebDAV(text string, data ...interface{}) string {
	return string(marshalJson(425, "Reserved for WebDAV", text, data...))
}

// 426 Upgrade Required
func UpgradeRequired(text string, data ...interface{}) string {
	return string(marshalJson(426, "Upgrade Required", text, data...))
}

// 428 Precondition Required
func PreconditionRequired(text string, data ...interface{}) string {
	return string(marshalJson(428, "Precondition Required", text, data...))
}

// 429 Too Many Requests
func TooManyRequests(text string, data ...interface{}) string {
	return string(marshalJson(429, "Too Many Requests", text, data...))
}

// 431 Request Header Fields Too Large
func RequestHeaderFieldsTooLarge(text string, data ...interface{}) string {
	return string(marshalJson(431, "Request Header Fields Too Large", text, data...))
}

// 444 No Response (Nginx)
func NoResponse(text string, data ...interface{}) string {
	return string(marshalJson(444, "No Response (Nginx)", text, data...))
}

// 449 Retry With (Microsoft)
func RetryWith(text string, data ...interface{}) string {
	return string(marshalJson(449, "Retry With (Microsoft)", text, data...))
}

// 450 Blocked by Windows Parental Controls (Microsoft)
func BlockedByWindowsParentalControls(text string, data ...interface{}) string {
	return string(marshalJson(450, "Blocked by Windows Parental Controls (Microsoft)", text, data...))
}

// 451 Unavailable For Legal Reasons
func UnavailableForLegalReasons(text string, data ...interface{}) string {
	return string(marshalJson(451, "Unavailable For Legal Reasons", text, data...))
}

// 499 Client Closed Request (Nginx)
func ClientClosedRequest(text string, data ...interface{}) string {
	return string(marshalJson(499, "Client Closed Request (Nginx)", text, data...))
}
