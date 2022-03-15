<template>
  <MainFrame>
    <h1><SettingsIndexView /></h1>
  </MainFrame>
</template>

<script>
import SettingsIndexView from "./../components/settings/index.vue";

export default {
  name: "Settings",
  components: { SettingsIndexView },
};
</script>

<style scoped>
.banner {
  height: 60vw;
  background-image: url("/image/kuaidi-xiaoge.jpg");
  background-size: 100vw;
  background-position-y: center;
}
</style>
