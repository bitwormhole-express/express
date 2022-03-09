<template>
  <!-- 选择运输方式 -->
  <div>
    <div>请选择要使用的运输渠道</div>

    <div class="selection">
      <el-select v-model="value" size="large" placeholder="请选择">
        <el-option
          v-for="item in options"
          :key="item.value"
          :label="item.label"
          :value="item.value"
        ></el-option>
      </el-select>
    </div>

    <div>
      <el-button type="primary" @click="nextstep">下一步</el-button>
    </div>
  </div>
</template>

<script>
export default {
  name: "SendStep2",
  data() {
    return {
      value: "",
      options: [
        { label: "本地直达", value: "localhost" },
        { label: "通过[阿里云]中转", value: "aliyun" },
        { label: "通过[腾讯云]中转", value: "qq" },
        { label: "通过[华为云]中转", value: "huawei" },
        { label: "通过[百度云]中转", value: "baidu" },
      ],
    };
  },

  methods: {
    nextstep() {
      this.$emit("nextstep");
    },
  },
};
</script>

<style scoped>
.selection {
  margin: 30px;
}
</style>