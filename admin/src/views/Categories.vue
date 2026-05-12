<script setup>
import { onMounted, ref } from 'vue'
import { categoryApi } from '../api'
import { useToastStore } from '../stores/toast'

const toast = useToastStore()

const items = ref([])
const keyword = ref('')
const showModal = ref(false)
const editing = ref(null)
const form = ref({ ten_danh_muc: '', mo_ta: '', trang_thai: true })

async function load() {
  const res = await categoryApi.list({ keyword: keyword.value })
  items.value = res.data
}

function openCreate() {
  editing.value = null
  form.value = { ten_danh_muc: '', mo_ta: '', trang_thai: true }
  showModal.value = true
}

function openEdit(item) {
  editing.value = item
  form.value = { ...item }
  showModal.value = true
}

async function save() {
  if (!form.value.ten_danh_muc) {
    toast.error('Vui lòng nhập tên danh mục')
    return
  }
  try {
    if (editing.value) {
      await categoryApi.update(editing.value.id, form.value)
      toast.success('Cập nhật thành công')
    } else {
      await categoryApi.create(form.value)
      toast.success('Thêm danh mục thành công')
    }
    showModal.value = false
    load()
  } catch (e) {
    toast.error(e.message)
  }
}

async function remove(item) {
  if (!confirm(`Bạn có chắc muốn xoá "${item.ten_danh_muc}"?`)) return
  try {
    await categoryApi.remove(item.id)
    toast.success('Đã xoá danh mục')
    load()
  } catch (e) {
    toast.error(e.message)
  }
}

onMounted(load)
</script>

<template>
  <div class="flex justify-between items-center mb-3">
    <h1>Quản lý danh mục</h1>
    <button class="btn btn-primary" @click="openCreate">+ Thêm danh mục</button>
  </div>

  <div class="card">
    <div class="toolbar">
      <input v-model="keyword" class="input" placeholder="Tìm theo tên..." @keyup.enter="load" />
      <button class="btn btn-outline" @click="load">Tìm kiếm</button>
    </div>

    <table class="table">
      <thead>
        <tr>
          <th>ID</th>
          <th>Tên danh mục</th>
          <th>Slug</th>
          <th>Mô tả</th>
          <th>Trạng thái</th>
          <th>Thao tác</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="it in items" :key="it.id">
          <td>{{ it.id }}</td>
          <td>{{ it.ten_danh_muc }}</td>
          <td>{{ it.slug }}</td>
          <td>{{ it.mo_ta }}</td>
          <td>
            <span :class="['badge', it.trang_thai ? 'badge-success' : 'badge-default']">
              {{ it.trang_thai ? 'Hiển thị' : 'Ẩn' }}
            </span>
          </td>
          <td>
            <button class="btn btn-sm btn-outline" @click="openEdit(it)">Sửa</button>
            <button class="btn btn-sm btn-danger" @click="remove(it)" style="margin-left:6px">Xoá</button>
          </td>
        </tr>
        <tr v-if="!items.length">
          <td colspan="6" class="text-center text-muted">Không có dữ liệu</td>
        </tr>
      </tbody>
    </table>
  </div>

  <div v-if="showModal" class="modal-overlay" @click.self="showModal = false">
    <div class="modal">
      <div class="modal-header">
        <h3>{{ editing ? 'Sửa danh mục' : 'Thêm danh mục' }}</h3>
        <button class="modal-close" @click="showModal = false">×</button>
      </div>
      <div class="modal-body">
        <div class="form-group">
          <label>Tên danh mục</label>
          <input v-model="form.ten_danh_muc" class="input" />
        </div>
        <div class="form-group">
          <label>Mô tả</label>
          <textarea v-model="form.mo_ta" rows="3"></textarea>
        </div>
        <div class="form-group">
          <label>
            <input type="checkbox" v-model="form.trang_thai" />
            Hiển thị
          </label>
        </div>
      </div>
      <div class="modal-footer">
        <button class="btn btn-outline" @click="showModal = false">Huỷ</button>
        <button class="btn btn-primary" @click="save">Lưu</button>
      </div>
    </div>
  </div>
</template>
