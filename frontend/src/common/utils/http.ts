import axios from 'axios';
import { HOST } from '../constants';

axios.defaults.baseURL = HOST;

axios.interceptors.response.use(res => {
	return {
		data: res.data
	}
})


export default axios;