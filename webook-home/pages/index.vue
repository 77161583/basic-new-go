<template>
    <div class="form-container">
        <Form ref="formItem" :rules="ruleValidate" :model="formItem" :label-width="80" @submit.prevent="onSubmit">
            <FormItem label="邮箱">
                <Input v-model="formItem.email" placeholder="请输入邮箱" />
            </FormItem>
            <FormItem label="密码">
                <Input type="password" v-model="formItem.password" placeholder="请输入密码" />
            </FormItem>
            <FormItem label="确认密码">
                <Input type="password" v-model="formItem.confirmPassword" placeholder="请确认密码" />
            </FormItem>
            <FormItem>
                <Button type="primary" @click="handleSubmit('formItem')">提交</Button>
                <Button @click="handleReset('formItem')" style="margin-left: 8px">重置</Button>
            </FormItem>
        </Form>
    </div>
</template>

<script>
export default {
  data() {
    return {
      formItem: {
        email: '',
        password: '',
        confirmPassword: ''
      },
      ruleValidate: {
          name: [
              { required: true, message: 'The name cannot be empty', trigger: 'blur' }
          ],
          password: [
              { required: true, message: 'The name cannot be empty', trigger: 'blur' }
          ],
          confirmPassword: [
              { required: true, message: 'The name cannot be empty', trigger: 'blur' }
          ],
      }
    };
  },
  methods: {
    handleSubmit (name) {
        this.$refs[name].validate((valid) => {
            if (valid) {
                this.$Message.success('Success!');
            } else {
                this.$Message.error('Fail!');
            }
        })
    },

    handleReset (name) {
        this.$refs[name].resetFields();
    },

    onSubmit() {
      // 验证邮箱格式
      const emailRegex = /^[a-z0-9_.+-]+@[a-z0-9-]+\.[a-z0-9-.]+$/;
      if (!emailRegex.test(this.formItem.email)) {
        this.$message.error('你的邮箱格式不正确');
        return;
      }

      // 验证密码长度和内容
      const passwordRegex = /^(?=.*[A-Z])(?=.*[a-z])(?=.*[0-9])(?=.*[!@#$%^&*()_+=-]).{8,}$/;
      if (!passwordRegex.test(this.formItem.password)) {
        this.$message.error('密码必须至少8个字符，并包含数字、特殊字符、大小写字母');
        return;
      }

      // 确认密码匹配
      if (this.formItem.password !== this.formItem.confirmPassword) {
        this.$message.error('两次密码输入不一致');
        return;
      }

      // 提交表单
      this.$message.success('注册成功');
      console.log('表单提交:', this.formItem);
    }
  }
};
</script>

<style scoped>
.form-container {
    width: 100%;
    max-width: 600px; /* 限制宽度为400像素，可以根据需要调整 */
    margin: 0 auto; /* 水平居中 */
    margin-top: 100px;
    padding: 20px; /* 内边距，保证一定的视觉空间 */
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1); /* 轻微的阴影效果 */
    background-color: #fff; /* 背景颜色 */
}
</style>
