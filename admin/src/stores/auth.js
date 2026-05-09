import { defineStore } from 'pinia'
import { authApi } from '../api'

export const useAuthStore = defineStore('auth', {
  state: () => ({
    token: localStorage.getItem('admin_token') || '',
    user: JSON.parse(localStorage.getItem('admin_user') || 'null')
  }),
  getters: {
    isLogin: (s) => !!s.token,
    isAdmin: (s) => s.user?.vai_tro === 'admin'
  },
  actions: {
    async login(ten_dang_nhap, mat_khau) {
      const res = await authApi.login({ ten_dang_nhap, mat_khau })
      const data = res.data
      if (data.user?.vai_tro !== 'admin') {
        throw { message: 'Tai khoan khong co quyen quan tri' }
      }
      this.token = data.token
      this.user = data.user
      localStorage.setItem('admin_token', data.token)
      localStorage.setItem('admin_user', JSON.stringify(data.user))
    },
    async fetchProfile() {
      const res = await authApi.me()
      this.user = res.data
      localStorage.setItem('admin_user', JSON.stringify(res.data))
    },
    logout() {
      this.token = ''
      this.user = null
      localStorage.removeItem('admin_token')
      localStorage.removeItem('admin_user')
    }
  }
})
