<template>
  <!-- 完成 -->
  <div>
    <h1>邮件发送成功！</h1>
    <el-button type="success" @click="handleClickDetail"
      >查看邮件详情</el-button
    >
  </div>
</template>

<script>
export default {
  name: "SendStep4",

  methods: {
    handleClickDetail() {
      let id = "todo_pack_id";
      this.$router.push("/packages/" + id);
    },
  },
};
</script>
