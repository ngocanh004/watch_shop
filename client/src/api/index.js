import http from './http'

export const authApi = {
  register: (data) => http.post('/auth/dang-ky', data),
  login: (data) => http.post('/auth/dang-nhap', data),
  me: () => http.get('/me'),
  updateProfile: (data) => http.put('/me', data),
  changePassword: (data) => http.post('/auth/doi-mat-khau', data)
}

export const productApi = {
  list: (params) => http.get('/products', { params }),
  get: (id) => http.get(`/products/${id}`)
}

export const categoryApi = {
  list: () => http.get('/categories')
}

export const brandApi = {
  list: () => http.get('/brands')
}

export const cartApi = {
  list: () => http.get('/cart'),
  add: (product_id, so_luong = 1) => http.post('/cart', { product_id, so_luong }),
  update: (id, so_luong) => http.put(`/cart/${id}`, { so_luong }),
  remove: (id) => http.delete(`/cart/${id}`),
  clear: () => http.delete('/cart')
}

export const wishlistApi = {
  list: () => http.get('/wishlist'),
  add: (product_id) => http.post('/wishlist', { product_id }),
  remove: (product_id) => http.delete(`/wishlist/${product_id}`)
}

export const orderApi = {
  checkout: (data) => http.post('/orders', data),
  myOrders: (params) => http.get('/orders/me', { params }),
  detail: (id) => http.get(`/orders/me/${id}`),
  cancel: (id) => http.put(`/orders/me/${id}/huy`)
}
