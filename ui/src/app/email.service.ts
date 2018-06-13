import { Injectable } from '@angular/core';
import { Email } from './email';
import { EMAILS } from './mock-emails';
import { Observable } from 'rxjs/Observable';
import { of } from 'rxjs/observable/of';
import { HttpClient, HttpHeaders } from '@angular/common/http';
import * as moment from 'moment';
import * as objectPath from 'object-path';

@Injectable()
export class EmailService {

  constructor(
    private http: HttpClient,
  ) { }

  // getEmails gets a collection of mock JSON emails
  getEmails(): Observable<Email[]> {
    return of(EMAILS);
  }

  // getEmail gets emailed from LocalStorage or the mock-email datastore
  getEmail(id: string): Observable<Email> {
    // match unbracketed server generated Message-IDs such as:
    // 6426946.1413.1301675117949.JavaMail.tomcat@osadmin02
    if (isNaN(+id)) {
      const localStorageEmail = localStorage.getItem(id);

      if (localStorageEmail === null) {
        return of(EMAILS.find(email => email.MessageId === `<${id}>`));
      }

      return of(JSON.parse(localStorageEmail));
    }
  }

  // extractMessageId extracts an email's Message-ID
  // from within its opening and closing brackets
  extractMessageId(id: string) {
    if (!Boolean(id.length) && typeof (id) === 'undefined') {
      return '';
    }

    return id.substring(1, id.length - 1);
  }

  // postEmail POSTs an email message to the server
  postEmail(email: string): Observable<Email> {
    const httpOptions = {
      headers: new HttpHeaders({
        'Content-Type': 'text/plain',
        'Accept': 'application/json',
      })
    };

    return this.http.post<Email>(
      'http://api-gmc.localhost/email',
      email,
      httpOptions
    );
  }

  // cutString cuts a section out of a string
  cutString(str: string, cutStart: number, cutEnd: number) {
    return str.substr(0, cutStart) + str.substr(cutEnd + 1);
  }

  // getInboxCounts get the number of inbox emails
  getInboxCount(): number {
    return (typeof (EMAILS) !== 'undefined' && Array.isArray(EMAILS)) ? EMAILS.length : 0;
  }

  // stripStyleElement strips the "<style>...</style>" element from an email body
  stripStyleElement(body: string): string {
    const styleStart = body.indexOf('<style');
    const styleEnd = body.indexOf('</style>');

    if (styleStart === -1 || styleEnd === -1) {
      return body;
    }

    return this.cutString(body, styleStart, styleEnd + 8);
  }

  // stripDoctype strips the <!DOCTYPE ...> declaration from an email body
  stripDoctype(body: string): string {
    return body.replace(/<!DOCTYPE[^>[]*(\[[^]]*\])?>/, '');
  }
}
