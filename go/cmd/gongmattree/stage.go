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
	__Node__000000_2_child := (&models.Node{Name: "2  child"}).Stage()
	__Node__000001_2_grand_child := (&models.Node{Name: "2 grand child"}).Stage()
	__Node__000002_Child_1 := (&models.Node{Name: "Child 1"}).Stage()
	__Node__000003_Child_2 := (&models.Node{Name: "Child 2"}).Stage()
	__Node__000004_Child_3 := (&models.Node{Name: "Child 3"}).Stage()
	__Node__000005_Grand_Child_1_1 := (&models.Node{Name: "Grand Child 1.1"}).Stage()
	__Node__000006_Parent := (&models.Node{Name: "Parent"}).Stage()
	__Node__000007_Parent_2 := (&models.Node{Name: "Parent 2"}).Stage()

	// Declarations of staged instances of Tree
	__Tree__000000_Tree_1 := (&models.Tree{Name: "Tree 1"}).Stage()
	__Tree__000001_Tree_2 := (&models.Tree{Name: "Tree 2"}).Stage()

	// Setup of values

	// Node 2  child values setup
	__Node__000000_2_child.Name = `2  child`
	__Node__000000_2_child.IsExpanded = true
	__Node__000000_2_child.HasCheckboxButton = true
	__Node__000000_2_child.IsChecked = true

	// Node 2 grand child values setup
	__Node__000001_2_grand_child.Name = `2 grand child`
	__Node__000001_2_grand_child.IsExpanded = false
	__Node__000001_2_grand_child.HasCheckboxButton = false
	__Node__000001_2_grand_child.IsChecked = false

	// Node Child 1 values setup
	__Node__000002_Child_1.Name = `Child 1`
	__Node__000002_Child_1.IsExpanded = true
	__Node__000002_Child_1.HasCheckboxButton = true
	__Node__000002_Child_1.IsChecked = false

	// Node Child 2 values setup
	__Node__000003_Child_2.Name = `Child 2`
	__Node__000003_Child_2.IsExpanded = false
	__Node__000003_Child_2.HasCheckboxButton = false
	__Node__000003_Child_2.IsChecked = false

	// Node Child 3 values setup
	__Node__000004_Child_3.Name = `Child 3`
	__Node__000004_Child_3.IsExpanded = false
	__Node__000004_Child_3.HasCheckboxButton = true
	__Node__000004_Child_3.IsChecked = false

	// Node Grand Child 1.1 values setup
	__Node__000005_Grand_Child_1_1.Name = `Grand Child 1.1`
	__Node__000005_Grand_Child_1_1.IsExpanded = false
	__Node__000005_Grand_Child_1_1.HasCheckboxButton = true
	__Node__000005_Grand_Child_1_1.IsChecked = false

	// Node Parent values setup
	__Node__000006_Parent.Name = `Parent`
	__Node__000006_Parent.IsExpanded = true
	__Node__000006_Parent.HasCheckboxButton = false
	__Node__000006_Parent.IsChecked = false

	// Node Parent 2 values setup
	__Node__000007_Parent_2.Name = `Parent 2`
	__Node__000007_Parent_2.IsExpanded = false
	__Node__000007_Parent_2.HasCheckboxButton = false
	__Node__000007_Parent_2.IsChecked = false

	// Tree Tree 1 values setup
	__Tree__000000_Tree_1.Name = `Tree 1`

	// Tree Tree 2 values setup
	__Tree__000001_Tree_2.Name = `Tree 2`

	// Setup of pointers
	__Node__000000_2_child.Children = append(__Node__000000_2_child.Children, __Node__000001_2_grand_child)
	__Node__000002_Child_1.Children = append(__Node__000002_Child_1.Children, __Node__000005_Grand_Child_1_1)
	__Node__000006_Parent.Children = append(__Node__000006_Parent.Children, __Node__000002_Child_1)
	__Node__000006_Parent.Children = append(__Node__000006_Parent.Children, __Node__000004_Child_3)
	__Node__000006_Parent.Children = append(__Node__000006_Parent.Children, __Node__000003_Child_2)
	__Tree__000000_Tree_1.RootNodes = append(__Tree__000000_Tree_1.RootNodes, __Node__000006_Parent)
	__Tree__000000_Tree_1.RootNodes = append(__Tree__000000_Tree_1.RootNodes, __Node__000007_Parent_2)
	__Tree__000001_Tree_2.RootNodes = append(__Tree__000001_Tree_2.RootNodes, __Node__000000_2_child)
}


