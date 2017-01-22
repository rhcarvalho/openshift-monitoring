import { Component } from '@angular/core';

@Component({
  selector: 'app-root',
  template:  `
        <h3>OpenShift - Updatemonitoring</h3>
        <app-deamon-overview></app-deamon-overview>
        <app-job></app-job>
        <simple-notifications></simple-notifications>
    `
})
export class AppComponent {
}
