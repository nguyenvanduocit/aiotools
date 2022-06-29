import { defineStore } from 'pinia'
import {Module} from "../../type";

interface ModuleState {
  modules: Module[],
}

export const useModuleStore = defineStore('modules', {
  state: () => {
    return {
      modules: [],
    } as ModuleState
  },
  actions: {
    addModule(module: Module) {
      this.modules.push(module)
    }
  }
})
