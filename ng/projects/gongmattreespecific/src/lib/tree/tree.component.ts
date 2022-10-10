import { Component, OnInit } from '@angular/core';
import { Observable, timer } from 'rxjs';

import { FlatTreeControl } from '@angular/cdk/tree';
import { MatTreeFlatDataSource, MatTreeFlattener } from '@angular/material/tree';

import * as gongmattree from 'gongmattree'

/**
 * Food data with nested structure.
 * Each node has a name and an optional list of children.
 */
interface Node {
  name: string;

  gongNode: gongmattree.NodeDB;
  children?: Node[];
}

/** Flat node with expandable and level information */
interface FlatNode {
  expandable: boolean;

  name: string;
  gongNode: gongmattree.NodeDB;
  level: number;
}

@Component({
  selector: 'lib-tree',
  templateUrl: './tree.component.html',
  styleUrls: ['./tree.component.css']
})
export class TreeComponent implements OnInit {

  private _transformer = (node: Node, level: number) => {
    return {
      expandable: !!node.children && node.children.length > 0,

      gongNode: node.gongNode,
      name: node.name,
      level: level,
    };
  };

  treeControl = new FlatTreeControl<FlatNode>(
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

  hasChild = (_: number, node: FlatNode) => node.expandable;

  public gongmattreeFrontRepo?: gongmattree.FrontRepo

  constructor(
    private gongmattreeFrontRepoService: gongmattree.FrontRepoService,
    private gongmattreeCommitNbService: gongmattree.CommitNbService,
    private gongmattreePushFromFrontNbService: gongmattree.PushFromFrontNbService,
    private gongmattreeNodeService: gongmattree.NodeService,

  ) {
    // this.dataSource.data = TREE_DATA;
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
        this.gongmattreeCommitNbService.getCommitNb().subscribe(
          commitNb => {
            if (this.lastCommitNb < commitNb) {

              console.log("last commit nb " + this.lastCommitNb + " new: " + commitNb)
              this.refresh()
              this.lastCommitNb = commitNb
            }
          }
        )

        // see above for the explanation
        this.gongmattreePushFromFrontNbService.getPushFromFrontNb().subscribe(
          pushFromFrontNb => {
            if (this.lastPushFromFrontNb < pushFromFrontNb) {

              console.log("last commit nb " + this.lastPushFromFrontNb + " new: " + pushFromFrontNb)
              this.refresh()
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


        if (this.gongmattreeFrontRepo.Trees_array.length != 1) {
          console.log("error: there should be exactly one tree")
          return
        }
        var treeSingloton = this.gongmattreeFrontRepo.Trees_array[0]

        for (var rootNode of treeSingloton.RootNodes!) {
          this.dataSource.data.concat(this.gongNodeToMatTreeNode(rootNode))
        }

        // expand nodes that were exapanded before
        this.treeControl.dataNodes?.forEach(
          node => {
            if (node.gongNode.IsExpanded) {
              this.treeControl.expand(node)
            }
          }
        )

      }
    )
  }

  gongNodeToMatTreeNode(nodeDB: gongmattree.NodeDB): Node {
    var matTreeNode: Node = { name: nodeDB.Name, gongNode: nodeDB, children: [] }
    if (nodeDB.Children != undefined) {
      matTreeNode.children = nodeDB.Children.map(child => this.gongNodeToMatTreeNode(child))
    }

    return matTreeNode
  }

  toggleNodeExpansion(node: FlatNode): void {
    console.log(node.name)

    node.gongNode.IsExpanded = !node.gongNode.IsExpanded

    this.gongmattreeNodeService.updateNode(node.gongNode).subscribe(
      gongmattreeNode => {
        console.log("updated node")
      }
    )
  }

  toggleNodeCheckbox(node: FlatNode): void {
    console.log(node.name)

    node.gongNode.IsChecked = !node.gongNode.IsChecked

    this.gongmattreeNodeService.updateNode(node.gongNode).subscribe(
      gongmattreeNode => {
        console.log("updated node")
      }
    )
  }
}
