<template>
  <SettingsFrame>
    <BucketsIndexView />
  </SettingsFrame>
</template>

<script>
import SettingsFrame from "./Frame.vue";
import BucketsIndexView from "./../buckets/index.vue";

export default {
  name: "SettingsBuckets",
  components: { BucketsIndexView, SettingsFrame },
};
</script>
