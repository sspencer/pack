package pack

import (
	"fmt"
	"testing"
)

func TestLayoutByWidth(t *testing.T) {

	blocks := Blocks{
		&Block{Name: "wide", Width: 20, Height: 25},
		&Block{Name: "widest", Width: 60, Height: 15},
		&Block{Name: "wider", Width: 40, Height: 5},
	}

	canvas := Fit(blocks, LayoutByWidth)
	fmt.Println(canvas)

	if canvas.Blocks[0].Name != "widest" || canvas.Blocks[1].Name != "wider" || canvas.Blocks[2].Name != "wide" {
		t.Error("Canvas not laid out by width")
	}
}

func TestLayoutByHeight(t *testing.T) {

	blocks := Blocks{
		&Block{Name: "high", Width: 40, Height: 5},
		&Block{Name: "highest", Width: 20, Height: 25},
		&Block{Name: "higher", Width: 60, Height: 15},
	}

	canvas := Fit(blocks, LayoutByHeight)
	fmt.Println(canvas)

	if canvas.Blocks[0].Name != "highest" || canvas.Blocks[1].Name != "higher" || canvas.Blocks[2].Name != "high" {
		t.Error("Canvas not laid out by height")
	}
}

func TestLayoutByArea(t *testing.T) {

	blocks := Blocks{
		&Block{Name: "bigger", Width: 30, Height: 15},
		&Block{Name: "biggest", Width: 20, Height: 25},
		&Block{Name: "big", Width: 60, Height: 5},
	}

	canvas := Fit(blocks, LayoutByArea)
	fmt.Println(canvas)

	if canvas.Blocks[0].Name != "biggest" || canvas.Blocks[1].Name != "bigger" || canvas.Blocks[2].Name != "big" {
		t.Error("Canvas not laid out by area")
	}
}

func TestLayoutByMax(t *testing.T) {

	blocks := Blocks{
		&Block{Name: "min", Width: 50, Height: 50},
		&Block{Name: "max", Width: 20, Height: 70},
		&Block{Name: "avg", Width: 60, Height: 5},
	}

	canvas := Fit(blocks, LayoutByMax)
	fmt.Println(canvas)

	if canvas.Blocks[0].Name != "max" || canvas.Blocks[1].Name != "avg" || canvas.Blocks[2].Name != "min" {
		t.Error("Canvas not laid out by max")
	}
}

func TestLayoutByBest(t *testing.T) {

	blocks := Blocks{
		&Block{Name: "2nd", Width: 50, Height: 50},
		&Block{Name: "3rd", Width: 39, Height: 60},
		&Block{Name: "1st", Width: 60, Height: 40},
	}

	canvas := BestFit(blocks)
	fmt.Println(canvas)

	if canvas.Blocks[0].Name != "1st" || canvas.Blocks[1].Name != "2nd" || canvas.Blocks[2].Name != "3rd" {
		t.Error("Canvas not laid out best")
	}
}
