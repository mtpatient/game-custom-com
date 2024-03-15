package main

import (
	_ "game-custom-com/internal/logic"
	_ "game-custom-com/internal/packed"
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"game-custom-com/internal/cmd"
)

// @title Game Custom API
// @version 1.0
// @description 个人游戏社区后端接口
func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
