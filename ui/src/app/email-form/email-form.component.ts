import { Component, OnInit } from '@angular/core';
import { Email } from '../email';

@Component({
  selector: 'app-email-form',
  templateUrl: './email-form.component.html',
  styleUrls: ['./email-form.component.css']
})

export class EmailFormComponent implements OnInit {
  fileToUpload: File;

  form: Email = {
    MessageId: '',
    Subject: '',
    To: '',
    From: '',
    Body: '',
    Date: '',
    ContentType: ''
  };

  constructor(
  ) { }

  ngOnInit() { }

  handleFileInput(files: FileList) {
    console.log('FILES', files.item(0));
    this.fileToUpload = files.item(0);
  }
}
