package safe

import (
	"github.com/safawo/fast/mvc"
)

type SafeSessionValidater struct {
}

func (this *SafeSessionValidater) Validate(reqData mvc.FastRequestInterface) (ok bool) {
	ok = true

	if reqData.GetReqSessionId() == "init" {
		return
	}

	if SessionMgr().IsValid(reqData.GetReqSessionId()) != true {
		ok = false
	}

	return
}
