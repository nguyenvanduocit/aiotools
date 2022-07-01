import {createApp} from 'vue'
import App from './App.vue'
import {createPinia} from "pinia";
import {router} from "./router";

import('element-plus/theme-chalk/dark/css-vars.css')
import('modern-normalize/modern-normalize.css')

import {install as installParseUrlPlugin} from "./modules/urlParser";
import {install as installTransformStruct} from "./modules/transformStruct";
import {install as installDnsAnalytics} from "./modules/dnsAnalytics";
import {install as installPackageInsights} from "./modules/packageInsights";
import {install as installGetJson} from "./modules/getJson";
import {install as installSetJson} from "./modules/setJson";

const pinia = createPinia()
const app = createApp(App)
  .use(pinia)

installParseUrlPlugin()
installDnsAnalytics()
installPackageInsights()
installTransformStruct()
installGetJson()
installSetJson()

app.use(router)
app.mount('#app')
