<template>
  <div class="app">
    <h2>连接状态</h2>
    <div class="info">
      <div class="label">连接网络:</div>
      <div class="value">{{ ipLabel }}</div>
    </div>
    <div class="info">
      <div class="label">网络状况:</div>
      <div class="value">{{ statusLabel }}</div>
    </div>
    <div class="info">
      <div class="label">伪装状态:</div>
      <div class="value">{{ setStaLabel }}</div>
    </div>
  </div>
  <a-divider :margin="8"></a-divider>
  <div class="buttons">
    <Button type="primary" @click="clearData">清空</Button>
    <Button type="primary" @click="checkStatus">检测</Button>
  </div>
</template>

<script setup lang="ts">
import Button from '@/components/Button/index.vue'
import { CheckNetWork, GetClientIP, RestoreNetWork } from '@/utils/bridge'
import { onMounted, ref } from 'vue'
import { onBeforeRouteLeave } from 'vue-router'

const statusLabel = ref('')
const setStaLabel = ref('')
const ipLabel = ref('')

const checkStatus = async () => {
  ipLabel.value = await GetClientIP()
  statusLabel.value = await CheckNetWork()
  setStaLabel.value = localStorage.getItem('setStatus') === 'false' ? '未设置' : '已设置'
}

const clearData = async () => {
  if (setStaLabel.value === 'false') {
    localStorage.clear()
  } else {
    await RestoreNetWork()
    localStorage.removeItem('netData')
  }
}

onMounted(() => {
  ipLabel.value = localStorage.getItem('ipLabel') ?? '未知'
  statusLabel.value = localStorage.getItem('statusLabel') ?? '未知'
  setStaLabel.value = localStorage.getItem('setStaLabel') ?? '未知'
})

onBeforeRouteLeave((to, from, next) => {
  localStorage.setItem('ipLabel', ipLabel.value)
  localStorage.setItem('statusLabel', statusLabel.value)
  localStorage.setItem('setStaLabel', setStaLabel.value)
  next()
})
</script>

<style lang="less" scoped>
.app {
  text-align: center;
}

.info {
  display: flex;
  align-items: center;
  margin-bottom: 10px;
}

.label {
  flex: 0 0 100px;
  text-align: right;
  font-weight: bold;
}

.value {
  margin-left: 10px;
}

.buttons {
  margin-top: 20px;
  display: flex;
  justify-content: center;
  gap: 10px;
}
</style>
