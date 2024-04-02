<script setup lang="ts">
import Radio from '@/components/Radio/index.vue'
import { useAppSettingsStore } from '@/stores/appSettings'
import { ClipboardSetText } from '@/utils/bridge'
import { useI18n } from 'vue-i18n'

const { t } = useI18n()
const appSettings = useAppSettingsStore()

const themes = [
  {
    label: 'settings.theme.dark',
    value: 'dark'
  },
  {
    label: 'settings.theme.light',
    value: 'light'
  },
  {
    label: 'settings.theme.auto',
    value: 'auto'
  }
]

const langs = [
  {
    label: 'settings.lang.zh',
    value: 'zh'
  },
  {
    label: 'settings.lang.en',
    value: 'en'
  }
]
</script>

<template>
  <div class="settings">
    <h3>{{ t('settings.theme.name') }}</h3>
    <Radio v-model="appSettings.app.themeMode" :options="themes" />
    <h3>{{ t('settings.lang.name') }}</h3>
    <Radio v-model="appSettings.app.lang" :options="langs" />
  </div>
  <a-divider :margin="10"></a-divider>
  <div class="show">
    <h3>@Fromsko 倾心制作</h3>
    <p>
      如有侵权请联系删除:
      <a
        href="#"
        @click="
          () => {
            ClipboardSetText('hnkong666@gmail.com')
          }
        "
        >hnkong666@gmail.com</a
      >
    </p>
    <p>该程序仅供研究学习 WAF 绕过技术，请于下载后24h内删除。</p>
  </div>
</template>

<style lang="less" scoped>
.about {
  padding: 22px 0 0 0;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  .appname {
    font-weight: bold;
    font-size: 16px;
    line-height: 4;
  }
}
</style>
