// generated by gong
import { Component, OnInit, AfterViewInit, ViewChild, Inject, Optional } from '@angular/core';
import { BehaviorSubject } from 'rxjs'
import { MatSort } from '@angular/material/sort';
import { MatPaginator } from '@angular/material/paginator';
import { MatTableDataSource } from '@angular/material/table';
import { MatButton } from '@angular/material/button'

import { MatDialogRef, MAT_DIALOG_DATA, MatDialog } from '@angular/material/dialog'
import { DialogData, FrontRepoService, FrontRepo, SelectionMode } from '../front-repo.service'
import { NullInt64 } from '../null-int64'
import { SelectionModel } from '@angular/cdk/collections';

const allowMultiSelect = true;

import { Router, RouterState } from '@angular/router';
import { NodeDB } from '../node-db'
import { NodeService } from '../node.service'

// insertion point for additional imports

// TableComponent is initilizaed from different routes
// TableComponentMode detail different cases 
enum TableComponentMode {
  DISPLAY_MODE,
  ONE_MANY_ASSOCIATION_MODE,
  MANY_MANY_ASSOCIATION_MODE,
}

// generated table component
@Component({
  selector: 'app-nodestable',
  templateUrl: './nodes-table.component.html',
  styleUrls: ['./nodes-table.component.css'],
})
export class NodesTableComponent implements OnInit {

  // mode at invocation
  mode: TableComponentMode = TableComponentMode.DISPLAY_MODE

  // used if the component is called as a selection component of Node instances
  selection: SelectionModel<NodeDB> = new (SelectionModel)
  initialSelection = new Array<NodeDB>()

  // the data source for the table
  nodes: NodeDB[] = []
  matTableDataSource: MatTableDataSource<NodeDB> = new (MatTableDataSource)

  // front repo, that will be referenced by this.nodes
  frontRepo: FrontRepo = new (FrontRepo)

  // displayedColumns is referenced by the MatTable component for specify what columns
  // have to be displayed and in what order
  displayedColumns: string[];

  // for sorting & pagination
  @ViewChild(MatSort)
  sort: MatSort | undefined
  @ViewChild(MatPaginator)
  paginator: MatPaginator | undefined;

  ngAfterViewInit() {

    // enable sorting on all fields (including pointers and reverse pointer)
    this.matTableDataSource.sortingDataAccessor = (nodeDB: NodeDB, property: string) => {
      switch (property) {
        case 'ID':
          return nodeDB.ID

        // insertion point for specific sorting accessor
        case 'Name':
          return nodeDB.Name;

        case 'IsExpanded':
          return nodeDB.IsExpanded?"true":"false";

        case 'HasCheckboxButton':
          return nodeDB.HasCheckboxButton?"true":"false";

        case 'IsChecked':
          return nodeDB.IsChecked?"true":"false";

        case 'IsCheckboxDisabled':
          return nodeDB.IsCheckboxDisabled?"true":"false";

        case 'HasAddChildButton':
          return nodeDB.HasAddChildButton?"true":"false";

        case 'HasEditButton':
          return nodeDB.HasEditButton?"true":"false";

        case 'IsInEditMode':
          return nodeDB.IsInEditMode?"true":"false";

        case 'HasDeleteButton':
          return nodeDB.HasDeleteButton?"true":"false";

        case 'Node_Children':
          if (this.frontRepo.Nodes.get(nodeDB.Node_ChildrenDBID.Int64) != undefined) {
            return this.frontRepo.Nodes.get(nodeDB.Node_ChildrenDBID.Int64)!.Name
          } else {
            return ""
          }

        case 'Tree_RootNodes':
          if (this.frontRepo.Trees.get(nodeDB.Tree_RootNodesDBID.Int64) != undefined) {
            return this.frontRepo.Trees.get(nodeDB.Tree_RootNodesDBID.Int64)!.Name
          } else {
            return ""
          }

        default:
          console.assert(false, "Unknown field")
          return "";
      }
    };

    // enable filtering on all fields (including pointers and reverse pointer, which is not done by default)
    this.matTableDataSource.filterPredicate = (nodeDB: NodeDB, filter: string) => {

      // filtering is based on finding a lower case filter into a concatenated string
      // the nodeDB properties
      let mergedContent = ""

      // insertion point for merging of fields
      mergedContent += nodeDB.Name.toLowerCase()
      if (nodeDB.Node_ChildrenDBID.Int64 != 0) {
        mergedContent += this.frontRepo.Nodes.get(nodeDB.Node_ChildrenDBID.Int64)!.Name.toLowerCase()
      }

      if (nodeDB.Tree_RootNodesDBID.Int64 != 0) {
        mergedContent += this.frontRepo.Trees.get(nodeDB.Tree_RootNodesDBID.Int64)!.Name.toLowerCase()
      }


      let isSelected = mergedContent.includes(filter.toLowerCase())
      return isSelected
    };

    this.matTableDataSource.sort = this.sort!
    this.matTableDataSource.paginator = this.paginator!
  }

  applyFilter(event: Event) {
    const filterValue = (event.target as HTMLInputElement).value;
    this.matTableDataSource.filter = filterValue.trim().toLowerCase();
  }

  constructor(
    private nodeService: NodeService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of node instances
    public dialogRef: MatDialogRef<NodesTableComponent>,
    @Optional() @Inject(MAT_DIALOG_DATA) public dialogData: DialogData,

    private router: Router,
  ) {

    // compute mode
    if (dialogData == undefined) {
      this.mode = TableComponentMode.DISPLAY_MODE
    } else {
      switch (dialogData.SelectionMode) {
        case SelectionMode.ONE_MANY_ASSOCIATION_MODE:
          this.mode = TableComponentMode.ONE_MANY_ASSOCIATION_MODE
          break
        case SelectionMode.MANY_MANY_ASSOCIATION_MODE:
          this.mode = TableComponentMode.MANY_MANY_ASSOCIATION_MODE
          break
        default:
      }
    }

    // observable for changes in structs
    this.nodeService.NodeServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getNodes()
        }
      }
    )
    if (this.mode == TableComponentMode.DISPLAY_MODE) {
      this.displayedColumns = ['ID', 'Edit', 'Delete', // insertion point for columns to display
        "Name",
        "IsExpanded",
        "HasCheckboxButton",
        "IsChecked",
        "IsCheckboxDisabled",
        "HasAddChildButton",
        "HasEditButton",
        "IsInEditMode",
        "HasDeleteButton",
        "Node_Children",
        "Tree_RootNodes",
      ]
    } else {
      this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
        "IsExpanded",
        "HasCheckboxButton",
        "IsChecked",
        "IsCheckboxDisabled",
        "HasAddChildButton",
        "HasEditButton",
        "IsInEditMode",
        "HasDeleteButton",
        "Node_Children",
        "Tree_RootNodes",
      ]
      this.selection = new SelectionModel<NodeDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    this.getNodes()
    this.matTableDataSource = new MatTableDataSource(this.nodes)
  }

  getNodes(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        this.nodes = this.frontRepo.Nodes_array;

        // insertion point for time duration Recoveries
        // insertion point for enum int Recoveries
        
        // in case the component is called as a selection component
        if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {
          for (let node of this.nodes) {
            let ID = this.dialogData.ID
            let revPointer = node[this.dialogData.ReversePointer as keyof NodeDB] as unknown as NullInt64
            if (revPointer.Int64 == ID) {
              this.initialSelection.push(node)
            }
            this.selection = new SelectionModel<NodeDB>(allowMultiSelect, this.initialSelection);
          }
        }

        if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

          let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, NodeDB>
          let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

          let sourceField = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]! as unknown as NodeDB[]
          for (let associationInstance of sourceField) {
            let node = associationInstance[this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as NodeDB
            this.initialSelection.push(node)
          }

          this.selection = new SelectionModel<NodeDB>(allowMultiSelect, this.initialSelection);
        }

        // update the mat table data source
        this.matTableDataSource.data = this.nodes
      }
    )
  }

  // newNode initiate a new node
  // create a new Node objet
  newNode() {
  }

  deleteNode(nodeID: number, node: NodeDB) {
    // list of nodes is truncated of node before the delete
    this.nodes = this.nodes.filter(h => h !== node);

    this.nodeService.deleteNode(nodeID).subscribe(
      node => {
        this.nodeService.NodeServiceChanged.next("delete")
      }
    );
  }

  editNode(nodeID: number, node: NodeDB) {

  }

  // display node in router
  displayNodeInRouter(nodeID: number) {
    this.router.navigate(["github_com_fullstack_lang_gongmattree_go-" + "node-display", nodeID])
  }

  // set editor outlet
  setEditorRouterOutlet(nodeID: number) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_gongmattree_go_editor: ["github_com_fullstack_lang_gongmattree_go-" + "node-detail", nodeID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(nodeID: number) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_gongmattree_go_presentation: ["github_com_fullstack_lang_gongmattree_go-" + "node-presentation", nodeID]
      }
    }]);
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.nodes.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.nodes.forEach(row => this.selection.select(row));
  }

  save() {

    if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {

      let toUpdate = new Set<NodeDB>()

      // reset all initial selection of node that belong to node
      for (let node of this.initialSelection) {
        let index = node[this.dialogData.ReversePointer as keyof NodeDB] as unknown as NullInt64
        index.Int64 = 0
        index.Valid = true
        toUpdate.add(node)

      }

      // from selection, set node that belong to node
      for (let node of this.selection.selected) {
        let ID = this.dialogData.ID as number
        let reversePointer = node[this.dialogData.ReversePointer as keyof NodeDB] as unknown as NullInt64
        reversePointer.Int64 = ID
        reversePointer.Valid = true
        toUpdate.add(node)
      }


      // update all node (only update selection & initial selection)
      for (let node of toUpdate) {
        this.nodeService.updateNode(node)
          .subscribe(node => {
            this.nodeService.NodeServiceChanged.next("update")
          });
      }
    }

    if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

      // get the source instance via the map of instances in the front repo
      let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, NodeDB>
      let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

      // First, parse all instance of the association struct and remove the instance
      // that have unselect
      let unselectedNode = new Set<number>()
      for (let node of this.initialSelection) {
        if (this.selection.selected.includes(node)) {
          // console.log("node " + node.Name + " is still selected")
        } else {
          console.log("node " + node.Name + " has been unselected")
          unselectedNode.add(node.ID)
          console.log("is unselected " + unselectedNode.has(node.ID))
        }
      }

      // delete the association instance
      let associationInstance = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]
      let node = associationInstance![this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as NodeDB
      if (unselectedNode.has(node.ID)) {
        this.frontRepoService.deleteService(this.dialogData.IntermediateStruct, associationInstance)


      }

      // is the source array is empty create it
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] == undefined) {
        (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] as unknown as Array<NodeDB>) = new Array<NodeDB>()
      }

      // second, parse all instance of the selected
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]) {
        this.selection.selected.forEach(
          node => {
            if (!this.initialSelection.includes(node)) {
              // console.log("node " + node.Name + " has been added to the selection")

              let associationInstance = {
                Name: sourceInstance["Name"] + "-" + node.Name,
              }

              let index = associationInstance[this.dialogData.IntermediateStructField + "ID" as keyof typeof associationInstance] as unknown as NullInt64
              index.Int64 = node.ID
              index.Valid = true

              let indexDB = associationInstance[this.dialogData.IntermediateStructField + "DBID" as keyof typeof associationInstance] as unknown as NullInt64
              indexDB.Int64 = node.ID
              index.Valid = true

              this.frontRepoService.postService(this.dialogData.IntermediateStruct, associationInstance)

            } else {
              // console.log("node " + node.Name + " is still selected")
            }
          }
        )
      }

      // this.selection = new SelectionModel<NodeDB>(allowMultiSelect, this.initialSelection);
    }

    // why pizza ?
    this.dialogRef.close('Pizza!');
  }
}
