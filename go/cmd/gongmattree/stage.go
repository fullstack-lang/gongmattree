package main

import (
	"time"

	"gongmattree/go/models"
)

func init() {
	var __Dummy_time_variable time.Time
	_ = __Dummy_time_variable
	InjectionGateway["stage"] = stageInjection
}

// stageInjection will stage objects of database "stage"
func stageInjection() {

	// Declaration of instances to stage

	// Declarations of staged instances of Button
	__Button__000000_Button1 := (&models.Button{Name: "Button1"}).Stage()

	// Declarations of staged instances of Node
	__Node__000000_Child_1 := (&models.Node{Name: "Child 1"}).Stage()
	__Node__000001_Child_2 := (&models.Node{Name: "Child 2"}).Stage()
	__Node__000002_Child_3 := (&models.Node{Name: "Child 3"}).Stage()
	__Node__000003_Grand_Child_1_1 := (&models.Node{Name: "Grand Child 1.1"}).Stage()
	__Node__000004_Parent := (&models.Node{Name: "Parent"}).Stage()

	// Declarations of staged instances of Tree
	__Tree__000000_Root := (&models.Tree{Name: "Root"}).Stage()

	// Setup of values

	// Button Button1 values setup
	__Button__000000_Button1.Name = `Button1`

	// Node Child 1 values setup
	__Node__000000_Child_1.Name = `Child 1`
	__Node__000000_Child_1.IsExpanded = true

	// Node Child 2 values setup
	__Node__000001_Child_2.Name = `Child 2`
	__Node__000001_Child_2.IsExpanded = false

	// Node Child 3 values setup
	__Node__000002_Child_3.Name = `Child 3`
	__Node__000002_Child_3.IsExpanded = false

	// Node Grand Child 1.1 values setup
	__Node__000003_Grand_Child_1_1.Name = `Grand Child 1.1`
	__Node__000003_Grand_Child_1_1.IsExpanded = false

	// Node Parent values setup
	__Node__000004_Parent.Name = `Parent`
	__Node__000004_Parent.IsExpanded = true

	// Tree Root values setup
	__Tree__000000_Root.Name = `Root`

	// Setup of pointers
	__Node__000000_Child_1.Children = append(__Node__000000_Child_1.Children, __Node__000003_Grand_Child_1_1)
	__Node__000000_Child_1.Button = __Button__000000_Button1
	__Node__000004_Parent.Children = append(__Node__000004_Parent.Children, __Node__000000_Child_1)
	__Node__000004_Parent.Children = append(__Node__000004_Parent.Children, __Node__000002_Child_3)
	__Node__000004_Parent.Children = append(__Node__000004_Parent.Children, __Node__000001_Child_2)
	__Tree__000000_Root.RootNode = __Node__000004_Parent
}


