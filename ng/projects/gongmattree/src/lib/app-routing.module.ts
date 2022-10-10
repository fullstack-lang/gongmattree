import { NgModule } from '@angular/core';
import { Routes, RouterModule } from '@angular/router';

// insertion point for imports
import { NodesTableComponent } from './nodes-table/nodes-table.component'
import { NodeDetailComponent } from './node-detail/node-detail.component'
import { NodePresentationComponent } from './node-presentation/node-presentation.component'

import { TreesTableComponent } from './trees-table/trees-table.component'
import { TreeDetailComponent } from './tree-detail/tree-detail.component'
import { TreePresentationComponent } from './tree-presentation/tree-presentation.component'


const routes: Routes = [ // insertion point for routes declarations
	{ path: 'github_com_fullstack_lang_gongmattree_go-nodes', component: NodesTableComponent, outlet: 'github_com_fullstack_lang_gongmattree_go_table' },
	{ path: 'github_com_fullstack_lang_gongmattree_go-node-adder', component: NodeDetailComponent, outlet: 'github_com_fullstack_lang_gongmattree_go_editor' },
	{ path: 'github_com_fullstack_lang_gongmattree_go-node-adder/:id/:originStruct/:originStructFieldName', component: NodeDetailComponent, outlet: 'github_com_fullstack_lang_gongmattree_go_editor' },
	{ path: 'github_com_fullstack_lang_gongmattree_go-node-detail/:id', component: NodeDetailComponent, outlet: 'github_com_fullstack_lang_gongmattree_go_editor' },
	{ path: 'github_com_fullstack_lang_gongmattree_go-node-presentation/:id', component: NodePresentationComponent, outlet: 'github_com_fullstack_lang_gongmattree_go_presentation' },
	{ path: 'github_com_fullstack_lang_gongmattree_go-node-presentation-special/:id', component: NodePresentationComponent, outlet: 'github_com_fullstack_lang_gongmattree_gonodepres' },

	{ path: 'github_com_fullstack_lang_gongmattree_go-trees', component: TreesTableComponent, outlet: 'github_com_fullstack_lang_gongmattree_go_table' },
	{ path: 'github_com_fullstack_lang_gongmattree_go-tree-adder', component: TreeDetailComponent, outlet: 'github_com_fullstack_lang_gongmattree_go_editor' },
	{ path: 'github_com_fullstack_lang_gongmattree_go-tree-adder/:id/:originStruct/:originStructFieldName', component: TreeDetailComponent, outlet: 'github_com_fullstack_lang_gongmattree_go_editor' },
	{ path: 'github_com_fullstack_lang_gongmattree_go-tree-detail/:id', component: TreeDetailComponent, outlet: 'github_com_fullstack_lang_gongmattree_go_editor' },
	{ path: 'github_com_fullstack_lang_gongmattree_go-tree-presentation/:id', component: TreePresentationComponent, outlet: 'github_com_fullstack_lang_gongmattree_go_presentation' },
	{ path: 'github_com_fullstack_lang_gongmattree_go-tree-presentation-special/:id', component: TreePresentationComponent, outlet: 'github_com_fullstack_lang_gongmattree_gotreepres' },

];

@NgModule({
	imports: [RouterModule.forRoot(routes)],
	exports: [RouterModule]
})
export class AppRoutingModule { }
