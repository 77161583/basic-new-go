import { NuxtAxiosInstance } from '@nuxtjs/axios'
export default ($axios: NuxtAxiosInstance) => ({

  // 注册
  getSignup(data: {}) {
    return $axios.$post('/go/users/signup', data)
  },

  // 登录
  getLogin(data: {}) {
    return $axios.$post('/go/users/login', data)
  }

})
