import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { EmailNavComponent } from './email-nav.component';

describe('EmailNavComponent', () => {
  let component: EmailNavComponent;
  let fixture: ComponentFixture<EmailNavComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ EmailNavComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(EmailNavComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
