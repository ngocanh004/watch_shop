export function formatMoney(n) {
  if (n == null) return '0 đ'
  return Number(n).toLocaleString('vi-VN') + ' đ'
}

export function formatDate(s) {
  if (!s) return ''
  return new Date(s).toLocaleString('vi-VN')
}

const ORDER_STATUS = {
  cho_xu_ly:    { label: 'Chờ xử lý',   class: 'badge-warning' },
  da_xac_nhan:  { label: 'Đã xác nhận', class: 'badge-info'    },
  dang_giao:    { label: 'Đang giao',   class: 'badge-info'    },
  hoan_thanh:   { label: 'Hoàn thành',  class: 'badge-success' },
  da_huy:       { label: 'Đã huỷ',      class: 'badge-danger'  }
}

export function getOrderStatus(key) {
  return ORDER_STATUS[key] || { label: key, class: 'badge-default' }
}

export function discountPercent(gia, gia_goc) {
  if (!gia_goc || gia_goc <= gia) return 0
  return Math.round(((gia_goc - gia) / gia_goc) * 100)
}
