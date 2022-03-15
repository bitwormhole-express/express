<template>
  <div>
    <el-table :data="tableData">
      <el-table-column label="名称" prop="name"></el-table-column>
      <el-table-column label="域名" prop="dn"></el-table-column>
      <el-table-column label="操作" prop="opt">
        <template #default>
          <el-button type="text">详情</el-button>
          <el-button type="text">修改</el-button>
          <el-button type="text">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
export default {
  name: "BucketTable",

  computed: {
    tableData() {
      return this.$store.getters["bucket/items"];
    },
  },
};
</script>
