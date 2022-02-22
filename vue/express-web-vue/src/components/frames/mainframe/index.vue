<template>
  <div>
    <navbar />
    <slot />
  </div>
</template>

<script>
import navbar from "./NavBar.vue";

export default {
  name: "MainFrameIndex",
  components: { navbar },
};
</script>