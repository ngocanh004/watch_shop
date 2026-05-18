import axios from 'axios'

const baseURL = import.meta.env.VITE_API_URL || 'http://localhost:8080/api'

const http = axios.create({
  baseURL,
  timeout: 20000
})

http.interceptors.request.use((cfg) => {
  const token = localStorage.getItem('client_token')
  if (token) cfg.headers.Authorization = `Bearer ${token}`
  return cfg
})

http.interceptors.response.use(
  (res) => res.data,
  (err) => {
    const data = err.response?.data || { message: 'Có lỗi xảy ra, vui lòng thử lại' }
    if (err.response?.status === 401) {
      const path = location.pathname
      const isAuthPage = path === '/login' || path === '/register'
      localStorage.removeItem('client_token')
      localStorage.removeItem('client_user')
      if (!isAuthPage) {
        location.href = '/login'
      }
    }
    return Promise.reject(data)
  }
)

export default http
export const PUBLIC_URL = import.meta.env.VITE_PUBLIC_URL || 'http://localhost:8080'
