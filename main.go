package main

import (
	"github.com/binz96/blog/model"
	"github.com/binz96/blog/router"
	"github.com/binz96/blog/setting"
)

func main() {
	setting.Load()
	model.Load()
	router.Run()
}
