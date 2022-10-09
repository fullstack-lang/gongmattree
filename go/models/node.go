package models

type Node struct {
	Name       string
	IsExpanded bool
	Children   []*Node
	Button     *Button
}
