<template>
  <!-- 验证码组件 -->

  <div class="vcode">
    <el-input v-model="vcodeText"></el-input>
    <el-button class="btn-fetch-code">{{ fetchCodeButtonText }}</el-button>
  </div>
</template>

<script>
export default {
  name: "verification-code",
  data() {
    return {
      fetchCodeButtonText: "获取验证码",
      vcodeText: "",
    };
  },
};
</script>

<style scoped>
.vcode {
  display: flex;
}

.btn-fetch-code {
  margin-left: 10px;
}
</style>