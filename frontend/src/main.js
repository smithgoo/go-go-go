import { createApp } from 'vue';
import App from './App.vue';
import Vant from 'vant';
import 'vant/lib/index.css'; // 引入 Vant 的 CSS 样式
import router  from './router';

createApp(App)
.use(router)
.use(Vant) 
.mount('#app');