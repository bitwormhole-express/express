<template>
  <MainFrame>
    <!-- <div><h1>Page: Login</h1></div> -->

    <el-form label-width="120px">
      <el-form-item label="我的邮箱地址"
        ><el-input v-model="myEmailAddress" type='email'></el-input
      ></el-form-item>
      <el-form-item label="验证码"><VeriCodeInput /></el-form-item>
    </el-form>

    <div>
      <el-button type="success">登录</el-button>
    </div>
  </MainFrame>
</template>

<script>
export default {
  name: "Login",

  data() {
    return {
      myEmailAddress: "",
    };
  },
};
</script>
