<script setup lang="ts">
import Button from '@/components/Button/index.vue'
import { APP_TITLE, APP_VERSION } from '@/utils/env'
import { Quit, WindowFullscreen, WindowMinimise, WindowSetSize } from '@wails/runtime'
import { ref } from 'vue'

const count = ref(0)

const handleClick = () => {
  count.value++
  console.log(count.value)

  if (count.value === 1) {
    WindowFullscreen()
  } else if (count.value === 2) {
    WindowSetSize(800, 550)
    count.value = 0
  }
}
</script>

<template>
  <div class="titlebar" style="--wails-draggable: drag">
    <img class="logo" draggable="false" src="@/assets/logo.png" />
    <div class="appname">{{ APP_TITLE }} {{ APP_VERSION }}</div>
    <div class="action">
      <Button @click="WindowMinimise" type="text"> - </Button>
      <Button @click="handleClick" type="text"> + </Button>
      <Button @click="Quit" type="text"> × </Button>
    </div>
  </div>
</template>

<style lang="less" scoped>
.titlebar {
  user-select: none;
  display: flex;
  padding: 4px 12px;
  align-items: center;
}
.logo {
  width: 32px;
  height: 32px;
  user-select: none;
}
.appname {
  font-size: 14px;
  margin-left: 4px;
  margin-top: 4px;
  font-weight: bold;
}

.action {
  display: flex;
  align-items: center;
  justify-content: center;
  margin-left: auto;
  font-size: 14px;
  &-btn {
    width: 32px;
    height: 32px;
    line-height: 32px;
    text-align: center;
    border-radius: 4px;
    &:hover {
      background-color: var(--hover-bg-color);
    }
  }
}
</style>
