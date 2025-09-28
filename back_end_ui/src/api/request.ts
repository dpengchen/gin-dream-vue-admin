import axios from 'axios'
import type { CreateAxiosDefaults } from 'axios'
import { message } from 'ant-design-vue'
import router from '@/router'

const config = {
  baseURL: '/api',
  timeout: 10000,
} as CreateAxiosDefaults

const instance = axios.create(config)

//请求拦截
instance.interceptors.request.use((config: any) => {
  const token = localStorage.getItem('token')
  if (token) {
    config.headers['token'] = token
  }
  return config
})

//响应拦截
instance.interceptors.response.use(
  (resp: any) => {
    const data = resp.data
    if (data.code !== 200) {
      //需要重新登录
      if (data.code === 401) {
        localStorage.clear()
        message.error(data.msg)
        router.push('/login')
        return Promise.reject(data)
      }

      //其他错误
      message.error(data.msg)
      return Promise.reject(data)
    }

    return resp.data
  },
  (error: any) => {
    message.error('系统错误！' + error.message)
    return Promise.reject(error)
  },
)

export default instance
