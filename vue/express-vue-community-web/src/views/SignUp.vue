<template>
  <!-- 注册账号 -->
  <MainFrame>
    <h1 class="title">创建新账号</h1>

    <div class="client-box">
      <div v-show="step.current == step.verify">
        <el-form>
          <el-form-item
            ><el-input
              v-model="myEmailAddress"
              type="email"
              placeholder="我的邮箱地址"
            ></el-input
          ></el-form-item>
          <el-form-item
            ><VeriCodeInput
              ref="inputVeriCode"
              v-model="vericode"
              :email="myEmailAddress"
          /></el-form-item>
        </el-form>
        <div>
          <el-button
            class="btn-verify"
            type="primary"
            @click="handleClickVerify"
            >验证</el-button
          >
        </div>
      </div>

      <div v-show="step.current == step.setpass">
        <el-form>
          <el-form-item
            ><el-input
              v-model="password1"
              type="password"
              placeholder="设置密码"
          /></el-form-item>
          <el-form-item
            ><el-input
              v-model="password2"
              type="password"
              placeholder="确认密码"
          /></el-form-item>
        </el-form>
        <div>
          <el-button
            class="btn-commit"
            type="primary"
            @click="handleClickCommit"
            >提交</el-button
          >
          <!-- <el-button type="default">跳过</el-button> -->
        </div>
      </div>

      <div v-show="step.current == step.finish">
        <h2>注册成功！</h2>
        <el-button
          class="btn-nav-to-login"
          type="success"
          @click="handleClickFinish"
        >
          去登录
        </el-button>
      </div>
    </div>

    <div class="other">
      <router-link to="/login">使用现有的账号登录</router-link>
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

<style scoped>
.client-box {
  margin-left: auto;
  margin-right: auto;
  box-shadow: 0px 0px 10px gray;
  max-width: 400px;
  padding: 10px;
  border-radius: 5px;
  margin-top: 50px;
}

.btn-verify {
  margin-top: 50px;
  width: 100%;
}

.title {
  font-size: 30px;
  color: gray;
}

.btn-commit {
  width: 100%;
  margin-top: 50px;
}

.btn-nav-to-login {
  width: 100%;
  margin-top: 50px;
}

.other {
  margin: 50px;
}
</style>