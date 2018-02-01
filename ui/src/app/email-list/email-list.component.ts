import { Component, OnInit } from '@angular/core';
import { Email } from '../email';
import { EmailService } from '../email.service';

@Component({
  selector: 'app-email-list',
  templateUrl: './email-list.component.html',
  styleUrls: ['./email-list.component.css']
})

export class EmailListComponent implements OnInit {
  emails: Email[];
  selectedEmail: Email;

  onSelect(email: Email): void {
    this.selectedEmail = email;
  }

  constructor(
    private emailService: EmailService
  ) { }

  ngOnInit() {
    this.getEmails();
  }

  getEmails(): void {
    this.emailService.getEmails()
      .subscribe(emails => this.emails = emails);
  }
}
