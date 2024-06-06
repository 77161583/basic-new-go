// 全局方法
import Vue from 'vue'
const utils = {
  install(Vue) {
    Vue.prototype.$say = function() {
      console.log('I love you')
    }
    Vue.prototype.$she = 'lijianhua'
    Vue.prototype.$html = function(str) {
      str = str
        ? str.replace(/&amp;/g, function(m) {
          return {
            '&amp;': '&'
          }[m]
        })
        : ''
      return str
        ? str.replace(/&(gt|lt|quot|amp|#39|nbsp);/g, function(m) {
          return {
            '&lt;': '<',
            '&amp;': '&',
            '&quot;': '"',
            '&gt;': '>',
            '&#39;': "'",
            '&nbsp;': ' '
          }[m]
        })
        : ''
    }
    Vue.prototype.$truncateText = function(text, length) {
      // 定义 $utils.truncateText 方法
      if (text && text.length > length) {
        return text.slice(0, length) + '...'
      } else {
        return text
      }
    }
  }
}

Vue.use(utils)
