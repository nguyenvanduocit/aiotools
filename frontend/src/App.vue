<script setup lang="ts">
import {useModuleStore} from "./store/modules";
import {onMounted} from "vue";
import {WindowShow} from "../wailsjs/runtime";
const { modules } = useModuleStore()
onMounted(() => {
  WindowShow()
})
</script>

<template>
  <ElContainer :class="$style.container">
    <ElAside width="200px" :class="$style.elAside">
      <ElScrollbar>
        <ElMenu :class="$style.elMenu" router>
          <ElMenuItem v-for="module in modules" :key="module.id" :index="module.id" :route="{ path: module.route }">{{ module.name }}</ElMenuItem>
        </ElMenu>
      </ElScrollbar>
    </ElAside>
    <ElMain :class="$style.main">
      <RouterView/>
    </ElMain>
  </ElContainer>
</template>

<style module lang="stylus">
.container
  height 100vh
  --el-menu-item-height: 40px
.elHeader
  position: relative;
  background-color: var(--el-color-primary-light-7);
  color: var(--el-text-color-primary);
  display flex
  justify-content center
  align-items center
.elAside
  border-right none;
.elMenu
  border-right none
.main
  background #000000
</style>
