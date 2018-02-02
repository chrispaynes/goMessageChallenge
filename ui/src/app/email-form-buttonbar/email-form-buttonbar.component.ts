import { Component, OnInit, Input } from '@angular/core';
import { Location } from '@angular/common';
import { Router } from '@angular/router';
import * as objectPath from 'object-path';
import { EmailService } from '../email.service';

@Component({
  selector: 'app-email-form-buttonbar',
  templateUrl: './email-form-buttonbar.component.html',
  styleUrls: ['./email-form-buttonbar.component.css']
})

export class EmailFormButtonbarComponent implements OnInit {
  response = {};

  @Input() postBody: string;
  @Input() postUpload: File;

  constructor(
    private emailService: EmailService,
    private location: Location,
    private router: Router,
  ) { }

  ngOnInit() {
  }

  post(body: string): void {
    localStorage.emails = [];
    body = body.trim();

    if (!body) { return; }

    this.emailService.postEmail(body)
      .subscribe(resp => {
        this.response = {};

        localStorage[resp.MessageId.substring(1, resp.MessageId.length - 1)] = JSON.stringify(resp);
        const regex1 = /Message-ID:\s{0,}<(\S+)>/;
        const bracketedId = regex1.exec(body);

        if (bracketedId === null || objectPath.get(bracketedId, '0', '') === '') {
          return;
        }

        const propName = /Message-ID:\s{0,}/;
        const id = objectPath.get(bracketedId, '0', '').replace(propName, '');

        this.router.navigateByUrl(`/inbox/${this.emailService.extractMessageId(id)}`);
      });
  }

}
