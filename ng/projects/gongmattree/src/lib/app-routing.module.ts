import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

// insertion point for imports
import { ButtonsTableComponent } from './buttons-table/buttons-table.component'
import { ButtonDetailComponent } from './button-detail/button-detail.component'
import { ButtonPresentationComponent } from './button-presentation/button-presentation.component'

import { NodesTableComponent } from './nodes-table/nodes-table.component'
import { NodeDetailComponent } from './node-detail/node-detail.component'
import { NodePresentationComponent } from './node-presentation/node-presentation.component'

import { TreesTableComponent } from './trees-table/trees-table.component'
import { TreeDetailComponent } from './tree-detail/tree-detail.component'
import { TreePresentationComponent } from './tree-presentation/tree-presentation.component'


const routes: Routes = [ // insertion point for routes declarations
	{ path: 'gongmattree_go-buttons', component: ButtonsTableComponent, outlet: 'gongmattree_go_table' },
	{ path: 'gongmattree_go-button-adder', component: ButtonDetailComponent, outlet: 'gongmattree_go_editor' },
	{ path: 'gongmattree_go-button-adder/:id/:originStruct/:originStructFieldName', component: ButtonDetailComponent, outlet: 'gongmattree_go_editor' },
	{ path: 'gongmattree_go-button-detail/:id', component: ButtonDetailComponent, outlet: 'gongmattree_go_editor' },
	{ path: 'gongmattree_go-button-presentation/:id', component: ButtonPresentationComponent, outlet: 'gongmattree_go_presentation' },
	{ path: 'gongmattree_go-button-presentation-special/:id', component: ButtonPresentationComponent, outlet: 'gongmattree_gobuttonpres' },

	{ path: 'gongmattree_go-nodes', component: NodesTableComponent, outlet: 'gongmattree_go_table' },
	{ path: 'gongmattree_go-node-adder', component: NodeDetailComponent, outlet: 'gongmattree_go_editor' },
	{ path: 'gongmattree_go-node-adder/:id/:originStruct/:originStructFieldName', component: NodeDetailComponent, outlet: 'gongmattree_go_editor' },
	{ path: 'gongmattree_go-node-detail/:id', component: NodeDetailComponent, outlet: 'gongmattree_go_editor' },
	{ path: 'gongmattree_go-node-presentation/:id', component: NodePresentationComponent, outlet: 'gongmattree_go_presentation' },
	{ path: 'gongmattree_go-node-presentation-special/:id', component: NodePresentationComponent, outlet: 'gongmattree_gonodepres' },

	{ path: 'gongmattree_go-trees', component: TreesTableComponent, outlet: 'gongmattree_go_table' },
	{ path: 'gongmattree_go-tree-adder', component: TreeDetailComponent, outlet: 'gongmattree_go_editor' },
	{ path: 'gongmattree_go-tree-adder/:id/:originStruct/:originStructFieldName', component: TreeDetailComponent, outlet: 'gongmattree_go_editor' },
	{ path: 'gongmattree_go-tree-detail/:id', component: TreeDetailComponent, outlet: 'gongmattree_go_editor' },
	{ path: 'gongmattree_go-tree-presentation/:id', component: TreePresentationComponent, outlet: 'gongmattree_go_presentation' },
	{ path: 'gongmattree_go-tree-presentation-special/:id', component: TreePresentationComponent, outlet: 'gongmattree_gotreepres' },

];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule]
})
export class AppRoutingModule { }
