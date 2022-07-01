<template>
  <div :class="$style.container">
    <ElForm :inline="true" :class="$style.form">
      <ElFormItem label="Path" :class="$style.flexGrow">
        <ElInput autofocus v-model="query" @keydown.enter="queryJson"/>
      </ElFormItem>
      <ElFormItem>
        <ElButton :class="$style.button" type="primary" @click="queryJson">Query</ElButton>
      </ElFormItem>
    </ElForm>
    <ElAlert v-if="error" :closable="false">{{error}}</ElAlert>
    <ElRow :gutter="15" :class="$style.codes">
      <ElCol :span="12">
        <Codemirror
            v-model="inputCode"
            :autofocus="true"
            :indent-with-tab="true"
            :tab-size="2"
            :extensions="extensions"
        ></Codemirror>
      </ElCol>
      <ElCol :span="12">
        <Codemirror
            v-model="outputCode"
            :indent-with-tab="true"
            :tab-size="2"
            :extensions="extensions"
        ></Codemirror>
      </ElCol>
    </ElRow>
  </div>
</template>

<script setup lang="ts">
import {getjson} from "../../../../wailsjs/go/models"
import {Get as getJson} from "../../../../wailsjs/go/getjson/GetJson"
import {json} from "@codemirror/lang-json"
import { Codemirror } from 'vue-codemirror'
import {oneDark} from "@codemirror/theme-one-dark"

const extensions = [json(), oneDark]

const query = ref('')
const inputCode = ref('')
const outputCode = ref('')
const error = ref()

const queryJson = () => {
  error.value = null
  const request: getjson.Input = {
    code: inputCode.value,
    query: query.value,
  }

  getJson(request).then((res) => {
    outputCode.value = res as string
  }).catch((err) => {
    error.value = err
  })
}

</script>

<style module lang="stylus">
.row
  display flex
  justify-content space-between
  align-items center
  flex-wrap nowrap
.container
  display flex
  flex-direction column
  justify-items center
  align-items stretch
  height 100%
.codes
  flex-grow 1
  :global(.cm-editor)
    min-height 100%
.form
  display flex
  justify-content space-between
  align-items center
  flex-wrap nowrap
.flexGrow
  flex-grow 1
</style>
