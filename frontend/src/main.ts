import {createApp} from 'vue'
import App from './App.vue'
import {createPinia} from "pinia";
import {router} from "./router";
import {install as installUrlParser} from "./modules/urlParser";
import {install as installJsonToGo} from "./modules/jsonToGo";
import 'element-plus/theme-chalk/dark/css-vars.css'
import('modern-normalize/modern-normalize.css')
const pinia = createPinia()
const app = createApp(App)
  .use(pinia)

installUrlParser()
installJsonToGo()

app.use(router)
app.mount('#app')
