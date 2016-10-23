package main

import (
	"fmt"

	"github.com/sspencer/pack"
)

func block(w, h int) *pack.Block {
	return &pack.Block{Width: w, Height: h}
}

func main() {
	blocks := pack.Blocks{
		block(50, 50),
		block(20, 70),
		block(60, 5),
	}

	fmt.Println(pack.BestFit(blocks))
}
