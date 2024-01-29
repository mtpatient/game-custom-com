package main

import (
	_ "game-custom-com/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"game-custom-com/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
