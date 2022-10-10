import { Injectable } from '@angular/core';
import { HttpClient, HttpHeaders } from '@angular/common/http';

import { Observable, combineLatest, BehaviorSubject } from 'rxjs';

// insertion point sub template for services imports 
import { NodeDB } from './node-db'
import { NodeService } from './node.service'

import { TreeDB } from './tree-db'
import { TreeService } from './tree.service'


// FrontRepo stores all instances in a front repository (design pattern repository)
export class FrontRepo { // insertion point sub template 
  Nodes_array = new Array<NodeDB>(); // array of repo instances
  Nodes = new Map<number, NodeDB>(); // map of repo instances
  Nodes_batch = new Map<number, NodeDB>(); // same but only in last GET (for finding repo instances to delete)
  Trees_array = new Array<TreeDB>(); // array of repo instances
  Trees = new Map<number, TreeDB>(); // map of repo instances
  Trees_batch = new Map<number, TreeDB>(); // same but only in last GET (for finding repo instances to delete)
}

//
// Store of all instances of the stack
//
export const FrontRepoSingloton = new (FrontRepo)

// the table component is called in different ways
//
// DISPLAY or ASSOCIATION MODE
//
// in ASSOCIATION MODE, it is invoked within a diaglo and a Dialog Data item is used to
// configure the component
// DialogData define the interface for information that is forwarded from the calling instance to 
// the select table
export class DialogData {
  ID: number = 0 // ID of the calling instance

  // the reverse pointer is the name of the generated field on the destination
  // struct of the ONE-MANY association
  ReversePointer: string = "" // field of {{Structname}} that serve as reverse pointer
  OrderingMode: boolean = false // if true, this is for ordering items

  // there are different selection mode : ONE_MANY or MANY_MANY
  SelectionMode: SelectionMode = SelectionMode.ONE_MANY_ASSOCIATION_MODE

  // used if SelectionMode is MANY_MANY_ASSOCIATION_MODE
  //
  // In Gong, a MANY-MANY association is implemented as a ONE-ZERO/ONE followed by a ONE_MANY association
  // 
  // in the MANY_MANY_ASSOCIATION_MODE case, we need also the Struct and the FieldName that are
  // at the end of the ONE-MANY association
  SourceStruct: string = ""  // The "Aclass"
  SourceField: string = "" // the "AnarrayofbUse"
  IntermediateStruct: string = "" // the "AclassBclassUse" 
  IntermediateStructField: string = "" // the "Bclass" as field
  NextAssociationStruct: string = "" // the "Bclass"
}

export enum SelectionMode {
  ONE_MANY_ASSOCIATION_MODE = "ONE_MANY_ASSOCIATION_MODE",
  MANY_MANY_ASSOCIATION_MODE = "MANY_MANY_ASSOCIATION_MODE",
}

//
// observable that fetch all elements of the stack and store them in the FrontRepo
//
@Injectable({
  providedIn: 'root'
})
export class FrontRepoService {

  httpOptions = {
    headers: new HttpHeaders({ 'Content-Type': 'application/json' })
  };

  constructor(
    private http: HttpClient, // insertion point sub template 
    private nodeService: NodeService,
    private treeService: TreeService,
  ) { }

  // postService provides a post function for each struct name
  postService(structName: string, instanceToBePosted: any) {
    let service = this[structName.toLowerCase() + "Service" + "Service" as keyof FrontRepoService]
    let servicePostFunction = service[("post" + structName) as keyof typeof service] as (instance: typeof instanceToBePosted) => Observable<typeof instanceToBePosted>

    servicePostFunction(instanceToBePosted).subscribe(
      instance => {
        let behaviorSubject = instanceToBePosted[(structName + "ServiceChanged") as keyof typeof instanceToBePosted] as unknown as BehaviorSubject<string>
        behaviorSubject.next("post")
      }
    );
  }

  // deleteService provides a delete function for each struct name
  deleteService(structName: string, instanceToBeDeleted: any) {
    let service = this[structName.toLowerCase() + "Service" as keyof FrontRepoService]
    let serviceDeleteFunction = service["delete" + structName as keyof typeof service] as (instance: typeof instanceToBeDeleted) => Observable<typeof instanceToBeDeleted>

    serviceDeleteFunction(instanceToBeDeleted).subscribe(
      instance => {
        let behaviorSubject = instanceToBeDeleted[(structName + "ServiceChanged") as keyof typeof instanceToBeDeleted] as unknown as BehaviorSubject<string>
        behaviorSubject.next("delete")
      }
    );
  }

  // typing of observable can be messy in typescript. Therefore, one force the type
  observableFrontRepo: [ // insertion point sub template 
    Observable<NodeDB[]>,
    Observable<TreeDB[]>,
  ] = [ // insertion point sub template 
      this.nodeService.getNodes(),
      this.treeService.getTrees(),
    ];

  //
  // pull performs a GET on all struct of the stack and redeem association pointers 
  //
  // This is an observable. Therefore, the control flow forks with
  // - pull() return immediatly the observable
  // - the observable observer, if it subscribe, is called when all GET calls are performs
  pull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest(
          this.observableFrontRepo
        ).subscribe(
          ([ // insertion point sub template for declarations 
            nodes_,
            trees_,
          ]) => {
            // Typing can be messy with many items. Therefore, type casting is necessary here
            // insertion point sub template for type casting 
            var nodes: NodeDB[]
            nodes = nodes_ as NodeDB[]
            var trees: TreeDB[]
            trees = trees_ as TreeDB[]

            // 
            // First Step: init map of instances
            // insertion point sub template for init 
            // init the array
            FrontRepoSingloton.Nodes_array = nodes

            // clear the map that counts Node in the GET
            FrontRepoSingloton.Nodes_batch.clear()

            nodes.forEach(
              node => {
                FrontRepoSingloton.Nodes.set(node.ID, node)
                FrontRepoSingloton.Nodes_batch.set(node.ID, node)
              }
            )

            // clear nodes that are absent from the batch
            FrontRepoSingloton.Nodes.forEach(
              node => {
                if (FrontRepoSingloton.Nodes_batch.get(node.ID) == undefined) {
                  FrontRepoSingloton.Nodes.delete(node.ID)
                }
              }
            )

            // sort Nodes_array array
            FrontRepoSingloton.Nodes_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });

            // init the array
            FrontRepoSingloton.Trees_array = trees

            // clear the map that counts Tree in the GET
            FrontRepoSingloton.Trees_batch.clear()

            trees.forEach(
              tree => {
                FrontRepoSingloton.Trees.set(tree.ID, tree)
                FrontRepoSingloton.Trees_batch.set(tree.ID, tree)
              }
            )

            // clear trees that are absent from the batch
            FrontRepoSingloton.Trees.forEach(
              tree => {
                if (FrontRepoSingloton.Trees_batch.get(tree.ID) == undefined) {
                  FrontRepoSingloton.Trees.delete(tree.ID)
                }
              }
            )

            // sort Trees_array array
            FrontRepoSingloton.Trees_array.sort((t1, t2) => {
              if (t1.Name > t2.Name) {
                return 1;
              }
              if (t1.Name < t2.Name) {
                return -1;
              }
              return 0;
            });


            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template for redeem 
            nodes.forEach(
              node => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Node.Children redeeming
                {
                  let _node = FrontRepoSingloton.Nodes.get(node.Node_ChildrenDBID.Int64)
                  if (_node) {
                    if (_node.Children == undefined) {
                      _node.Children = new Array<NodeDB>()
                    }
                    _node.Children.push(node)
                    if (node.Node_Children_reverse == undefined) {
                      node.Node_Children_reverse = _node
                    }
                  }
                }
                // insertion point for slice of pointer field Tree.RootNodes redeeming
                {
                  let _tree = FrontRepoSingloton.Trees.get(node.Tree_RootNodesDBID.Int64)
                  if (_tree) {
                    if (_tree.RootNodes == undefined) {
                      _tree.RootNodes = new Array<NodeDB>()
                    }
                    _tree.RootNodes.push(node)
                    if (node.Tree_RootNodes_reverse == undefined) {
                      node.Tree_RootNodes_reverse = _tree
                    }
                  }
                }
              }
            )
            trees.forEach(
              tree => {
                // insertion point sub sub template for ONE-/ZERO-ONE associations pointers redeeming

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
          }
        )
      }
    )
  }

  // insertion point for pull per struct 

  // NodePull performs a GET on Node of the stack and redeem association pointers 
  NodePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.nodeService.getNodes()
        ]).subscribe(
          ([ // insertion point sub template 
            nodes,
          ]) => {
            // init the array
            FrontRepoSingloton.Nodes_array = nodes

            // clear the map that counts Node in the GET
            FrontRepoSingloton.Nodes_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            nodes.forEach(
              node => {
                FrontRepoSingloton.Nodes.set(node.ID, node)
                FrontRepoSingloton.Nodes_batch.set(node.ID, node)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
                // insertion point for slice of pointer field Node.Children redeeming
                {
                  let _node = FrontRepoSingloton.Nodes.get(node.Node_ChildrenDBID.Int64)
                  if (_node) {
                    if (_node.Children == undefined) {
                      _node.Children = new Array<NodeDB>()
                    }
                    _node.Children.push(node)
                    if (node.Node_Children_reverse == undefined) {
                      node.Node_Children_reverse = _node
                    }
                  }
                }
                // insertion point for slice of pointer field Tree.RootNodes redeeming
                {
                  let _tree = FrontRepoSingloton.Trees.get(node.Tree_RootNodesDBID.Int64)
                  if (_tree) {
                    if (_tree.RootNodes == undefined) {
                      _tree.RootNodes = new Array<NodeDB>()
                    }
                    _tree.RootNodes.push(node)
                    if (node.Tree_RootNodes_reverse == undefined) {
                      node.Tree_RootNodes_reverse = _tree
                    }
                  }
                }
              }
            )

            // clear nodes that are absent from the GET
            FrontRepoSingloton.Nodes.forEach(
              node => {
                if (FrontRepoSingloton.Nodes_batch.get(node.ID) == undefined) {
                  FrontRepoSingloton.Nodes.delete(node.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
          }
        )
      }
    )
  }

  // TreePull performs a GET on Tree of the stack and redeem association pointers 
  TreePull(): Observable<FrontRepo> {
    return new Observable<FrontRepo>(
      (observer) => {
        combineLatest([
          this.treeService.getTrees()
        ]).subscribe(
          ([ // insertion point sub template 
            trees,
          ]) => {
            // init the array
            FrontRepoSingloton.Trees_array = trees

            // clear the map that counts Tree in the GET
            FrontRepoSingloton.Trees_batch.clear()

            // 
            // First Step: init map of instances
            // insertion point sub template 
            trees.forEach(
              tree => {
                FrontRepoSingloton.Trees.set(tree.ID, tree)
                FrontRepoSingloton.Trees_batch.set(tree.ID, tree)

                // insertion point for redeeming ONE/ZERO-ONE associations

                // insertion point for redeeming ONE-MANY associations
              }
            )

            // clear trees that are absent from the GET
            FrontRepoSingloton.Trees.forEach(
              tree => {
                if (FrontRepoSingloton.Trees_batch.get(tree.ID) == undefined) {
                  FrontRepoSingloton.Trees.delete(tree.ID)
                }
              }
            )

            // 
            // Second Step: redeem pointers between instances (thanks to maps in the First Step)
            // insertion point sub template 

            // hand over control flow to observer
            observer.next(FrontRepoSingloton)
          }
        )
      }
    )
  }
}

// insertion point for get unique ID per struct 
export function getNodeUniqueID(id: number): number {
  return 31 * id
}
export function getTreeUniqueID(id: number): number {
  return 37 * id
}
