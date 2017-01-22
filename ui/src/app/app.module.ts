import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HttpModule } from '@angular/http';

import { AppComponent } from './app.component';
import { DeamonOverviewComponent } from './deamon-overview/deamon-overview.component';
import {SocketService} from "./socket.service";
import {SimpleNotificationsModule} from 'angular2-notifications';
import { JobComponent } from './job/job.component';

@NgModule({
  declarations: [
    AppComponent,
    DeamonOverviewComponent,
    JobComponent
  ],
  imports: [
    BrowserModule,
    FormsModule,
    HttpModule,
    SimpleNotificationsModule
  ],
  providers: [SocketService],
  bootstrap: [AppComponent]
})
export class AppModule { }
