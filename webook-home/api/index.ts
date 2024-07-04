import { NuxtAxiosInstance } from '@nuxtjs/axios'
export default ($axios: NuxtAxiosInstance) => ({

  // 资讯相关接口
  getSignup(data: {}) {
    return $axios.$post('/go/users/signup', data)
  },

  // 登录
  getLogin(data: {}) {
    return $axios.$post('/go/users/login', data)
  }

})
