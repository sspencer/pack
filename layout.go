package pack

// Layout is used to sort sprites in different ways to achieve differnt packs.
//go:generate stringer -type=Layout
type Layout int

// Layout algorithms.
const (
	LayoutByWidth Layout = iota
	LayoutByHeight
	LayoutByArea
	LayoutByMax
)
