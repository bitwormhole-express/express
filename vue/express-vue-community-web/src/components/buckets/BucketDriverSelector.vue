<template>
  <div>
    <el-select>
      <el-option v-for="d in drivers" :key="d.key"></el-option>
    </el-select>
  </div>
</template>

<script>
export default {
  name: "BucketDriverSelector",

  computed: {
    drivers() {
      return this.$store.getters["bucket/drivers"];
    },
  },
};
</script>
