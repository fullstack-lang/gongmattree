import { Component, OnInit } from '@angular/core';
import { Observable, timer } from 'rxjs';

import * as gongmattree from 'gongmattree'

@Component({
  selector: 'lib-tree',
  templateUrl: './tree.component.html',
  styleUrls: ['./tree.component.css']
})
export class TreeComponent implements OnInit {


  public gongmattreeFrontRepo?: gongmattree.FrontRepo

  constructor(
    private gongmattreeFrontRepoService: gongmattree.FrontRepoService,
    private gongmattreeCommitNbService: gongmattree.CommitNbService,
    private gongmattreePushFromFrontNbService: gongmattree.PushFromFrontNbService,

  ) { }


  // the component is refreshed when modification are performed in the back repo 
  // 
  // the checkCommitNbTimer polls the commit number of the back repo
  // if the commit number has increased, it pulls the front repo and redraw the diagram

  checkCommitNbTimer: Observable<number> = timer(500, 500);
  lastCommitNb = -1
  lastPushFromFrontNb = -1
  currTime: number = 0

  ngOnInit(): void {

    this.checkCommitNbTimer.subscribe(
      currTime => {
        this.currTime = currTime

        // see above for the explanation
        this.gongmattreeCommitNbService.getCommitNb().subscribe(
          commitNb => {
            if (this.lastCommitNb < commitNb) {

              console.log("last commit nb " + this.lastCommitNb + " new: " + commitNb)
              // this.refresh()
              this.lastCommitNb = commitNb
            }
          }
        )

        // see above for the explanation
        this.gongmattreePushFromFrontNbService.getPushFromFrontNb().subscribe(
          pushFromFrontNb => {
            if (this.lastPushFromFrontNb < pushFromFrontNb) {

              console.log("last commit nb " + this.lastPushFromFrontNb + " new: " + pushFromFrontNb)
              // this.refresh()
              this.lastPushFromFrontNb = pushFromFrontNb
            }
          }
        )
      }
    )
  }

  refresh(): void {

    this.gongmattreeFrontRepoService.pull().subscribe(
      gongmattreesFrontRepo => {
        this.gongmattreeFrontRepo = gongmattreesFrontRepo

      }

    )
  }

}
