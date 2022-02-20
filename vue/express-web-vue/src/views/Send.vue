<template>
  <MainFrame>
    <!-- <div><h1>Page: Send</h1></div> -->

    <el-form class="form" label-width="80px">
      <el-form-item label="收件人">
        <el-input v-model="receiver" type="email"></el-input>
      </el-form-item>
      <el-form-item label="寄件人">
        <el-input v-model="sender" type="email"></el-input>
      </el-form-item>
      <el-form-item label="备注">
        <el-input v-model="comment" type="textarea"></el-input>
      </el-form-item>
    </el-form>

    <div><el-button class="btn-commit" type="success">提交</el-button></div>
  </MainFrame>
</template>

<script>
export default {
  name: "Send",

  data() {
    return {
      sender: "",
      receiver: "",
      comment: "",
    };
  },
};
</script>

<style scoped>
.form {
  padding-bottom: 50px;
}

.btn-commit {
  padding: 20px 50px 20px 50px;
}
</style>