<template>
  <div>
    <el-select :modelValue="modelValue" @update:modelValue="fireInput">
      <el-option
        v-for="d in drivers"
        :key="d.driver"
        :label="d.label"
        :value="d.driver"
      ></el-option>
    </el-select>
  </div>
</template>

<script>
export default {
  name: "BucketDriverSelector",

  props: {
    modelValue: String,
  },

  data() {
    return {};
  },

  watch: {
    sel() {},
  },

  computed: {
    drivers() {
      return this.$store.getters["bucket/drivers"];
    },
  },

  mounted() {
    this.init();
  },

  methods: {
    init() {
      this.$store.dispatch("bucket/loadDrivers");
    },

    fireInput(value) {
      this.$emit("update:modelValue", value);
    },
  },
};
</script>
