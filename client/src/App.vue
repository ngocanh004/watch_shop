<script setup>
import { onMounted } from 'vue'
import { useToastStore } from './stores/toast'
import { useAuthStore } from './stores/auth'
import { useCartStore } from './stores/cart'

const toast = useToastStore()
const auth = useAuthStore()
const cart = useCartStore()

onMounted(() => {
  if (auth.isLogin) {
    cart.load().catch(() => {})
  }
})
</script>

<template>
  <router-view />
  <div class="toast-container">
    <div v-for="t in toast.items" :key="t.id" :class="['toast', 'toast-' + t.type]">
      {{ t.message }}
    </div>
  </div>
</template>
