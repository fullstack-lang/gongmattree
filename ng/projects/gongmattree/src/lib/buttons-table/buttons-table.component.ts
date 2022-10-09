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
import { ButtonDB } from '../button-db'
import { ButtonService } from '../button.service'

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
  selector: 'app-buttonstable',
  templateUrl: './buttons-table.component.html',
  styleUrls: ['./buttons-table.component.css'],
})
export class ButtonsTableComponent implements OnInit {

  // mode at invocation
  mode: TableComponentMode = TableComponentMode.DISPLAY_MODE

  // used if the component is called as a selection component of Button instances
  selection: SelectionModel<ButtonDB> = new (SelectionModel)
  initialSelection = new Array<ButtonDB>()

  // the data source for the table
  buttons: ButtonDB[] = []
  matTableDataSource: MatTableDataSource<ButtonDB> = new (MatTableDataSource)

  // front repo, that will be referenced by this.buttons
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
    this.matTableDataSource.sortingDataAccessor = (buttonDB: ButtonDB, property: string) => {
      switch (property) {
        case 'ID':
          return buttonDB.ID

        // insertion point for specific sorting accessor
        case 'Name':
          return buttonDB.Name;

        default:
          console.assert(false, "Unknown field")
          return "";
      }
    };

    // enable filtering on all fields (including pointers and reverse pointer, which is not done by default)
    this.matTableDataSource.filterPredicate = (buttonDB: ButtonDB, filter: string) => {

      // filtering is based on finding a lower case filter into a concatenated string
      // the buttonDB properties
      let mergedContent = ""

      // insertion point for merging of fields
      mergedContent += buttonDB.Name.toLowerCase()

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
    private buttonService: ButtonService,
    private frontRepoService: FrontRepoService,

    // not null if the component is called as a selection component of button instances
    public dialogRef: MatDialogRef<ButtonsTableComponent>,
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
    this.buttonService.ButtonServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.getButtons()
        }
      }
    )
    if (this.mode == TableComponentMode.DISPLAY_MODE) {
      this.displayedColumns = ['ID', 'Edit', 'Delete', // insertion point for columns to display
        "Name",
      ]
    } else {
      this.displayedColumns = ['select', 'ID', // insertion point for columns to display
        "Name",
      ]
      this.selection = new SelectionModel<ButtonDB>(allowMultiSelect, this.initialSelection);
    }

  }

  ngOnInit(): void {
    this.getButtons()
    this.matTableDataSource = new MatTableDataSource(this.buttons)
  }

  getButtons(): void {
    this.frontRepoService.pull().subscribe(
      frontRepo => {
        this.frontRepo = frontRepo

        this.buttons = this.frontRepo.Buttons_array;

        // insertion point for time duration Recoveries
        // insertion point for enum int Recoveries
        
        // in case the component is called as a selection component
        if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {
          for (let button of this.buttons) {
            let ID = this.dialogData.ID
            let revPointer = button[this.dialogData.ReversePointer as keyof ButtonDB] as unknown as NullInt64
            if (revPointer.Int64 == ID) {
              this.initialSelection.push(button)
            }
            this.selection = new SelectionModel<ButtonDB>(allowMultiSelect, this.initialSelection);
          }
        }

        if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

          let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, ButtonDB>
          let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

          let sourceField = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]! as unknown as ButtonDB[]
          for (let associationInstance of sourceField) {
            let button = associationInstance[this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as ButtonDB
            this.initialSelection.push(button)
          }

          this.selection = new SelectionModel<ButtonDB>(allowMultiSelect, this.initialSelection);
        }

        // update the mat table data source
        this.matTableDataSource.data = this.buttons
      }
    )
  }

  // newButton initiate a new button
  // create a new Button objet
  newButton() {
  }

  deleteButton(buttonID: number, button: ButtonDB) {
    // list of buttons is truncated of button before the delete
    this.buttons = this.buttons.filter(h => h !== button);

    this.buttonService.deleteButton(buttonID).subscribe(
      button => {
        this.buttonService.ButtonServiceChanged.next("delete")
      }
    );
  }

  editButton(buttonID: number, button: ButtonDB) {

  }

  // display button in router
  displayButtonInRouter(buttonID: number) {
    this.router.navigate(["github_com_fullstack_lang_gongmattree_go-" + "button-display", buttonID])
  }

  // set editor outlet
  setEditorRouterOutlet(buttonID: number) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_gongmattree_go_editor: ["github_com_fullstack_lang_gongmattree_go-" + "button-detail", buttonID]
      }
    }]);
  }

  // set presentation outlet
  setPresentationRouterOutlet(buttonID: number) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_gongmattree_go_presentation: ["github_com_fullstack_lang_gongmattree_go-" + "button-presentation", buttonID]
      }
    }]);
  }

  /** Whether the number of selected elements matches the total number of rows. */
  isAllSelected() {
    const numSelected = this.selection.selected.length;
    const numRows = this.buttons.length;
    return numSelected === numRows;
  }

  /** Selects all rows if they are not all selected; otherwise clear selection. */
  masterToggle() {
    this.isAllSelected() ?
      this.selection.clear() :
      this.buttons.forEach(row => this.selection.select(row));
  }

  save() {

    if (this.mode == TableComponentMode.ONE_MANY_ASSOCIATION_MODE) {

      let toUpdate = new Set<ButtonDB>()

      // reset all initial selection of button that belong to button
      for (let button of this.initialSelection) {
        let index = button[this.dialogData.ReversePointer as keyof ButtonDB] as unknown as NullInt64
        index.Int64 = 0
        index.Valid = true
        toUpdate.add(button)

      }

      // from selection, set button that belong to button
      for (let button of this.selection.selected) {
        let ID = this.dialogData.ID as number
        let reversePointer = button[this.dialogData.ReversePointer as keyof ButtonDB] as unknown as NullInt64
        reversePointer.Int64 = ID
        reversePointer.Valid = true
        toUpdate.add(button)
      }


      // update all button (only update selection & initial selection)
      for (let button of toUpdate) {
        this.buttonService.updateButton(button)
          .subscribe(button => {
            this.buttonService.ButtonServiceChanged.next("update")
          });
      }
    }

    if (this.mode == TableComponentMode.MANY_MANY_ASSOCIATION_MODE) {

      // get the source instance via the map of instances in the front repo
      let mapOfSourceInstances = this.frontRepo[this.dialogData.SourceStruct + "s" as keyof FrontRepo] as Map<number, ButtonDB>
      let sourceInstance = mapOfSourceInstances.get(this.dialogData.ID)!

      // First, parse all instance of the association struct and remove the instance
      // that have unselect
      let unselectedButton = new Set<number>()
      for (let button of this.initialSelection) {
        if (this.selection.selected.includes(button)) {
          // console.log("button " + button.Name + " is still selected")
        } else {
          console.log("button " + button.Name + " has been unselected")
          unselectedButton.add(button.ID)
          console.log("is unselected " + unselectedButton.has(button.ID))
        }
      }

      // delete the association instance
      let associationInstance = sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]
      let button = associationInstance![this.dialogData.IntermediateStructField as keyof typeof associationInstance] as unknown as ButtonDB
      if (unselectedButton.has(button.ID)) {
        this.frontRepoService.deleteService(this.dialogData.IntermediateStruct, associationInstance)


      }

      // is the source array is empty create it
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] == undefined) {
        (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance] as unknown as Array<ButtonDB>) = new Array<ButtonDB>()
      }

      // second, parse all instance of the selected
      if (sourceInstance[this.dialogData.SourceField as keyof typeof sourceInstance]) {
        this.selection.selected.forEach(
          button => {
            if (!this.initialSelection.includes(button)) {
              // console.log("button " + button.Name + " has been added to the selection")

              let associationInstance = {
                Name: sourceInstance["Name"] + "-" + button.Name,
              }

              let index = associationInstance[this.dialogData.IntermediateStructField + "ID" as keyof typeof associationInstance] as unknown as NullInt64
              index.Int64 = button.ID
              index.Valid = true

              let indexDB = associationInstance[this.dialogData.IntermediateStructField + "DBID" as keyof typeof associationInstance] as unknown as NullInt64
              indexDB.Int64 = button.ID
              index.Valid = true

              this.frontRepoService.postService(this.dialogData.IntermediateStruct, associationInstance)

            } else {
              // console.log("button " + button.Name + " is still selected")
            }
          }
        )
      }

      // this.selection = new SelectionModel<ButtonDB>(allowMultiSelect, this.initialSelection);
    }

    // why pizza ?
    this.dialogRef.close('Pizza!');
  }
}
