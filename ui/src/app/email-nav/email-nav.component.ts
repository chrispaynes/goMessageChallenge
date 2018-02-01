import { Component, OnInit } from '@angular/core';
import { EmailNav } from '../email-nav';
import { EmailService } from '../email.service';

@Component({
  selector: 'app-email-nav',
  templateUrl: './email-nav.component.html',
  styleUrls: ['./email-nav.component.css']
})

export class EmailNavComponent implements OnInit {
  nav: EmailNav = {
    TotalEmails: this.emailService.getInboxCount(),
  };

  constructor(
    private emailService: EmailService,
  ) { }

  ngOnInit() { }

}
