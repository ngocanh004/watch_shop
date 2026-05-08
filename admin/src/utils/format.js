export function formatMoney(n) {
  if (n == null) return '0 đ'
  return Number(n).toLocaleString('vi-VN') + ' đ'
}

export function formatDate(s) {
  if (!s) return ''
  const d = new Date(s)
  return d.toLocaleString('vi-VN')
}

export function formatDateShort(s) {
  if (!s) return ''
  const d = new Date(s)
  return d.toLocaleDateString('vi-VN')
}

const ORDER_STATUS = {
  cho_xu_ly: { label: 'Chờ xử lý', class: 'badge-warning' },
  da_xac_nhan: { label: 'Đã xác nhận', class: 'badge-info' },
  dang_giao: { label: 'Đang giao', class: 'badge-info' },
  hoan_thanh: { label: 'Hoàn thành', class: 'badge-success' },
  da_huy: { label: 'Đã huỷ', class: 'badge-danger' }
}

export function getOrderStatus(key) {
  return ORDER_STATUS[key] || { label: key, class: 'badge-default' }
}

export const ORDER_STATUS_LIST = Object.entries(ORDER_STATUS).map(([k, v]) => ({
  value: k, label: v.label
}))
