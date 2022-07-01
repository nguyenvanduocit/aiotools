import {router} from "../../router";
import {useModuleStore} from "../../store/modules";

export const install = () => {
  router.addRoute({
    path: "/modules/get-json",
    component: () => import("./view/ViewGetJson.vue"),
  })

  const moduleStore = useModuleStore();
  moduleStore.addModule(
    {
      id: 'get-json',
      name: 'Get JSON',
      route: '/modules/get-json'
    }
  )
}
