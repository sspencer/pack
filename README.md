# Packer

Packer is a binary tree algorithm used to find the most efficient way
to pack rectangles into the smallest space.  The core algorithm stands alone,
no images required. This utility is written in Go and the algorithm is based on
Jake Gordon's [bin-packing](https://github.com/jakesgordon/bin-packing)
JavaScript project on github.

The goal for this project is to create a command line Sprite Packer program.
For now, sample images are created using the fit command.

    cd fit
    go run main.go

One of the examples generated from fit is complex.png:

<img src="doc/complex.png?raw=true" />

To pack images (or just plain ol' rectangles), the code is as simple as this:

```go
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
```

To run this program:

````
$ cd example
$ go run main.go
|========== Canvas {80 x 70} ==========|
|    {{0, 0} {20 x 70}}
|    {{20, 0} {60 x 5}}
|    {{20, 5} {50 x 50}}
|======================================|
````

## About the Code

Channels are used for an "embarrassingly parallel" problem ... pack the
images in 4 different ways by sorting the images differently:

* sort by width
* sort by height
* sort by area
* sort by max side (width or height)

By changing the sort order of the images, an occasional advantage can
be realized.  Each algorithm is attempted and the empty space is
calculated.  The algorithm that generates the least empty space wins.

For example:

	==== Packing complex ====
	LayoutByWidth <650x730> has wasted 194700 pixels
	LayoutByArea <650x650> has wasted 142700 pixels
	LayoutByHeight <530x530> has wasted 1100 pixels
	LayoutByMax <730x400> has wasted 12200 pixels
	>>>> RETURNING  LayoutByHeight
