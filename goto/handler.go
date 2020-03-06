package function

import (
	"fmt"
)

// Handle a serverless requestdf
func Handle(req []byte) string {
	return fmt.Sprintf("Hedlldfo, Gdo. You said: %s", string(req))
}
