<template>
  <div>
    <el-menu :default-active="currentPagePath" router>
      <el-menu-item index="/settings">首页</el-menu-item>
      <el-menu-item index="/settings/account">我的信息</el-menu-item>
      <el-menu-item index="/settings/account-security">账号安全</el-menu-item>
      <el-menu-item index="/settings/buckets">存储桶管理</el-menu-item>
    </el-menu>
  </div>
</template>

<script>
export default {
  name: "SettingsNavMenu",

  computed: {
    currentPagePath() {
      let path = this.$route.path;
      return path;
    },
  },
};
</script>
