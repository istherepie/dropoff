import { Component } from '@angular/core';
import { UploadService } from "./upload-service.service"

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent {
  title = 'website';
  file: File = null; // Variable to store file

  constructor(private service: UploadService) { }

  onChange(files: File[]) {

    console.log(files)
    this.file = files[0];
  }

  onUpload() {
    const formdata = new FormData()

    formdata.append('uploadfile', this.file, this.file.name)

    console.log(formdata)
    this.service.upload(formdata)
  }



  ngOnInit(): void {
  }
}
