<script setup>
import { onMounted, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { orderApi } from '../api'
import { useToastStore } from '../stores/toast'
import { formatMoney, formatDate, getOrderStatus } from '../utils/format'

const route = useRoute()
const router = useRouter()
const toast = useToastStore()

const order = ref(null)

async function load() {
  const res = await orderApi.detail(route.params.id)
  order.value = res.data
}

async function cancel() {
  if (!confirm('Bạn chắc chắn muốn huỷ đơn hàng này?')) return
  try {
    await orderApi.cancel(order.value.id)
    toast.success('Đã huỷ đơn hàng')
    load()
  } catch (e) {
    toast.error(e.message)
  }
}

onMounted(load)
</script>

<template>
  <div v-if="order" class="container" style="padding-top:20px">
    <div class="flex-bar">
      <h2 class="section-title" style="margin:0">Chi tiết đơn hàng {{ order.ma_don_hang }}</h2>
      <button class="btn btn-outline" @click="router.back()">← Quay lại</button>
    </div>

    <div class="grid-2">
      <div class="box">
        <h3>Thông tin đơn hàng</h3>
        <div class="info-row"><span>Mã đơn:</span> <strong>{{ order.ma_don_hang }}</strong></div>
        <div class="info-row"><span>Ngày đặt:</span> {{ formatDate(order.ngay_dat_hang) }}</div>
        <div class="info-row"><span>Thanh toán:</span> {{ order.phuong_thuc?.toUpperCase() }}</div>
        <div class="info-row">
          <span>Trạng thái:</span>
          <span :class="['badge', getOrderStatus(order.trang_thai).class]">
            {{ getOrderStatus(order.trang_thai).label }}
          </span>
        </div>
        <div class="info-row"><span>Tổng tiền:</span> <strong style="color:#dc2626">{{ formatMoney(order.tong_tien) }}</strong></div>
      </div>

      <div class="box">
        <h3>Giao hàng đến</h3>
        <div class="info-row"><span>Người nhận:</span> <strong>{{ order.ho_ten }}</strong></div>
        <div class="info-row"><span>SĐT:</span> {{ order.so_dien_thoai }}</div>
        <div class="info-row"><span>Địa chỉ:</span> {{ order.dia_chi }}</div>
        <div v-if="order.ghi_chu" class="info-row"><span>Ghi chú:</span> {{ order.ghi_chu }}</div>
      </div>
    </div>

    <div class="box">
      <h3 class="mb-3">Sản phẩm</h3>
      <div v-for="it in order.chi_tiet_don_hang" :key="it.id" class="line">
        <img :src="it.hinh_anh" />
        <div class="flex-1">
          <div style="font-weight:500">{{ it.ten_san_pham }}</div>
          <div class="text-muted" style="font-size:13px">{{ formatMoney(it.gia) }} × {{ it.so_luong }}</div>
        </div>
        <div style="font-weight:700; color:#dc2626">{{ formatMoney(it.thanh_tien) }}</div>
      </div>
      <div class="total-row">
        <span>Tổng cộng:</span>
        <strong>{{ formatMoney(order.tong_tien) }}</strong>
      </div>
    </div>

    <div v-if="order.trang_thai === 'cho_xu_ly'" class="text-right">
      <button class="btn btn-danger" @click="cancel">Huỷ đơn hàng</button>
    </div>
  </div>
</template>

<style scoped>
.flex-bar { display: flex; justify-content: space-between; align-items: center; margin-bottom: 16px; }
.grid-2 {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 16px;
}
.box {
  background: #fff;
  padding: 20px;
  border-radius: 8px;
  margin-bottom: 16px;
}
.box h3 { margin-bottom: 14px; }
.info-row {
  display: flex;
  justify-content: space-between;
  padding: 8px 0;
  border-bottom: 1px dashed #e5e7eb;
}
.info-row span:first-child { color: #6b7280; }
.line {
  display: flex;
  align-items: center;
  gap: 14px;
  padding: 10px 0;
  border-bottom: 1px solid #e5e7eb;
}
.line img { width: 60px; height: 60px; object-fit: cover; border-radius: 4px; }
.flex-1 { flex: 1; }
.total-row {
  display: flex;
  justify-content: space-between;
  padding: 14px 0 0;
  font-size: 16px;
  color: #dc2626;
}
@media (max-width: 768px) { .grid-2 { grid-template-columns: 1fr; } }
</style>
