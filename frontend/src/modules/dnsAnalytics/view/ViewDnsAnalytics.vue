<template>
  <div :class="$style.container">
    <div :class="$style.row">
      <ElInput placeholder="12bit.vn" autofocus v-model="host" @keydown.enter="lookup"/>
      <ElButton :class="$style.button" type="primary" @click="lookup" :disabled="!canParse">Analyze</ElButton>
    </div>

    <ElAlert  type="error" v-if="error" :closable="false">{{error}}</ElAlert>
    <div :class="$style.code">
      <Codemirror
          v-if="code"
          v-model="code"
          :autofocus="true"
          :indent-with-tab="true"
          :tab-size="2"
          :extensions="extensions"
      ></Codemirror>
      <ElEmpty v-else></ElEmpty>
    </div>
  </div>
</template>

<script setup lang="ts">
import {dnslookup} from "../../../../wailsjs/go/models";
import {Lookup} from "../../../../wailsjs/go/dnslookup/DnsLookup";
import {Codemirror} from "vue-codemirror";
import {computed} from "vue";
import {json} from "@codemirror/lang-json";
import {oneDark} from "@codemirror/theme-one-dark";
const data = ref();
const host = ref("")
const extensions = [json(), oneDark]
const error = ref()

const code = computed(() => {
  return JSON.stringify(data.value, null, "\t")
})

const lookup = () => {
  error.value = null
  const input: dnslookup.Input = {
    host: host.value,
  }
  Lookup(input).then(res => {
    const result = res as dnslookup.Output
    data.value = result
  }).catch(err => {
    error.value = err
  })
}

const canParse = computed(() => host.value.length)

</script>

<style module lang="stylus">
.container
  display flex
  flex-direction column
  justify-items center
  align-items stretch
  height 100%
.code
  flex-grow 1
  margin-top 20px
  :global(.cm-editor)
    height 100%
.button
  margin-left 10px
.row
  display flex
  justify-content space-between
  align-items center
  flex-wrap nowrap
</style>
