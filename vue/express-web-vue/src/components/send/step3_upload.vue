<template>
  <!-- 发送(上传)文件 -->

  <div>
    <el-upload
      class="upload-demo"
      drag
      action="https://jsonplaceholder.typicode.com/posts/"
      multiple
    >
      <el-icon class="el-icon--upload"><upload-filled /></el-icon>
      <div class="el-upload__text">把文件拖放到这里，或 <em>点击上传</em></div>
      <template #tip>
        <div class="el-upload__tip">各种类型的文件，小于 999MB</div>
        <div class="el-upload__tip">
          大于999MB? 请使用<el-button type="text">客户端</el-button>发送。
        </div>
      </template>
    </el-upload>

    <div class="manifest">
      <h3>发件清单</h3>
      <el-table :data="items">
        <el-table-column prop="filename" label="文件名"></el-table-column>
        <el-table-column prop="size" label="大小"></el-table-column>
        <el-table-column prop="contenttype" label="类型"></el-table-column>
        <el-table-column prop="opt" label="操作">
          <template #default>
            <el-button type="text">设置</el-button>
            <el-button type="text">移除</el-button>
          </template>
        </el-table-column>
      </el-table>
    </div>

    <div>
      <el-checkbox> 向收件人发送一封email，提醒他（她）收件。 </el-checkbox>
    </div>

    <el-button type="primary" @click="handleClickCommit">提交</el-button>
  </div>
</template>

<script>
import { UploadFilled } from "@element-plus/icons-vue";

export default {
  name: "SendStep3",
  components: { UploadFilled },
  data() {
    return {
      items: [
        {
          filename: "demo1.txt",
          contenttype: "text/plain",
          size: 1024,
        },
      ],
    };
  },
  methods: {
    handleClickCommit() {
      this.nextstep();
    },
    nextstep() {
      this.$emit("nextstep");
    },
  },
};
</script>

<style scoped>
.manifest {
  text-align: left;
  padding: 20px;
}
</style>
