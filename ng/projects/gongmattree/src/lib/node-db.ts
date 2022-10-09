// insertion point for imports
import { ButtonDB } from './button-db'

// usefull for managing pointer ID values that can be nullable
import { NullInt64 } from './null-int64'

export class NodeDB {
	CreatedAt?: string
	DeletedAt?: string
	ID: number = 0

	// insertion point for basic fields declarations
	Name: string = ""
	IsExpanded: boolean = false
	HasCheckboxButton: boolean = false

	// insertion point for other declarations
	Children?: Array<NodeDB>
	Button?: ButtonDB
	ButtonID: NullInt64 = new NullInt64 // if pointer is null, Button.ID = 0

	Node_ChildrenDBID: NullInt64 = new NullInt64
	Node_ChildrenDBID_Index: NullInt64  = new NullInt64 // store the index of the node instance in Node.Children
	Node_Children_reverse?: NodeDB 

}
