import {router} from "../../router";
import {useModuleStore} from "../../store/modules";

export const install = () => {
  router.addRoute({
    path: "/modules/npm-insight",
    component: () => import("./view/ViewNpmInsight.vue"),
  })

  const moduleStore = useModuleStore();
  moduleStore.addModule(
    {
      id: 'npm-insight',
      name: 'NPM Insight',
      route: '/modules/npm-insight'
    }
  )
}
