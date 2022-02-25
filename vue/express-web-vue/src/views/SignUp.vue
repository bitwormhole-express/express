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
          ><VeriCodeInput v-model="vericode"
        /></el-form-item>
      </el-form>
      <div>
        <el-button type="primary">验证</el-button>
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
        <el-button type="primary" @click="handleClickCommit" >提交</el-button>
        <!-- <el-button type="default">跳过</el-button> -->
      </div>
    </div>

    <div v-show="step.current == step.finish">
      <h2>注册成功！</h2>
      <el-button type="success"> 去登录 </el-button>
    </div>
  </MainFrame>
</template>

<script>
export default {
  name: "SignUp",

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
    };
  },
};
</script>
