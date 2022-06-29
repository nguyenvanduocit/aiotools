<template>
  <div :class="$style.container">
    <div :class="$style.input">
      <ElInput autofocus v-model="input"/>
      <ElButton type="primary" :class="$style.button" @click="parse" :disabled="!canParse">Parse</ElButton>
    </div>
    <div :class="$style.attribute">
      <ElTable :data="tableData">
        <ElTableColumn prop="param" label="Params"></ElTableColumn>
        <ElTableColumn prop="value" label="Value"></ElTableColumn>
      </ElTable>
    </div>
  </div>
</template>

<script setup lang="ts">
import {ParseUrl} from "../../../../wailsjs/go/app/App";
import {app} from "../../../../wailsjs/go/models";
import {ElMessage} from "element-plus";

const tableData = ref([{}])

const input = ref("")
const parse = () => {
  ParseUrl(input.value).then(res => {
    const result = res as app.ParseResult
    tableData.value = Object.entries(result).map(([key, value]) => {
      return {
        param: key,
        value: value
      }
    })
  }).catch(err => {
    ElMessage({
      message: err,
      type: 'error',
    })
  })
}

const canParse = computed(() => input.value.length)
</script>

<style module lang="stylus">
.container
  height 100%
  width 100%
  grid-gap 20px
.attribute
  margin-top 20px
.input
  display flex
.button
  margin-left 10px
</style>
