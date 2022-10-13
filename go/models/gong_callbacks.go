package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for generic get functions
	case *Node:
		if stage.OnAfterNodeCreateCallback != nil {
			stage.OnAfterNodeCreateCallback.OnAfterCreate(stage, target)
		}
	case Tree:
	}
}

// AfterUpdateFromFront is called after a update from front
func AfterUpdateFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for generic get functions
	case *Node:
		if stage.OnAfterNodeUpdateCallback != nil {
			stage.OnAfterNodeUpdateCallback.OnAfterUpdate(stage, target)
		}
	case Tree:
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for generic get functions
	case *Node:
		if stage.OnAfterNodeDeleteCallback != nil {
			stage.OnAfterNodeDeleteCallback.OnAfterDelete(stage, target)
		}
	case Tree:
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for generic get functions
	case *Node:
		if stage.OnAfterNodeReadCallback != nil {
			stage.OnAfterNodeReadCallback.OnAfterRead(stage, target)
		}
	case Tree:
	}
}

// SetCallbackAfterUpdateFromFront is a function to set up callback that is robust to refactoring
func SetCallbackAfterUpdateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterUpdateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
	// insertion point for generic get functions
	case Node:
		stage.OnAfterNodeUpdateCallback = any(callback).(OnAfterUpdateInterface[Node])
	case Tree:
	}
}
