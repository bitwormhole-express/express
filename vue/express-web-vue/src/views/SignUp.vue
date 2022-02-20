<template>
  <MainFrame>
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
        <el-button type="success">注册</el-button>
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
        <el-button type="success">确定</el-button>
        <el-button type="default">跳过</el-button>
      </div>
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
        setpass: 2,
      },
      myEmailAddress: "",
      password1: "",
      password2: "",
      vericode: "",
    };
  },
};
</script>
