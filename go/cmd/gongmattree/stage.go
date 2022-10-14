package main

import (
	"time"

	"github.com/fullstack-lang/gongmattree/go/models"
)

func init() {
	var __Dummy_time_variable time.Time
	_ = __Dummy_time_variable
	InjectionGateway["stage"] = stageInjection
}

// stageInjection will stage objects of database "stage"
func stageInjection() {

	// Declaration of instances to stage

	// Declarations of staged instances of Node
	__Node__000000_Child_1 := (&models.Node{Name: "Child 1"}).Stage()

	// Declarations of staged instances of Tree
	__Tree__000000_Tree_1 := (&models.Tree{Name: "Tree 1"}).Stage()

	// Setup of values

	// Node Child 1 values setup
	__Node__000000_Child_1.Name = `Child 1`
	__Node__000000_Child_1.IsExpanded = false
	__Node__000000_Child_1.HasCheckboxButton = true
	__Node__000000_Child_1.IsChecked = true

	// Tree Tree 1 values setup
	__Tree__000000_Tree_1.Name = `Tree 1`

	// Setup of pointers
	__Tree__000000_Tree_1.RootNodes = append(__Tree__000000_Tree_1.RootNodes, __Node__000000_Child_1)
}


