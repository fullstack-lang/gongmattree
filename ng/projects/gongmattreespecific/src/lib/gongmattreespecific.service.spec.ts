import { TestBed } from '@angular/core/testing';

import { GongmattreespecificService } from './gongmattreespecific.service';

describe('GongmattreespecificService', () => {
  let service: GongmattreespecificService;

  beforeEach(() => {
    TestBed.configureTestingModule({});
    service = TestBed.inject(GongmattreespecificService);
  });

  it('should be created', () => {
    expect(service).toBeTruthy();
  });
});
