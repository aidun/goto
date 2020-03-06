package function

import (
	"fmt"
)

// Handle a serverless requestdfddf
func Handle(req []byte) string {
	return fmt.Sprintf("Hedlldfo, Gdo. You said: %s", string(req))
}
