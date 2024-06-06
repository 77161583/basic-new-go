import { NuxtAxiosInstance } from '@nuxtjs/axios'
import qs from 'qs'

export default ($axios: NuxtAxiosInstance) => ({
  // 列表
  getCourseList(data: {}) {
    return $axios.$post('/gp61/gp6/lms/stu/trainplanCourseHandle/ableSelect_course', data)
  },

  // 获取专业分类
  getCategory(data: {}) {
    return $axios.$post('/gp61/gp6/lms/stu/trainplanCourseHandle/trainplanCourseTypes', data)
  },

  // 专业课，公需课详情
  getClassDetail(data:{}) {
    const params = qs.stringify(data)
    return $axios.$get('/gp61/gp6/lms/stu/course/courseDetail?' + params)
  },

  // 文章详情
  getCourseDetail(data: {}) {
    const params = qs.stringify(data)
    return $axios.$get('/gp61/gp6/resources/stu/pharmacistContent/detail?' + params)
  },

})
