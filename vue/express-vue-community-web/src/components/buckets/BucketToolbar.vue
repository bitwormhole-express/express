<template>
  <div class="toolbar">
    <el-button type="default" @click="handleClickRefresh">刷新</el-button>
    <el-button type="primary">添加</el-button>
  </div>
</template>

<script>
export default {
  name: "BucketToolbar",
  components: {},

  mounted() {
    this.fetch();
  },

  methods: {
    fetch() {
      this.$store.dispatch("bucket/fetch");
    },

    handleClickRefresh() {
      this.fetch();
    },
  },
};
</script>

<style scoped>
.toolbar {
  text-align: right;
}
</style>
