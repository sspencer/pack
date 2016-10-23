package pack

import (
	"sort"
	"testing"
)

func getBlocks() Blocks {
	blocks := Blocks{
		&Block{Width: 5, Height: 30},
		&Block{Width: 20, Height: 15},
		&Block{Width: 25, Height: 10},
	}

	return blocks
}

func TestSortByWidth(t *testing.T) {
	blocks := getBlocks()
	sort.Sort(BlocksByWidth(blocks))
	if blocks[0].Width != 25 {
		t.Error("Sort failed - first item")
	}

	if blocks[2].Width != 5 {
		t.Error("Sort failed - last item")
	}
}

func TestSortByHeight(t *testing.T) {
	blocks := getBlocks()
	sort.Sort(BlocksByHeight(blocks))
	if blocks[0].Height != 30 {
		t.Error("Sort failed - first item")
	}

	if blocks[2].Height != 10 {
		t.Error("Sort failed - last item")
	}
}

func TestSortByArea(t *testing.T) {
	blocks := getBlocks()
	sort.Sort(BlocksByArea(blocks))
	if blocks[0].Width != 20 {
		t.Error("Sort failed - first item")
	}

	if blocks[2].Width != 5 {
		t.Error("Sort failed - last item")
	}
}

func TestSortByMax(t *testing.T) {
	blocks := getBlocks()
	sort.Sort(BlocksByMax(blocks))
	if blocks[0].Height != 30 {
		t.Error("Sort failed - first item")
	}

	if blocks[1].Width != 25 {
		t.Error("Sort failed - last item")
	}
}
