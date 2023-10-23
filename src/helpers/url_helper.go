package helpers

import (
	"strings"
)

// GetUrl constructs a URL with the given base API URL, API key, and additional parameters.
func GetUrl(apiUrl string, apiKey string, params map[string]string) string {
	var builder strings.Builder

	builder.WriteString(apiUrl)
	builder.WriteString("/?")

	builder.WriteString("apikey=")
	builder.WriteString(apiKey)

	for key, value := range params {
		builder.WriteString("&")
		builder.WriteString(key)
		builder.WriteString("=")
		builder.WriteString(value)
	}

	url := builder.String()

	return url
}
