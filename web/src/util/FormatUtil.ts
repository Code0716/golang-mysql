import * as moment from 'moment';

export class FormatUtil {
  formatedDate = (dete: Date): string => {
    return moment(dete).format('YYYY/MM/DD');
  };
}
