<script setup>
import { onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'
import { useCartStore } from '../stores/cart'
import { useAuthStore } from '../stores/auth'
import { useToastStore } from '../stores/toast'
import { orderApi } from '../api'
import { formatMoney } from '../utils/format'

const router = useRouter()
const cart = useCartStore()
const auth = useAuthStore()
const toast = useToastStore()

const form = ref({
  ho_ten: '',
  so_dien_thoai: '',
  dia_chi: '',
  ghi_chu: '',
  phuong_thuc: 'cod'
})
const loading = ref(false)
const errMsg = ref('')

onMounted(async () => {
  await cart.load()
  if (!cart.items.length) {
    router.push('/cart')
    return
  }
  if (auth.user) {
    form.value.ho_ten = auth.user.ho_ten || ''
    form.value.so_dien_thoai = auth.user.so_dien_thoai || ''
    form.value.dia_chi = auth.user.dia_chi || ''
  }
})

async function submit() {
  errMsg.value = ''
  if (!form.value.ho_ten || !form.value.so_dien_thoai || !form.value.dia_chi) {
    errMsg.value = 'Vui lòng nhập đầy đủ thông tin giao hàng'
    return
  }
  if (!/^0\d{9}$/.test(form.value.so_dien_thoai)) {
    errMsg.value = 'Số điện thoại không hợp lệ (10 số, bắt đầu bằng 0)'
    return
  }
  loading.value = true
  try {
    const res = await orderApi.checkout(form.value)
    await cart.load()
    toast.success('Đặt hàng thành công!')
    router.push(`/orders/${res.data.id}`)
  } catch (e) {
    errMsg.value = e.message
  } finally {
    loading.value = false
  }
}
</script>

<template>
  <div class="container" style="padding-top:20px">
    <h2 class="section-title">Thanh toán</h2>

    <div class="checkout-wrap">
      <div class="checkout-form">
        <div class="box">
          <h3 class="mb-3">Thông tin giao hàng</h3>
          <div class="form-group">
            <label>Họ và tên *</label>
            <input v-model="form.ho_ten" class="input" />
          </div>
          <div class="form-group">
            <label>Số điện thoại *</label>
            <input v-model="form.so_dien_thoai" class="input" placeholder="0123456789" />
          </div>
          <div class="form-group">
            <label>Địa chỉ nhận hàng *</label>
            <textarea v-model="form.dia_chi" rows="2" placeholder="Số nhà, đường, phường/xã, quận/huyện, tỉnh/thành"></textarea>
          </div>
          <div class="form-group">
            <label>Ghi chú</label>
            <textarea v-model="form.ghi_chu" rows="2" placeholder="Ghi chú thêm cho đơn hàng (tuỳ chọn)"></textarea>
          </div>
        </div>

        <div class="box">
          <h3 class="mb-3">Phương thức thanh toán</h3>
          <label class="payment-option">
            <input type="radio" v-model="form.phuong_thuc" value="cod" />
            <div>
              <strong>Thanh toán khi nhận hàng (COD)</strong>
              <div class="text-muted" style="font-size:13px">Trả tiền mặt khi nhận hàng</div>
            </div>
          </label>
          <label class="payment-option">
            <input type="radio" v-model="form.phuong_thuc" value="bank" />
            <div>
              <strong>Chuyển khoản ngân hàng</strong>
              <div class="text-muted" style="font-size:13px">Liên hệ shop để nhận thông tin tài khoản</div>
            </div>
          </label>
        </div>

        <div v-if="errMsg" class="text-error">{{ errMsg }}</div>
      </div>

      <aside class="checkout-summary">
        <h3 class="mb-3">Đơn hàng của bạn</h3>
        <div v-for="it in cart.items" :key="it.id" class="sum-item">
          <img :src="it.product.hinh_anh" />
          <div class="flex-1">
            <div style="font-size:13px">{{ it.product.ten_san_pham }}</div>
            <div class="text-muted" style="font-size:12px">SL: {{ it.so_luong }}</div>
          </div>
          <div style="font-weight:600">{{ formatMoney(it.product.gia * it.so_luong) }}</div>
        </div>

        <div class="sum-row">
          <span>Tạm tính</span>
          <strong>{{ formatMoney(cart.tong_tien) }}</strong>
        </div>
        <div class="sum-row">
          <span>Phí vận chuyển</span>
          <strong>Miễn phí</strong>
        </div>
        <div class="sum-row total">
          <span>Tổng cộng</span>
          <strong>{{ formatMoney(cart.tong_tien) }}</strong>
        </div>

        <button class="btn btn-primary btn-block mt-3" style="padding:12px" :disabled="loading" @click="submit">
          {{ loading ? 'Đang xử lý...' : 'Đặt hàng' }}
        </button>
      </aside>
    </div>
  </div>
</template>

<style scoped>
.checkout-wrap {
  display: grid;
  grid-template-columns: 1fr 360px;
  gap: 24px;
  align-items: start;
}
.box {
  background: #fff;
  padding: 20px;
  border-radius: 8px;
  margin-bottom: 16px;
}
.payment-option {
  display: flex;
  gap: 12px;
  align-items: center;
  padding: 12px 14px;
  border: 1px solid #d1d5db;
  border-radius: 6px;
  margin-bottom: 10px;
  cursor: pointer;
}
.payment-option:hover { border-color: #111827; }
.checkout-summary {
  background: #fff;
  padding: 20px;
  border-radius: 8px;
  position: sticky;
  top: 100px;
}
.sum-item {
  display: flex;
  gap: 10px;
  align-items: center;
  padding: 10px 0;
  border-bottom: 1px dashed #e5e7eb;
}
.sum-item img { width: 50px; height: 50px; object-fit: cover; border-radius: 4px; }
.flex-1 { flex: 1; }
.sum-row {
  display: flex;
  justify-content: space-between;
  padding: 8px 0;
  font-size: 14px;
}
.sum-row.total {
  border-top: 1px solid #e5e7eb;
  margin-top: 8px;
  padding-top: 14px;
  font-size: 16px;
  color: #dc2626;
}
@media (max-width: 768px) {
  .checkout-wrap { grid-template-columns: 1fr; }
}
</style>
