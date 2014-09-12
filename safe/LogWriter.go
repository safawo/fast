package safe

import (
	"think/fast/mvc"
)

type SecLogInterface interface {
	writeLog(rspMsg mvc.FastResponseInterface, object, content string)
}
