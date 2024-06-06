import env from './.env.js'
export default {
  // Global page headers: https://go.nuxtjs.dev/config-head
  head: {
    titleTemplate: 'Golang学习页面',
    htmlAttrs: {
      lang: 'zh-CN'
    },
    meta: [
      { charset: 'utf-8' },
      { name: 'viewport', content: 'width=device-width, initial-scale=1' },
      { name: 'keywords', content: 'Golang学习页面' },
      { hid: 'description', name: 'description', content: '“Golang学习页面' },
      { name: 'format-detection', content: 'telephone=no' },
    ],
    link: [
      { rel: 'icon', type: 'image/x-icon', href: '/favicon.ico' }
    ],
    script: [
      { src: '/js/jquery-3.3.1.min.js' },
      { src: '/js/jquery.SuperSlide.2.1.3.js' },
      { src: '/js/ckplayer/Ckplayer.js' },
    ]
  },

  // 全局CSS
  css: [
    '@/assets/css/global.scss',
  ],
  styleResources: {
    scss: './assets/css/variables.scss'
  },

  // Plugins to run before rendering page: https://go.nuxtjs.dev/config-plugins
  plugins: [
    { src: '@/plugins/ajaxJsonp' },
    { src: '@/plugins/baidu' },
    { src: '@/plugins/api' },
    { src: '@/plugins/iview' },
    { src: '@/plugins/utils' },
    { src: '@/directive/slider', ssr: false },
    { src: '@/plugins/element-ui', ssr: true }
  ],

  // 自动导入组件: https://go.nuxtjs.dev/config-components
  components: true,

  // Modules for dev and build (recommended): https://go.nuxtjs.dev/config-modules
  buildModules: [
    // https://go.nuxtjs.dev/typescript
    '@nuxt/typescript-build',
    '@nuxtjs/style-resources'
  ],

  // Modules: https://go.nuxtjs.dev/config-modules
  modules: [
    // https://go.nuxtjs.dev/axios
    '@nuxtjs/axios',
    'cookie-universal-nuxt'
  ],

  // Axios module configuration: https://go.nuxtjs.dev/config-axios
  axios: {
    // Workaround to avoid enforcing hard-coded localhost:3000: https://github.com/nuxt-community/axios-module/issues/308
    // baseURL: env[process.env.NODE_ENV].BASE_API,
    proxy: true, // 表示开启代理
    // prefix: '', // 表示给请求url加个前缀 /api
    // credentials: true,
    // withCredentials: true
  },

  proxy: {
    '/gp61': {
      target: env[process.env.NODE_ENV].BASE_API, // 目标接口域名
      changeOrigin: true, // 表示是否跨域
      withCredentials: true,
      pathRewrite: {
        '^/gp61': '' // 把 /api 替换成 ''
      }
    },
    '/web': {
      target: env[process.env.NODE_ENV].WEB_API, // 目标接口域名
      changeOrigin: true, // 表示是否跨域
      withCredentials: true,
      pathRewrite: {
        '^/web': '' // 把 /api 替换成 ''
      }
    }
  },

  // 定义`构建时`所需的环境变量
  env: {
    baseUrl: env[process.env.NODE_ENV].BASE_API,
    ...env[process.env.NODE_ENV],
  },
  // 默认的`.nuxt`大概率会被当成隐藏目录
  buildDir: 'nuxt-dist',

  router: {
    // 覆写 <nuxt-link> 默认的ExactActiveClass
    linkExactActiveClass: 'current'
  },

  // 配置进度条
  loading: {
    color: '#4285f4'
  },

  server: {
    port: env[process.env.NODE_ENV].NODE_PORT,
    host: '0.0.0.0'
  }
}
