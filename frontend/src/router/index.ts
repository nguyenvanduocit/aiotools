import {createRouter, createWebHashHistory} from "vue-router";

const routes = [
  {
    path: "/",
    name: "home",
    component: () => import("../views/ViewHome.vue"),
  }
]

export const router = createRouter({
  history: createWebHashHistory(),
  routes
})
