<template>
  <TitleBar />
  <div class="main">
    <div v-if="isTargetDate" class="maintenance-message">
      <h3>公告</h3>
      <p>🎉 程序已停止维护</p>
      <p>如需调研程序请联系{{ contactInfo }}</p>
    </div>
    <div v-else>
      <NavigationBar />
      <MainPage />
    </div>
  </div>
</template>

<script setup lang="ts">
import MainPage from '@/components/MainPage.vue'
import NavigationBar from '@/components/NavigationBar.vue'
import TitleBar from '@/components/TitleBar.vue'

const currentDate = new Date()
const targetDate = new Date('2024-5-1')

const isTargetDate = currentDate.toDateString() === targetDate.toDateString()
const contactInfo = 'hnkong666@gmail.com'
</script>

<style scoped>
.main {
  flex: 1;
  overflow-y: auto;
  display: flex;
  flex-direction: column;
  padding: 8px;
  background: var(--bg-color);
}

.maintenance-message {
  display: flex;
  justify-content: center;
  align-items: center;
  margin-top: 20px;
  font-size: 20px;
}
</style>
