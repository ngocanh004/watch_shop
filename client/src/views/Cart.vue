<script setup>
import { onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { useCartStore } from '../stores/cart'
import { useToastStore } from '../stores/toast'
import { formatMoney } from '../utils/format'

const router = useRouter()
const cart = useCartStore()
const toast = useToastStore()

async function changeQty(item, delta) {
  const newQty = item.so_luong + delta
  if (newQty < 1) return
  if (newQty > item.product.so_luong) {
    toast.error('Vượt quá tồn kho')
    return
  }
  try {
    await cart.update(item.id, newQty)
  } catch (e) {
    toast.error(e.message)
  }
}

async function remove(item) {
  if (!confirm(`Xoá "${item.product.ten_san_pham}" khỏi giỏ hàng?`)) return
  try {
    await cart.remove(item.id)
    toast.success('Đã xoá khỏi giỏ hàng')
  } catch (e) {
    toast.error(e.message)
  }
}

async function clearAll() {
  if (!cart.items.length) return
  if (!confirm('Xoá toàn bộ giỏ hàng?')) return
  await cart.clear()
  toast.info('Đã xoá toàn bộ giỏ hàng')
}

function checkout() {
  if (!cart.items.length) {
    toast.error('Giỏ hàng đang trống')
    return
  }
  router.push('/checkout')
}

onMounted(() => cart.load())
</script>

<template>
  <div class="container" style="padding-top:20px">
    <h2 class="section-title">Giỏ hàng của bạn</h2>

    <div v-if="cart.loading" class="text-center text-muted" style="padding:40px">Đang tải...</div>

    <div v-else-if="!cart.items.length" class="empty-state">
      <div class="icon">🛒</div>
      <h3>Giỏ hàng đang trống</h3>
      <p class="mb-3">Hãy thêm sản phẩm vào giỏ để tiếp tục mua sắm</p>
      <router-link to="/products" class="btn btn-primary">Tiếp tục mua sắm</router-link>
    </div>

    <div v-else class="cart-wrap">
      <div>
        <div v-for="item in cart.items" :key="item.id" class="cart-item">
          <img :src="item.product.hinh_anh" :alt="item.product.ten_san_pham" />
          <div>
            <router-link :to="`/products/${item.product_id}`" class="cart-name">
              {{ item.product.ten_san_pham }}
            </router-link>
            <div class="text-muted" style="font-size:13px">{{ item.product.brand?.ten_thuong_hieu }}</div>
            <div class="cart-price">{{ formatMoney(item.product.gia) }}</div>
          </div>
          <div class="qty-box">
            <button @click="changeQty(item, -1)">−</button>
            <input :value="item.so_luong" readonly />
            <button @click="changeQty(item, 1)">+</button>
          </div>
          <div class="cart-subtotal">
            {{ formatMoney(item.product.gia * item.so_luong) }}
          </div>
          <button class="btn-trash" @click="remove(item)">🗑️</button>
        </div>

        <button class="btn btn-outline mt-3" @click="clearAll">Xoá toàn bộ giỏ hàng</button>
      </div>

      <aside class="cart-summary">
        <h3>Tóm tắt đơn hàng</h3>
        <div class="sum-row">
          <span>Tạm tính ({{ cart.count }} sản phẩm)</span>
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
        <button class="btn btn-primary btn-block mt-3" style="padding:12px" @click="checkout">
          Tiến hành đặt hàng →
        </button>
        <router-link to="/products" class="btn btn-outline btn-block mt-3">
          ← Tiếp tục mua sắm
        </router-link>
      </aside>
    </div>
  </div>
</template>

<style scoped>
.cart-wrap {
  display: grid;
  grid-template-columns: 1fr 320px;
  gap: 24px;
  align-items: start;
}
.cart-name { font-weight: 600; color: #1f2937; }
.cart-price { color: #dc2626; font-weight: 600; margin-top: 4px; }
.cart-subtotal { font-weight: 700; color: #dc2626; min-width: 100px; text-align: right; }
.btn-trash {
  background: none;
  border: none;
  font-size: 18px;
  cursor: pointer;
  padding: 4px 8px;
}
.cart-summary {
  background: #fff;
  padding: 20px;
  border-radius: 8px;
  position: sticky;
  top: 100px;
}
.cart-summary h3 { margin-bottom: 14px; }
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
  .cart-wrap { grid-template-columns: 1fr; }
}
</style>
