package pack

import (
	"fmt"
	"strings"
)

// Canvas contains location information for all the sprites
type Canvas struct {
	Root   *Block
	Blocks Blocks
	layout Layout
}

func (c *Canvas) String() string {
	arr := []string{}
	arr = append(arr, fmt.Sprintf("|========== Canvas {%d x %d} ==========|", c.Root.Width, c.Root.Height))

	for _, b := range c.Blocks {
		arr = append(arr, fmt.Sprintf("|    %s", b.String()))
	}

	arr = append(arr, fmt.Sprintf("|%s|", strings.Repeat("=", len(arr[0])-2)))
	return strings.Join(arr, "\n")
}
