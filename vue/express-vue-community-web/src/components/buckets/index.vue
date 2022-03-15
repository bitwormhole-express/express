<template>
  <div>
    <h2 class="title">存储桶管理</h2>
    <BucketToolbar />
    <BucketTable />
  </div>
</template>

<script>
import BucketTable from "./BucketTable.vue";
import BucketToolbar from "./BucketToolbar.vue";

export default {
  name: "BucketsIndexView",
  components: { BucketTable, BucketToolbar },
};
</script>

<style scoped>
.title {
  text-align: left;
}
</style>
