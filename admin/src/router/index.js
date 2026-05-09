import { createRouter, createWebHistory } from 'vue-router'
import { useAuthStore } from '../stores/auth'

const routes = [
  {
    path: '/login',
    component: () => import('../views/Login.vue'),
    meta: { layout: 'blank' }
  },
  {
    path: '/',
    component: () => import('../layouts/AdminLayout.vue'),
    meta: { requiresAuth: true },
    children: [
      { path: '', redirect: '/dashboard' },
      { path: 'dashboard', component: () => import('../views/Dashboard.vue') },
      { path: 'products', component: () => import('../views/Products.vue') },
      { path: 'categories', component: () => import('../views/Categories.vue') },
      { path: 'brands', component: () => import('../views/Brands.vue') },
      { path: 'orders', component: () => import('../views/Orders.vue') },
      { path: 'orders/:id', component: () => import('../views/OrderDetail.vue') },
      { path: 'users', component: () => import('../views/Users.vue') },
      { path: 'profile', component: () => import('../views/Profile.vue') }
    ]
  },
  { path: '/:pathMatch(.*)*', redirect: '/dashboard' }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

router.beforeEach((to) => {
  const auth = useAuthStore()
  if (to.meta.requiresAuth && !auth.isLogin) {
    return '/login'
  }
  if (to.path === '/login' && auth.isLogin) {
    return '/dashboard'
  }
})

export default router
