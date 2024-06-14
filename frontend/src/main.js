import { createApp } from 'vue';
import App from './App.vue';
import { createRouter, createWebHistory } from 'vue-router';

// 导入你的组件
import UserLogin from './components/UserLogin.vue';
import UserRegister from './components/UserRegister.vue';

const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: '/login', component: UserLogin },
    { path: '/register', component: UserRegister },
    // 可以根据需要添加其他路由
  ],
});

createApp(App).use(router).mount('#app');
