package files

import (
	"net/http"
)

func isBinary(content string) bool {
	return http.DetectContentType([]byte(content))[0:4] != "text"
}
