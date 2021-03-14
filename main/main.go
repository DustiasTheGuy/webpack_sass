package main

import "github.com/DustiasTheGuy/webpack_sass/webpack_sass"

func main() {
	cfg := webpack_sass.Initalize()

	cfg.InitWebpack()
}
