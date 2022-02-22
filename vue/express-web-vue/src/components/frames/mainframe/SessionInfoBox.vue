<template>
  <div class="session-info-box">
    <div v-show="!session.authenticated">
      <el-button type="primary" plain @click="handleClickLogin">登录</el-button>
      <el-button type="primary" @click="handleClickSignUp">注册</el-button>
    </div>
    <div v-show="session.authenticated">
      <el-dropdown>
        <span>
          <el-avatar size="default" :src="avatar"></el-avatar>
        </span>
        <template #dropdown>
          <div class="account-info-panel">
            <div>{{ session.nickname }}</div>
            <div>{{ session.email }}</div>
            <div></div>
            <div>
              <el-button type="text" @click="handleClickLogout">退出</el-button>
            </div>
          </div>
        </template>
      </el-dropdown>
    </div>
    <div></div>
  </div>
</template>

<script>
export default {
  name: "SessionInfoBox",

  data() {
    return {};
  },

  computed: {
    avatar() {
      let url = this.session.avatar;
      if (url == null || url == "") {
        url =
          "https://cube.elemecdn.com/0/88/03b0d39583f48206768a7534e55bcpng.png";
      }
      return url;
    },

    session() {
      return this.$store.getters["session/info"];
    },
  },

  mounted() {
    this.fetch();
  },

  methods: {
    fetch() {
      this.$store.dispatch("session/fetch");
    },

    handleClickLogin() {
      this.$router.push("/login");
    },

    handleClickLogout() {
      this.$store.dispatch("session/logout");
    },

    handleClickSignUp() {
      this.$router.push("/sign-up");
    },
  },
};
</script>

<style scoped>
.session-info-box {
  display: flex;
}

.account-info-panel {
  padding: 20px;
}
</style>