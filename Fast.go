package fast

import (
	_ "github.com/safawo/fast/backup"

	_ "github.com/safawo/fast/safe"

	_ "github.com/safawo/fast/mvc"

	_ "github.com/safawo/fast/msg"

	_ "github.com/safawo/fast/utils"

	_ "github.com/safawo/fast/comm"

	_ "github.com/safawo/fast/ds"

	_ "github.com/lxn/go-pgsql"

	_ "github.com/safawo/fast/run"

	"fmt"
)

func init() {
	fmt.Println("Start Fast PlatForm")

	initControl()
}
