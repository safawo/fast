package fast

import (
	"github.com/safawo/fast/mvc"
)

func initControl() {
	mvc.Router("/github.com/safawo/fast/msg/loadList", &LoadMsgListAction{})
}
