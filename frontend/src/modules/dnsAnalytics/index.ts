import {router} from "../../router";
import {useModuleStore} from "../../store/modules";

export const install = () => {
  router.addRoute({
    path: "/modules/dns-analytics",
    component: () => import("./view/ViewDnsAnalytics.vue"),
  })

  const moduleStore = useModuleStore();
  moduleStore.addModule(
    {
      id: 'dns-analytics',
      name: 'DNS Analytics',
      route: '/modules/dns-analytics'
    }
  )
}
