import { computed, ref } from 'vue'
import { defineStore } from 'pinia'
import axios from '../axios'

export const useAuthStore = defineStore('auth', () => {
  const token = ref<string | null>(localStorage.getItem('token'))
  const isAuthenticated = computed(() => !!token.value)
  const login = async (username: string, password: string) => {
    try {
      const response = await axios.post('/auth/login', { username, password })
      token.value = response.data.token
      localStorage.setItem('token', token.value || '')
    } catch (err) {
      throw new Error(`Login failed: ${err}`)
    }
  }

  const register = async (username: string, password: string) => {
    try {
      const response = await axios.post('/auth/register', { username, password })
      token.value = response.data.token
      localStorage.setItem('token', token.value || '')
    } catch (err) {
      throw new Error(`Register failed: ${err}`)
    }
  }

  const logout = () => {
    token.value = null
    localStorage.removeItem('token')
  }

  return { token, isAuthenticated, login, register, logout }
})
