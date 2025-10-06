import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '@/views/HomeView.vue'

const routes: RouteRecordRaw[] = [{ path: '/', name: 'Home', component: HomeView }]

const router = createRouter({
  history: createWebHistory(),
  routes: routes,
})

export default router
