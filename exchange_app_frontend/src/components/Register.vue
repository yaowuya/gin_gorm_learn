<script setup lang="ts">
import { ref } from 'vue'
import { useAuthStore } from '@/stores/auth.ts'
import { ElMessage } from 'element-plus'
import { useRouter } from 'vue-router'

const form = ref({ username: '', password: '' })
const authStore = useAuthStore()
const router = useRouter()
const register = async () => {
  try {
    await authStore.register(form.value.username, form.value.password)
    await router.push({ name: 'News' })
  } catch {
    ElMessage.error('注册失败，请重试')
  }
}
</script>

<template>
  <div class="auth-container">
    <el-form :model="form" class="auth-form" @submit.prevent="register">
      <el-form-item label="用户名" label-width="80px">
        <el-input v-model="form.username" placeholder="请输入用户名" />
      </el-form-item>
      <el-form-item label="密码" label-width="80px">
        <el-input v-model="form.password" type="password" placeholder="请输入密码" />
      </el-form-item>
      <el-button type="primary" native-type="submit">注册</el-button>
    </el-form>
  </div>
</template>

<style scoped>
.auth-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background-color: #f5f5f5;
  padding: 20px;
  box-sizing: border-box;
}

.auth-form {
  width: 100%;
  max-width: 360px;
  padding: 20px;
  background-color: #fff;
  border-radius: 4px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
}
</style>
