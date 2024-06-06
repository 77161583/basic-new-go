import Vue from 'vue'
const $ajaxJsonp = {
  // install (Vue) {
  //     Vue.prototype.$ajaxJsonp =  {
  get(apiUrl, params, successFun, jsonpCallback) {
    if (params == null) {
      const jp = {}
      jp.jsonp = '1'
      params = jp
    } else {
      params.jsonp = '1'
    }
    $.support.cors = true
    $.ajax({
      dataType: 'jsonp',
      jsonp: 'jsoncallback',
      jsonpCallback: jsonpCallback || 'success_jsonpCallback',
      type: 'get',
      data: params,
      async: false,
      url: apiUrl,
      success(data) {
        successFun(data)
      },
      error(XMLHttpRequest, textStatus, errorThrown) {
        console.log(XMLHttpRequest, errorThrown)
      }
    })
  }
  //     }

  // }
}
export default $ajaxJsonp
Vue.use($ajaxJsonp)
