<template>
  <div>
    <el-drawer v-model="drawer" :title="title" :with-header="false">
      <div class="title-bar">
        <h3>存储桶属性</h3>
        <div class="gap" />
        <el-button type="default" @click="drawer = false">关闭</el-button>
      </div>

      <el-form label-width="100px">
        <el-form-item label="ID">
          <span class="value-box"> {{ item.id }} </span>
        </el-form-item>
        <el-form-item label="名称">
          <span class="value-box"> {{ item.name }} </span>
        </el-form-item>
        <el-form-item label="域名">
          <span class="value-box"> {{ item.domain }} </span>
        </el-form-item>
        <el-form-item label="提供商">
          <span class="value-box"> {{ item.driver }} </span>
        </el-form-item>

        <el-form-item v-for="(k, v) in properties" :label="k" :key="k">
          <span class="value-box"> {{ v }} </span>
        </el-form-item>
      </el-form>
    </el-drawer>
  </div>
</template>

<script>
export default {
  name: "BucketDetail",

  data() {
    return {
      title: "编辑存储桶配置",
      drawer: false,
      item: {},
    };
  },

  computed: {
    properties() {
      let p = this.item.properties;
      if (p == null) {
        p = {};
      }
      return p;
    },
  },

  methods: {
    displayItem(item) {
      this.item = item;
      this.drawer = true;
    },
  },
};
</script>

<style scoped>
.title-bar {
  display: flex;
  flex-direction: row;
}
.gap {
  flex-grow: 1;
}
.value-box {
  background-color: lightgray;
  width: 100%;
  text-align: left;
  padding-left: 10px;
  padding-right: 10px;
}
</style>
