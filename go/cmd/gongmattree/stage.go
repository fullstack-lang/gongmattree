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
	__Node__000001_Child_2 := (&models.Node{Name: "Child 2"}).Stage()
	__Node__000002_Child_3 := (&models.Node{Name: "Child 3"}).Stage()
	__Node__000003_Grand_Child_1_1 := (&models.Node{Name: "Grand Child 1.1"}).Stage()
	__Node__000004_Parent := (&models.Node{Name: "Parent"}).Stage()
	__Node__000005_Parent_2 := (&models.Node{Name: "Parent 2"}).Stage()

	// Declarations of staged instances of Tree
	__Tree__000000_Root := (&models.Tree{Name: "Root"}).Stage()

	// Setup of values

	// Node Child 1 values setup
	__Node__000000_Child_1.Name = `Child 1`
	__Node__000000_Child_1.IsExpanded = true
	__Node__000000_Child_1.HasCheckboxButton = true
	__Node__000000_Child_1.IsChecked = false

	// Node Child 2 values setup
	__Node__000001_Child_2.Name = `Child 2`
	__Node__000001_Child_2.IsExpanded = false
	__Node__000001_Child_2.HasCheckboxButton = false
	__Node__000001_Child_2.IsChecked = false

	// Node Child 3 values setup
	__Node__000002_Child_3.Name = `Child 3`
	__Node__000002_Child_3.IsExpanded = false
	__Node__000002_Child_3.HasCheckboxButton = true
	__Node__000002_Child_3.IsChecked = false

	// Node Grand Child 1.1 values setup
	__Node__000003_Grand_Child_1_1.Name = `Grand Child 1.1`
	__Node__000003_Grand_Child_1_1.IsExpanded = false
	__Node__000003_Grand_Child_1_1.HasCheckboxButton = true
	__Node__000003_Grand_Child_1_1.IsChecked = false

	// Node Parent values setup
	__Node__000004_Parent.Name = `Parent`
	__Node__000004_Parent.IsExpanded = true
	__Node__000004_Parent.HasCheckboxButton = false
	__Node__000004_Parent.IsChecked = false

	// Node Parent 2 values setup
	__Node__000005_Parent_2.Name = `Parent 2`
	__Node__000005_Parent_2.IsExpanded = false
	__Node__000005_Parent_2.HasCheckboxButton = false
	__Node__000005_Parent_2.IsChecked = false

	// Tree Root values setup
	__Tree__000000_Root.Name = `Root`

	// Setup of pointers
	__Node__000000_Child_1.Children = append(__Node__000000_Child_1.Children, __Node__000003_Grand_Child_1_1)
	__Node__000004_Parent.Children = append(__Node__000004_Parent.Children, __Node__000000_Child_1)
	__Node__000004_Parent.Children = append(__Node__000004_Parent.Children, __Node__000002_Child_3)
	__Node__000004_Parent.Children = append(__Node__000004_Parent.Children, __Node__000001_Child_2)
	__Tree__000000_Root.RootNodes = append(__Tree__000000_Root.RootNodes, __Node__000004_Parent)
	__Tree__000000_Root.RootNodes = append(__Tree__000000_Root.RootNodes, __Node__000005_Parent_2)
}


