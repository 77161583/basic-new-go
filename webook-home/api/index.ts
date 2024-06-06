import { NuxtAxiosInstance } from '@nuxtjs/axios'
export default ($axios: NuxtAxiosInstance) => ({

  // 资讯相关接口
  getNewsList(data: {}) {
    return $axios.$post('/gp61/gp6/resources/stu/pharmacistContent/list', data)
  },

  // 登录
  getLogin(data: {}) {
    return $axios.$post('/web/loginValid?t=' + Date.now(), data)
    // const params = qs.stringify(data)
    // return $axios.$post('/web/loginValid/third_portal_login?' + params)
  }

})
