package models

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
