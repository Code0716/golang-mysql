import axios from 'axios';
import config from '../config';

class HttpRequest {
  get(data) {
    return axios
      .get(config.API_SERVER_URL + data.url)
      .then(response => {
        console.log(response);
        const payload = JSON.stringify(response.data);
        return { payload };
      })
      .catch(error => {
        return { error };
      });
  }

  post(data) {
    const { url, json } = data.payload;
    const switchObject = JSON.parse(json);
    return axios
      .post(apiPrefix + url, { data: switchObject })
      .then(response => {
        const payload = response.data;
        return { payload };
      })
      .catch(error => {
        return { error };
      });
  }
}

export default new HttpRequest();
