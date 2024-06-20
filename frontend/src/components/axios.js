import axios from 'axios';

const instance = axios.create({
    baseURL: 'http://localhost:9000', // Set your base URL and port here
});

export default instance;
