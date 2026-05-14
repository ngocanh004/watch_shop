<script setup>
import { onMounted, reactive, ref } from 'vue'
import { orderApi } from '../api'
import { formatMoney, formatDate, getOrderStatus, ORDER_STATUS_LIST } from '../utils/format'

const items = ref([])
const total = ref(0)

const query = reactive({
  keyword: '',
  trang_thai: '',
  page: 1,
  page_size: 10
})

async function load() {
  const res = await orderApi.list(query)
  items.value = res.data
  total.value = res.total
}

function changePage(p) {
  query.page = p
  load()
}

onMounted(load)
</script>

<template>
  <h1 class="mb-3">Quản lý đơn hàng</h1>

  <div class="card">
    <div class="toolbar">
      <input v-model="query.keyword" class="input" placeholder="Mã đơn, tên, SĐT..." @keyup.enter="load" />
      <select v-model="query.trang_thai" class="select" @change="load">
        <option value="">-- Tất cả trạng thái --</option>
        <option v-for="s in ORDER_STATUS_LIST" :key="s.value" :value="s.value">{{ s.label }}</option>
      </select>
      <button class="btn btn-outline" @click="load">Lọc</button>
    </div>

    <table class="table">
      <thead>
        <tr>
          <th>Mã đơn</th>
          <th>Khách hàng</th>
          <th>SĐT</th>
          <th>Tổng tiền</th>
          <th>Phương thức</th>
          <th>Trạng thái</th>
          <th>Ngày đặt</th>
          <th>Chi tiết</th>
        </tr>
      </thead>
      <tbody>
        <tr v-for="o in items" :key="o.id">
          <td><strong>{{ o.ma_don_hang }}</strong></td>
          <td>{{ o.ho_ten }}</td>
          <td>{{ o.so_dien_thoai }}</td>
          <td>{{ formatMoney(o.tong_tien) }}</td>
          <td style="text-transform: uppercase">{{ o.phuong_thuc }}</td>
          <td>
            <span :class="['badge', getOrderStatus(o.trang_thai).class]">
              {{ getOrderStatus(o.trang_thai).label }}
            </span>
          </td>
          <td>{{ formatDate(o.ngay_dat_hang) }}</td>
          <td>
            <router-link :to="`/orders/${o.id}`" class="btn btn-sm btn-outline">Xem</router-link>
          </td>
        </tr>
        <tr v-if="!items.length">
          <td colspan="8" class="text-center text-muted">Không có đơn hàng</td>
        </tr>
      </tbody>
    </table>

    <div class="pagination">
      <button :disabled="query.page === 1" @click="changePage(query.page - 1)">‹</button>
      <span class="text-muted" style="padding: 6px 12px;">Trang {{ query.page }} / {{ Math.ceil(total / query.page_size) || 1 }}</span>
      <button :disabled="query.page * query.page_size >= total" @click="changePage(query.page + 1)">›</button>
    </div>
  </div>
</template>
