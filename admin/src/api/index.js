import http from './http'

export const authApi = {
  login: (data) => http.post('/auth/dang-nhap', data),
  me: () => http.get('/me'),
  changePassword: (data) => http.post('/auth/doi-mat-khau', data)
}

export const userApi = {
  list: (params) => http.get('/admin/users', { params }),
  get: (id) => http.get(`/admin/users/${id}`),
  update: (id, data) => http.put(`/admin/users/${id}`, data),
  remove: (id) => http.delete(`/admin/users/${id}`),
  resetPassword: (id, data) => http.post(`/admin/users/${id}/reset-password`, data)
}

export const categoryApi = {
  list: (params) => http.get('/categories', { params }),
  create: (data) => http.post('/admin/categories', data),
  update: (id, data) => http.put(`/admin/categories/${id}`, data),
  remove: (id) => http.delete(`/admin/categories/${id}`)
}

export const brandApi = {
  list: (params) => http.get('/brands', { params }),
  create: (data) => http.post('/admin/brands', data),
  update: (id, data) => http.put(`/admin/brands/${id}`, data),
  remove: (id) => http.delete(`/admin/brands/${id}`)
}

export const productApi = {
  list: (params) => http.get('/products', { params: { ...params, admin: 1 } }),
  get: (id) => http.get(`/products/${id}`),
  create: (data) => http.post('/admin/products', data),
  update: (id, data) => http.put(`/admin/products/${id}`, data),
  remove: (id) => http.delete(`/admin/products/${id}`)
}

export const orderApi = {
  list: (params) => http.get('/admin/orders', { params }),
  detail: (id) => http.get(`/admin/orders/${id}`),
  updateStatus: (id, trang_thai) => http.put(`/admin/orders/${id}/trang-thai`, { trang_thai })
}

export const uploadApi = {
  image: (file) => {
    const fd = new FormData()
    fd.append('file', file)
    return http.post('/admin/upload', fd, { headers: { 'Content-Type': 'multipart/form-data' } })
  }
}

export const dashboardApi = {
  stats: () => http.get('/admin/thong-ke')
}
