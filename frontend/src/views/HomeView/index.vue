<script setup lang="ts">
import Button from '@/components/Button/index.vue'
import { ClipboardSetText, RestoreNetWork, SetNetWork, StartScan } from '@/utils/bridge'
import { Notification } from '@arco-design/web-vue'
import { onMounted, ref, watchEffect } from 'vue'
import { useI18n } from 'vue-i18n'
import { onBeforeRouteLeave } from 'vue-router'

// 更新net接口
interface net {
  mac: string
  ip: string
  ttl: string
  using: boolean // 新增字段表示是否正在使用
}

const loading = ref(false)
const scaned = ref(false)
const { t } = useI18n()
const netData = ref(new Array<net>().map((_) => ({ ..._, using: false })))

const notify = async () => {
  Notification.success({
    title: '扫描状态',
    content: '小狐狸已经扫描完毕!😀',
    closable: false,
    style: { width: '300px', height: '75px' }
  })
}

const handleClick = async () => {
  loading.value = true
  let results: any = await StartScan()

  if (results !== null) {
    netData.value.push(...results)

    setTimeout(async () => {
      // 事件完成后解除加载状态
      loading.value = false
      scaned.value = true
      await notify()
    }, 2000)
  }
}

const onUseNetwork = async (index: number, networkItem: net) => {
  // 如果当前网络项正在使用（using为true）
  if (networkItem.using) {
    // 执行恢复操作
    let restoreInfo = await RestoreNetWork()

    // 恢复成功后，将所有项的 using 属性设为 false
    netData.value.forEach((item) => {
      item.using = false
    })

    if (restoreInfo === '恢复成功') {
      // 将当前项的 using 设为 false，因为它已被释放
      networkItem.using = false
      localStorage.setItem('setStatus', 'false')
    }
  } else {
    // 对于未使用的项，执行设置操作
    let setInfo = await SetNetWork(networkItem.ip)

    if (setInfo === '伪装失败') {
      // 移除当前失败的数据源
      netData.value.splice(index, 1)
    } else {
      // 恢复成功后，将所有项的 using 属性设为 false
      netData.value.forEach((item) => {
        item.using = false
      })

      // 设置成功，将当前网络项的 using 设为 true
      networkItem.using = true
      localStorage.setItem('setStatus', 'true')
    }
  }

  // 触发视图刷新
  netData.value = [...netData.value]
}

const onRemoveItem = (index: number) => {
  netData.value.splice(index, 1)
}

// 初始化时尝试从localStorage中恢复状态
onMounted(() => {
  const savedNetData = localStorage.getItem('netData')
  const savedScaned = localStorage.getItem('scaned')

  if (savedNetData && savedScaned) {
    netData.value = JSON.parse(savedNetData) // 取消硬编码的using: false
    scaned.value = JSON.parse(savedScaned)

    // 检查初始化时netData是否为空，为空则重置scaned
    if (netData.value.length === 0) {
      scaned.value = false
    }
  } else {
    // 如果localStorage中没有数据，也应确保scaned为false以显示"快速启动"
    scaned.value = false
  }
})

// 在离开路由之前将状态保存到localStorage
onBeforeRouteLeave((to, from, next) => {
  localStorage.setItem('netData', JSON.stringify(netData.value))
  localStorage.setItem('scaned', JSON.stringify(scaned.value))
  next()
})

watchEffect(() => {
  if (netData.value.length === 0) {
    scaned.value = false
  }
})

const handleItemClick = (ip: string) => {
  ClipboardSetText(ip)
  Notification.success({
    title: '',
    content: '已复制',
    closable: false,
    style: { width: '170px', height: '45px' }
  })
}
</script>

<template>
  <keep-alive>
    <div class="overview" v-show="!scaned">
      <div class="quick-start">
        <p>{{ t('home.noProfile') }}</p>
        <Button type="primary" @click="handleClick" :loading="loading">
          {{ t('home.quickStart') }}
        </Button>
      </div>
    </div>

    <div class="quick-start" v-show="scaned">
      <a-back-top
        target-container="#list-top"
        :style="{ position: 'fixed', right: '50px', bottom: '50px' }"
      />
      <a-list>
        <a-list-item v-for="(item, index) in netData" :key="index" id="list-top">
          <a-list-item-meta
            :title="`局域网地址: ${item.ip}`"
            :description="`网络时延: ${item.ttl}`"
            @click="handleItemClick(item.ip)"
          />

          <template #actions>
            <a-button-group shape="round" type="primary" size="small">
              <a-button
                :status="item.using ? 'warning' : 'success'"
                @click="onUseNetwork(index, item)"
                >{{ item.using ? '恢复' : '使用' }}</a-button
              >
              <a-button status="danger" :disabled="item.using" @click="onRemoveItem(index)">
                移除</a-button
              >
            </a-button-group>
          </template>
        </a-list-item>
      </a-list>
    </div>
  </keep-alive>
</template>

<style lang="less" scoped>
.overview {
  .quick-start {
    display: flex;
    min-height: 66vh;
    flex-direction: column;
    justify-content: center;
    align-items: center;

    p {
      margin-bottom: 20px;
    }
  }
}
</style>
