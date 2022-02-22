<template>
  <MainFrame>
    <div class="login-box">
      <el-tabs type="border-card">
        <el-tab-pane label="通过邮件验证登录"><LoginWithEmail /></el-tab-pane>
        <el-tab-pane label="通过密码登录"><LoginWithPassword /></el-tab-pane>
      </el-tabs>
    </div>
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

<style scoped>
.login-box {
  margin-top: 100px;
  margin-left: auto;
  margin-right: auto;
  width: 100vw;
  max-width: 380px;
}
</style>
