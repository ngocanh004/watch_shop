<script setup>
import { useToastStore } from './stores/toast'
const toast = useToastStore()
</script>

<template>
  <router-view />
  <div class="toast-container">
    <div v-for="t in toast.items" :key="t.id" :class="['toast', 'toast-' + t.type]">
      {{ t.message }}
    </div>
  </div>
</template>

<style>
.toast-container {
  position: fixed;
  top: 20px;
  right: 20px;
  display: flex;
  flex-direction: column;
  gap: 8px;
  z-index: 9999;
}
</style>
