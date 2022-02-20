<template>
  <div class="session-info-box">
    <el-button type="primary" plain @click="handleClickLogin">登录</el-button>
    <el-button type="primary" @click="handleClickSignUp">注册</el-button>
  </div>
</template>

<script>
export default {
  name: "SessionInfoBox",

  methods: {
    handleClickLogin() {
      this.$router.push("/login");
    },

    handleClickSignUp() {
      this.$router.push("/sign-up");
    },
  },
};
</script>

<style scoped>
.session-info-box {
  display: flex;
}
</style>