package function

import (
	"fmt"
)

// Handle a serverless requestdfd
func Handle(req []byte) string {
	return fmt.Sprintf("Hedlldfo, Gdo. You said: %s", string(req))
}
