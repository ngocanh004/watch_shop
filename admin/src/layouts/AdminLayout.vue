<script setup>
import { computed } from 'vue'
import { useRouter, useRoute } from 'vue-router'
import { useAuthStore } from '../stores/auth'

const router = useRouter()
const route = useRoute()
const auth = useAuthStore()

const menu = [
  { path: '/dashboard', label: 'Tổng quan', icon: '📊' },
  { path: '/products', label: 'Sản phẩm', icon: '⌚' },
  { path: '/categories', label: 'Danh mục', icon: '📁' },
  { path: '/brands', label: 'Thương hiệu', icon: '🏷️' },
  { path: '/orders', label: 'Đơn hàng', icon: '🛒' },
  { path: '/users', label: 'Người dùng', icon: '👥' }
]

const currentPath = computed(() => route.path)

function logout() {
  auth.logout()
  router.push('/login')
}
</script>

<template>
  <div class="admin-wrap">
    <aside class="sidebar">
      <div class="brand">
        <span class="brand-icon">⌚</span>
        <div>
          <div class="brand-name">Watch Shop</div>
          <div class="brand-sub">Trang quản trị</div>
        </div>
      </div>
      <nav class="menu">
        <router-link
          v-for="m in menu"
          :key="m.path"
          :to="m.path"
          :class="['menu-item', { active: currentPath.startsWith(m.path) }]"
        >
          <span class="menu-icon">{{ m.icon }}</span>
          <span>{{ m.label }}</span>
        </router-link>
      </nav>
    </aside>

    <div class="content">
      <header class="topbar">
        <div></div>
        <div class="user-info">
          <router-link to="/profile" class="user-name">
            👤 {{ auth.user?.ho_ten }}
          </router-link>
          <button class="btn btn-sm btn-outline" @click="logout">Đăng xuất</button>
        </div>
      </header>
      <main class="main">
        <router-view />
      </main>
    </div>
  </div>
</template>

<style scoped>
.admin-wrap { display: flex; min-height: 100vh; }
.sidebar {
  width: 240px;
  background: #1e293b;
  color: #fff;
  display: flex;
  flex-direction: column;
}
.brand {
  display: flex;
  align-items: center;
  gap: 10px;
  padding: 20px;
  border-bottom: 1px solid #334155;
}
.brand-icon { font-size: 28px; }
.brand-name { font-weight: 700; font-size: 16px; }
.brand-sub { font-size: 12px; color: #94a3b8; }
.menu { padding: 12px 0; }
.menu-item {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 10px 20px;
  color: #cbd5e1;
  text-decoration: none;
  transition: all .15s;
}
.menu-item:hover { background: #334155; color: #fff; text-decoration: none; }
.menu-item.active { background: #2563eb; color: #fff; }
.menu-icon { font-size: 18px; width: 20px; }
.content { flex: 1; display: flex; flex-direction: column; }
.topbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 14px 24px;
  background: #fff;
  border-bottom: 1px solid #e5e7eb;
}
.user-info { display: flex; align-items: center; gap: 12px; }
.user-name { color: #1f2937; text-decoration: none; font-weight: 500; }
.main { padding: 24px; flex: 1; }
</style>
