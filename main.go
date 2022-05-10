package main

import (
	_ "gf-bugs/packed"

	"github.com/gogf/gf/v2/os/gctx"
	"gf-bugs/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
