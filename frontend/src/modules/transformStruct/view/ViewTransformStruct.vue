<template>
  <ElRow :gutter="30" justify="space-around" align="middle">
    <ElSelect v-model="inputLanguage">
      <ElOption value="json">JSON</ElOption>
      <ElOption value="yaml">YAML</ElOption>
    </ElSelect>
    <ElButton type="primary" @click="convert" :disabled="!canTransform">Transform</ElButton>
    <ElSelect v-model="outputLanguage">
      <ElOption value="go">Golang</ElOption>
      <ElOption value="php">PHP</ElOption>
    </ElSelect>
  </ElRow>
  <ElDivider/>
  <ElRow :gutter="30">
    <ElCol :span="12">
      <Codemirror
          :class="$style.container"
          v-model="inputCode"
          :autofocus="true"
          :indent-with-tab="true"
          :tab-size="2"
          :extensions="extensions"
      ></Codemirror>
    </ElCol>
    <ElCol :span="12">
      <Codemirror
          :class="$style.container"
          v-model="outputCode"
          :autofocus="false"
          :indent-with-tab="true"
          :tab-size="2"
          :extensions="extensions"
      ></Codemirror>
    </ElCol>
  </ElRow>
</template>

<script setup lang="ts">
import { Codemirror } from 'vue-codemirror'
import {TransformStruct} from "../../../../wailsjs/go/app/App";
import {ElMessage} from "element-plus";
import { json } from '@codemirror/lang-json'
import { oneDark } from '@codemirror/theme-one-dark'
import {computed} from "vue";

const extensions = [json(), oneDark]

const inputLanguage = ref('json')
const outputLanguage = ref('go')

const inputCode = ref("");
const outputCode = ref("");

const canTransform = computed(() => inputCode.value.length > 0)

const convert = () => {
  TransformStruct({
    SourceCode: inputCode.value,
  }).then((destinationCoude) => {
    outputCode.value = destinationCoude;
  }).catch((err) => {
    ElMessage({
      message: err,
      type: 'error',
    })
  });
}
</script>

<style lang="stylus" module>
.container
 :global(.cm-editor)
  min-height 380px
.input
  height 100%
</style>
