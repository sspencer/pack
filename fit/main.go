package main

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"strconv"
	"strings"

	"github.com/sspencer/pack"
)

type blockdef []string

var (
	palette = []color.Color{
		color.RGBA{253, 243, 171, 255},
		color.RGBA{80, 153, 160, 255},
		color.RGBA{26, 87, 142, 255},
		color.RGBA{68, 61, 113, 255},
		color.RGBA{215, 10, 67, 255},
		color.RGBA{251, 90, 46, 255},
		color.RGBA{139, 0, 65, 255},
		color.RGBA{54, 117, 34, 255},
	}

	examples = map[string]blockdef{

		"icons": {
			"283x75x2",
			"270x45",
			"238x43x2",
			"86x32",
			"45x45x14",
			"28x20x4",
		},

		"simple": {
			"500x200",
			"250x200",
			"50x50x20",
		},

		"square": {
			"50x50x100",
		},

		"tall": {
			"50x400x2",
			"50x300x5",
			"50x200x10",
			"50x100x20",
			"50x50x40",
		},

		"wide": {
			"400x50x2",
			"300x50x5",
			"200x50x10",
			"100x50x20",
			"50x50x40",
		},

		"tallAndWide": {
			"400x100",
			"100x400",
			"400x100",
			"100x400",
			"400x100",
			"100x400",
		},

		"powersOf2": {
			"2x2x256",
			"4x4x128",
			"8x8x64",
			"16x16x32",
			"32x32x16",
			"64x64x8",
			"128x128x4",
			"256x256x2",
		},

		"oddAndEven": {
			"50x50x20",
			"47x31x20",
			"23x17x20",
			"109x42x20",
			"42x109x20",
			"17x33x20",
		},

		"complex": {
			"100x100x3",
			"60x60x3",
			"50x20x20",
			"20x50x20",
			"250x250",
			"250x100",
			"100x250",
			"400x80",
			"80x400",
			"10x10x100",
			"5x5x500",
		},
	}
)

func main() {
	for name, blocks := range examples {
		fmt.Println("==== Packing", name, "====")
		canvas := packer(blocks)
		render(name, canvas)
	}
}

// pack blocks into canvas
func packer(rects []string) *pack.Canvas {
	blocks := make(pack.Blocks, 0)
	for _, b := range rects {
		w, h, n := parse(b)
		for i := 0; i < n; i++ {
			blocks = append(blocks, &pack.Block{Width: w, Height: h})
		}
	}

	return pack.BestFit(blocks)
}

// parse blockdef "(width)x(height)" or "(width)x(height)x(num)"
func parse(block string) (int, int, int) {
	s := strings.Split(block, "x")
	n := 1
	w, _ := strconv.Atoi(s[0])
	h, _ := strconv.Atoi(s[1])

	if len(s) > 2 {
		n, _ = strconv.Atoi(s[2])
	}

	return w, h, n
}

func render(name string, c *pack.Canvas) {
	img := image.NewRGBA(image.Rect(0, 0, c.Root.Width, c.Root.Height))

	for n, s := range c.Blocks {
		rect := image.Rect(s.X, s.Y, s.X+s.Width, s.Y+s.Height)
		col := palette[n%len(palette)]
		draw.Draw(img, rect, &image.Uniform{col}, image.ZP, draw.Src)
	}

	w, _ := os.Create(name + ".png")
	defer w.Close()
	png.Encode(w, img)
}
