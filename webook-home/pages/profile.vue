<template>
  <h1>123</h1>
</template>
<script>
export default {
  data() {
    return {

    }
  },
  mounted() {
    this.getData()
  },
  methods: {
    async getData() {
      const params = {
        id: 1
      }
      try {
        const res = await this.$api.course.getClassDetail(params)
        console.log(res)
      } catch (error) {
        // 处理请求错误
        console.error('An error occurred during signup:', error)
      }
    }
  }
}
</script>
