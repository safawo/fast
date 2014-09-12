package fast

import (
	"think/fast/mvc"
)

func initControl() {
	mvc.Router("/think/fast/msg/loadList", &LoadMsgListAction{})
}
