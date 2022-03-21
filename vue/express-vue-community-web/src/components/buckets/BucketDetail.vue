<template>
  <div>
    <el-drawer v-model="drawer" :title="title" :with-header="false" size="95%">
      <div class="title-bar">
        <h3>存储桶属性</h3>
        <div class="gap" />
        <el-button type="default" @click="drawer = false">关闭</el-button>
      </div>

      <el-form label-width="200px">
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

        <el-form-item v-for="p in properties" :label="p.name" :key="p.name">
          <span class="value-box"> {{ p.value }} </span>
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

      properties: [{ name: "", value: "" }],
    };
  },

  computed: {
    // properties() {
    // },
  },

  methods: {
    displayItem(item) {
      this.item = item;
      this.drawer = true;
      this.updateProperties(item.properties);
    },

    updateProperties(src) {
      if (src == null) {
        return;
      }
      let dst = [];
      for (var k in src) {
        let v = src[k];
        dst.push({ name: k, value: v });
      }
      this.properties = dst;
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
