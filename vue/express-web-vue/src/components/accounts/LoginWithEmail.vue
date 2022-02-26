<template>
  <div>
    <el-form label-width="120px">
      <el-form-item label="我的邮箱地址"
        ><el-input v-model="myEmailAddress" type="email"></el-input
      ></el-form-item>
      <el-form-item label="验证码"
        ><VeriCodeInput ref="inputVeriCode" :email="myEmailAddress"
      /></el-form-item>
    </el-form>

    <div>
      <el-button type="success" @click="handleClickLogin">登录</el-button>
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
