<script setup>
import { onMounted, ref } from 'vue'
import { productApi, brandApi, categoryApi } from '../api'
import ProductCard from '../components/ProductCard.vue'

const featured = ref([])
const newest = ref([])
const brands = ref([])
const categories = ref([])

async function load() {
  const [f, n, b, c] = await Promise.all([
    productApi.list({ noi_bat: 1, page_size: 8 }),
    productApi.list({ sort: 'moi_nhat', page_size: 8 }),
    brandApi.list(),
    categoryApi.list()
  ])
  featured.value = f.data
  newest.value = n.data
  brands.value = b.data
  categories.value = c.data
}

onMounted(load)
</script>

<template>
  <section class="hero">
    <div class="container hero-content">
      <div>
        <h1>Đồng hồ chính hãng - Giá tốt nhất</h1>
        <p>Hàng trăm mẫu đồng hồ từ các thương hiệu nổi tiếng. Bảo hành chính hãng, miễn phí giao hàng toàn quốc.</p>
        <router-link to="/products" class="btn">Mua sắm ngay →</router-link>
      </div>
      <div>
        <img src="https://images.unsplash.com/photo-1523275335684-37898b6baf30?w=600" alt="Watch" />
      </div>
    </div>
  </section>

  <div class="container">
    <h2 class="section-title">Danh mục sản phẩm</h2>
    <div class="cat-grid">
      <router-link
        v-for="c in categories"
        :key="c.id"
        :to="{ path: '/products', query: { category_id: c.id } }"
        class="cat-card"
      >
        <div class="cat-icon">⌚</div>
        <div>{{ c.ten_danh_muc }}</div>
      </router-link>
    </div>

    <h2 class="section-title">Sản phẩm nổi bật</h2>
    <div class="product-grid">
      <ProductCard v-for="p in featured" :key="p.id" :product="p" />
    </div>

    <h2 class="section-title">Sản phẩm mới</h2>
    <div class="product-grid">
      <ProductCard v-for="p in newest" :key="p.id" :product="p" />
    </div>

    <h2 class="section-title">Thương hiệu</h2>
    <div class="brand-grid">
      <router-link
        v-for="b in brands"
        :key="b.id"
        :to="{ path: '/products', query: { brand_id: b.id } }"
        class="brand-card"
      >
        <div class="brand-name-large">{{ b.ten_thuong_hieu }}</div>
        <div class="brand-country">{{ b.quoc_gia }}</div>
      </router-link>
    </div>
  </div>
</template>

<style scoped>
.cat-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(160px, 1fr));
  gap: 16px;
  margin-bottom: 30px;
}
.cat-card {
  background: #fff;
  border: 1px solid #e5e7eb;
  border-radius: 8px;
  padding: 24px;
  text-align: center;
  transition: all .2s;
}
.cat-card:hover { border-color: #111827; transform: translateY(-2px); }
.cat-icon { font-size: 36px; margin-bottom: 8px; }
.brand-grid {
  display: grid;
  grid-template-columns: repeat(auto-fit, minmax(160px, 1fr));
  gap: 14px;
  margin-bottom: 30px;
}
.brand-card {
  background: #fff;
  border: 1px solid #e5e7eb;
  border-radius: 8px;
  padding: 18px;
  text-align: center;
}
.brand-card:hover { border-color: #111827; }
.brand-name-large { font-weight: 700; font-size: 16px; color: #111827; }
.brand-country { font-size: 12px; color: #6b7280; margin-top: 4px; }
</style>
