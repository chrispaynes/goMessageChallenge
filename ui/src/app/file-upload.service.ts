import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/Observable';
import * as request from 'superagent';
import { Router } from '@angular/router';
import * as objectPath from 'object-path';
import { EmailService } from './email.service';

@Injectable()
export class FileUploadService {
  constructor(
    private router: Router,
    private emailService: EmailService,
  ) { }

  // postFile reads and POSTS an attached msg file to the server
  postFile(fileToUpload: File): void {
    const fileReader = new FileReader();

    fileReader.onload = () => {
      // TODO: move various service extraction methods to separate service?
      const id = this.emailService.extractMessageId(
        this.extractMessageIDfromBody(fileReader.result));

      // TODO: add error handling for reading and ensure file reading is 100% done.
      request
        .post('http://localhost:3000/email')
        .set('Content-Type', 'text/plain')
        .set('Accept', 'application/json')
        .send(fileReader.result)
        .end(() => {
          setTimeout(() => {
            this.router.navigateByUrl(`/inbox/${id}`);
          }, 2000);
        });
    };

    fileReader.readAsText(fileToUpload);
  }

  // extractMessageIDfromBody extracts a bracketed message id from an email body
  extractMessageIDfromBody(body: string) {
    const regex1 = /Message-ID:\s{0,}<(\S+)>/i;
    const bracketedId = regex1.exec(body);
    if (bracketedId === null || objectPath.get(bracketedId, '0', '') === '') {
      return;
    }
    const propName = /Message-ID:\s{0,}/i;
    const id = objectPath.get(bracketedId, '0', '').replace(propName, '');

    return id;
  }
}
