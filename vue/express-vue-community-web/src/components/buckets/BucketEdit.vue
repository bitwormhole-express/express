<template>
  <div>
    <el-drawer v-model="drawer" :title="title" :with-header="false">
      <span> 存储桶：{{ item.name }} </span>

      <el-form label-width="100px">
        <el-form-item label="id">
          <el-input v-model="item.id"></el-input>
        </el-form-item>
        <el-form-item label="domain">
          <el-input v-model="item.domain"></el-input>
        </el-form-item>
        <el-form-item label="driver">
          <!-- <el-input v-model="item.driver"></el-input> -->
          <BucketDriverSelector />
        </el-form-item>
        <el-form-item label="name">
          <el-input v-model="item.name"></el-input>
        </el-form-item>

        <el-form-item
          v-for="p in extendsProperties"
          :key="p.key"
          :label="p.label"
        >
          <el-input v-model="p.value"></el-input>
        </el-form-item>
      </el-form>
    </el-drawer>
  </div>
</template>

<script>
import BucketDriverSelector from "./BucketDriverSelector.vue";

export default {
  name: "BucketEdit",

  components: { BucketDriverSelector },

  data() {
    return {
      title: "编辑存储桶配置",
      drawer: false,
      item: {
        properties: { k: "v" },
      },
      driver: "",
      extendsProperties: [
        { key: "ex1", label: "ex1", value: "1" },
        { key: "ex2", label: "ex2", value: "2" },
      ],
    };
  },

  methods: {
    displayItem(item, method) {
      this.item = item;
      this.drawer = true;
      this.method = method;
    },
  },
};
</script>
