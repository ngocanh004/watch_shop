<script setup>
import { ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { useAuthStore } from '../stores/auth'
import { useCartStore } from '../stores/cart'
import { useToastStore } from '../stores/toast'

const route = useRoute()
const router = useRouter()
const auth = useAuthStore()
const cart = useCartStore()
const toast = useToastStore()

const form = ref({ ten_dang_nhap: '', mat_khau: '' })
const errMsg = ref('')
const loading = ref(false)

async function submit() {
  errMsg.value = ''
  if (!form.value.ten_dang_nhap || !form.value.mat_khau) {
    errMsg.value = 'Vui lòng nhập đầy đủ thông tin!'
    return
  }
  loading.value = true
  try {
    await auth.login(form.value.ten_dang_nhap, form.value.mat_khau)
    await cart.load().catch(() => {})
    toast.success('Đăng nhập thành công')
    router.push(route.query.redirect || '/')
  } catch (e) {
    errMsg.value = e.message || 'Tên đăng nhập hoặc mật khẩu không đúng!'
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="auth-wrap">
    <form class="auth-box" @submit.prevent="submit">
      <div class="auth-head">
        <router-link to="/" class="logo">⌚ WATCH SHOP</router-link>
        <h2>Đăng nhập</h2>
        <p>Đăng nhập để mua sắm và theo dõi đơn hàng</p>
      </div>

      <div class="form-group">
        <label>Tên đăng nhập</label>
        <input v-model="form.ten_dang_nhap" class="input" placeholder="Nhập tên đăng nhập" />
      </div>
      <div class="form-group">
        <label>Mật khẩu</label>
        <input v-model="form.mat_khau" type="password" class="input" placeholder="Nhập mật khẩu" />
      </div>

      <div v-if="errMsg" class="text-error">{{ errMsg }}</div>

      <button type="submit" class="btn btn-primary btn-block" :disabled="loading" style="padding:12px; margin-top:8px">
        {{ loading ? 'Đang xử lý...' : 'Đăng nhập' }}
      </button>

      <p class="auth-bottom">
        Bạn chưa có tài khoản?
        <router-link to="/register">Đăng ký</router-link>
      </p>
    </form>
  </div>
</template>

<style scoped>
.auth-wrap {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: #f3f4f6;
  padding: 20px;
}
.auth-box {
  background: #fff;
  padding: 32px;
  border-radius: 12px;
  width: 100%;
  max-width: 420px;
  box-shadow: 0 4px 20px rgba(0,0,0,.08);
}
.auth-head { text-align: center; margin-bottom: 24px; }
.logo { font-size: 22px; font-weight: 700; color: #111827; display: inline-block; margin-bottom: 14px; }
.auth-head h2 { color: #111827; }
.auth-head p { color: #6b7280; font-size: 13px; margin-top: 4px; }
.auth-bottom { text-align: center; margin-top: 18px; color: #6b7280; font-size: 14px; }
.auth-bottom a { color: #111827; font-weight: 600; }
</style>
