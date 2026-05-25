<script setup>
import { onMounted, ref } from 'vue'
import { authApi } from '../api'
import { useAuthStore } from '../stores/auth'
import { useToastStore } from '../stores/toast'

const auth = useAuthStore()
const toast = useToastStore()

const profile = ref({ ho_ten: '', so_dien_thoai: '', email: '', dia_chi: '' })
const pwd = ref({ mat_khau_cu: '', mat_khau_moi: '', xac_nhan_mat_khau: '' })
const pwdErr = ref('')

async function load() {
  await auth.fetchProfile()
  profile.value = {
    ho_ten: auth.user?.ho_ten || '',
    so_dien_thoai: auth.user?.so_dien_thoai || '',
    email: auth.user?.email || '',
    dia_chi: auth.user?.dia_chi || ''
  }
}

async function saveProfile() {
  try {
    await authApi.updateProfile(profile.value)
    toast.success('Cập nhật thành công')
    auth.fetchProfile()
  } catch (e) {
    toast.error(e.message)
  }
}

async function changePassword() {
  pwdErr.value = ''
  if (!pwd.value.mat_khau_cu || !pwd.value.mat_khau_moi || !pwd.value.xac_nhan_mat_khau) {
    pwdErr.value = 'Vui lòng nhập đầy đủ thông tin'
    return
  }
  if (pwd.value.mat_khau_moi !== pwd.value.xac_nhan_mat_khau) {
    pwdErr.value = 'Mật khẩu xác nhận không khớp'
    return
  }
  try {
    await authApi.changePassword(pwd.value)
    toast.success('Đổi mật khẩu thành công')
    pwd.value = { mat_khau_cu: '', mat_khau_moi: '', xac_nhan_mat_khau: '' }
  } catch (e) {
    pwdErr.value = e.message
  }
}

onMounted(load)
</script>

<template>
  <div class="container" style="padding-top:20px">
    <h2 class="section-title">Tài khoản của tôi</h2>

    <div class="grid-2">
      <div class="box">
        <h3 class="mb-3">Thông tin cá nhân</h3>
        <div class="form-group">
          <label>Họ và tên</label>
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
          <textarea v-model="profile.dia_chi" rows="2"></textarea>
        </div>
        <button class="btn btn-primary" @click="saveProfile">Lưu thay đổi</button>
      </div>

      <div class="box">
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
          <label>Xác nhận mật khẩu mới</label>
          <input v-model="pwd.xac_nhan_mat_khau" type="password" class="input" />
        </div>
        <div v-if="pwdErr" class="text-error mb-3">{{ pwdErr }}</div>
        <button class="btn btn-primary" @click="changePassword">Đổi mật khẩu</button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.grid-2 {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 16px;
}
.box {
  background: #fff;
  padding: 20px;
  border-radius: 8px;
}
@media (max-width: 768px) { .grid-2 { grid-template-columns: 1fr; } }
</style>
