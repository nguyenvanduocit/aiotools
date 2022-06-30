import {createApp} from 'vue'
import App from './App.vue'
import {createPinia} from "pinia";
import {router} from "./router";

import('element-plus/theme-chalk/dark/css-vars.css')
import('modern-normalize/modern-normalize.css')

import {install as installUrlParser} from "./modules/urlParser";
import {install as installTransformStruct} from "./modules/transformStruct";
import {install as installDnsAnalytics} from "./modules/dnsAnalytics";
import {install as installNpmInsight} from "./modules/npmInsight";

const pinia = createPinia()
const app = createApp(App)
  .use(pinia)

installUrlParser()
installTransformStruct()
installDnsAnalytics()
installNpmInsight()

app.use(router)
app.mount('#app')
