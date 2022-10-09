import { Component, OnInit } from '@angular/core';
import { FormBuilder, FormControl, FormGroup } from '@angular/forms';

import { ButtonDB } from '../button-db'
import { ButtonService } from '../button.service'

import { FrontRepoService, FrontRepo } from '../front-repo.service'

import { Router, RouterState, ActivatedRoute } from '@angular/router';

// insertion point for additional imports

export interface buttonDummyElement {
}

const ELEMENT_DATA: buttonDummyElement[] = [
];

@Component({
	selector: 'app-button-presentation',
	templateUrl: './button-presentation.component.html',
	styleUrls: ['./button-presentation.component.css'],
})
export class ButtonPresentationComponent implements OnInit {

	// insertion point for additionnal time duration declarations
	// insertion point for additionnal enum int field declarations

	displayedColumns: string[] = []
	dataSource = ELEMENT_DATA

	button: ButtonDB = new (ButtonDB)

	// front repo
	frontRepo: FrontRepo = new (FrontRepo)
 
	constructor(
		private buttonService: ButtonService,
		private frontRepoService: FrontRepoService,
		private route: ActivatedRoute,
		private router: Router,
	) {
		this.router.routeReuseStrategy.shouldReuseRoute = function () {
			return false;
		};
	}

	ngOnInit(): void {
		this.getButton();

		// observable for changes in 
		this.buttonService.ButtonServiceChanged.subscribe(
			message => {
				if (message == "update") {
					this.getButton()
				}
			}
		)
	}

	getButton(): void {
		const id = +this.route.snapshot.paramMap.get('id')!
		this.frontRepoService.pull().subscribe(
			frontRepo => {
				this.frontRepo = frontRepo

				this.button = this.frontRepo.Buttons.get(id)!

				// insertion point for recovery of durations
				// insertion point for recovery of enum tint
			}
		);
	}

	// set presentation outlet
	setPresentationRouterOutlet(structName: string, ID: number) {
		this.router.navigate([{
			outlets: {
				github_com_fullstack_lang_gongmattree_go_presentation: ["github_com_fullstack_lang_gongmattree_go-" + structName + "-presentation", ID]
			}
		}]);
	}

	// set editor outlet
	setEditorRouterOutlet(ID: number) {
		this.router.navigate([{
			outlets: {
				github_com_fullstack_lang_gongmattree_go_editor: ["github_com_fullstack_lang_gongmattree_go-" + "button-detail", ID]
			}
		}]);
	}
}
