import { createApp } from 'vue';
import App from './App.vue';
import router from './routers';
import http from '@common/utils/http';

const app = createApp(App);

app.config.globalProperties.$http = http;
app.use(router);
app.mount('#app');
