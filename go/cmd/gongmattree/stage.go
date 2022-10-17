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
	__Node__000000_child_1 := (&models.Node{Name: "child 1"}).Stage()
	__Node__000001_child_1_1 := (&models.Node{Name: "child 1.1"}).Stage()
	__Node__000002_child_1_1_1 := (&models.Node{Name: "child 1.1.1"}).Stage()
	__Node__000003_child_1_1_1_1 := (&models.Node{Name: "child 1.1.1.1"}).Stage()
	__Node__000004_child_1_1_1_2 := (&models.Node{Name: "child 1.1.1.2"}).Stage()
	__Node__000005_child_2 := (&models.Node{Name: "child 2"}).Stage()
	__Node__000006_child_2_2 := (&models.Node{Name: "child 2.2"}).Stage()

	// Declarations of staged instances of Tree
	__Tree__000000_Tree_1 := (&models.Tree{Name: "Tree 1"}).Stage()

	// Setup of values

	// Node child 1 values setup
	__Node__000000_child_1.Name = `child 1`
	__Node__000000_child_1.IsExpanded = true
	__Node__000000_child_1.HasCheckboxButton = true
	__Node__000000_child_1.IsChecked = true
	__Node__000000_child_1.IsCheckboxDisabled = false
	__Node__000000_child_1.HasAddChildButton = true
	__Node__000000_child_1.HasEditButton = true
	__Node__000000_child_1.IsInEditMode = false
	__Node__000000_child_1.HasDeleteButton = false

	// Node child 1.1 values setup
	__Node__000001_child_1_1.Name = `child 1.1`
	__Node__000001_child_1_1.IsExpanded = true
	__Node__000001_child_1_1.HasCheckboxButton = false
	__Node__000001_child_1_1.IsChecked = false
	__Node__000001_child_1_1.IsCheckboxDisabled = false
	__Node__000001_child_1_1.HasAddChildButton = true
	__Node__000001_child_1_1.HasEditButton = true
	__Node__000001_child_1_1.IsInEditMode = false
	__Node__000001_child_1_1.HasDeleteButton = false

	// Node child 1.1.1 values setup
	__Node__000002_child_1_1_1.Name = `child 1.1.1`
	__Node__000002_child_1_1_1.IsExpanded = true
	__Node__000002_child_1_1_1.HasCheckboxButton = false
	__Node__000002_child_1_1_1.IsChecked = false
	__Node__000002_child_1_1_1.IsCheckboxDisabled = false
	__Node__000002_child_1_1_1.HasAddChildButton = true
	__Node__000002_child_1_1_1.HasEditButton = true
	__Node__000002_child_1_1_1.IsInEditMode = false
	__Node__000002_child_1_1_1.HasDeleteButton = true

	// Node child 1.1.1.1 values setup
	__Node__000003_child_1_1_1_1.Name = `child 1.1.1.1`
	__Node__000003_child_1_1_1_1.IsExpanded = false
	__Node__000003_child_1_1_1_1.HasCheckboxButton = false
	__Node__000003_child_1_1_1_1.IsChecked = false
	__Node__000003_child_1_1_1_1.IsCheckboxDisabled = false
	__Node__000003_child_1_1_1_1.HasAddChildButton = false
	__Node__000003_child_1_1_1_1.HasEditButton = true
	__Node__000003_child_1_1_1_1.IsInEditMode = false
	__Node__000003_child_1_1_1_1.HasDeleteButton = false

	// Node child 1.1.1.2 values setup
	__Node__000004_child_1_1_1_2.Name = `child 1.1.1.2`
	__Node__000004_child_1_1_1_2.IsExpanded = false
	__Node__000004_child_1_1_1_2.HasCheckboxButton = false
	__Node__000004_child_1_1_1_2.IsChecked = false
	__Node__000004_child_1_1_1_2.IsCheckboxDisabled = false
	__Node__000004_child_1_1_1_2.HasAddChildButton = false
	__Node__000004_child_1_1_1_2.HasEditButton = true
	__Node__000004_child_1_1_1_2.IsInEditMode = false
	__Node__000004_child_1_1_1_2.HasDeleteButton = false

	// Node child 2 values setup
	__Node__000005_child_2.Name = `child 2`
	__Node__000005_child_2.IsExpanded = true
	__Node__000005_child_2.HasCheckboxButton = true
	__Node__000005_child_2.IsChecked = false
	__Node__000005_child_2.IsCheckboxDisabled = false
	__Node__000005_child_2.HasAddChildButton = true
	__Node__000005_child_2.HasEditButton = true
	__Node__000005_child_2.IsInEditMode = false
	__Node__000005_child_2.HasDeleteButton = false

	// Node child 2.2 values setup
	__Node__000006_child_2_2.Name = `child 2.2`
	__Node__000006_child_2_2.IsExpanded = false
	__Node__000006_child_2_2.HasCheckboxButton = false
	__Node__000006_child_2_2.IsChecked = false
	__Node__000006_child_2_2.IsCheckboxDisabled = false
	__Node__000006_child_2_2.HasAddChildButton = true
	__Node__000006_child_2_2.HasEditButton = true
	__Node__000006_child_2_2.IsInEditMode = false
	__Node__000006_child_2_2.HasDeleteButton = false

	// Tree Tree 1 values setup
	__Tree__000000_Tree_1.Name = `Tree 1`

	// Setup of pointers
	__Node__000000_child_1.Children = append(__Node__000000_child_1.Children, __Node__000001_child_1_1)
	__Node__000001_child_1_1.Children = append(__Node__000001_child_1_1.Children, __Node__000002_child_1_1_1)
	__Node__000002_child_1_1_1.Children = append(__Node__000002_child_1_1_1.Children, __Node__000003_child_1_1_1_1)
	__Node__000002_child_1_1_1.Children = append(__Node__000002_child_1_1_1.Children, __Node__000004_child_1_1_1_2)
	__Node__000005_child_2.Children = append(__Node__000005_child_2.Children, __Node__000006_child_2_2)
	__Tree__000000_Tree_1.RootNodes = append(__Tree__000000_Tree_1.RootNodes, __Node__000005_child_2)
	__Tree__000000_Tree_1.RootNodes = append(__Tree__000000_Tree_1.RootNodes, __Node__000000_child_1)
}


