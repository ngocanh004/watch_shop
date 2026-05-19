import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '../stores/auth'

const routes = [
  {
    path: '/',
    component: () => import('../layouts/ShopLayout.vue'),
    children: [
      { path: '', component: () => import('../views/Home.vue') },
      { path: 'products', component: () => import('../views/Products.vue') },
      { path: 'products/:id', component: () => import('../views/ProductDetail.vue') },
      { path: 'cart', component: () => import('../views/Cart.vue'), meta: { requiresAuth: true } },
      { path: 'checkout', component: () => import('../views/Checkout.vue'), meta: { requiresAuth: true } },
      { path: 'wishlist', component: () => import('../views/Wishlist.vue'), meta: { requiresAuth: true } },
      { path: 'orders', component: () => import('../views/Orders.vue'), meta: { requiresAuth: true } },
      { path: 'orders/:id', component: () => import('../views/OrderDetail.vue'), meta: { requiresAuth: true } },
      { path: 'profile', component: () => import('../views/Profile.vue'), meta: { requiresAuth: true } }
    ]
  },
  { path: '/login', component: () => import('../views/Login.vue') },
  { path: '/register', component: () => import('../views/Register.vue') },
  { path: '/:pathMatch(.*)*', redirect: '/' }
]

const router = createRouter({
  history: createWebHistory(),
  routes,
  scrollBehavior() {
    return { top: 0 }
  }
})

router.beforeEach((to) => {
  const auth = useAuthStore()
  if (to.meta.requiresAuth && !auth.isLogin) {
    return { path: '/login', query: { redirect: to.fullPath } }
  }
})

export default router
