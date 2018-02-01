import { Component, OnInit, Input } from '@angular/core';
import { ActivatedRoute } from '@angular/router';
import { Location } from '@angular/common';
import { PipeTransform } from '@angular/core/src/change_detection/pipe_transform';
import { Email } from '../email';
import { EmailService } from '../email.service';
import { Sanitizer } from '@angular/core/src/security';
import * as objectPath from 'object-path';

@Component({
  selector: 'app-email-details',
  templateUrl: './email-details.component.html',
  styleUrls: ['./email-details.component.css']
})

export class EmailDetailsComponent implements OnInit {
  formattedDate: string;
  isTextPlain = false;

  @Input() email: Email;

  constructor(
    private route: ActivatedRoute,
    private location: Location,
    private emailService: EmailService,
  ) { }

  ngOnInit() {
    this.getEmail();
  }

  getEmail(): void {
    const id = this.route.snapshot.paramMap.get('id');

    if (!id) { return; }

    this.emailService.getEmail(id)
      .subscribe(email => {
        this.isTextPlain = objectPath.get(email, 'ContentType', '').includes('text/plain')
          || objectPath.get(email, 'ContentType', '').includes('multipart');

        // strip Doctype and Style from html content types
        if (!this.isTextPlain) {
          const body = objectPath.get(email, 'Body', '');

          if (!body) { return Email; }

          const doctypelessBody = this.emailService.stripDoctype(body);
          const unstyledBody = this.emailService.stripStyleElement(doctypelessBody);

          email.Body = unstyledBody;
        }

        this.email = email;
      });
  }
}
