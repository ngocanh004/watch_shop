import { defineStore } from 'pinia'
import { authApi } from '../api'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    token: localStorage.getItem('client_token') || '',
    user: JSON.parse(localStorage.getItem('client_user') || 'null')
  }),
  getters: {
    isLogin: (s) => !!s.token
  },
  actions: {
    async register(data) {
      return authApi.register(data)
    },
    async login(ten_dang_nhap, mat_khau) {
      const res = await authApi.login({ ten_dang_nhap, mat_khau })
      this.token = res.data.token
      this.user = res.data.user
      localStorage.setItem('client_token', this.token)
      localStorage.setItem('client_user', JSON.stringify(this.user))
    },
    async fetchProfile() {
      const res = await authApi.me()
      this.user = res.data
      localStorage.setItem('client_user', JSON.stringify(res.data))
    },
    logout() {
      this.token = ''
      this.user = null
      localStorage.removeItem('client_token')
      localStorage.removeItem('client_user')
    }
  }
})
