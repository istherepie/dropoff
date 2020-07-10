import { Injectable } from '@angular/core';
import { HttpClient, HttpEventType } from '@angular/common/http';
import { environment } from "../environments/environment"

@Injectable({
  providedIn: 'root'
})
export class UploadService {

  constructor(private http: HttpClient) { }

  list() {
    return this.http.get(environment.apiUrl)
  }

  upload(formData) {

    let url = environment.apiUrl + "/state/upload"

    this.http.post(url, formData, {
      reportProgress: true,
      observe: 'events'
    })

    .subscribe(resp => {
      if (resp.type === HttpEventType.Response) {
          console.log('Upload complete');
      }
      if (resp.type === HttpEventType.UploadProgress) {
          const percentDone = Math.round(100 * resp.loaded / resp.total);
          console.log('Progress ' + percentDone + '%');
      }
  });
  }
}
