import axios from 'axios';
import config from '../config';

class HttpRequest {
  get(url) {
    return axios
      .get(config.API_SERVER_URL + url)
      .then(response => {
        const payload = response.data;
        return payload;
      })
      .catch(error => {
        return { error };
      });
  }

  post(data) {
    const { url, json } = data.payload;
    const switchObject = JSON.parse(json);
    return axios
      .post(config.API_SERVER_URL + url, { data: switchObject })
      .then(response => {
        const payload = response.data;
        return { payload };
      })
      .catch(error => {
        return { error };
      });
  }

  postImg(url, submitData) {
    axios.post(config.API_SERVER_URL + url, submitData, {
      headers: {
        'content-type': 'multipart/form-data',
      },
    });
  }
}

export default new HttpRequest();
