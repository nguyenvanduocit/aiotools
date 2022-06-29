import {router} from "../../router";
import {useModuleStore} from "../../store/modules";

export const install = () => {
  router.addRoute({
    path: "/modules/transform",
    component: () => import("./view/ViewJsonToGo.vue"),
  })

  const moduleStore = useModuleStore();
  moduleStore.addModule(
    {
      id: 'transform',
      name: 'Transform struct',
      route: '/modules/transform'
    }
  )
}
