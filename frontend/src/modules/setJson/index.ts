import {router} from "../../router";
import {useModuleStore} from "../../store/modules";

export const install = () => {
  router.addRoute({
    path: "/modules/set-json",
    component: () => import("./view/ViewSetJson.vue"),
  })

  const moduleStore = useModuleStore();
  moduleStore.addModule(
    {
      id: 'set-json',
      name: 'Set JSON',
      route: '/modules/set-json'
    }
  )
}
