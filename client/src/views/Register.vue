<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useAuthStore } from '../stores/auth'
import { useToastStore } from '../stores/toast'

const router = useRouter()
const auth = useAuthStore()
const toast = useToastStore()

const form = ref({
  ho_ten: '',
  so_dien_thoai: '',
  email: '',
  ten_dang_nhap: '',
  mat_khau: '',
  xac_nhan_mat_khau: ''
})
const errMsg = ref('')
const loading = ref(false)

function validate() {
  if (!form.value.ho_ten || !form.value.so_dien_thoai ||
      !form.value.ten_dang_nhap || !form.value.mat_khau || !form.value.xac_nhan_mat_khau) {
    return 'Vui lòng nhập đầy đủ thông tin!'
  }
  if (!/^[a-zA-ZÀ-ỹ\s]+$/.test(form.value.ho_ten) || form.value.ho_ten.length > 50) {
    return 'Họ tên chỉ được chứa chữ và không quá 50 ký tự'
  }
  if (!/^0\d{9}$/.test(form.value.so_dien_thoai)) {
    return 'Số điện thoại không hợp lệ. Định dạng: 0********* (10 số)'
  }
  if (!/^[a-zA-Z0-9]+$/.test(form.value.ten_dang_nhap) ||
      !/[a-zA-Z]/.test(form.value.ten_dang_nhap) ||
      !/[0-9]/.test(form.value.ten_dang_nhap)) {
    return 'Tên đăng nhập phải chứa cả chữ và số. (Không chứa ký tự đặc biệt)'
  }
  if (form.value.mat_khau.length < 8 || form.value.mat_khau.length > 20 ||
      !/[A-Z]/.test(form.value.mat_khau) ||
      !/[a-z]/.test(form.value.mat_khau) ||
      !/[0-9]/.test(form.value.mat_khau) ||
      !/[^a-zA-Z0-9]/.test(form.value.mat_khau)) {
    return 'Mật khẩu phải gồm chữ hoa, chữ thường, số, ký tự đặc biệt (8–20 ký tự)'
  }
  if (form.value.mat_khau !== form.value.xac_nhan_mat_khau) {
    return 'Mật khẩu xác nhận không khớp'
  }
  return ''
}

async function submit() {
  errMsg.value = validate()
  if (errMsg.value) return

  loading.value = true
  try {
    await auth.register(form.value)
    toast.success('Đăng ký thành công! Vui lòng đăng nhập.')
    router.push('/login')
  } catch (e) {
    errMsg.value = e.message
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
        <h2>Đăng ký tài khoản</h2>
        <p>Tạo tài khoản để bắt đầu mua sắm</p>
      </div>

      <div class="form-group">
        <label>Họ và tên</label>
        <input v-model="form.ho_ten" class="input" placeholder="Nguyễn Văn A" />
      </div>
      <div class="form-group">
        <label>Số điện thoại</label>
        <input v-model="form.so_dien_thoai" class="input" placeholder="0123456789" />
      </div>
      <div class="form-group">
        <label>Email (tuỳ chọn)</label>
        <input v-model="form.email" type="email" class="input" />
      </div>
      <div class="form-group">
        <label>Tên đăng nhập</label>
        <input v-model="form.ten_dang_nhap" class="input" placeholder="user123" />
      </div>
      <div class="form-group">
        <label>Mật khẩu</label>
        <input v-model="form.mat_khau" type="password" class="input" />
      </div>
      <div class="form-group">
        <label>Xác nhận mật khẩu</label>
        <input v-model="form.xac_nhan_mat_khau" type="password" class="input" />
      </div>

      <div v-if="errMsg" class="text-error">{{ errMsg }}</div>

      <button type="submit" class="btn btn-primary btn-block" :disabled="loading" style="padding:12px; margin-top:8px">
        {{ loading ? 'Đang xử lý...' : 'Đăng ký' }}
      </button>

      <p class="auth-bottom">
        Đã có tài khoản?
        <router-link to="/login">Đăng nhập</router-link>
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
  max-width: 440px;
  box-shadow: 0 4px 20px rgba(0,0,0,.08);
}
.auth-head { text-align: center; margin-bottom: 20px; }
.logo { font-size: 22px; font-weight: 700; color: #111827; display: inline-block; margin-bottom: 12px; }
.auth-head h2 { color: #111827; }
.auth-head p { color: #6b7280; font-size: 13px; margin-top: 4px; }
.auth-bottom { text-align: center; margin-top: 16px; color: #6b7280; font-size: 14px; }
.auth-bottom a { color: #111827; font-weight: 600; }
</style>
