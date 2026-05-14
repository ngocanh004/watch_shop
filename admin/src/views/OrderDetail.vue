<script setup>
import { onMounted, ref } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { orderApi } from '../api'
import { useToastStore } from '../stores/toast'
import { formatMoney, formatDate, getOrderStatus, ORDER_STATUS_LIST } from '../utils/format'

const route = useRoute()
const router = useRouter()
const toast = useToastStore()

const order = ref(null)
const newStatus = ref('')

async function load() {
  const res = await orderApi.detail(route.params.id)
  order.value = res.data
  newStatus.value = res.data.trang_thai
}

async function update() {
  try {
    await orderApi.updateStatus(order.value.id, newStatus.value)
    toast.success('Cập nhật trạng thái thành công')
    load()
  } catch (e) {
    toast.error(e.message)
  }
}

onMounted(load)
</script>

<template>
  <div v-if="order">
    <div class="flex justify-between items-center mb-3">
      <h1>Chi tiết đơn hàng {{ order.ma_don_hang }}</h1>
      <button class="btn btn-outline" @click="router.back()">← Quay lại</button>
    </div>

    <div class="grid">
      <div class="card">
        <h3 class="mb-3">Thông tin đơn hàng</h3>
        <div class="info-row"><span>Mã đơn:</span> <strong>{{ order.ma_don_hang }}</strong></div>
        <div class="info-row"><span>Ngày đặt:</span> {{ formatDate(order.ngay_dat_hang) }}</div>
        <div class="info-row"><span>Phương thức:</span> {{ order.phuong_thuc?.toUpperCase() }}</div>
        <div class="info-row"><span>Tổng tiền:</span> <strong style="color:#dc2626">{{ formatMoney(order.tong_tien) }}</strong></div>
        <div class="info-row">
          <span>Trạng thái:</span>
          <span :class="['badge', getOrderStatus(order.trang_thai).class]">
            {{ getOrderStatus(order.trang_thai).label }}
          </span>
        </div>
      </div>

      <div class="card">
        <h3 class="mb-3">Thông tin giao hàng</h3>
        <div class="info-row"><span>Người nhận:</span> <strong>{{ order.ho_ten }}</strong></div>
        <div class="info-row"><span>SĐT:</span> {{ order.so_dien_thoai }}</div>
        <div class="info-row"><span>Địa chỉ:</span> {{ order.dia_chi }}</div>
        <div class="info-row" v-if="order.ghi_chu"><span>Ghi chú:</span> {{ order.ghi_chu }}</div>
      </div>
    </div>

    <div class="card">
      <h3 class="mb-3">Sản phẩm trong đơn</h3>
      <table class="table">
        <thead>
          <tr>
            <th>Ảnh</th>
            <th>Sản phẩm</th>
            <th>Đơn giá</th>
            <th>Số lượng</th>
            <th>Thành tiền</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="it in order.chi_tiet_don_hang" :key="it.id">
            <td><img :src="it.hinh_anh" class="thumb" /></td>
            <td>{{ it.ten_san_pham }}</td>
            <td>{{ formatMoney(it.gia) }}</td>
            <td>{{ it.so_luong }}</td>
            <td><strong>{{ formatMoney(it.thanh_tien) }}</strong></td>
          </tr>
        </tbody>
        <tfoot>
          <tr>
            <td colspan="4" class="text-right"><strong>Tổng cộng:</strong></td>
            <td><strong style="color:#dc2626">{{ formatMoney(order.tong_tien) }}</strong></td>
          </tr>
        </tfoot>
      </table>
    </div>

    <div class="card">
      <h3 class="mb-3">Cập nhật trạng thái</h3>
      <div class="flex gap-3 items-center">
        <select v-model="newStatus" class="select" style="max-width:240px">
          <option v-for="s in ORDER_STATUS_LIST" :key="s.value" :value="s.value">{{ s.label }}</option>
        </select>
        <button class="btn btn-primary" @click="update">Cập nhật</button>
      </div>
    </div>
  </div>
</template>

<style scoped>
.grid {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 16px;
}
.info-row {
  display: flex;
  justify-content: space-between;
  padding: 8px 0;
  border-bottom: 1px dashed #e5e7eb;
}
.info-row span:first-child { color: #6b7280; }
.thumb { width: 50px; height: 50px; object-fit: cover; border-radius: 4px; }
</style>
