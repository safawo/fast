package safe

import (
	"github.com/safawo/fast/mvc"
)

type SecLogInterface interface {
	writeLog(rspMsg mvc.FastResponseInterface, object, content string)
}
