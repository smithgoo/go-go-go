// storage.js
class StorageService {
    constructor(storageKey) {
      this.storageKey = storageKey;
    }
  
    setItem(data) {
      localStorage.setItem(this.storageKey, JSON.stringify(data));
    }
  
    getItem() {
      const data = localStorage.getItem(this.storageKey);
      return data ? JSON.parse(data) : null;
    }
  
    removeItem() {
      localStorage.removeItem(this.storageKey);
    }
  
    clear() {
      localStorage.clear();
    }
  }
  
  // 导出一个实例，使用指定的 key
  const userStorage = new StorageService('user');
  
  export default userStorage;
  