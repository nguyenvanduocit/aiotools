import { defineStore } from 'pinia'
import {Module} from "../../type";

interface ModuleState {
  modules: Module[],
}

export const useModuleStore = defineStore('modules', {
  state: () => {
    return {
      modules: [
        {
          id: 'yaml-to-json',
          name: 'YAML to JSON',
          route: '/modules/yaml-to-json'
        }
      ],
    } as ModuleState
  },
  actions: {
    addModule(module: Module) {
      this.modules.push(module)
    }
  }
})
