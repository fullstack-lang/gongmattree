import { Component, OnInit } from '@angular/core';
import { Router, RouterState } from '@angular/router';

import { BehaviorSubject, Subscription } from 'rxjs';

import { FlatTreeControl } from '@angular/cdk/tree';
import { MatTreeFlatDataSource, MatTreeFlattener } from '@angular/material/tree';

import { FrontRepoService, FrontRepo } from '../front-repo.service'
import { CommitNbService } from '../commitnb.service'
import { GongstructSelectionService } from '../gongstruct-selection.service'

// insertion point for per struct import code
import { ButtonService } from '../button.service'
import { getButtonUniqueID } from '../front-repo.service'
import { NodeService } from '../node.service'
import { getNodeUniqueID } from '../front-repo.service'
import { TreeService } from '../tree.service'
import { getTreeUniqueID } from '../front-repo.service'

/**
 * Types of a GongNode / GongFlatNode
 */
export enum GongNodeType {
  STRUCT = "STRUCT",
  INSTANCE = "INSTANCE",
  ONE__ZERO_ONE_ASSOCIATION = 'ONE__ZERO_ONE_ASSOCIATION',
  ONE__ZERO_MANY_ASSOCIATION = 'ONE__ZERO_MANY_ASSOCIATION',
}

/**
 * GongNode is the "data" node
 */
interface GongNode {
  name: string; // if STRUCT, the name of the struct, if INSTANCE the name of the instance
  children: GongNode[];
  type: GongNodeType;
  structName: string;
  associationField: string;
  associatedStructName: string;
  id: number;
  uniqueIdPerStack: number;
}


/** 
 * GongFlatNode is the dynamic visual node with expandable and level information
 * */
interface GongFlatNode {
  expandable: boolean;
  name: string;
  level: number;
  type: GongNodeType;
  structName: string;
  associationField: string;
  associatedStructName: string;
  id: number;
  uniqueIdPerStack: number;
}


@Component({
  selector: 'app-gongmattree-sidebar',
  templateUrl: './sidebar.component.html',
  styleUrls: ['./sidebar.component.css'],
})
export class SidebarComponent implements OnInit {

  /**
  * _transformer generated a displayed node from a data node
  *
  * @param node input data noe
  * @param level input level
  *
  * @returns an ExampleFlatNode
  */
  private _transformer = (node: GongNode, level: number) => {
    return {

      /**
      * in javascript, The !! ensures the resulting type is a boolean (true or false).
      *
      * !!node.children will evaluate to true is the variable is defined
      */
      expandable: !!node.children && node.children.length > 0,
      name: node.name,
      level: level,
      type: node.type,
      structName: node.structName,
      associationField: node.associationField,
      associatedStructName: node.associatedStructName,
      id: node.id,
      uniqueIdPerStack: node.uniqueIdPerStack,
    }
  }

  /**
   * treeControl is passed as the paramter treeControl in the "mat-tree" selector
   *
   * Flat tree control. Able to expand/collapse a subtree recursively for flattened tree.
   *
   * Construct with flat tree data node functions getLevel and isExpandable.
  constructor(
    getLevel: (dataNode: T) => number,
    isExpandable: (dataNode: T) => boolean, 
    options?: FlatTreeControlOptions<T, K> | undefined);
   */
  treeControl = new FlatTreeControl<GongFlatNode>(
    node => node.level,
    node => node.expandable
  );

  /**
   * from mat-tree documentation
   *
   * Tree flattener to convert a normal type of node to node with children & level information.
   */
  treeFlattener = new MatTreeFlattener(
    this._transformer,
    node => node.level,
    node => node.expandable,
    node => node.children
  );

  /**
   * data is the other paramter to the "mat-tree" selector
   * 
   * strangely, the dataSource declaration has to follow the treeFlattener declaration
   */
  dataSource = new MatTreeFlatDataSource(this.treeControl, this.treeFlattener);

  /**
   * hasChild is used by the selector for expandable nodes
   * 
   *  <mat-tree-node *matTreeNodeDef="let node;when: hasChild" matTreeNodePadding>
   * 
   * @param _ 
   * @param node 
   */
  hasChild = (_: number, node: GongFlatNode) => node.expandable;

  // front repo
  frontRepo: FrontRepo = new (FrontRepo)
  commitNb: number = 0

  // "data" tree that is constructed during NgInit and is passed to the mat-tree component
  gongNodeTree = new Array<GongNode>();

  // SelectedStructChanged is the behavior subject that will emit
  // the selected gong struct whose table has to be displayed in the table outlet
  SelectedStructChanged: BehaviorSubject<string> = new BehaviorSubject("");

  subscription: Subscription = new Subscription

  constructor(
    private router: Router,
    private frontRepoService: FrontRepoService,
    private commitNbService: CommitNbService,
    private gongstructSelectionService: GongstructSelectionService,

    // insertion point for per struct service declaration
    private buttonService: ButtonService,
    private nodeService: NodeService,
    private treeService: TreeService,
  ) { }

  ngOnDestroy() {
    // prevent memory leak when component destroyed
    this.subscription.unsubscribe();
  }

  ngOnInit(): void {

    this.subscription = this.gongstructSelectionService.gongtructSelected$.subscribe(
      gongstructName => {
        // console.log("sidebar gongstruct selected " + gongstructName)

        this.setTableRouterOutlet(gongstructName.toLowerCase() + "s")
      });

    this.refresh()

    this.SelectedStructChanged.subscribe(
      selectedStruct => {
        if (selectedStruct != "") {
          this.setTableRouterOutlet(selectedStruct.toLowerCase() + "s")
        }
      }
    )

    // insertion point for per struct observable for refresh trigger
    // observable for changes in structs
    this.buttonService.ButtonServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.nodeService.NodeServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
    // observable for changes in structs
    this.treeService.TreeServiceChanged.subscribe(
      message => {
        if (message == "post" || message == "update" || message == "delete") {
          this.refresh()
        }
      }
    )
  }

  refresh(): void {
    this.frontRepoService.pull().subscribe(frontRepo => {
      this.frontRepo = frontRepo

      // use of a Gödel number to uniquely identfy nodes : 2 * node.id + 3 * node.level
      let memoryOfExpandedNodes = new Map<number, boolean>()
      let nonInstanceNodeId = 1

      this.treeControl.dataNodes?.forEach(
        node => {
          if (this.treeControl.isExpanded(node)) {
            memoryOfExpandedNodes.set(node.uniqueIdPerStack, true)
          } else {
            memoryOfExpandedNodes.set(node.uniqueIdPerStack, false)
          }
        }
      )

      // reset the gong node tree
      this.gongNodeTree = new Array<GongNode>();

      // insertion point for per struct tree construction
      /**
      * fill up the Button part of the mat tree
      */
      let buttonGongNodeStruct: GongNode = {
        name: "Button",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Button",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(buttonGongNodeStruct)

      this.frontRepo.Buttons_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Buttons_array.forEach(
        buttonDB => {
          let buttonGongNodeInstance: GongNode = {
            name: buttonDB.Name,
            type: GongNodeType.INSTANCE,
            id: buttonDB.ID,
            uniqueIdPerStack: getButtonUniqueID(buttonDB.ID),
            structName: "Button",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          buttonGongNodeStruct.children!.push(buttonGongNodeInstance)

          // insertion point for per field code
        }
      )

      /**
      * fill up the Node part of the mat tree
      */
      let nodeGongNodeStruct: GongNode = {
        name: "Node",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Node",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(nodeGongNodeStruct)

      this.frontRepo.Nodes_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Nodes_array.forEach(
        nodeDB => {
          let nodeGongNodeInstance: GongNode = {
            name: nodeDB.Name,
            type: GongNodeType.INSTANCE,
            id: nodeDB.ID,
            uniqueIdPerStack: getNodeUniqueID(nodeDB.ID),
            structName: "Node",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          nodeGongNodeStruct.children!.push(nodeGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the slide of pointer Children
          */
          let ChildrenGongNodeAssociation: GongNode = {
            name: "(Node) Children",
            type: GongNodeType.ONE__ZERO_MANY_ASSOCIATION,
            id: nodeDB.ID,
            uniqueIdPerStack: 19 * nonInstanceNodeId,
            structName: "Node",
            associationField: "Children",
            associatedStructName: "Node",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          nodeGongNodeInstance.children.push(ChildrenGongNodeAssociation)

          nodeDB.Children?.forEach(nodeDB => {
            let nodeNode: GongNode = {
              name: nodeDB.Name,
              type: GongNodeType.INSTANCE,
              id: nodeDB.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                7 * getNodeUniqueID(nodeDB.ID)
                + 11 * getNodeUniqueID(nodeDB.ID),
              structName: "Node",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            ChildrenGongNodeAssociation.children.push(nodeNode)
          })

          /**
          * let append a node for the association Button
          */
          let ButtonGongNodeAssociation: GongNode = {
            name: "(Button) Button",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: nodeDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "Node",
            associationField: "Button",
            associatedStructName: "Button",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          nodeGongNodeInstance.children!.push(ButtonGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation Button
            */
          if (nodeDB.Button != undefined) {
            let nodeGongNodeInstance_Button: GongNode = {
              name: nodeDB.Button.Name,
              type: GongNodeType.INSTANCE,
              id: nodeDB.Button.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getNodeUniqueID(nodeDB.ID)
                + 5 * getButtonUniqueID(nodeDB.Button.ID),
              structName: "Button",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            ButtonGongNodeAssociation.children.push(nodeGongNodeInstance_Button)
          }

        }
      )

      /**
      * fill up the Tree part of the mat tree
      */
      let treeGongNodeStruct: GongNode = {
        name: "Tree",
        type: GongNodeType.STRUCT,
        id: 0,
        uniqueIdPerStack: 13 * nonInstanceNodeId,
        structName: "Tree",
        associationField: "",
        associatedStructName: "",
        children: new Array<GongNode>()
      }
      nonInstanceNodeId = nonInstanceNodeId + 1
      this.gongNodeTree.push(treeGongNodeStruct)

      this.frontRepo.Trees_array.sort((t1, t2) => {
        if (t1.Name > t2.Name) {
          return 1;
        }
        if (t1.Name < t2.Name) {
          return -1;
        }
        return 0;
      });

      this.frontRepo.Trees_array.forEach(
        treeDB => {
          let treeGongNodeInstance: GongNode = {
            name: treeDB.Name,
            type: GongNodeType.INSTANCE,
            id: treeDB.ID,
            uniqueIdPerStack: getTreeUniqueID(treeDB.ID),
            structName: "Tree",
            associationField: "",
            associatedStructName: "",
            children: new Array<GongNode>()
          }
          treeGongNodeStruct.children!.push(treeGongNodeInstance)

          // insertion point for per field code
          /**
          * let append a node for the association RootNode
          */
          let RootNodeGongNodeAssociation: GongNode = {
            name: "(Node) RootNode",
            type: GongNodeType.ONE__ZERO_ONE_ASSOCIATION,
            id: treeDB.ID,
            uniqueIdPerStack: 17 * nonInstanceNodeId,
            structName: "Tree",
            associationField: "RootNode",
            associatedStructName: "Node",
            children: new Array<GongNode>()
          }
          nonInstanceNodeId = nonInstanceNodeId + 1
          treeGongNodeInstance.children!.push(RootNodeGongNodeAssociation)

          /**
            * let append a node for the instance behind the asssociation RootNode
            */
          if (treeDB.RootNode != undefined) {
            let treeGongNodeInstance_RootNode: GongNode = {
              name: treeDB.RootNode.Name,
              type: GongNodeType.INSTANCE,
              id: treeDB.RootNode.ID,
              uniqueIdPerStack: // godel numbering (thank you kurt)
                3 * getTreeUniqueID(treeDB.ID)
                + 5 * getNodeUniqueID(treeDB.RootNode.ID),
              structName: "Node",
              associationField: "",
              associatedStructName: "",
              children: new Array<GongNode>()
            }
            RootNodeGongNodeAssociation.children.push(treeGongNodeInstance_RootNode)
          }

        }
      )


      this.dataSource.data = this.gongNodeTree

      // expand nodes that were exapanded before
      this.treeControl.dataNodes?.forEach(
        node => {
          if (memoryOfExpandedNodes.get(node.uniqueIdPerStack)) {
            this.treeControl.expand(node)
          }
        }
      )
    });

    // fetch the number of commits
    this.commitNbService.getCommitNb().subscribe(
      commitNb => {
        this.commitNb = commitNb
      }
    )
  }

  /**
   * 
   * @param path for the outlet selection
   */
  setTableRouterOutlet(path: string) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_gongmattree_go_table: ["github_com_fullstack_lang_gongmattree_go-" + path]
      }
    }]);
  }

  /**
   * 
   * @param path for the outlet selection
   */
  setTableRouterOutletFromTree(path: string, type: GongNodeType, structName: string, id: number) {

    if (type == GongNodeType.STRUCT) {
      this.router.navigate([{
        outlets: {
          github_com_fullstack_lang_gongmattree_go_table: ["github_com_fullstack_lang_gongmattree_go-" + path.toLowerCase()]
        }
      }]);
    }

    if (type == GongNodeType.INSTANCE) {
      this.router.navigate([{
        outlets: {
          github_com_fullstack_lang_gongmattree_go_presentation: ["github_com_fullstack_lang_gongmattree_go-" + structName.toLowerCase() + "-presentation", id]
        }
      }]);
    }
  }

  setEditorRouterOutlet(path: string) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_gongmattree_go_editor: ["github_com_fullstack_lang_gongmattree_go-" + path.toLowerCase()]
      }
    }]);
  }

  setEditorSpecialRouterOutlet(node: GongFlatNode) {
    this.router.navigate([{
      outlets: {
        github_com_fullstack_lang_gongmattree_go_editor: ["github_com_fullstack_lang_gongmattree_go-" + node.associatedStructName.toLowerCase() + "-adder", node.id, node.structName, node.associationField]
      }
    }]);
  }
}