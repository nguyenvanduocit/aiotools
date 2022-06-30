<template>
  <ElContainer direction="vertical">
    <div :class="$style.row">
      <ElInput autofocus v-model="input"/>
      <ElButton :class="$style.button" type="primary" @click="parse" :disabled="!canParse">Get Insights</ElButton>
    </div>
    <div :class="$style.attribute">
      <ElTable :data="data">
        <ElTableColumn width="150px" prop="collected.metadata.name" label="Name"></ElTableColumn>
        <ElTableColumn width="100px" prop="collected.metadata.version" label="Version"></ElTableColumn>
        <ElTableColumn width="100px" prop="collected.metadata.license" label="License"></ElTableColumn>
        <ElTableColumn width="100px" prop="collected.github.starsCount" label="Stars"></ElTableColumn>
        <ElTableColumn width="200px" prop="collected.github.issues.openCount" label="Issues"></ElTableColumn>
      </ElTable>
    </div>
  </ElContainer>
</template>

<script setup lang="ts">
import {GetPackageInsight, ParseUrl} from "../../../../wailsjs/go/app/App";
import {app} from "../../../../wailsjs/go/models";
import {ElMessage} from "element-plus";
import GetPackageInsightsRequest = app.GetPackageInsightsRequest;

const data = ref();
const input = ref("")
const parse = () => {
  const request: GetPackageInsightsRequest = {
    package_name: input.value,
  }

  GetPackageInsight(request).then(res => {
    const result = res as app.GetPackageInsightsResponse
    console.log(result)
  }).catch(err => {
    console.log(err)
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
