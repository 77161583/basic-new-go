import { Plugin } from '@nuxt/types'
import { NuxtAxiosInstance } from '@nuxtjs/axios'

import indexModule from '@/api/index'
import courseModule from '@/api/course'

const apiMap = ($axios: NuxtAxiosInstance) => ({
  index: indexModule($axios),
  course: courseModule($axios),
})

type ApiType = ReturnType<typeof apiMap>

declare module 'vue/types/vue' {
    interface Vue {
        $api: ApiType
    }
}

declare module '@nuxt/types' {
    interface Context {
        $api: ApiType
    }
}

declare module 'vuex/types/index' {
    interface Store<S> {
        $api: ApiType
    }
}

// https://axios.nuxtjs.org/extend
// https://typescript.nuxtjs.org/zh-hant/cookbook/plugins/
const axiosPlugin: Plugin = ({ $axios, store, isDev, redirect, error: nuxtError }, inject) => {
  $axios.onRequest((config) => {
    const token = store.state.user.tokenVal
    if (token) {
      config.headers.hrttoken = token
    }
    if (isDev) {
      console.log(`[${process.client ? 'client' : 'server'} side] ${config.method} ${config.url}`)
    }
  })

  $axios.onError((error: any) => {
    const statusCode = parseInt(error.response && error.response.status)
    // 用户token失效时，主动清除掉本地相关信息
    if (statusCode === 401) {
      store.dispatch('user/logout')
      if (process.client) {
        window.location.reload()
      }
    }
    // 开发模式下直接打印异常原因便于调试，否则转到异常页面
    if (isDev) {
      console.log(error)
    } else {
      nuxtError({
        statusCode: error.response.status,
        message: error.message,
      })
    }
  })

  // Inject to context as $api
  inject('api', apiMap($axios))
}

export default axiosPlugin
