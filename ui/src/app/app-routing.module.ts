import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import { EmailListComponent } from './email-list/email-list.component';
import { EmailFormComponent } from './email-form/email-form.component';
import { EmailDetailsComponent } from './email-details/email-details.component';

const routes: Routes = [
  { path: '', redirectTo: '/inbox', pathMatch: 'full' },
  { path: 'inbox', component: EmailListComponent },
  { path: 'inbox/:id', component: EmailDetailsComponent },
  { path: 'compose', component: EmailFormComponent },
  { path: 'post', component: EmailDetailsComponent },
];

@NgModule({
  imports: [ RouterModule.forRoot(routes) ],
  exports: [
    RouterModule
  ]
})

export class AppRoutingModule {
}
