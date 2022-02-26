<template>
  <!-- 忘记密码，重设 -->
  <MainFrame>
    <h1>重设密码</h1>

    <div v-show="step.current == step.verify">
      <el-form label-width="120px">
        <el-form-item label="我的邮箱地址"
          ><el-input v-model="myEmailAddress" type="email"></el-input
        ></el-form-item>
        <el-form-item label="验证码"
          ><VeriCodeInput
            v-model="vericode"
            :email="myEmailAddress"
            ref="inputVeriCode"
        /></el-form-item>
      </el-form>
      <div>
        <el-button type="primary" @click="handleClickVerify">验证</el-button>
      </div>
    </div>

    <div v-show="step.current == step.setpass">
      <el-form label-width="120px">
        <el-form-item label="设置密码"
          ><el-input v-model="password1" type="password"
        /></el-form-item>
        <el-form-item label="确认密码"
          ><el-input v-model="password2" type="password"
        /></el-form-item>
      </el-form>
      <div>
        <el-button type="primary" @click="handleClickCommit">提交</el-button>
      </div>
    </div>

    <div v-show="step.current == step.finish">
      <h2>设置成功！</h2>
      <el-button type="success" @click="handleClickNavToLogin"
        >去登录</el-button
      >
    </div>
  </MainFrame>
</template>

<script>
import PasswordUtil from "@/utils/password.js";

export default {
  name: "ForgetPassword",

  data() {
    return {
      step: {
        current: 1,
        verify: 1,
        setpass: 1,
        finish: 1,
      },
      myEmailAddress: "",
      password1: "",
      password2: "",
      vericode: "",
      verification: {},
    };
  },

  methods: {
    getVerification() {
      return this.$refs.inputVeriCode.getVerificationDTO();
    },

    handleClickCommit() {
      let pass1 = this.password1;
      let pass2 = this.password2;
      if (pass1 != pass2) {
        return this.showAlert("两个密码不一致");
      }
      if (!this.isPasswordRegular(pass1)) {
        return this.showAlert("密码不合格");
      }
      let verification = this.getVerification();
      let newpassword = PasswordUtil.encodePassword(pass1);
      let data = {
        password: {
          newpassword,
          verification,
        },
      };
      this.resetPassword(data);
    },

    isPasswordRegular(passwd) {
      if (passwd == null) {
        return false;
      }
      let len = ("" + passwd).length;
      return len >= 8;
    },

    handleClickNavToLogin() {
      this.$router.push("/login");
    },

    handleClickVerify() {
      // todo ...
    },

    resetPassword(data) {
      let p = {
        method: "PUT",
        url: "/api/v1/password/",
        data,
      };
      this.$store
        .dispatch("axios/execute", p)
        .then(() => {
          this.showAlert("成功");
        })
        .catch(() => {
          this.showAlert("失败");
        });
    },

    showAlert(msg) {
      this.$alert(msg);
    },
  },
};
</script>
