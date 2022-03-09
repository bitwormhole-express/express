<template>
  <div>
    <el-form>
      <el-form-item
        ><el-input
          v-model="myEmailAddress"
          type="email"
          placeholder="我的邮箱地址"
        ></el-input
      ></el-form-item>
      <el-form-item
        ><VeriCodeInput ref="inputVeriCode" :email="myEmailAddress"
      /></el-form-item>
    </el-form>

    <div>
      <el-button class="btn-login" type="success" @click="handleClickLogin"
        >登录</el-button
      >
    </div>
  </div>
</template>

<script>
import PasswordUtil from "./../../utils/password.js";

export default {
  name: "LoginWithEmail",

  data() {
    return {
      myEmailAddress: "",
    };
  },

  methods: {
    handleClickLogin() {
      this.$refs.inputVeriCode.verify().then(() => {
        this.login();
      });
    },

    login() {
      let email = this.myEmailAddress;
      let verification = this.$refs.inputVeriCode.getVerificationDTO();
      let secret = this.makeSecret(verification);
      let data = {
        auth: {
          account: email,
          mechanism: "email",
          secret,
        },
      };
      let p = {
        method: "POST",
        url: "/api/v1/auth/",
        data,
      };
      this.$store.dispatch("axios/execute", p).then(() => {
        this.navToHome();
      });
    },

    makeSecret(verification) {
      let json = JSON.stringify(verification);
      return PasswordUtil.encodePassword(json);
    },

    navToHome() {
      this.$router.push("/");
    },
  },
};
</script>

<style scoped>
.btn-login {
  width: 100%;
  margin-top: 50px;
}
</style>