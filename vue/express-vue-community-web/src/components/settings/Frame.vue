<template>
  <!-- 这是所有设置页面的通用框架 -->

  <div>
    <el-container>
      <el-aside class='left-box' >
        <SettingsNavMenu class="settings-nav-menu" />
      </el-aside>
      <el-main>
        <slot />
      </el-main>
    </el-container>
  </div>
</template>

<script>
import SettingsNavMenu from "./NavMenu.vue";

export default {
  name: "SettingsFrame",
  components: { SettingsNavMenu },
};
</script>

<style scoped>
.settings-nav-menu {
  margin-top: 50px;
  padding-top: 50px;
}

.left-box {
  width: 180px;
}

</style>
