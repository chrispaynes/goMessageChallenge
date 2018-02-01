import { Pipe, PipeTransform } from '@angular/core';

@Pipe({
  name: 'stripMessageId'
})

export class StripMessageIdPipe implements PipeTransform {

  transform(id: string): string {
    if (!Boolean(id.length) && typeof(id) === 'undefined') {
      return '';
    }

    return id.substring(1, id.length - 1);
  }

}
