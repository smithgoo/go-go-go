import axios from 'axios';

const instance = axios.create({
    baseURL: 'http://192.168.1.119:9000', // Set your base URL and port here
});

export default instance;
