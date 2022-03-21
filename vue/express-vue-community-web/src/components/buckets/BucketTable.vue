<template>
  <div>
    <el-table :data="tableData">
      <el-table-column label="名称" prop="name"></el-table-column>
      <el-table-column label="端点" prop="domain"></el-table-column>
      <el-table-column label="提供商" prop="driver"></el-table-column>
      <el-table-column label="操作" prop="opt">
        <template #default="scope">
          <el-button type="text" @click="handleClickItemDetail(scope.row)"
            >详情</el-button
          >
          <el-button type="text" @click="handleClickItemEdit(scope.row)"
            >修改</el-button
          >
          <el-button type="text" @click="handleClickItemDelete(scope.row)"
            >删除</el-button
          >
        </template>
      </el-table-column>
    </el-table>

    <BucketDetail ref="theBucketDetail" />
    <BucketEdit ref="theBucketEdit" />
  </div>
</template>

<script>
import { ElMessageBox } from "element-plus";
import BucketDetail from "./BucketDetail.vue";
import BucketEdit from "./BucketEdit.vue";

export default {
  name: "BucketTable",
  components: { BucketDetail, BucketEdit },

  computed: {
    tableData() {
      return this.$store.getters["bucket/items"];
    },
  },

  methods: {
    doDeleteItem(item) {
      this.$store.dispatch("bucket/delete", item);
    },

    handleClickItemDelete(item) {
      let name = item.name;
      let text = "真的要删除这个对【" + name + "】这个存储桶的引用吗？";
      ElMessageBox.confirm(text, "注意", {
        confirmButtonText: "确定",
        cancelButtonText: "取消",
        type: "warning",
      }).then(() => {
        this.doDeleteItem(item);
      });
    },

    handleClickItemDetail(item) {
      this.$refs.theBucketDetail.displayItem(item);
    },

    handleClickItemEdit(item) {
      this.$refs.theBucketEdit.displayItem(item);
    },
  },
};
</script>
