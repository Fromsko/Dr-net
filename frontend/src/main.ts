import ArcoVue, { Notification } from '@arco-design/web-vue';
import '@arco-design/web-vue/dist/arco.css';
import ArcoVueIcon from '@arco-design/web-vue/es/icon';
import { createPinia } from 'pinia';
import { createApp } from 'vue';
import './assets/main.less';

import App from './App.vue';
import i18n from './lang';
import router from './router';

const app = createApp(App)

app.use(ArcoVue)
app.use(ArcoVueIcon);
app.use(createPinia())
app.use(router)
app.use(i18n)

Notification._context = app._context;

app.mount('#app')
