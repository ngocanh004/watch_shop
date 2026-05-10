<script setup>
import { onMounted, ref } from 'vue'
import { dashboardApi, orderApi } from '../api'
import { formatMoney, formatDate, getOrderStatus } from '../utils/format'

const stats = ref({})
const recent = ref([])
const loading = ref(true)

async function load() {
  loading.value = true
  try {
    const [s, o] = await Promise.all([
      dashboardApi.stats(),
      orderApi.list({ page: 1, page_size: 5 })
    ])
    stats.value = s.data
    recent.value = o.data
  } finally {
    loading.value = false
  }
}

onMounted(load)
</script>

<template>
  <h1 class="page-title">Tổng quan</h1>

  <div v-if="loading" class="text-muted">Đang tải...</div>

  <div v-else>
    <div class="stat-grid">
      <div class="stat-card stat-blue">
        <div class="stat-icon">🛒</div>
        <div>
          <div class="stat-label">Tổng đơn hàng</div>
          <div class="stat-value">{{ stats.tong_don_hang }}</div>
        </div>
      </div>
      <div class="stat-card stat-green">
        <div class="stat-icon">💰</div>
        <div>
          <div class="stat-label">Doanh thu (hoàn thành)</div>
          <div class="stat-value">{{ formatMoney(stats.doanh_thu) }}</div>
        </div>
      </div>
      <div class="stat-card stat-orange">
        <div class="stat-icon">⌚</div>
        <div>
          <div class="stat-label">Sản phẩm</div>
          <div class="stat-value">{{ stats.tong_san_pham }}</div>
        </div>
      </div>
      <div class="stat-card stat-purple">
        <div class="stat-icon">👥</div>
        <div>
          <div class="stat-label">Người dùng</div>
          <div class="stat-value">{{ stats.tong_nguoi_dung }}</div>
        </div>
      </div>
    </div>

    <div class="card">
      <h3 class="mb-3">Đơn hàng theo trạng thái</h3>
      <div class="status-grid">
        <div v-for="s in stats.theo_trang_thai" :key="s.trang_thai" class="status-item">
          <span :class="['badge', getOrderStatus(s.trang_thai).class]">
            {{ getOrderStatus(s.trang_thai).label }}
          </span>
          <strong>{{ s.so_luong }}</strong>
        </div>
      </div>
    </div>

    <div class="card">
      <h3 class="mb-3">Đơn hàng gần đây</h3>
      <table class="table">
        <thead>
          <tr>
            <th>Mã đơn</th>
            <th>Khách hàng</th>
            <th>Tổng tiền</th>
            <th>Trạng thái</th>
            <th>Ngày đặt</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="o in recent" :key="o.id">
            <td><router-link :to="`/orders/${o.id}`">{{ o.ma_don_hang }}</router-link></td>
            <td>{{ o.ho_ten }}</td>
            <td>{{ formatMoney(o.tong_tien) }}</td>
            <td>
              <span :class="['badge', getOrderStatus(o.trang_thai).class]">
                {{ getOrderStatus(o.trang_thai).label }}
              </span>
            </td>
            <td>{{ formatDate(o.ngay_dat_hang) }}</td>
          </tr>
          <tr v-if="!recent.length">
            <td colspan="5" class="text-center text-muted">Chưa có đơn hàng nào</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<style scoped>
.page-title { margin-bottom: 20px; }
.stat-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(220px, 1fr));
  gap: 16px;
  margin-bottom: 20px;
}
.stat-card {
  background: #fff;
  padding: 20px;
  border-radius: 8px;
  display: flex;
  align-items: center;
  gap: 16px;
  box-shadow: 0 1px 3px rgba(0,0,0,.06);
}
.stat-icon {
  width: 56px; height: 56px;
  border-radius: 12px;
  display: flex; align-items: center; justify-content: center;
  font-size: 28px;
}
.stat-blue .stat-icon { background: #dbeafe; }
.stat-green .stat-icon { background: #dcfce7; }
.stat-orange .stat-icon { background: #fed7aa; }
.stat-purple .stat-icon { background: #ede9fe; }
.stat-label { color: #6b7280; font-size: 13px; }
.stat-value { font-size: 22px; font-weight: 700; color: #111827; }
.status-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(180px, 1fr));
  gap: 12px;
}
.status-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px 14px;
  background: #f9fafb;
  border-radius: 6px;
}
</style>
