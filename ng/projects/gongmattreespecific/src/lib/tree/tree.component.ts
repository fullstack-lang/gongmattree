import { Component, OnInit } from '@angular/core';
import { Observable, timer } from 'rxjs';

import { FlatTreeControl } from '@angular/cdk/tree';
import { MatTreeFlatDataSource, MatTreeFlattener } from '@angular/material/tree';

import * as gongmattree from 'gongmattree'

/**
 * Food data with nested structure.
 * Each node has a name and an optional list of children.
 */
interface FoodNode {
  name: string;
  children?: FoodNode[];
}

/** Flat node with expandable and level information */
interface ExampleFlatNode {
  expandable: boolean;
  name: string;
  level: number;
}

const TREE_DATA: FoodNode[] = [
  {
    name: 'Fruit',
    children: [{ name: 'Apple' }, { name: 'Banana' }, { name: 'Fruit loops' }],
  },
  {
    name: 'Vegetables',
    children: [
      {
        name: 'Green',
        children: [{ name: 'Broccoli' }, { name: 'Brussels sprouts' }],
      },
      {
        name: 'Orange',
        children: [{ name: 'Pumpkins' }, { name: 'Carrots' }],
      },
    ],
  },
];

@Component({
  selector: 'lib-tree',
  templateUrl: './tree.component.html',
  styleUrls: ['./tree.component.css']
})
export class TreeComponent implements OnInit {

  private _transformer = (node: FoodNode, level: number) => {
    return {
      expandable: !!node.children && node.children.length > 0,
      name: node.name,
      level: level,
    };
  };

  treeControl = new FlatTreeControl<ExampleFlatNode>(
    node => node.level,
    node => node.expandable,
  );

  treeFlattener = new MatTreeFlattener(
    this._transformer,
    node => node.level,
    node => node.expandable,
    node => node.children,
  );

  dataSource = new MatTreeFlatDataSource(this.treeControl, this.treeFlattener);

  hasChild = (_: number, node: ExampleFlatNode) => node.expandable;

  public gongmattreeFrontRepo?: gongmattree.FrontRepo

  constructor(
    private gongmattreeFrontRepoService: gongmattree.FrontRepoService,
    private gongmattreeCommitNbService: gongmattree.CommitNbService,
    private gongmattreePushFromFrontNbService: gongmattree.PushFromFrontNbService,

  ) {
    this.dataSource.data = TREE_DATA;
  }


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
        // this.gongmattreeCommitNbService.getCommitNb().subscribe(
        //   commitNb => {
        //     if (this.lastCommitNb < commitNb) {

        //       console.log("last commit nb " + this.lastCommitNb + " new: " + commitNb)
        //       // this.refresh()
        //       this.lastCommitNb = commitNb
        //     }
        //   }
        // )

        // see above for the explanation
        // this.gongmattreePushFromFrontNbService.getPushFromFrontNb().subscribe(
        //   pushFromFrontNb => {
        //     if (this.lastPushFromFrontNb < pushFromFrontNb) {

        //       console.log("last commit nb " + this.lastPushFromFrontNb + " new: " + pushFromFrontNb)
        //       // this.refresh()
        //       this.lastPushFromFrontNb = pushFromFrontNb
        //     }
        //   }
        // )
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
