<template>
  <MainFrame>
    <!-- <div><h1>Page: Receive</h1></div> -->

    <el-form class="form" label-width="100px">
      <el-form-item label="运单号"
        ><el-input v-model="id" type="text"></el-input
      ></el-form-item>
      <el-form-item label="收件人"
        ><el-input v-model="receiver" type="text"></el-input
      ></el-form-item>
      <el-form-item label="寄件人"
        ><el-input v-model="sender" type="text"></el-input
      ></el-form-item>
      <el-form-item label="时间戳"
        ><el-input v-model="timestamp" type="text"></el-input
      ></el-form-item>
      <el-form-item label="备注"
        ><el-input v-model="comment" type="textarea"></el-input
      ></el-form-item>
      <el-form-item label="快件详情"
        ><el-input v-model="detail" type="textarea"></el-input
      ></el-form-item>
    </el-form>

    <div><el-button class="btn-commit" type="success">取件</el-button></div>
  </MainFrame>
</template>

<script>
export default {
  name: "Receive",

  data() {
    return {
      id: "",
      receiver: "",
      sender: "",
      timestamp: "",
      comment: "",
      detail: "",
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
