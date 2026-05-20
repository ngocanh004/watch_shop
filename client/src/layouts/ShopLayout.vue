<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '../stores/auth'
import { useCartStore } from '../stores/cart'

const router = useRouter()
const auth = useAuthStore()
const cart = useCartStore()
const keyword = ref('')
const showMenu = ref(false)

function search() {
  if (keyword.value.trim()) {
    router.push({ path: '/products', query: { keyword: keyword.value } })
  }
}

function logout() {
  auth.logout()
  cart.reset()
  router.push('/')
}
</script>

<template>
  <header class="header">
    <div class="header-top">
      <div class="container">
        Hotline: 0888 841 904 · Miễn phí giao hàng từ 2 triệu · Bảo hành chính hãng
      </div>
    </div>
    <div class="container">
      <div class="header-main">
        <router-link to="/" class="brand">
          <span class="brand-icon">⌚</span>
          WATCH SHOP
        </router-link>

        <div class="search-box">
          <input
            v-model="keyword"
            placeholder="Tìm kiếm đồng hồ..."
            @keyup.enter="search"
          />
          <button @click="search">🔍</button>
        </div>

        <div class="header-actions">
          <router-link v-if="auth.isLogin" to="/wishlist">
            ❤️ <span class="hide-sm">Yêu thích</span>
          </router-link>
          <router-link v-if="auth.isLogin" to="/cart" class="icon-badge">
            🛒 <span class="hide-sm">Giỏ hàng</span>
            <span v-if="cart.count" class="cart-count">{{ cart.count }}</span>
          </router-link>

          <div v-if="auth.isLogin" class="user-menu" @mouseenter="showMenu = true" @mouseleave="showMenu = false">
            <a href="#">👤 {{ auth.user?.ho_ten?.split(' ').slice(-1)[0] }} ▾</a>
            <div v-if="showMenu" class="dropdown">
              <router-link to="/profile">Tài khoản</router-link>
              <router-link to="/orders">Đơn hàng</router-link>
              <router-link to="/wishlist">Yêu thích</router-link>
              <a href="#" @click.prevent="logout">Đăng xuất</a>
            </div>
          </div>

          <template v-else>
            <router-link to="/login">Đăng nhập</router-link>
            <router-link to="/register">Đăng ký</router-link>
          </template>
        </div>
      </div>

      <nav class="nav">
        <ul class="nav-list">
          <li><router-link to="/" exact-active-class="router-link-active">Trang chủ</router-link></li>
          <li><router-link to="/products">Sản phẩm</router-link></li>
          <li><router-link :to="{ path: '/products', query: { gioi_tinh: 'Nam' }}">Đồng hồ nam</router-link></li>
          <li><router-link :to="{ path: '/products', query: { gioi_tinh: 'Nữ' }}">Đồng hồ nữ</router-link></li>
          <li><router-link :to="{ path: '/products', query: { noi_bat: '1' }}">Nổi bật</router-link></li>
        </ul>
      </nav>
    </div>
  </header>

  <main class="main">
    <router-view />
  </main>

  <footer class="footer">
    <div class="container">
      <div class="footer-grid">
        <div>
          <h4>⌚ Watch Shop</h4>
          <p style="font-size:13px; line-height:1.6">
            Cửa hàng đồng hồ chính hãng với hàng trăm mẫu mã đa dạng từ các thương hiệu nổi tiếng thế giới.
          </p>
        </div>
        <div>
          <h4>Hỗ trợ khách hàng</h4>
          <ul>
            <li><a href="#">Hướng dẫn mua hàng</a></li>
            <li><a href="#">Chính sách bảo hành</a></li>
            <li><a href="#">Chính sách đổi trả</a></li>
            <li><a href="#">Vận chuyển</a></li>
          </ul>
        </div>
        <div>
          <h4>Tài khoản</h4>
          <ul>
            <li><router-link to="/profile">Thông tin cá nhân</router-link></li>
            <li><router-link to="/orders">Đơn hàng của tôi</router-link></li>
            <li><router-link to="/wishlist">Yêu thích</router-link></li>
          </ul>
        </div>
        <div>
          <h4>Liên hệ</h4>
          <ul>
            <li>Hotline: 0888 841 904</li>
            <li>Email: support@watchshop.vn</li>
            <li>Địa chỉ: 175 Tây Sơn, Đống Đa, Hà Nội</li>
          </ul>
        </div>
      </div>
      <div class="footer-bottom">
        © 2026 Watch Shop. All rights reserved.
      </div>
    </div>
  </footer>
</template>

<style scoped>
.user-menu { position: relative; }
.dropdown {
  position: absolute;
  top: 100%;
  right: 0;
  background: #fff;
  border: 1px solid #e5e7eb;
  border-radius: 6px;
  box-shadow: 0 4px 12px rgba(0,0,0,.08);
  min-width: 160px;
  padding: 8px 0;
  z-index: 50;
}
.dropdown a {
  display: block;
  padding: 8px 16px;
  color: #374151;
  font-size: 14px;
}
.dropdown a:hover { background: #f9fafb; }
@media (max-width: 768px) {
  .hide-sm { display: none; }
}
</style>
