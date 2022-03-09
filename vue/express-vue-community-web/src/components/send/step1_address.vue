<template>
  <div class="client-box">
    <el-form class="form" label-width="80px">
      <el-form-item label="寄件人">
        <el-input
          v-model="sender"
          type="email"
          placeholder="请输入寄件人邮箱"
        ></el-input>
      </el-form-item>

      <el-form-item label="收件人">
        <el-input
          v-model="receiver"
          type="email"
          placeholder="请输入收件人邮箱"
        ></el-input>
      </el-form-item>
      <el-form-item label="备注">
        <el-input v-model="comment" type="textarea"></el-input>
      </el-form-item>
    </el-form>

    <div>
      <el-button class="btn-commit" type="primary" @click="handleClickCommit"
        >下一步</el-button
      >
    </div>
  </div>
</template>

<script>
export default {
  name: "SendStep1",

  data() {
    return {
      sender: "from@example.com",
      receiver: "to@example.com",
      comment: "msg...",
    };
  },

  methods: {
    commit(data) {
      let p = {
        method: "POST",
        url: "/api/v1/packages",
        data,
      };
      this.$store.dispatch("axios/execute", p).then(() => {
        this.nextstep();
      });
    },
    handleClickCommit() {
      let from_address = this.sender;
      let to_address = this.receiver;
      let comment = this.comment;
      let status = "";
      let transport = "";
      let data = {
        packages: [
          {
            from_address,
            to_address,
            comment,
            status,
            transport,
          },
        ],
      };
      this.commit(data);
    },

    nextstep() {
      this.$emit("nextstep");
    },
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

.client-box {
  margin-left: auto;
  margin-right: auto;
  max-width: 768px;
}
</style>
