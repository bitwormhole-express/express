<template>
  <div class="navbar">
    <router-link class="logo" to="/"
      ><h1 class="logo">某通快递</h1></router-link
    >

    <div class="menu-items">
      <router-link class="menu-item" to="/query"> 查询 </router-link>
      <router-link class="menu-item" to="/send"> 我要寄件 </router-link>
    </div>

    <div class="padding-space" />

    <div class="append-items">
      <sessionbox class="session-box" />
    </div>
  </div>
</template>

<script>
import sessionbox from "./session-info-box.vue";

export default {
  name: "NavBar",
  components: { sessionbox },
};
</script>

<style scoped>
.navbar {
  display: flex;
  /* background-color: gray; */
  align-items: center;
  color: gray;
}

.logo {
  padding: 0px;
  margin: 0px 20px 0px 20px;
  text-decoration-line: none;
}

.padding-space {
  flex-grow: 1;
}

.append-items {
  display: inline-block;
  margin: 0px 20px 0px 20px;
}

.menu-items {
  display: flex;
  flex-direction: row;
  align-content: center;
  align-items: center;
  text-align: center;
  justify-self: center;
}

.menu-item {
  text-decoration-line: none;
  padding: 20px;
  cursor: pointer;
  align-content: center;
  align-items: center;
  text-align: center;
  justify-self: center;

  border-bottom-width: 3px;
  border-bottom-color: #00000000;
  border-bottom-style: solid;
}

.menu-item:hover {
  cursor: pointer;
  color: white;
  background-color: rgba(211, 211, 211, 0.411);
  color: black;
  border-bottom-width: 3px;
  border-bottom-color: darkblue;
  border-bottom-style: solid;
}
</style>