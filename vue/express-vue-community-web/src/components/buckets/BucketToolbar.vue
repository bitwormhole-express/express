<template>
  <div class="toolbar">
    <el-button type="default" @click="handleClickRefresh">刷新</el-button>
    <el-button type="primary" @click="handleClickInsert">添加</el-button>

    <div><BucketEdit ref="theBucketEdit" @onok="handleBucketEditOK" /></div>
  </div>
</template>

<script>
import BucketEdit from "./BucketEdit.vue";

export default {
  name: "BucketToolbar",
  components: { BucketEdit },

  mounted() {
    this.fetch();
  },

  methods: {
    fetch() {
      this.$store.dispatch("bucket/fetch");
    },

    handleBucketEditOK(item) {
      this.insert(item);
    },

    handleClickInsert() {
      let item = {};
      this.$refs.theBucketEdit.displayItem(item, "POST");
    },

    handleClickRefresh() {
      this.fetch();
    },

    insert(item) {
      // let b = {
      //   driver: "baidu",
      //   properties: {
      //     a: "1",
      //     b: "2",
      //   },
      // };
      this.$store.dispatch("bucket/insert", item);
    },
  },
};
</script>

<style scoped>
.toolbar {
  text-align: right;
}
</style>
