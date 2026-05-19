import { defineStore } from 'pinia'
import { cartApi } from '../api'

export const useCartStore = defineStore('cart', {
  state: () => ({
    items: [],
    tong_tien: 0,
    loading: false
  }),
  getters: {
    count: (s) => s.items.reduce((a, b) => a + b.so_luong, 0)
  },
  actions: {
    async load() {
      this.loading = true
      try {
        const res = await cartApi.list()
        this.items = res.data.items || []
        this.tong_tien = res.data.tong_tien || 0
      } finally {
        this.loading = false
      }
    },
    async add(product_id, so_luong = 1) {
      await cartApi.add(product_id, so_luong)
      await this.load()
    },
    async update(id, so_luong) {
      await cartApi.update(id, so_luong)
      await this.load()
    },
    async remove(id) {
      await cartApi.remove(id)
      await this.load()
    },
    async clear() {
      await cartApi.clear()
      this.items = []
      this.tong_tien = 0
    },
    reset() {
      this.items = []
      this.tong_tien = 0
    }
  }
})
