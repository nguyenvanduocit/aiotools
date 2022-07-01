<template>

</template>

<script setup lang="ts">
import { Codemirror } from 'vue-codemirror'
import { json } from '@codemirror/lang-json'
import { oneDark } from '@codemirror/theme-one-dark'
const extensions = [json(), oneDark]

defineProps({
  input: {
    type: String,
    default: '',
  },
  output: {
    type: String,
    default: '',
  },
})

const inputCodeChanged = () => {
  console.log('inputCodeChanged')
}

</script>

<style lang="stylus" module>
.container
 :global(.cm-editor)
  min-height 380px
.input
  height 100%
</style>
