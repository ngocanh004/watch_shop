<script setup>
import { onMounted, ref } from 'vue'
import { wishlistApi } from '../api'
import { useCartStore } from '../stores/cart'
import { useToastStore } from '../stores/toast'
import { formatMoney } from '../utils/format'

const cart = useCartStore()
const toast = useToastStore()

const items = ref([])

async function load() {
  const res = await wishlistApi.list()
  items.value = res.data
}

async function remove(it) {
  try {
    await wishlistApi.remove(it.product_id)
    toast.success('Đã xoá khỏi yêu thích')
    load()
  } catch (e) {
    toast.error(e.message)
  }
}

async function addCart(it) {
  try {
    await cart.add(it.product_id, 1)
    toast.success('Đã thêm vào giỏ hàng')
  } catch (e) {
    toast.error(e.message)
  }
}

onMounted(load)
</script>

<template>
  <div class="container" style="padding-top:20px">
    <h2 class="section-title">Danh sách yêu thích</h2>

    <div v-if="!items.length" class="empty-state">
      <div class="icon">❤️</div>
      <h3>Chưa có sản phẩm yêu thích</h3>
      <p class="mb-3">Lưu lại những mẫu đồng hồ bạn thích để xem sau</p>
      <router-link to="/products" class="btn btn-primary">Khám phá sản phẩm</router-link>
    </div>

    <div v-else class="product-grid">
      <div v-for="it in items" :key="it.id" class="product-card">
        <router-link :to="`/products/${it.product_id}`">
          <img :src="it.product.hinh_anh" class="product-img" />
        </router-link>
        <div class="product-info">
          <div class="product-brand">{{ it.product.brand?.ten_thuong_hieu }}</div>
          <router-link :to="`/products/${it.product_id}`" class="product-name">
            {{ it.product.ten_san_pham }}
          </router-link>
          <div class="product-price mb-3">
            <span class="price-main">{{ formatMoney(it.product.gia) }}</span>
          </div>
          <div class="wish-actions">
            <button class="btn btn-sm btn-primary" @click="addCart(it)">🛒 Thêm vào giỏ</button>
            <button class="btn btn-sm btn-outline" @click="remove(it)">Xoá</button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<style scoped>
.wish-actions { display: flex; gap: 6px; }
.wish-actions .btn { flex: 1; }
</style>
