<template>
  <div>
    <el-drawer v-model="drawer" :title="title" :with-header="false" size="90%">
      <span> {{ title }}：{{ item.name }} </span>

      <div class="toolbar">
        <el-button type="success" @click="handleClickSave">保存</el-button>
        <el-button type="default" @click="handleClickCancel">取消</el-button>
        <el-button type="default" @click="handleClickTest">测试</el-button>
      </div>

      <el-form label-width="100px">

        <!-- <el-form-item label="id">
          <el-input v-model="item.id"></el-input>
        </el-form-item> -->

        <!-- <el-form-item label="domain">
          <el-input v-model="item.domain"></el-input>
        </el-form-item> -->

        <el-form-item label="方案">
          <!-- <el-input v-model="item.driver"></el-input> -->
          <BucketDriverSelector v-model="driver" />
        </el-form-item>

        <!-- <el-form-item label="name">
          <el-input v-model="item.name"></el-input>
        </el-form-item> -->

        <el-form-item
          v-for="p in extendsProperties"
          :key="p.name"
          :label="p.label"
        >
          <el-input v-model="p.value" :readonly="p.c"></el-input>
        </el-form-item>
      </el-form>
    </el-drawer>
  </div>
</template>

<script>
import BucketDrivers from "./../../utils/bucket-drivers.js";
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
      bucketID: 0,
      driver: "aliyun",
      extendsProperties: [
        { name: "ex1", label: "ex1", value: "1" },
        { name: "ex2", label: "ex2", value: "2", description: "" },
      ],
    };
  },

  watch: {
    driver() {
      this.updateExtendsProperties();
    },
  },

  methods: {
    displayItem(item, method) {
      this.item = item;
      this.drawer = true;
      this.bucketID = item.id;
      this.method = method;
      if (method == "POST") {
        this.title = "添加存储桶";
      } else {
        this.title = "修改存储桶配置";
      }
      this.driver = item.driver;
      this.updateExtendsProperties();
    },

    fireCancel() {
      this.$emit("oncancel");
    },

    fireOK() {
      this.$emit("onok");
    },

    handleClickCancel() {
      this.fireCancel();
      this.drawer = false;
    },

    handleClickSave() {
      // this.fireOK();

      let b = this.getBucketDTO();
      let id = b.id;
      let action = "bucket/update";
      if (id == null || id == 0) {
        action = "bucket/insert";
      }
      this.$store
        .dispatch(action, b)
        .then(() => {})
        .catch(() => {});
    },

    handleClickTest() {
      let b = this.getBucketDTO();
      this.$store
        .dispatch("bucket/test", b)
        .then(() => {})
        .catch(() => {});
    },

    getBucketDTO() {
      let properties = this.getBucketProperties();
      let driver = this.driver;
      let id = this.bucketID;
      let b = {
        id,
        name: "",
        domain: "",
        driver,
        properties,
      };
      return b;
    },

    getBucketProperties() {
      return BucketDrivers.propertiesListToTable(this.extendsProperties);
    },

    updateExtendsProperties() {
      let item = this.item;
      if (item == null) {
        return;
      }
      let table = item.properties;
      let p = BucketDrivers.propertiesTableToList(table, this.driver);
      this.extendsProperties = p;
    },
  },
};
</script>

<style scoped>
.toolbar {
  text-align: left;
  margin-bottom: 30px;
}
</style>
