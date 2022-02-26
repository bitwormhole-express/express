<template>
  <!-- 注册账号 -->
  <MainFrame>
    <h1>创建新账号</h1>
    <div v-show="step.current == step.verify">
      <el-form label-width="120px">
        <el-form-item label="我的邮箱地址"
          ><el-input v-model="myEmailAddress" type="email"></el-input
        ></el-form-item>
        <el-form-item label="验证码"
          ><VeriCodeInput
            ref="inputVeriCode"
            v-model="vericode"
            :email="myEmailAddress"
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
        <!-- <el-button type="default">跳过</el-button> -->
      </div>
    </div>

    <div v-show="step.current == step.finish">
      <h2>注册成功！</h2>
      <el-button type="success" @click="handleClickFinish"> 去登录 </el-button>
    </div>
  </MainFrame>
</template>

<script>
import PasswordUtil from "./../utils/password.js";

export default {
  name: "SignUp",

  data() {
    return {
      step: {
        current: 1,
        verify: 1,
        setpass: 2,
        finish: 3,
      },
      myEmailAddress: "",
      password1: "",
      password2: "",
    };
  },

  methods: {
    commitAndCreatePassword(password) {
      let newpassword = PasswordUtil.encodePassword(password);
      let verification = this.getVerification();
      let data = {
        password: {
          verification,
          newpassword,
        },
      };
      let p = {
        method: "POST",
        url: "/api/v1/password/",
        data,
      };
      this.$store
        .dispatch("axios/execute", p)
        .then(() => {
          this.showAlert("成功");
          this.nextStep();
        })
        .catch(() => {
          this.showAlert("失败");
        });
    },

    getVerification() {
      return this.$refs.inputVeriCode.getVerificationDTO();
    },

    handleClickCommit() {
      let pass1 = this.password1;
      let pass2 = this.password2;
      if (!this.isPasswordRegular(pass1)) {
        return this.showAlert("密码不合格！");
      }
      if (pass1 != pass2) {
        return this.showAlert("两个密码不一致。");
      }
      this.commitAndCreatePassword(pass1);
    },

    handleClickFinish() {
      this.$router.push("/login");
    },

    handleClickVerify() {
      this.$refs.inputVeriCode.verify().then(() => {
        this.nextStep();
      });
    },

    isPasswordRegular(password) {
      if (password == null) {
        return false;
      }
      let text = "" + password;
      return text.length >= 8;
    },

    nextStep() {
      this.step.current++;
    },

    showAlert(msg) {
      this.$alert(msg);
    },
  },
};
</script>
