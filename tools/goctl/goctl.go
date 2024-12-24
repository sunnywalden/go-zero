package main

import (
	"github.com/sunnywalden/go-zero/core/load"
	"github.com/sunnywalden/go-zero/core/logx"
	"github.com/sunnywalden/go-zero/tools/goctl/cmd"
)

func main() {
	logx.Disable()
	load.Disable()
	cmd.Execute()
}
