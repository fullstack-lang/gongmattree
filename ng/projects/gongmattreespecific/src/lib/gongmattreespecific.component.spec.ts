import { ComponentFixture, TestBed } from '@angular/core/testing';

import { GongmattreespecificComponent } from './gongmattreespecific.component';

describe('GongmattreespecificComponent', () => {
  let component: GongmattreespecificComponent;
  let fixture: ComponentFixture<GongmattreespecificComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ GongmattreespecificComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(GongmattreespecificComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
