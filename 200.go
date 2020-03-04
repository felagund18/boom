package boom

// 200 Ok
func Ok(text string, data ...interface{}) string {
	return string(marshalJson(200, "Ok", text, data...))
}
