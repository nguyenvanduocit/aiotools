<template>
  <ElContainer direction="vertical">
    <div :class="$style.row">
      <ElInput placeholder="package name: vue" autofocus v-model="input" @keydown.enter="parse"/>
      <ElButton :class="$style.button" type="primary" @click="parse" :disabled="!canParse" :loading="isLoading">Get Insights</ElButton>
    </div>
    <ElAlert type="error" v-if="error" :closable="false">{{error}}</ElAlert>
    <template v-if="data">
      <div v-if="data" :class="$style.attribute">
        <ElDescriptions
            direction="vertical"
            border
            :column="4"
        >
          <ElDescriptionsItem label="Name">{{data.collected.metadata.name}}</ElDescriptionsItem>
          <ElDescriptionsItem label="Version">{{data.collected.metadata.version}}</ElDescriptionsItem>
          <ElDescriptionsItem label="Updated">{{data.collected.metadata.date}}</ElDescriptionsItem>
          <ElDescriptionsItem label="License">{{data.collected.metadata.license}}</ElDescriptionsItem>
          <template v-if="data.collected.github">
            <ElDescriptionsItem label="Stars">{{data.collected.github.starsCount}}</ElDescriptionsItem>
            <ElDescriptionsItem label="Forks">{{data.collected.github.forksCount}}</ElDescriptionsItem>
            <ElDescriptionsItem label="Open Issues">{{data.collected.github.issues.openCount}}</ElDescriptionsItem>
          </template>
          <ElDescriptionsItem label="Links">
            <ElButton @click="openUrl(data.collected.metadata.links.npm)" size="small">NPM</ElButton>
            <ElButton @click="openUrl(data.collected.metadata.links.homepage)" size="small">Home</ElButton>
            <ElButton @click="openUrl(data.collected.metadata.links.bugs)" size="small">Repository</ElButton>
            <ElButton @click="openUrl(data.collected.metadata.links.repository)" size="small">Bugs</ElButton>
          </ElDescriptionsItem>
        </ElDescriptions>

        <ElDivider/>

        <ElDescriptions
            direction="vertical"
            border
            :column="4"
        >
          <ElDescriptionsItem label="Carefulness">{{data.evaluation.quality.carefulness}}</ElDescriptionsItem>
          <ElDescriptionsItem label="Tests">{{data.evaluation.quality.tests}}</ElDescriptionsItem>
          <ElDescriptionsItem label="Health">{{data.evaluation.quality.health}}</ElDescriptionsItem>
          <ElDescriptionsItem label="Branding">{{data.evaluation.quality.branding}}</ElDescriptionsItem>
          <ElDescriptionsItem label="Interest">{{data.evaluation.popularity.communityInterest}}</ElDescriptionsItem>
          <ElDescriptionsItem label="Downloads">{{data.evaluation.popularity.downloadsCount}}</ElDescriptionsItem>
        </ElDescriptions>
      </div>
    </template>
    <ElEmpty v-else-if="!error" description="Type something"></ElEmpty>
  </ElContainer>
</template>

<script setup lang="ts">
import {GetPackageInsight} from "../../../../wailsjs/go/packageinsights/PackageInsights";
import {packageinsights} from "../../../../wailsjs/go/models";
import {onMounted} from "vue";
import {BrowserOpenURL} from "../../../../wailsjs/runtime";

const data = ref<packageinsights.Response|null>();
const input = ref("")
const isLoading = ref(false)
const error = ref()

const openUrl = (url: string) => {
  window.open(url, '_blank')
  BrowserOpenURL(url)
}

const parse = () => {
  error.value = null
  const request: packageinsights.Request = {
    package_name: input.value,
  }

  isLoading.value = true

  GetPackageInsight(request).then(res => {
    console.log(res)
    data.value = res as packageinsights.Response
  }).catch(err => {
    error.value = err
    data.value = null
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
.description
  margin-top 20px
  margin-bottom 10px
  font-size 14px
</style>
