import {router} from "../../router";
import {useModuleStore} from "../../store/modules";

export const install = () => {
  router.addRoute({
    path: "/modules/package-insights",
    component: () => import("./view/ViewPackageInsights.vue"),
  })

  const moduleStore = useModuleStore();
  moduleStore.addModule(
    {
      id: 'package-insights',
      name: 'NPM Packages',
      route: '/modules/package-insights'
    }
  )
}
