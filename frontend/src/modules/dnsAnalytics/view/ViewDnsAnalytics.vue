<template>
  <ElContainer direction="vertical">
    <div :class="$style.row">
      <ElInput autofocus v-model="input"/>
      <ElButton :class="$style.button" type="primary" @click="parse" :disabled="!canParse">Analyze</ElButton>
    </div>
    <div :class="$style.attribute">
      <ElTable :data="data">
        <ElTableColumn width="150px" prop="param" label="Params"></ElTableColumn>
        <ElTableColumn prop="value" label="Value"></ElTableColumn>
      </ElTable>
    </div>
  </ElContainer>
</template>

<script setup lang="ts">
import {ParseUrl} from "../../../../wailsjs/go/app/App";
import {app} from "../../../../wailsjs/go/models";
import {ElMessage} from "element-plus";

const data = ref();
const input = ref("")
const parse = () => {
  ParseUrl(input.value).then(res => {
    const result = res as app.ParseResult
    data.value = Object.entries(result).map(([key, value]) => {
      return {
        param: key,
        value: value as string
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
.attribute
  margin-top 20px
.button
  margin-left 10px
.row
  display flex
  justify-content space-between
  align-items center
  flex-wrap nowrap
</style>
