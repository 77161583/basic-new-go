<template>
  <div class="form-container">
    <el-form
      ref="ruleForm"
      style="max-width: 600px"
      :model="ruleForm"
      :rules="rules"
      label-width="auto"
      class="demo-ruleForm"
      status-icon
    >
      <el-form-item label="邮箱" prop="email">
        <el-input v-model="ruleForm.email" type="email" />
      </el-form-item>
      <el-form-item label="密码" prop="password">
        <el-input v-model="ruleForm.password" type="password" />
      </el-form-item>
      <el-form-item label="重复密码" prop="confirmPassword">
        <el-input v-model="ruleForm.confirmPassword" type="password" />
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="submitForm('ruleForm')">注册</el-button>
        <el-button @click="resetForm('ruleForm')">重置</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>
<script>
export default {
  data() {
    return {
      ruleForm: {
        email: '123@qq.com',
        password: 'Qwe!12345',
        confirmPassword: 'Qwe!12345',
      },
      rules: {
        email: [
          { required: true, message: '请输入邮箱', trigger: 'blur' },
          { type: 'email', message: '请输入有效的邮箱地址', trigger: 'blur' },
        ],
        password: [
          { required: true, message: '请输入密码', trigger: 'blur' },
          { min: 8, message: '密码长度不能少于 8 位', trigger: 'blur' },
          { pattern: /^(?=.*[A-Z])(?=.*[a-z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$/, message: '密码必须以大写字母开头，包含字母、数字和特殊字符', trigger: 'blur' },
        ],
        confirmPassword: [
          { required: true, message: '请输入重复密码', trigger: 'blur' },
          { min: 8, message: '密码长度不能少于 8 位', trigger: 'blur' },
          { pattern: /^(?=.*[A-Z])(?=.*[a-z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$/, message: '密码必须以大写字母开头，包含字母、数字和特殊字符', trigger: 'blur' },
          {
            validator: (_rule, value, callback) => {
              if (value !== this.ruleForm.password) {
                callback(new Error('两次输入的密码不一致'))
              } else {
                callback()
              }
            },
            trigger: 'blur'
          },
        ],
      }
    }
  },
  methods: {
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          // alert('submit!')
          this.getSubmit()
        } else {
          console.log('error submit!!')
          return false
        }
      })
    },
    resetForm(formName) {
      this.$refs[formName].resetFields()
    },
    async getSubmit() {
      const params = {
        email: this.ruleForm.email,
        password: this.ruleForm.password,
        confirmPassword: this.ruleForm.confirmPassword
      }
      try {
        const res = await this.$api.index.getSignup(params)
        if (res.code === 200) {
          this.$message({
            message: res.message,
            type: 'success'
          })
          setTimeout(() => {
            window.location.href = '/login'
          }, 500)
        } else {
          this.$message.error(res)
        }
      } catch (error) {
        // 处理请求错误
        console.error('An error occurred during signup:', error)
      }
    }
  }
}
</script>
<style>
/* 定义一个表单容器 */
.form-container {
  width: 100%;  /* 设置宽度为100%，以适应父容器 */
  max-width: 600px;  /* 限制最大宽度为600像素，可以根据需要调整 */
  margin: 0 auto;  /* 水平居中 */
  margin-top: 100px;  /* 顶部边距为100像素 */
  padding: 20px;  /* 内边距为20像素，保证一定的视觉空间 */
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);  /* 轻微的阴影效果 */
  background-color: #fff;  /* 背景颜色为白色 */
}
</style>
