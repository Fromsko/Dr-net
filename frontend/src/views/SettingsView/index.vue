<script setup lang="ts">
import Tabs from '@/components/Tabs/index.vue'
import AboutView from '@/views/AboutView.vue'
import { ref } from 'vue'
import GeneralSetting from './components/GeneralSetting.vue'

const settings = [
  { key: 'general', tab: 'settings.general' },
  { key: 'about', tab: 'router.about' }
]

const activeKey = ref(settings[0].key)
</script>

<template>
  <Tabs v-model:active-key="activeKey" :items="settings">
    <template #general>
      <GeneralSetting />
    </template>

    <template #about>
      <AboutView />
    </template>
  </Tabs>
</template>

<style lang="less" scoped></style>
