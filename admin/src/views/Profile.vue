<script setup>
import { onMounted, ref } from 'vue'
import { authApi } from '../api'
import { useAuthStore } from '../stores/auth'
import { useToastStore } from '../stores/toast'

const auth = useAuthStore()
const toast = useToastStore()

const profile = ref({ ho_ten: '', so_dien_thoai: '', email: '', dia_chi: '' })
const pwd = ref({ mat_khau_cu: '', mat_khau_moi: '', xac_nhan_mat_khau: '' })

async function load() {
  await auth.fetchProfile()
  profile.value = { ...auth.user }
}

async function saveProfile() {
  try {
    await import('../api').then(m => m.userApi)
  } catch {}
  try {
    const http = (await import('../api/http')).default
    await http.put('/me', profile.value)
    toast.success('Cập nhật thành công')
    auth.fetchProfile()
  } catch (e) {
    toast.error(e.message)
  }
}

async function changePassword() {
  if (!pwd.value.mat_khau_cu || !pwd.value.mat_khau_moi || !pwd.value.xac_nhan_mat_khau) {
    toast.error('Vui lòng nhập đầy đủ thông tin')
    return
  }
  if (pwd.value.mat_khau_moi !== pwd.value.xac_nhan_mat_khau) {
    toast.error('Mật khẩu xác nhận không khớp')
    return
  }
  try {
    await authApi.changePassword(pwd.value)
    toast.success('Đổi mật khẩu thành công')
    pwd.value = { mat_khau_cu: '', mat_khau_moi: '', xac_nhan_mat_khau: '' }
  } catch (e) {
    toast.error(e.message)
  }
}

onMounted(load)
</script>

<template>
  <h1 class="mb-3">Thông tin cá nhân</h1>

  <div class="grid">
    <div class="card">
      <h3 class="mb-3">Cập nhật thông tin</h3>
      <div class="form-group">
        <label>Họ tên</label>
        <input v-model="profile.ho_ten" class="input" />
      </div>
      <div class="form-group">
        <label>Số điện thoại</label>
        <input v-model="profile.so_dien_thoai" class="input" />
      </div>
      <div class="form-group">
        <label>Email</label>
        <input v-model="profile.email" class="input" />
      </div>
      <div class="form-group">
        <label>Địa chỉ</label>
        <input v-model="profile.dia_chi" class="input" />
      </div>
      <button class="btn btn-primary" @click="saveProfile">Lưu thay đổi</button>
    </div>

    <div class="card">
      <h3 class="mb-3">Đổi mật khẩu</h3>
      <div class="form-group">
        <label>Mật khẩu cũ</label>
        <input v-model="pwd.mat_khau_cu" type="password" class="input" />
      </div>
      <div class="form-group">
        <label>Mật khẩu mới</label>
        <input v-model="pwd.mat_khau_moi" type="password" class="input" />
      </div>
      <div class="form-group">
        <label>Xác nhận mật khẩu</label>
        <input v-model="pwd.xac_nhan_mat_khau" type="password" class="input" />
      </div>
      <button class="btn btn-warning" @click="changePassword">Đổi mật khẩu</button>
    </div>
  </div>
</template>

<style scoped>
.grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 16px;
}
</style>
