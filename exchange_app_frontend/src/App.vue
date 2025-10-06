<script setup lang="ts">
import { useRoute, useRouter } from 'vue-router'
import { ref, watch } from 'vue'
import { useAuthStore } from '@/stores/auth.ts'

const router = useRouter()
const route = useRoute()
const activeIndex = ref(route.name?.toString() || 'home')
const authStore = useAuthStore()
watch(route, (newRoute) => {
  activeIndex.value = newRoute.name?.toString() || 'home'
})

const handleSelect = (key: string) => {
  if (key == 'logout') {
    authStore.logout()
    router.push({ name: 'home' })
  } else {
    router.push({ name: key.charAt(0).toUpperCase() + key.slice(1) })
  }
}
</script>

<template>
  <el-container>
    <el-header>
      <el-menu
        :default-active="activeIndex"
        class="el-menu-demo"
        mode="horizontal"
        :ellipsis="true"
        @select="handleSelect"
      >
        <el-menu-item index="home">首页</el-menu-item>
        <el-menu-item index="currencyExchange">兑换货币</el-menu-item>
        <el-menu-item index="news">查看新闻</el-menu-item>
        <el-menu-item index="login" v-if="!authStore.isAuthenticated">登录</el-menu-item>
        <el-menu-item index="register" v-if="!authStore.isAuthenticated">注册</el-menu-item>
        <el-menu-item index="logout" v-if="authStore.isAuthenticated">退出</el-menu-item>
      </el-menu>
    </el-header>
    <el-main>
      <router-view></router-view>
    </el-main>
  </el-container>
</template>

<style scoped>
.el-menu-demo {
  line-height: 60px;
}
</style>
