import request from '@/utils/request'

export function uplink(data) {
  return request({
    url: '/uplink',
    method: 'post',
    headers: {
      'Content-Type': 'multipart/form-data'
    },
    data
  })
}

// getMoutaiInfo
export function getMoutaiInfo(data) {
  return request({
    url: '/getMoutaiInfo',
    method: 'post',
    headers: {
      'Content-Type': 'multipart/form-data'
    },
    data
  })
}

// getMoutaiList
export function getMoutaiList(data) {
  return request({
    url: '/getMoutaiList',
    method: 'post',
    headers: {
      'Content-Type': 'multipart/form-data'
    },
    data
  })
}

// getAllMoutaiInfo
export function getAllMoutaiInfo(data) {
  return request({
    url: '/getAllMoutaiInfo',
    method: 'post',
    headers: {
      'Content-Type': 'multipart/form-data'
    },
    data
  })
}

// getMoutaiHistory
export function getMoutaiHistory(data) {
  return request({
    url: '/getMoutaiHistory',
    method: 'post',
    headers: {
      'Content-Type': 'multipart/form-data'
    },
    data
  })
}

