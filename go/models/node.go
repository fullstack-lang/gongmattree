package models

type Node struct {
	Name              string
	IsExpanded        bool
	HasCheckboxButton bool
	Children          []*Node
	Button            *Button
}
