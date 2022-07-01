<template>
  <div :class="$style.container">
    <ElRow :gutter="30" justify="space-around" align="middle">
      <ElSelect v-model="inputLanguage">
        <ElOption value="json">JSON</ElOption>
        <ElOption value="yaml">YAML</ElOption>
      </ElSelect>
      <ElButton type="primary" @click="convert" :disabled="!canTransform">Transform</ElButton>
      <ElSelect v-model="outputLanguage">
        <ElOption value="go">Golang</ElOption>
        <ElOption value="php">PHP</ElOption>
        <ElOption value="go-bson">Go BSON</ElOption>
      </ElSelect>
    </ElRow>

    <ElRow :gutter="15" :class="$style.code">
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
            :autofocus="false"
            :indent-with-tab="true"
            :tab-size="2"
            :extensions="extensions"
        ></Codemirror>
      </ElCol>
    </ElRow>
  </div>
</template>

<script setup lang="ts">
import { Codemirror } from 'vue-codemirror'
import {ElMessage} from "element-plus";
import { json } from '@codemirror/lang-json'
import { oneDark } from '@codemirror/theme-one-dark'
import {computed} from "vue";
import {TransformStruct} from "../../../../wailsjs/go/transformstruct/TransformStruct";

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
  height 100%
  display flex
  flex-direction column
  justify-content center
  align-items stretch
.code
  margin-top 18px
  flex-grow 1
 :global(.cm-editor)
    height 100%
</style>
