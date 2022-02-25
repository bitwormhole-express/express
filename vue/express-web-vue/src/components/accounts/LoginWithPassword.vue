<template>
  <div>
    <el-form label-width="120px">
      <el-form-item label="用户名"
        ><el-input
          v-model="myUserID"
          type="email"
          placeholder="请在此输入用户名、邮箱地址、或电话号码等身份标识"
        ></el-input
      ></el-form-item>
      <el-form-item label="密码"
        ><el-input v-model="myPassword" type="password"
      /></el-form-item>
    </el-form>

    <div>
      <el-button type="success" @click="handleClickLogin" :disabled="working">
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
      myUserID: "123",
      myPassword: "123",
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
        })
        .catch(() => {
          this.working = false;
          this.showLoginErrorMessage();
        });
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
