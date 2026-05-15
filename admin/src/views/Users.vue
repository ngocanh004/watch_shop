<script setup>
import { onMounted, reactive, ref } from 'vue'
import { userApi } from '../api'
import { useToastStore } from '../stores/toast'
import { formatDate } from '../utils/format'

const toast = useToastStore()
const items = ref([])
const total = ref(0)

const query = reactive({
  keyword: '',
  vai_tro: '',
  page: 1,
  page_size: 10
})

const showEdit = ref(false)
const editing = ref(null)
const form = ref({})

const showReset = ref(false)
const resetTarget = ref(null)
const newPwd = ref('')

async function load() {
  const res = await userApi.list(query)
  items.value = res.data
  total.value = res.total
}

function openEdit(item) {
  editing.value = item
  form.value = { ...item }
  showEdit.value = true
}

async function save() {
  try {
    await userApi.update(editing.value.id, form.value)
    toast.success('Cập nhật thành công')
    showEdit.value = false
    load()
  } catch (e) {
    toast.error(e.message)
  }
}

async function remove(u) {
  if (u.vai_tro === 'admin') {
    toast.error('Không thể xoá tài khoản admin')
    return
  }
  if (!confirm(`Xoá tài khoản "${u.ten_dang_nhap}"?`)) return
  try {
    await userApi.remove(u.id)
    toast.success('Đã xoá')
    load()
  } catch (e) {
    toast.error(e.message)
  }
}

function openReset(u) {
  resetTarget.value = u
  newPwd.value = ''
  showReset.value = true
}

async function doReset() {
  try {
    await userApi.resetPassword(resetTarget.value.id, { mat_khau_moi: newPwd.value })
    toast.success('Đặt lại mật khẩu thành công')
    showReset.value = false
  } catch (e) {
    toast.error(e.message)
  }
}

function changePage(p) {
  query.page = p
  load()
}

onMounted(load)
</script>

<template>
  <h1 class="mb-3">Quản lý người dùng</h1>

  <div class="card">
    <div class="toolbar">
      <input v-model="query.keyword" class="input" placeholder="Tên, username, email..." @keyup.enter="load" />
      <select v-model="query.vai_tro" class="select" @change="load">
        <option value="">-- Tất cả vai trò --</option>
        <option value="admin">Quản trị viên</option>
        <option value="user">Người dùng</option>
      </select>
      <button class="btn btn-outline" @click="load">Lọc</button>
    </div>

    <table class="table">
      <thead>
        <tr>
          <th>ID</th>
          <th>Họ tên</th>
          <th>Tên đăng nhập</th>
          <th>SĐT</th>
          <th>Email</th>
          <th>Vai trò</th>
          <th>Trạng thái</th>
          <th>Tham gia</th>
          <th>Thao tác</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="u in items" :key="u.id">
          <td>{{ u.id }}</td>
          <td>{{ u.ho_ten }}</td>
          <td>{{ u.ten_dang_nhap }}</td>
          <td>{{ u.so_dien_thoai }}</td>
          <td>{{ u.email }}</td>
          <td>
            <span :class="['badge', u.vai_tro === 'admin' ? 'badge-info' : 'badge-default']">
              {{ u.vai_tro === 'admin' ? 'Admin' : 'Người dùng' }}
            </span>
          </td>
          <td>
            <span :class="['badge', u.trang_thai ? 'badge-success' : 'badge-danger']">
              {{ u.trang_thai ? 'Hoạt động' : 'Khoá' }}
            </span>
          </td>
          <td>{{ formatDate(u.created_at) }}</td>
          <td>
            <button class="btn btn-sm btn-outline" @click="openEdit(u)">Sửa</button>
            <button class="btn btn-sm btn-warning" @click="openReset(u)" style="margin-left:4px">Reset MK</button>
            <button class="btn btn-sm btn-danger" @click="remove(u)" style="margin-left:4px">Xoá</button>
          </td>
        </tr>
      </tbody>
    </table>

    <div class="pagination">
      <button :disabled="query.page === 1" @click="changePage(query.page - 1)">‹</button>
      <span class="text-muted" style="padding: 6px 12px;">Trang {{ query.page }} / {{ Math.ceil(total / query.page_size) || 1 }}</span>
      <button :disabled="query.page * query.page_size >= total" @click="changePage(query.page + 1)">›</button>
    </div>
  </div>

  <div v-if="showEdit" class="modal-overlay" @click.self="showEdit = false">
    <div class="modal">
      <div class="modal-header">
        <h3>Sửa người dùng</h3>
        <button class="modal-close" @click="showEdit = false">×</button>
      </div>
      <div class="modal-body">
        <div class="form-group">
          <label>Họ tên</label>
          <input v-model="form.ho_ten" class="input" />
        </div>
        <div class="form-row">
          <div class="form-group">
            <label>SĐT</label>
            <input v-model="form.so_dien_thoai" class="input" />
          </div>
          <div class="form-group">
            <label>Email</label>
            <input v-model="form.email" class="input" />
          </div>
        </div>
        <div class="form-group">
          <label>Địa chỉ</label>
          <input v-model="form.dia_chi" class="input" />
        </div>
        <div class="form-row">
          <div class="form-group">
            <label>Vai trò</label>
            <select v-model="form.vai_tro" class="select">
              <option value="user">Người dùng</option>
              <option value="admin">Quản trị viên</option>
            </select>
          </div>
          <div class="form-group">
            <label>Trạng thái</label>
            <select v-model="form.trang_thai" class="select">
              <option :value="true">Hoạt động</option>
              <option :value="false">Khoá</option>
            </select>
          </div>
        </div>
      </div>
      <div class="modal-footer">
        <button class="btn btn-outline" @click="showEdit = false">Huỷ</button>
        <button class="btn btn-primary" @click="save">Lưu</button>
      </div>
    </div>
  </div>

  <div v-if="showReset" class="modal-overlay" @click.self="showReset = false">
    <div class="modal">
      <div class="modal-header">
        <h3>Đặt lại mật khẩu</h3>
        <button class="modal-close" @click="showReset = false">×</button>
      </div>
      <div class="modal-body">
        <p class="mb-3">Người dùng: <strong>{{ resetTarget?.ten_dang_nhap }}</strong></p>
        <div class="form-group">
          <label>Mật khẩu mới</label>
          <input v-model="newPwd" type="text" class="input" placeholder="8-20 ký tự, có hoa, thường, số, ký tự đặc biệt" />
        </div>
      </div>
      <div class="modal-footer">
        <button class="btn btn-outline" @click="showReset = false">Huỷ</button>
        <button class="btn btn-warning" @click="doReset">Đặt lại</button>
      </div>
    </div>
  </div>
</template>
