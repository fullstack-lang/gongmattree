import { Component, OnInit } from '@angular/core';

import { Observable, timer } from 'rxjs'

import * as gongdoc from 'gongdoc'
import * as gongmattree from 'gongmattree'

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
})
export class AppComponent implements OnInit {

  // choices for the top radio button
  view = 'Tree'
  tree = 'Tree'
  default = 'Default view'
  diagrams = 'UML view'
  meta = 'Meta view'
  views: string[] = [this.tree, this.default, this.diagrams, this.meta];

  // variable that enables pooling of selected gongstruct
  obsTimer: Observable<number> = timer(1000, 1000)
  gongdocCommandSingloton: gongdoc.GongdocCommandDB = new gongdoc.GongdocCommandDB
  lastSelectionDate: string = ''

  constructor(private gongdocCommandService: gongdoc.GongdocCommandService,
    private gongstructSelectionService: gongmattree.GongstructSelectionService
  ) {

  }

  ngOnInit(): void {

    // pool the gongdoc command and check wether a gongstruct has been selected
    this.obsTimer.subscribe(
      currTime => {

        // fetch the command singloton
        this.gongdocCommandService.getGongdocCommands().subscribe(
          gongdocCommands => {
            for (let gongdocCommand_ of gongdocCommands) {
              this.gongdocCommandSingloton = gongdocCommand_
            }

            // check the type of command
            if (this.gongdocCommandSingloton.Command == gongdoc.GongdocCommandType.DIAGRAM_GONGSTRUCT_SELECT) {
              if (this.lastSelectionDate != this.gongdocCommandSingloton.Date) {
                // console.log("New user selection of gongstruct " + this.gongdocCommandSingloton.StructName)

                this.gongstructSelectionService.gongstructSelected(this.gongdocCommandSingloton.StructName)
                this.lastSelectionDate = this.gongdocCommandSingloton.Date
              }
            }
          }
        )
      }
    )
  }
}
