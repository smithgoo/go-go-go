import { createApp } from 'vue';
import App from './App.vue';
import Vant, { Toast } from 'vant';
import 'vant/lib/index.css';
import { createRouter, createWebHistory } from 'vue-router';

// 导入你的组件
import UserLogin from './components/UserLogin.vue';
import UserRegister from './components/UserRegister.vue';

const router = createRouter({
    history: createWebHistory(),
    routes: [
        { path: '/', redirect: '/register' },
        { path: '/register', component: UserRegister },
        { path: '/login', component: UserLogin },
        // 可以根据需要添加其他路由
    ],
});

const app = createApp(App);

app.use(router);
app.use(Vant);
app.use(Toast);

app.mount('#app');
