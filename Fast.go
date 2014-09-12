package fast

import (
	_ "think/fast/backup"

	_ "think/fast/safe"

	_ "think/fast/mvc"

	_ "think/fast/msg"

	_ "think/fast/utils"

	_ "think/fast/comm"

	_ "think/fast/ds"

	_ "github.com/lxn/go-pgsql"

	_ "think/fast/run"

	"fmt"
)

func init() {
	fmt.Println("Start Fast PlatForm")

	initControl()
}
