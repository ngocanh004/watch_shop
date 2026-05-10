<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '../stores/auth'
import { useToastStore } from '../stores/toast'

const router = useRouter()
const auth = useAuthStore()
const toast = useToastStore()

const form = ref({ ten_dang_nhap: '', mat_khau: '' })
const loading = ref(false)
const errMsg = ref('')

async function submit() {
  errMsg.value = ''
  if (!form.value.ten_dang_nhap || !form.value.mat_khau) {
    errMsg.value = 'Vui lòng nhập đầy đủ thông tin!'
    return
  }
  loading.value = true
  try {
    await auth.login(form.value.ten_dang_nhap, form.value.mat_khau)
    toast.success('Đăng nhập thành công')
    router.push('/dashboard')
  } catch (e) {
    errMsg.value = e.message || 'Đăng nhập thất bại'
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="login-wrap">
    <form class="login-box" @submit.prevent="submit">
      <div class="login-head">
        <div class="logo">⌚</div>
        <h2>Watch Shop</h2>
        <p>Đăng nhập trang quản trị</p>
      </div>

      <div class="form-group">
        <label>Tên đăng nhập</label>
        <input v-model="form.ten_dang_nhap" class="input" placeholder="admin01" />
      </div>
      <div class="form-group">
        <label>Mật khẩu</label>
        <input v-model="form.mat_khau" type="password" class="input" placeholder="••••••••" />
      </div>

      <div v-if="errMsg" class="text-error">{{ errMsg }}</div>

      <button type="submit" class="btn btn-primary login-btn" :disabled="loading">
        {{ loading ? 'Đang xử lý...' : 'Đăng nhập' }}
      </button>

      <p class="hint">Mặc định: admin01 / Admin@123</p>
    </form>
  </div>
</template>

<style scoped>
.login-wrap {
  min-height: 100vh;
  display: flex;
  align-items: center;
  justify-content: center;
  background: linear-gradient(135deg, #1e293b 0%, #2563eb 100%);
}
.login-box {
  background: #fff;
  padding: 32px;
  border-radius: 12px;
  width: 100%;
  max-width: 380px;
  box-shadow: 0 20px 40px rgba(0,0,0,.15);
}
.login-head { text-align: center; margin-bottom: 24px; }
.logo { font-size: 48px; }
.login-head h2 { margin-top: 8px; color: #1e293b; }
.login-head p { color: #6b7280; margin-top: 4px; }
.login-btn { width: 100%; padding: 12px; margin-top: 8px; }
.hint { text-align: center; color: #9ca3af; font-size: 12px; margin-top: 14px; }
</style>
