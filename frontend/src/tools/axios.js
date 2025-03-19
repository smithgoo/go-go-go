import axios from 'axios';

const instance = axios.create({
    baseURL: 'http://localhost:9000', // 设置基础 URL
    timeout: 10000, // 可选: 设置请求超时时间
    // headers: {
    //     'Content-Type': 'application/json', // 可选: 设置默认请求头
    // }
});

// 获取 token（可以从 localStorage, Vuex, 或其他存储方式中获取）
const getToken = () => {
    return localStorage.getItem('user') || ''; // 假设 token 存在 localStorage 中
};

// 请求拦截器，统一添加 token
instance.interceptors.request.use(
    config => {
        const token = getToken(); // 获取 token
        const cleanedToken = token.replace(/"/g, ''); // 去除多余的引号
        if (token) {
            config.headers['Authorization'] = `${cleanedToken}`;// 如果存在 token，则添加到请求头
        }
        return config;
    },
    error => {
        return Promise.reject(error); // 处理错误
    }
);

// 封装 GET 方法
export const getRequest = (url, params = {}, withToken = true) => {
    return new Promise((resolve, reject) => {
        // 判断是否需要 token
        if (!withToken) {
            delete instance.defaults.headers.common['Authorization']; // 如果不需要 token，移除 Authorization 头
        }

        instance.get(url, { params })
            .then(response => {
                resolve(response.data);  // 返回实际数据
            })
            .catch(error => {
                console.error('GET request failed:', error);
                reject(error);  // 抛出错误，供调用方处理
            });
    });
};

// 封装 POST 方法
export const postRequest = (url, data = {}, withToken = true) => {
    return new Promise((resolve, reject) => {
        // 判断是否需要 token
        if (!withToken) {
            delete instance.defaults.headers.common['Authorization']; // 如果不需要 token，移除 Authorization 头
        }

        instance.post(url, data)
            .then(response => {
                resolve(response.data);  // 返回实际数据
            })
            .catch(error => {
                console.error('POST request failed:', error);
                reject(error);  // 抛出错误，供调用方处理
            });
    });
};

 

export default axios;
