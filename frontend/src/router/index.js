import { createRouter, createWebHistory } from 'vue-router';

const routes = [
    {
        path: '/',
        name: 'index',
        component: () => import('../components/pages/UserLogin.vue')  // 懒加载路由组件
    },
    {
        path: '/login',
        name: 'login',
        component: () => import('../components/pages/UserLogin.vue')  // 懒加载路由组件
    },
    {
        path: '/register',
        name: 'register',
        component: () => import('../components/pages/UserRegister.vue')  // 懒加载路由组件
    },
    {
        path: '/home',
        name: 'home',
        component: () => import('../components/pages/UserHome.vue')  // 懒加载路由组件
    },
    {
        path: '/setting',
        name: 'setting',
        component: () => import('../components/pages/UserSetting.vue')  // 懒加载路由组件
    },
    {
        path: '/player',
        name: 'player',
        component: () => import('../components/player/playerContainer.vue')  // 懒加载路由组件
    },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL), // 使用 HTML5 模式
  routes
});

export default router;
