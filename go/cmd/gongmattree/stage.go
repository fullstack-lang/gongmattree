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
	__Node__000001_Child_1_1 := (&models.Node{Name: "Child 1.1"}).Stage()
	__Node__000002_Child_1_1_1 := (&models.Node{Name: "Child 1.1.1"}).Stage()
	__Node__000003_Child2 := (&models.Node{Name: "Child2"}).Stage()
	__Node__000004_new_node := (&models.Node{Name: "new node"}).Stage()
	__Node__000005_new_node := (&models.Node{Name: "new node"}).Stage()
	__Node__000006_new_node := (&models.Node{Name: "new node"}).Stage()
	__Node__000007_new_node := (&models.Node{Name: "new node"}).Stage()

	// Declarations of staged instances of Tree
	__Tree__000000_Tree_1 := (&models.Tree{Name: "Tree 1"}).Stage()

	// Setup of values

	// Node Child 1 values setup
	__Node__000000_Child_1.Name = `Child 1`
	__Node__000000_Child_1.IsExpanded = true
	__Node__000000_Child_1.HasCheckboxButton = true
	__Node__000000_Child_1.IsChecked = false
	__Node__000000_Child_1.IsCheckboxDisabled = true
	__Node__000000_Child_1.HasAddChildButton = false
	__Node__000000_Child_1.HasEditButton = false

	// Node Child 1.1 values setup
	__Node__000001_Child_1_1.Name = `Child 1.1`
	__Node__000001_Child_1_1.IsExpanded = true
	__Node__000001_Child_1_1.HasCheckboxButton = true
	__Node__000001_Child_1_1.IsChecked = true
	__Node__000001_Child_1_1.IsCheckboxDisabled = false
	__Node__000001_Child_1_1.HasAddChildButton = true
	__Node__000001_Child_1_1.HasEditButton = true

	// Node Child 1.1.1 values setup
	__Node__000002_Child_1_1_1.Name = `Child 1.1.1`
	__Node__000002_Child_1_1_1.IsExpanded = false
	__Node__000002_Child_1_1_1.HasCheckboxButton = false
	__Node__000002_Child_1_1_1.IsChecked = false
	__Node__000002_Child_1_1_1.IsCheckboxDisabled = false
	__Node__000002_Child_1_1_1.HasAddChildButton = false
	__Node__000002_Child_1_1_1.HasEditButton = false

	// Node Child2 values setup
	__Node__000003_Child2.Name = `Child2`
	__Node__000003_Child2.IsExpanded = false
	__Node__000003_Child2.HasCheckboxButton = true
	__Node__000003_Child2.IsChecked = false
	__Node__000003_Child2.IsCheckboxDisabled = false
	__Node__000003_Child2.HasAddChildButton = false
	__Node__000003_Child2.HasEditButton = false

	// Node new node values setup
	__Node__000004_new_node.Name = `new node`
	__Node__000004_new_node.IsExpanded = false
	__Node__000004_new_node.HasCheckboxButton = false
	__Node__000004_new_node.IsChecked = false
	__Node__000004_new_node.IsCheckboxDisabled = false
	__Node__000004_new_node.HasAddChildButton = false
	__Node__000004_new_node.HasEditButton = false

	// Node new node values setup
	__Node__000005_new_node.Name = `new node`
	__Node__000005_new_node.IsExpanded = false
	__Node__000005_new_node.HasCheckboxButton = false
	__Node__000005_new_node.IsChecked = false
	__Node__000005_new_node.IsCheckboxDisabled = false
	__Node__000005_new_node.HasAddChildButton = false
	__Node__000005_new_node.HasEditButton = false

	// Node new node values setup
	__Node__000006_new_node.Name = `new node`
	__Node__000006_new_node.IsExpanded = false
	__Node__000006_new_node.HasCheckboxButton = false
	__Node__000006_new_node.IsChecked = false
	__Node__000006_new_node.IsCheckboxDisabled = false
	__Node__000006_new_node.HasAddChildButton = false
	__Node__000006_new_node.HasEditButton = false

	// Node new node values setup
	__Node__000007_new_node.Name = `new node`
	__Node__000007_new_node.IsExpanded = true
	__Node__000007_new_node.HasCheckboxButton = false
	__Node__000007_new_node.IsChecked = false
	__Node__000007_new_node.IsCheckboxDisabled = false
	__Node__000007_new_node.HasAddChildButton = true
	__Node__000007_new_node.HasEditButton = false

	// Tree Tree 1 values setup
	__Tree__000000_Tree_1.Name = `Tree 1`

	// Setup of pointers
	__Node__000000_Child_1.Children = append(__Node__000000_Child_1.Children, __Node__000001_Child_1_1)
	__Node__000001_Child_1_1.Children = append(__Node__000001_Child_1_1.Children, __Node__000002_Child_1_1_1)
	__Node__000001_Child_1_1.Children = append(__Node__000001_Child_1_1.Children, __Node__000007_new_node)
	__Node__000001_Child_1_1.Children = append(__Node__000001_Child_1_1.Children, __Node__000004_new_node)
	__Node__000007_new_node.Children = append(__Node__000007_new_node.Children, __Node__000005_new_node)
	__Node__000007_new_node.Children = append(__Node__000007_new_node.Children, __Node__000006_new_node)
	__Tree__000000_Tree_1.RootNodes = append(__Tree__000000_Tree_1.RootNodes, __Node__000003_Child2)
	__Tree__000000_Tree_1.RootNodes = append(__Tree__000000_Tree_1.RootNodes, __Node__000000_Child_1)
}


