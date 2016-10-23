package pack

import (
	"fmt"
	"strings"
)

// Block holds describes the location of each and every image within the canvas.
type Block struct {
	Name   string
	X      int
	Y      int
	Width  int
	Height int
	used   bool
	fit    *Block
	right  *Block
	down   *Block
}

// Blocks is a slice of Blocks
type Blocks []*Block

func newBlock(name string, x, y, w, h int) *Block {
	return &Block{Name: name, X: x, Y: y, Width: w, Height: h}
}

func (b *Block) String() string {
	if b.Name == "" {
		return fmt.Sprintf("{{%d, %d} {%d x %d}}", b.X, b.Y, b.Width, b.Height)
	}
	return fmt.Sprintf("{%q: {%d, %d} {%d x %d}}", b.Name, b.X, b.Y, b.Width, b.Height)
}

func (a Blocks) String() string {
	arr := []string{}
	for _, b := range a {
		arr = append(arr, b.String())
	}

	return strings.Join(arr, "\n")
}
