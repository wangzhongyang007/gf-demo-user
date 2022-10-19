package main

import (
	"github.com/gogf/gf-demo-user/v2/internal/cmd"
	_ "github.com/gogf/gf-demo-user/v2/internal/logic"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.Run(gctx.New())
}
