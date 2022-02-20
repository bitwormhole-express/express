<template>
  <MainFrame>
    <!-- <div><h1>Page: Query</h1></div> -->

    <el-form label-width="120px">
      <el-form-item label="我的邮箱地址">
        <el-input v-model="myEmailAddress" type="email"></el-input>
      </el-form-item>
      <el-form-item label="时间戳">
        <el-date-picker
          v-model="dateRange"
          type="daterange"
          range-separator="到"
          start-placeholder="选择最早时间"
          end-placeholder="选择最晚时间"
        ></el-date-picker>
      </el-form-item>
      <el-form-item label="类型">
        <el-radio-group v-model="direction">
          <el-radio-button label="我寄的" />
          <el-radio-button label="我收的" />
          <el-radio-button label="全部" />
        </el-radio-group>
      </el-form-item>
    </el-form>

    <div><el-button type="primary">查询</el-button></div>

    <el-table>
      <el-table-column label="运单号"></el-table-column>
      <el-table-column label="发件人"></el-table-column>
      <el-table-column label="收件人"></el-table-column>
      <el-table-column label="时间戳"></el-table-column>
      <el-table-column label="状态"></el-table-column>
    </el-table>
  </MainFrame>
</template>

<script>
export default {
  name: "Query",

  data() {
    return {
      myEmailAddress: "",
      dateRange: ["", ""],
      direction: "",
    };
  },
};
</script>
