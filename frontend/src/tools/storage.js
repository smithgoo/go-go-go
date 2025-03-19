class StorageService {
  constructor() {}

  // setItem 接受 key 和 data
  setItem(key, data) {
    localStorage.setItem(key, JSON.stringify(data));
  }

  // getItem 接受 key
  getItem(key) {
    const data = localStorage.getItem(key);
    return data ? JSON.parse(data) : null;
  }

  // removeItem 接受 key
  removeItem(key) {
    localStorage.removeItem(key);
  }

  clear() {
    localStorage.clear();
  }
}

// 导出一个通用的实例
const storageService = new StorageService();
export default storageService;
