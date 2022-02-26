<template>
  <div>
    <el-form>
      <el-form-item
        ><el-input
          v-model="myUserID"
          type="email"
          placeholder="用户名或邮箱"
        ></el-input
      ></el-form-item>
      <el-form-item
        ><el-input
          v-model="myPassword"
          type="password"
          placeholder="用户的密码"
      /></el-form-item>
    </el-form>

    <div>
      <div><router-link to="/forget-password"> 忘记密码？ </router-link></div>
    </div>

    <div>
      <el-button
        class="btn-login"
        type="success"
        @click="handleClickLogin"
        :disabled="working"
      >
        <el-icon class="is-loading" v-show="working">
          <Loading />
        </el-icon>
        登录</el-button
      >
    </div>
  </div>
</template>

<script>
import PasswordUtil from "@/utils/password.js";

import { Loading } from "@element-plus/icons-vue";
import { ElMessageBox, ElMessage } from "element-plus";

function encodePassword(password) {
  return PasswordUtil.encodePassword(password);
}

export default {
  name: "LoginWithPassword",

  components: { Loading },

  data() {
    return {
      myUserID: "",
      myPassword: "",
      working: false,
    };
  },

  methods: {
    handleClickLogin() {
      let account = this.myUserID;
      let secret = encodePassword(this.myPassword);
      let data = {
        auth: {
          account,
          secret,
          mechanism: "password",
        },
      };
      let p = {
        data,
      };
      this.working = true;
      this.$store
        .dispatch("session/login", p)
        .then(() => {
          this.working = false;
          this.showLoginOkMessage();
          this.navToHome();
        })
        .catch(() => {
          this.working = false;
          this.showLoginErrorMessage();
        });
    },

    navToHome() {
      this.$router.push("/");
    },

    showLoginErrorMessage() {
      ElMessageBox.alert("可能是用户名或者密码错误", "登录失败", {
        confirmButtonText: "确定",
        type: "error",
      });
    },

    showLoginOkMessage() {
      ElMessage({
        message: "登录成功",
        type: "success",
      });
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