<script setup>
import { onMounted, ref, watch } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { productApi, wishlistApi } from '../api'
import { useCartStore } from '../stores/cart'
import { useAuthStore } from '../stores/auth'
import { useToastStore } from '../stores/toast'
import { formatMoney, discountPercent } from '../utils/format'
import ProductCard from '../components/ProductCard.vue'

const route = useRoute()
const router = useRouter()
const cart = useCartStore()
const auth = useAuthStore()
const toast = useToastStore()

const product = ref(null)
const qty = ref(1)
const tab = ref('mota')
const related = ref([])

async function load() {
  const res = await productApi.get(route.params.id)
  product.value = res.data
  qty.value = 1
  if (res.data.brand_id) {
    const rel = await productApi.list({ brand_id: res.data.brand_id, page_size: 4 })
    related.value = rel.data.filter(p => p.id !== res.data.id).slice(0, 4)
  }
}

async function addToCart() {
  if (!auth.isLogin) {
    toast.info('Vui lòng đăng nhập để mua hàng')
    router.push({ path: '/login', query: { redirect: route.fullPath } })
    return
  }
  try {
    await cart.add(product.value.id, qty.value)
    toast.success('Đã thêm vào giỏ hàng')
  } catch (e) {
    toast.error(e.message)
  }
}

async function buyNow() {
  if (!auth.isLogin) {
    router.push({ path: '/login', query: { redirect: route.fullPath } })
    return
  }
  try {
    await cart.add(product.value.id, qty.value)
    router.push('/checkout')
  } catch (e) {
    toast.error(e.message)
  }
}

async function addWish() {
  if (!auth.isLogin) {
    router.push({ path: '/login', query: { redirect: route.fullPath } })
    return
  }
  try {
    await wishlistApi.add(product.value.id)
    toast.success('Đã thêm vào danh sách yêu thích')
  } catch (e) {
    toast.error(e.message)
  }
}

function incr() {
  if (qty.value < product.value.so_luong) qty.value++
}
function decr() {
  if (qty.value > 1) qty.value--
}

watch(() => route.params.id, () => { if (route.params.id) load() })
onMounted(load)
</script>

<template>
  <div v-if="product" class="container" style="padding-top: 20px">
    <div class="breadcrumb">
      <router-link to="/">Trang chủ</router-link> /
      <router-link to="/products">Sản phẩm</router-link> /
      <span>{{ product.ten_san_pham }}</span>
    </div>

    <div class="detail-wrap">
      <div class="image-box">
        <img :src="product.hinh_anh" :alt="product.ten_san_pham" />
        <span v-if="discountPercent(product.gia, product.gia_goc) > 0" class="discount-big">
          -{{ discountPercent(product.gia, product.gia_goc) }}%
        </span>
      </div>

      <div class="info-box">
        <div class="info-brand">{{ product.brand?.ten_thuong_hieu }} · {{ product.brand?.quoc_gia }}</div>
        <h1 class="info-name">{{ product.ten_san_pham }}</h1>
        <div class="info-meta">
          <span>Đã bán: <strong>{{ product.da_ban }}</strong></span>
          <span>Còn lại: <strong>{{ product.so_luong }}</strong></span>
        </div>

        <div class="price-box">
          <div class="price-now">{{ formatMoney(product.gia) }}</div>
          <div v-if="product.gia_goc > product.gia" class="price-old-big">{{ formatMoney(product.gia_goc) }}</div>
        </div>

        <div class="specs">
          <div class="spec-row"><span>Loại máy:</span><strong>{{ product.loai_may }}</strong></div>
          <div class="spec-row"><span>Loại dây:</span><strong>{{ product.day_deo }}</strong></div>
          <div class="spec-row"><span>Mặt kính:</span><strong>{{ product.mat_kinh }}</strong></div>
          <div class="spec-row"><span>Đường kính:</span><strong>{{ product.duong_kinh }}</strong></div>
          <div class="spec-row"><span>Kháng nước:</span><strong>{{ product.khang_nuoc }}</strong></div>
          <div class="spec-row"><span>Giới tính:</span><strong>{{ product.gioi_tinh }}</strong></div>
          <div class="spec-row"><span>Xuất xứ:</span><strong>{{ product.xuat_xu }}</strong></div>
          <div class="spec-row"><span>Bảo hành:</span><strong>{{ product.bao_hanh }}</strong></div>
        </div>

        <div class="qty-wrap">
          <span>Số lượng:</span>
          <div class="qty-box">
            <button @click="decr">−</button>
            <input v-model.number="qty" type="number" min="1" />
            <button @click="incr">+</button>
          </div>
        </div>

        <div class="action-row">
          <button class="btn btn-outline" @click="addToCart">🛒 Thêm vào giỏ</button>
          <button class="btn btn-primary" @click="buyNow">Mua ngay</button>
          <button class="btn-icon" @click="addWish" title="Yêu thích">❤️</button>
        </div>

        <div class="promo">
          <div>🚚 Miễn phí giao hàng toàn quốc</div>
          <div>🔒 Cam kết hàng chính hãng 100%</div>
          <div>↩️ Đổi trả trong 7 ngày</div>
        </div>
      </div>
    </div>

    <div class="tabs-box">
      <div class="tabs-head">
        <button :class="{ active: tab === 'mota' }" @click="tab = 'mota'">Mô tả sản phẩm</button>
        <button :class="{ active: tab === 'thongtin' }" @click="tab = 'thongtin'">Thông số kỹ thuật</button>
      </div>
      <div class="tabs-body">
        <div v-if="tab === 'mota'">
          <p style="white-space: pre-line">{{ product.mo_ta }}</p>
        </div>
        <div v-else>
          <table class="spec-table">
            <tr><th>Thương hiệu</th><td>{{ product.brand?.ten_thuong_hieu }}</td></tr>
            <tr><th>Loại máy</th><td>{{ product.loai_may }}</td></tr>
            <tr><th>Loại dây</th><td>{{ product.day_deo }}</td></tr>
            <tr><th>Mặt kính</th><td>{{ product.mat_kinh }}</td></tr>
            <tr><th>Đường kính</th><td>{{ product.duong_kinh }}</td></tr>
            <tr><th>Kháng nước</th><td>{{ product.khang_nuoc }}</td></tr>
            <tr><th>Giới tính</th><td>{{ product.gioi_tinh }}</td></tr>
            <tr><th>Xuất xứ</th><td>{{ product.xuat_xu }}</td></tr>
            <tr><th>Bảo hành</th><td>{{ product.bao_hanh }}</td></tr>
          </table>
        </div>
      </div>
    </div>

    <div v-if="related.length">
      <h2 class="section-title">Sản phẩm cùng thương hiệu</h2>
      <div class="product-grid">
        <ProductCard v-for="p in related" :key="p.id" :product="p" />
      </div>
    </div>
  </div>
</template>

<style scoped>
.breadcrumb { padding: 8px 0 16px; color: #6b7280; font-size: 13px; }
.breadcrumb a { color: #374151; }

.detail-wrap {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 30px;
  background: #fff;
  padding: 24px;
  border-radius: 8px;
  margin-bottom: 24px;
}
.image-box {
  position: relative;
  background: #f3f4f6;
  border-radius: 8px;
  overflow: hidden;
  aspect-ratio: 1;
}
.image-box img { width: 100%; height: 100%; object-fit: cover; }
.discount-big {
  position: absolute;
  top: 16px; left: 16px;
  background: #dc2626;
  color: #fff;
  padding: 4px 12px;
  border-radius: 4px;
  font-weight: 600;
}
.info-brand { font-size: 13px; color: #6b7280; text-transform: uppercase; }
.info-name { font-size: 24px; margin: 8px 0; line-height: 1.3; }
.info-meta {
  display: flex;
  gap: 20px;
  color: #6b7280;
  font-size: 13px;
  padding-bottom: 16px;
  border-bottom: 1px solid #e5e7eb;
}
.price-box {
  display: flex;
  align-items: baseline;
  gap: 12px;
  padding: 16px;
  background: #fef3f2;
  border-radius: 6px;
  margin: 16px 0;
}
.price-now { font-size: 28px; font-weight: 700; color: #dc2626; }
.price-old-big { font-size: 15px; color: #9ca3af; text-decoration: line-through; }
.specs { margin: 16px 0; }
.spec-row {
  display: flex;
  justify-content: space-between;
  padding: 6px 0;
  font-size: 14px;
  border-bottom: 1px dashed #e5e7eb;
}
.spec-row span { color: #6b7280; }

.qty-wrap {
  display: flex;
  align-items: center;
  gap: 14px;
  margin: 20px 0;
}
.action-row {
  display: flex;
  gap: 10px;
  margin-bottom: 20px;
}
.action-row .btn { flex: 1; padding: 12px; }
.btn-icon {
  width: 46px;
  border: 1px solid #d1d5db;
  background: #fff;
  border-radius: 6px;
  cursor: pointer;
  font-size: 18px;
}
.promo {
  background: #f9fafb;
  padding: 14px;
  border-radius: 6px;
  font-size: 13px;
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.tabs-box { background: #fff; border-radius: 8px; margin-bottom: 24px; overflow: hidden; }
.tabs-head { display: flex; border-bottom: 1px solid #e5e7eb; }
.tabs-head button {
  padding: 14px 24px;
  background: none;
  border: none;
  font-size: 14px;
  font-weight: 600;
  color: #6b7280;
  cursor: pointer;
  border-bottom: 2px solid transparent;
}
.tabs-head button.active { color: #111827; border-color: #111827; }
.tabs-body { padding: 24px; }
.spec-table { width: 100%; border-collapse: collapse; }
.spec-table th, .spec-table td {
  padding: 10px 14px;
  text-align: left;
  border-bottom: 1px solid #e5e7eb;
}
.spec-table th { width: 200px; background: #f9fafb; color: #374151; }

@media (max-width: 768px) {
  .detail-wrap { grid-template-columns: 1fr; padding: 16px; }
}
</style>
