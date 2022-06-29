import {router} from "../../router";
import {useModuleStore} from "../../store/modules";

export const install = () => {
  router.addRoute({
    path: "/modules/url-parser",
    component: () => import("./view/ViewUrlParser.vue"),
  })

  const moduleStore = useModuleStore();
  moduleStore.addModule(
    {
      id: 'url-parser',
      name: 'URL Parser',
      route: '/modules/url-parser'
    }
  )
}
