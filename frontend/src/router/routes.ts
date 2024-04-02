import HomeView from '@/views/HomeView/index.vue'
import SettingsView from '@/views/SettingsView/index.vue'
import { type RouteRecordRaw } from 'vue-router'

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    name: 'Home',
    component: HomeView,
    meta: {
      name: 'router.overview'
    }
  },
  {
    path: '/settings',
    name: 'Settings',
    component: SettingsView,
    meta: {
      name: 'router.settings'
    }
  }
]

export default routes
