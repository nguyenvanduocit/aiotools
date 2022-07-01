<template>
  <ElContainer direction="vertical">
    <div :class="$style.row">
      <ElInput autofocus v-model="input" @keydown.enter="parse"/>
      <ElButton :loading="isLoading" :class="$style.button" type="primary" @click="parse" :disabled="!canParse">Parse URL</ElButton>
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
import {ParseUrl} from "../../../../wailsjs/go/parseurl/ParseUrl";
import {ElMessage} from "element-plus";
import {parseurl} from "../../../../wailsjs/go/models";

const data = ref();
const input = ref('')
const isLoading = ref(false)

const parse = () => {
  isLoading.value = true
  ParseUrl(input.value).then(res => {
    const result = res as parseurl.Response
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
  }).finally(() => {
    isLoading.value = false
  })
}

const canParse = computed(() => input.value.length)
</script>

<style module lang="stylus">
.attribute
  margin-top 20px
.input
  display flex
.button
  margin-left 10px
.row
  display flex
  justify-content space-between
  align-items center
  flex-wrap nowrap
</style>
