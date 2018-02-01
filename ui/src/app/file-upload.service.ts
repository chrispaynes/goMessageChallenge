import { Injectable } from '@angular/core';
import { Observable } from 'rxjs/Observable';
import { HttpClient, HttpHeaders } from '@angular/common/http';
// import { map } from 'rxjs/observable/of';


@Injectable()
export class FileUploadService {

  constructor(
    private http: HttpClient,
  ) { }

  postFile(fileToUpload: File): Observable<Object> {
    console.log('fileToUpload', fileToUpload);
    console.log('fileToUpload.name', fileToUpload.name);

    const endpoint = 'http://localhost:3000/email';
    const formData: FormData = new FormData();

    formData.append('fileKey', fileToUpload, fileToUpload.name);

    console.log('FORM DATA', formData);

    return this.http
      .post(endpoint, formData);
      // .map(() => true);
      // .catch((e) => this.handleError(e));
  }
}
