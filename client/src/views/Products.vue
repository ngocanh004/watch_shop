<script setup>
import { onMounted, reactive, ref, watch } from 'vue'
import { useRoute } from 'vue-router'
import { productApi, brandApi, categoryApi } from '../api'
import ProductCard from '../components/ProductCard.vue'

const route = useRoute()
const items = ref([])
const total = ref(0)
const totalPages = ref(0)
const brands = ref([])
const categories = ref([])
const loading = ref(false)

const filter = reactive({
  keyword: '',
  category_id: '',
  brand_id: '',
  loai_may: '',
  day_deo: '',
  gioi_tinh: '',
  gia_min: '',
  gia_max: '',
  noi_bat: '',
  sort: 'moi_nhat',
  page: 1,
  page_size: 12
})

function applyQuery() {
  const q = route.query
  filter.keyword = q.keyword || ''
  filter.category_id = q.category_id || ''
  filter.brand_id = q.brand_id || ''
  filter.gioi_tinh = q.gioi_tinh || ''
  filter.noi_bat = q.noi_bat || ''
  filter.page = 1
}

async function load() {
  loading.value = true
  try {
    const params = {}
    for (const k in filter) if (filter[k] !== '') params[k] = filter[k]
    const res = await productApi.list(params)
    items.value = res.data
    total.value = res.total
    totalPages.value = res.total_pages
  } finally {
    loading.value = false
  }
}

async function loadOptions() {
  const [b, c] = await Promise.all([brandApi.list(), categoryApi.list()])
  brands.value = b.data
  categories.value = c.data
}

function resetFilter() {
  filter.keyword = ''
  filter.category_id = ''
  filter.brand_id = ''
  filter.loai_may = ''
  filter.day_deo = ''
  filter.gioi_tinh = ''
  filter.gia_min = ''
  filter.gia_max = ''
  filter.noi_bat = ''
  filter.sort = 'moi_nhat'
  filter.page = 1
  load()
}

function changePage(p) {
  filter.page = p
  load()
  window.scrollTo({ top: 200, behavior: 'smooth' })
}

watch(() => route.query, () => {
  applyQuery()
  load()
})

onMounted(() => {
  applyQuery()
  loadOptions()
  load()
})
</script>

<template>
  <div class="container" style="padding-top: 20px">
    <h2 class="section-title">Sản phẩm</h2>

    <div class="filter-wrap">
      <aside class="filter-sidebar">
        <h3 class="mb-3">Bộ lọc</h3>

        <div class="filter-section">
          <h4>Tìm kiếm</h4>
          <input v-model="filter.keyword" class="input" placeholder="Nhập từ khoá..." @keyup.enter="load" />
        </div>

        <div class="filter-section">
          <h4>Danh mục</h4>
          <label v-for="c in categories" :key="c.id">
            <input type="radio" :value="String(c.id)" v-model="filter.category_id" @change="load" />
            {{ c.ten_danh_muc }}
          </label>
          <label>
            <input type="radio" value="" v-model="filter.category_id" @change="load" /> Tất cả
          </label>
        </div>

        <div class="filter-section">
          <h4>Thương hiệu</h4>
          <label v-for="b in brands" :key="b.id">
            <input type="radio" :value="String(b.id)" v-model="filter.brand_id" @change="load" />
            {{ b.ten_thuong_hieu }}
          </label>
          <label>
            <input type="radio" value="" v-model="filter.brand_id" @change="load" /> Tất cả
          </label>
        </div>

        <div class="filter-section">
          <h4>Giới tính</h4>
          <select v-model="filter.gioi_tinh" class="select" @change="load">
            <option value="">Tất cả</option>
            <option>Nam</option>
            <option>Nữ</option>
            <option>Unisex</option>
          </select>
        </div>

        <div class="filter-section">
          <h4>Loại máy</h4>
          <select v-model="filter.loai_may" class="select" @change="load">
            <option value="">Tất cả</option>
            <option>Quartz</option>
            <option>Automatic</option>
            <option>Mechanical</option>
            <option>Eco-Drive</option>
          </select>
        </div>

        <div class="filter-section">
          <h4>Loại dây</h4>
          <select v-model="filter.day_deo" class="select" @change="load">
            <option value="">Tất cả</option>
            <option>Da</option>
            <option>Kim loại</option>
            <option>Cao su</option>
            <option>Vải</option>
          </select>
        </div>

        <div class="filter-section">
          <h4>Khoảng giá</h4>
          <input v-model="filter.gia_min" type="number" class="input mb-3" placeholder="Từ..." />
          <input v-model="filter.gia_max" type="number" class="input" placeholder="Đến..." />
          <button class="btn btn-primary btn-block mt-3" @click="load">Áp dụng</button>
        </div>

        <button class="btn btn-outline btn-block" @click="resetFilter">Xoá bộ lọc</button>
      </aside>

      <div>
        <div class="list-toolbar">
          <div class="text-muted">Tìm thấy <strong>{{ total }}</strong> sản phẩm</div>
          <select v-model="filter.sort" class="select" style="max-width:180px" @change="load">
            <option value="moi_nhat">Mới nhất</option>
            <option value="gia_tang">Giá tăng dần</option>
            <option value="gia_giam">Giá giảm dần</option>
            <option value="ban_chay">Bán chạy</option>
          </select>
        </div>

        <div v-if="loading" class="text-center text-muted" style="padding:40px">Đang tải...</div>

        <div v-else-if="!items.length" class="empty-state">
          <div class="icon">⌚</div>
          <h3>Không tìm thấy sản phẩm</h3>
          <p>Vui lòng thử bộ lọc khác</p>
        </div>

        <div v-else class="product-grid">
          <ProductCard v-for="p in items" :key="p.id" :product="p" />
        </div>

        <div v-if="totalPages > 1" class="pagination-bar">
          <button :disabled="filter.page === 1" @click="changePage(filter.page - 1)">‹</button>
          <button
            v-for="p in totalPages"
            :key="p"
            :class="{ active: p === filter.page }"
            @click="changePage(p)"
          >{{ p }}</button>
          <button :disabled="filter.page === totalPages" @click="changePage(filter.page + 1)">›</button>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.list-toolbar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 16px;
  background: #fff;
  padding: 10px 14px;
  border-radius: 8px;
}
.pagination-bar {
  display: flex;
  gap: 6px;
  justify-content: center;
  margin-top: 24px;
}
.pagination-bar button {
  padding: 8px 14px;
  background: #fff;
  border: 1px solid #d1d5db;
  border-radius: 4px;
  cursor: pointer;
}
.pagination-bar button.active { background: #111827; color: #fff; border-color: #111827; }
.pagination-bar button:disabled { opacity: .4; cursor: not-allowed; }
</style>
