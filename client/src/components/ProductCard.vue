<script setup>
import { formatMoney, discountPercent } from '../utils/format'

defineProps({
  product: { type: Object, required: true }
})
</script>

<template>
  <router-link :to="`/products/${product.id}`" class="product-card">
    <div style="position:relative">
      <img :src="product.hinh_anh" :alt="product.ten_san_pham" class="product-img" loading="lazy" />
      <span v-if="discountPercent(product.gia, product.gia_goc) > 0" class="discount-tag">
        -{{ discountPercent(product.gia, product.gia_goc) }}%
      </span>
    </div>
    <div class="product-info">
      <div class="product-brand">{{ product.brand?.ten_thuong_hieu }}</div>
      <div class="product-name">{{ product.ten_san_pham }}</div>
      <div class="product-price">
        <span class="price-main">{{ formatMoney(product.gia) }}</span>
        <span v-if="product.gia_goc > product.gia" class="price-old">{{ formatMoney(product.gia_goc) }}</span>
      </div>
    </div>
  </router-link>
</template>

<style scoped>
.discount-tag {
  position: absolute;
  top: 10px; left: 10px;
  background: #dc2626;
  color: #fff;
  font-size: 12px;
  font-weight: 600;
  padding: 2px 8px;
  border-radius: 4px;
}
</style>
