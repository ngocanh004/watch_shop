import { defineStore } from 'pinia'

export const useToastStore = defineStore('toast', {
  state: () => ({
    items: []
  }),
  actions: {
    push(message, type = 'info') {
      const id = Date.now() + Math.random()
      this.items.push({ id, message, type })
      setTimeout(() => this.remove(id), 3000)
    },
    success(m) { this.push(m, 'success') },
    error(m) { this.push(m, 'error') },
    info(m) { this.push(m, 'info') },
    remove(id) {
      this.items = this.items.filter(i => i.id !== id)
    }
  }
})
