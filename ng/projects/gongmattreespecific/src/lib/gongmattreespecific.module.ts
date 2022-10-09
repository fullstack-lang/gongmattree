import { NgModule } from '@angular/core';

import { GongmattreespecificComponent } from './gongmattreespecific.component';

import { TreeComponent } from './tree/tree.component';

import { MatSliderModule } from '@angular/material/slider';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatInputModule } from '@angular/material/input';
import { MatSelectModule } from '@angular/material/select'
import { MatDatepickerModule } from '@angular/material/datepicker'
import { MatTableModule } from '@angular/material/table'
import { MatSortModule } from '@angular/material/sort'
import { MatPaginatorModule } from '@angular/material/paginator'
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatButtonModule } from '@angular/material/button';
import { MatIconModule } from '@angular/material/icon';
import { MatToolbarModule } from '@angular/material/toolbar'
import { MatListModule } from '@angular/material/list'
import { MatExpansionModule } from '@angular/material/expansion';
import { MatDialogModule, MatDialogRef } from '@angular/material/dialog';
import { MatGridListModule } from '@angular/material/grid-list';
import { MatTreeModule } from '@angular/material/tree';
import { DragDropModule } from '@angular/cdk/drag-drop';


@NgModule({
  declarations: [
    GongmattreespecificComponent,
    TreeComponent
  ],
  imports: [
    MatTreeModule,
    MatIconModule,
    MatButtonModule
  ],
  exports: [
    GongmattreespecificComponent,
    TreeComponent
  ]
})
export class GongmattreespecificModule { }