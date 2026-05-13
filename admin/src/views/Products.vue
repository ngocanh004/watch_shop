<script setup>
import { onMounted, reactive, ref } from 'vue'
import { productApi, categoryApi, brandApi, uploadApi, PUBLIC_URL } from '../api'
import { useToastStore } from '../stores/toast'
import { formatMoney } from '../utils/format'

const toast = useToastStore()
const items = ref([])
const total = ref(0)
const categories = ref([])
const brands = ref([])

const query = reactive({
  keyword: '',
  category_id: '',
  brand_id: '',
  page: 1,
  page_size: 10
})

const showModal = ref(false)
const editing = ref(null)
const saving = ref(false)
const emptyForm = () => ({
  ten_san_pham: '', mo_ta: '', gia: 0, gia_goc: 0, so_luong: 0,
  category_id: '', brand_id: '', hinh_anh: '',
  loai_may: 'Quartz', day_deo: 'Da', mat_kinh: 'Mineral',
  duong_kinh: '40mm', khang_nuoc: '3ATM', gioi_tinh: 'Nam',
  xuat_xu: '', bao_hanh: '12 thang',
  noi_bat: false, trang_thai: true
})
const form = ref(emptyForm())

async function load() {
  const res = await productApi.list(query)
  items.value = res.data
  total.value = res.total
}

async function loadOptions() {
  const [c, b] = await Promise.all([categoryApi.list(), brandApi.list()])
  categories.value = c.data
  brands.value = b.data
}

function openCreate() {
  editing.value = null
  form.value = emptyForm()
  showModal.value = true
}

async function openEdit(item) {
  const res = await productApi.get(item.id)
  editing.value = res.data
  form.value = {
    ...emptyForm(),
    ...res.data,
    category_id: res.data.category_id,
    brand_id: res.data.brand_id
  }
  showModal.value = true
}

async function onUpload(e) {
  const file = e.target.files[0]
  if (!file) return
  try {
    const res = await uploadApi.image(file)
    form.value.hinh_anh = PUBLIC_URL + res.data.url
    toast.success('Tải ảnh thành công')
  } catch (err) {
    toast.error(err.message)
  }
}

async function save() {
  if (!form.value.ten_san_pham || form.value.gia <= 0) {
    toast.error('Vui lòng nhập tên và giá sản phẩm')
    return
  }
  if (!form.value.category_id || !form.value.brand_id) {
    toast.error('Vui lòng chọn danh mục và thương hiệu')
    return
  }
  saving.value = true
  try {
    const payload = {
      ...form.value,
      gia: Number(form.value.gia),
      gia_goc: Number(form.value.gia_goc),
      so_luong: Number(form.value.so_luong),
      category_id: Number(form.value.category_id),
      brand_id: Number(form.value.brand_id)
    }
    if (editing.value) {
      await productApi.update(editing.value.id, payload)
      toast.success('Cập nhật thành công')
    } else {
      await productApi.create(payload)
      toast.success('Thêm sản phẩm thành công')
    }
    showModal.value = false
    load()
  } catch (e) {
    toast.error(e.message)
  } finally {
    saving.value = false
  }
}

async function remove(item) {
  if (!confirm(`Xoá "${item.ten_san_pham}"?`)) return
  try {
    await productApi.remove(item.id)
    toast.success('Đã xoá sản phẩm')
    load()
  } catch (e) {
    toast.error(e.message)
  }
}

function changePage(p) {
  query.page = p
  load()
}

onMounted(() => {
  loadOptions()
  load()
})
</script>

<template>
  <div class="flex justify-between items-center mb-3">
    <h1>Quản lý sản phẩm</h1>
    <button class="btn btn-primary" @click="openCreate">+ Thêm sản phẩm</button>
  </div>

  <div class="card">
    <div class="toolbar">
      <input v-model="query.keyword" class="input" placeholder="Tìm theo tên..." @keyup.enter="load" />
      <select v-model="query.category_id" class="select" @change="load">
        <option value="">-- Tất cả danh mục --</option>
        <option v-for="c in categories" :key="c.id" :value="c.id">{{ c.ten_danh_muc }}</option>
      </select>
      <select v-model="query.brand_id" class="select" @change="load">
        <option value="">-- Tất cả thương hiệu --</option>
        <option v-for="b in brands" :key="b.id" :value="b.id">{{ b.ten_thuong_hieu }}</option>
      </select>
      <button class="btn btn-outline" @click="load">Lọc</button>
    </div>

    <table class="table">
      <thead>
        <tr>
          <th>Ảnh</th>
          <th>Tên sản phẩm</th>
          <th>Thương hiệu</th>
          <th>Giá</th>
          <th>Kho</th>
          <th>Đã bán</th>
          <th>Trạng thái</th>
          <th>Thao tác</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="it in items" :key="it.id">
          <td>
            <img v-if="it.hinh_anh" :src="it.hinh_anh" class="thumb" />
            <div v-else class="thumb-empty">N/A</div>
          </td>
          <td>
            <div><strong>{{ it.ten_san_pham }}</strong></div>
            <div class="text-muted" style="font-size:12px">{{ it.loai_may }} · {{ it.day_deo }}</div>
          </td>
          <td>{{ it.brand?.ten_thuong_hieu }}</td>
          <td>{{ formatMoney(it.gia) }}</td>
          <td>{{ it.so_luong }}</td>
          <td>{{ it.da_ban }}</td>
          <td>
            <span :class="['badge', it.trang_thai ? 'badge-success' : 'badge-default']">
              {{ it.trang_thai ? 'Đang bán' : 'Ẩn' }}
            </span>
          </td>
          <td>
            <button class="btn btn-sm btn-outline" @click="openEdit(it)">Sửa</button>
            <button class="btn btn-sm btn-danger" @click="remove(it)" style="margin-left:6px">Xoá</button>
          </td>
        </tr>
        <tr v-if="!items.length">
          <td colspan="8" class="text-center text-muted">Không có sản phẩm</td>
        </tr>
      </tbody>
    </table>

    <div class="pagination">
      <button :disabled="query.page === 1" @click="changePage(query.page - 1)">‹</button>
      <span class="text-muted" style="padding: 6px 12px;">Trang {{ query.page }} / {{ Math.ceil(total / query.page_size) || 1 }}</span>
      <button :disabled="query.page * query.page_size >= total" @click="changePage(query.page + 1)">›</button>
    </div>
  </div>

  <div v-if="showModal" class="modal-overlay" @click.self="showModal = false">
    <div class="modal modal-large">
      <div class="modal-header">
        <h3>{{ editing ? 'Sửa sản phẩm' : 'Thêm sản phẩm' }}</h3>
        <button class="modal-close" @click="showModal = false">×</button>
      </div>
      <div class="modal-body">
        <div class="form-row">
          <div class="form-group">
            <label>Tên sản phẩm *</label>
            <input v-model="form.ten_san_pham" class="input" />
          </div>
          <div class="form-group">
            <label>Hình ảnh</label>
            <input type="file" accept="image/*" @change="onUpload" />
            <input v-model="form.hinh_anh" class="input mt-2" placeholder="Hoặc dán URL" />
            <img v-if="form.hinh_anh" :src="form.hinh_anh" class="preview mt-2" />
          </div>
        </div>

        <div class="form-row">
          <div class="form-group">
            <label>Danh mục *</label>
            <select v-model="form.category_id" class="select">
              <option value="">-- Chọn --</option>
              <option v-for="c in categories" :key="c.id" :value="c.id">{{ c.ten_danh_muc }}</option>
            </select>
          </div>
          <div class="form-group">
            <label>Thương hiệu *</label>
            <select v-model="form.brand_id" class="select">
              <option value="">-- Chọn --</option>
              <option v-for="b in brands" :key="b.id" :value="b.id">{{ b.ten_thuong_hieu }}</option>
            </select>
          </div>
        </div>

        <div class="form-row">
          <div class="form-group">
            <label>Giá bán *</label>
            <input v-model.number="form.gia" type="number" class="input" />
          </div>
          <div class="form-group">
            <label>Giá gốc</label>
            <input v-model.number="form.gia_goc" type="number" class="input" />
          </div>
        </div>

        <div class="form-row">
          <div class="form-group">
            <label>Số lượng</label>
            <input v-model.number="form.so_luong" type="number" class="input" />
          </div>
          <div class="form-group">
            <label>Giới tính</label>
            <select v-model="form.gioi_tinh" class="select">
              <option>Nam</option>
              <option>Nữ</option>
              <option>Unisex</option>
            </select>
          </div>
        </div>

        <div class="form-row">
          <div class="form-group">
            <label>Loại máy</label>
            <select v-model="form.loai_may" class="select">
              <option>Quartz</option>
              <option>Automatic</option>
              <option>Mechanical</option>
              <option>Eco-Drive</option>
              <option>Solar</option>
            </select>
          </div>
          <div class="form-group">
            <label>Loại dây</label>
            <select v-model="form.day_deo" class="select">
              <option>Da</option>
              <option>Kim loại</option>
              <option>Cao su</option>
              <option>Vải</option>
            </select>
          </div>
        </div>

        <div class="form-row">
          <div class="form-group">
            <label>Mặt kính</label>
            <input v-model="form.mat_kinh" class="input" />
          </div>
          <div class="form-group">
            <label>Đường kính</label>
            <input v-model="form.duong_kinh" class="input" />
          </div>
        </div>

        <div class="form-row">
          <div class="form-group">
            <label>Kháng nước</label>
            <input v-model="form.khang_nuoc" class="input" />
          </div>
          <div class="form-group">
            <label>Xuất xứ</label>
            <input v-model="form.xuat_xu" class="input" />
          </div>
        </div>

        <div class="form-row">
          <div class="form-group">
            <label>Bảo hành</label>
            <input v-model="form.bao_hanh" class="input" />
          </div>
          <div class="form-group">
            <label>Tuỳ chọn</label>
            <label style="display:block">
              <input type="checkbox" v-model="form.noi_bat" /> Sản phẩm nổi bật
            </label>
            <label style="display:block">
              <input type="checkbox" v-model="form.trang_thai" /> Đang bán
            </label>
          </div>
        </div>

        <div class="form-group">
          <label>Mô tả</label>
          <textarea v-model="form.mo_ta" rows="4"></textarea>
        </div>
      </div>
      <div class="modal-footer">
        <button class="btn btn-outline" @click="showModal = false">Huỷ</button>
        <button class="btn btn-primary" @click="save" :disabled="saving">
          {{ saving ? 'Đang lưu...' : 'Lưu' }}
        </button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.thumb { width: 50px; height: 50px; object-fit: cover; border-radius: 4px; }
.thumb-empty {
  width: 50px; height: 50px; background: #e5e7eb; border-radius: 4px;
  display: flex; align-items: center; justify-content: center;
  color: #9ca3af; font-size: 11px;
}
.preview { width: 120px; height: 120px; object-fit: cover; border-radius: 6px; border: 1px solid #e5e7eb; }
</style>
