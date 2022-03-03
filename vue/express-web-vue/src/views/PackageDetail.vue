<template>
  <MainFrame>
    <div><DetailView /></div>
  </MainFrame>
</template>

<script>
import DetailView from "./../components/packages/PackageDetail.vue";

export default {
  name: "PackageDetail",
  components: { DetailView },
};
</script>
