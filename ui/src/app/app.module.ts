import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HttpClientModule } from '@angular/common/http';

import { AppComponent } from './app.component';
import { EmailListComponent } from './email-list/email-list.component';
import { EmailDetailsComponent } from './email-details/email-details.component';
import { EmailComponent } from './email/email.component';
import { EmailNavComponent } from './email-nav/email-nav.component';
import { EmailFormComponent } from './email-form/email-form.component';
import { EmailFormButtonbarComponent } from './email-form-buttonbar/email-form-buttonbar.component';
import { EmailService } from './email.service';
import { FileUploadService } from './file-upload.service';
import { AppRoutingModule } from './app-routing.module';
import { StripMessageIdPipe } from './strip-message-id.pipe';
import { FormatDatePipe } from './format-date.pipe';

@NgModule({
  declarations: [
    AppComponent,
    EmailListComponent,
    EmailDetailsComponent,
    EmailComponent,
    EmailNavComponent,
    EmailFormComponent,
    EmailFormButtonbarComponent,
    StripMessageIdPipe,
    FormatDatePipe
  ],
  imports: [
    BrowserModule,
    FormsModule,
    AppRoutingModule,
    HttpClientModule,
  ],
  providers: [EmailService, FileUploadService],
  bootstrap: [AppComponent]
})
export class AppModule { }
