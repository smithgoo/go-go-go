const { defineConfig } = require('@vue/cli-service')
module.exports = defineConfig({
  transpileDependencies: true
})

module.exports = {
  devServer: {
    proxy: {
      '/user': {
        target: 'http://192.168.1.90:9000',  // 确保代理指向正确的服务器
        changeOrigin: true
      }
    }
  }
};

