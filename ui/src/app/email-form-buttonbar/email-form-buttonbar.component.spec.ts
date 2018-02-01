import { async, ComponentFixture, TestBed } from '@angular/core/testing';

import { EmailFormButtonbarComponent } from './email-form-buttonbar.component';

describe('EmailFormButtonbarComponent', () => {
  let component: EmailFormButtonbarComponent;
  let fixture: ComponentFixture<EmailFormButtonbarComponent>;

  beforeEach(async(() => {
    TestBed.configureTestingModule({
      declarations: [ EmailFormButtonbarComponent ]
    })
    .compileComponents();
  }));

  beforeEach(() => {
    fixture = TestBed.createComponent(EmailFormButtonbarComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
