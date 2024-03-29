package main

import (
	"goe/app"
	"goe/app/common"
	"os"
)

/**
 * @command: go run main.go dev ||  go run main.go
 * @user: Mr.LiuQH
 * @date 2021-02-04 16:01:54
 */
func main() {
	appServer := &app.App{}
	arg := os.Args
	if len(arg) < 2 {
		appServer.Env = common.EnvDev
	} else {
		appServer.Env = arg[1]
	}
	appServer.Start()
}
