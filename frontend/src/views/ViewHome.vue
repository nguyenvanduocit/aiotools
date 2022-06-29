<script setup lang="ts">
import { useMainStore } from "../store"
const mainStore = useMainStore()

</script>

<template>
  <ElButton @click="mainStore.increment">count is: {{ mainStore.counter }}</ElButton>
</template>
