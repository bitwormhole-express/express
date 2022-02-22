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
      <el-button type="success" @click="handleClickLogin">登录</el-button>
    </div>
  </div>
</template>

<script>
export default {
  name: "LoginWithPassword",

  data() {
    return {
      myUserID: "123",
      myPassword: "123",
    };
  },

  methods: {
    handleClickLogin() {
      let account = this.myUserID;
      let secret = this.myPassword;
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
      this.$store.dispatch("session/login", p);
    },
  },
};
</script>
