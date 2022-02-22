<template>
  <MainFrame>
    <el-tabs type="border-card">
      <el-tab-pane label="通过邮件验证登录"><LoginWithEmail /></el-tab-pane>
      <el-tab-pane label="通过密码登录"><LoginWithPassword /></el-tab-pane>
    </el-tabs>
  </MainFrame>
</template>

<script>
import LoginWithEmail from "./../components/accounts/LoginWithEmail.vue";
import LoginWithPassword from "./../components/accounts/LoginWithPassword.vue";

export default {
  name: "Login",
  components: { LoginWithEmail, LoginWithPassword },
};
</script>
