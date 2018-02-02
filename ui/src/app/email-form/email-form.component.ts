import { Component, OnInit } from '@angular/core';
import { Email } from '../email';
import { FileUploadService } from '../file-upload.service';

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
    private fileUploadService: FileUploadService,
  ) { }

  ngOnInit() { }

  handleFileInput(files: FileList) {
    this.fileToUpload = files.item(0);
  }

  uploadFile(file: File) {
    this.fileUploadService.postFile(file);
  }
}
