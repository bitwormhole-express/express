<template>
  <div>
    <navbar />
    <slot />
  </div>
</template>

<script>
import navbar from "./navbar.vue";

export default {
  name: "MainFrameIndex",
  components: { navbar },
};
</script>