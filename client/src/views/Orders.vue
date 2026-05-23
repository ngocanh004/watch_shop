<script setup>
import { onMounted, ref } from 'vue'
import { orderApi } from '../api'
import { formatMoney, formatDate, getOrderStatus } from '../utils/format'

const items = ref([])
const status = ref('')

async function load() {
  const res = await orderApi.myOrders({ trang_thai: status.value, page: 1, page_size: 20 })
  items.value = res.data
}

onMounted(load)
</script>

<template>
  <div class="container" style="padding-top:20px">
    <h2 class="section-title">Đơn hàng của tôi</h2>

    <div class="status-tabs">
      <button :class="{ active: !status }" @click="status = ''; load()">Tất cả</button>
      <button :class="{ active: status === 'cho_xu_ly' }" @click="status = 'cho_xu_ly'; load()">Chờ xử lý</button>
      <button :class="{ active: status === 'da_xac_nhan' }" @click="status = 'da_xac_nhan'; load()">Đã xác nhận</button>
      <button :class="{ active: status === 'dang_giao' }" @click="status = 'dang_giao'; load()">Đang giao</button>
      <button :class="{ active: status === 'hoan_thanh' }" @click="status = 'hoan_thanh'; load()">Hoàn thành</button>
      <button :class="{ active: status === 'da_huy' }" @click="status = 'da_huy'; load()">Đã huỷ</button>
    </div>

    <div v-if="!items.length" class="empty-state">
      <div class="icon">📦</div>
      <h3>Chưa có đơn hàng nào</h3>
      <router-link to="/products" class="btn btn-primary mt-3">Mua sắm ngay</router-link>
    </div>

    <div v-else>
      <div v-for="o in items" :key="o.id" class="order-card">
        <div class="order-head">
          <div>
            <strong>{{ o.ma_don_hang }}</strong>
            <span class="text-muted" style="margin-left:10px">{{ formatDate(o.ngay_dat_hang) }}</span>
          </div>
          <span :class="['badge', getOrderStatus(o.trang_thai).class]">
            {{ getOrderStatus(o.trang_thai).label }}
          </span>
        </div>
        <div class="order-body">
          <div v-for="it in o.chi_tiet_don_hang" :key="it.id" class="order-item">
            <img :src="it.hinh_anh" />
            <div class="flex-1">
              <div>{{ it.ten_san_pham }}</div>
              <div class="text-muted" style="font-size:13px">x{{ it.so_luong }}</div>
            </div>
            <div>{{ formatMoney(it.thanh_tien) }}</div>
          </div>
        </div>
        <div class="order-foot">
          <div>Tổng: <strong style="color:#dc2626">{{ formatMoney(o.tong_tien) }}</strong></div>
          <router-link :to="`/orders/${o.id}`" class="btn btn-sm btn-outline">Xem chi tiết</router-link>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.status-tabs {
  display: flex;
  gap: 8px;
  background: #fff;
  padding: 6px;
  border-radius: 8px;
  margin-bottom: 16px;
  overflow-x: auto;
}
.status-tabs button {
  padding: 8px 16px;
  background: none;
  border: none;
  border-radius: 6px;
  cursor: pointer;
  font-size: 14px;
  white-space: nowrap;
}
.status-tabs button.active { background: #111827; color: #fff; }
.order-card {
  background: #fff;
  border-radius: 8px;
  margin-bottom: 14px;
  overflow: hidden;
}
.order-head {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 14px 18px;
  border-bottom: 1px solid #e5e7eb;
}
.order-body { padding: 12px 18px; }
.order-item {
  display: flex;
  gap: 12px;
  align-items: center;
  padding: 8px 0;
}
.order-item img { width: 60px; height: 60px; object-fit: cover; border-radius: 4px; }
.flex-1 { flex: 1; }
.order-foot {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 14px 18px;
  background: #f9fafb;
}
</style>
