package function

import (
	"fmt"
)

// Handle a serverless requestdf
func Handle(req []byte) string {
	return fmt.Sprintf("Helldfo, Gdo. You said: %s", string(req))
}
