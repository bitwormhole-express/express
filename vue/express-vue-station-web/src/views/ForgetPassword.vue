<template>
  <!-- 忘记密码，重设 -->
  <MainFrame>
    <h1 class="title">重设密码</h1>

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
              v-model="vericode"
              :email="myEmailAddress"
              ref="inputVeriCode"
          /></el-form-item>
        </el-form>
        <div>
          <el-button class="btn-all" type="primary" @click="handleClickVerify"
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
          <el-button class="btn-all" type="primary" @click="handleClickCommit"
            >提交</el-button
          >
        </div>
      </div>

      <div v-show="step.current == step.finish">
        <h2>设置成功！</h2>
        <el-button class="btn-all" type="success" @click="handleClickNavToLogin"
          >去登录</el-button
        >
      </div>
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
        setpass: 2,
        finish: 3,
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
      this.$refs.inputVeriCode.verify().then(() => {
        this.nextStep();
      });
    },

    nextStep() {
      this.step.current++;
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
          // this.showAlert("成功");
          this.nextStep();
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

<style scoped>
.client-box {
  margin-left: auto;
  margin-right: auto;
  margin-top: 50px;
  box-shadow: 0px 0px 10px gray;
  max-width: 400px;
  padding: 10px;
  border-radius: 5px;
}

.btn-all {
  width: 100%;
  margin-top: 50px;
}

.title {
  font-size: 30px;
  color: gray;
}
</style>