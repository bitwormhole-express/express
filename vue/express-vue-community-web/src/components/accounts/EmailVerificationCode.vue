<template>
  <!-- 验证码组件 -->

  <div class="vcode">
    <el-input
      class="input-vcode"
      v-model="vcodeText"
      placeholder="邮件中的验证码"
    ></el-input>
    <el-button
      class="btn-fetch-code"
      @click="handleClickFetchVeriCode"
      :disabled="fetchCodeButtonBusy"
      >{{ fetchCodeButtonText }}</el-button
    >

    <!-- <div>{{ email }}</div> -->
  </div>
</template>

<script>
function RetryTimer(ctx) {
  this.timeout = 0;
  this.context = ctx;
}

RetryTimer.prototype = {
  startWithTimeout(tout) {
    this.timeout = tout;
    this.onStep();
  },

  onStep() {
    let sec = this.timeout;
    if (sec <= 0) {
      this.context.reset();
      return;
    }
    let msg = sec + "秒后重试";
    this.timeout = sec - 1;
    this.context.setFetchButtonText(msg);
    this.doStep();
  },

  doStep() {
    setTimeout(() => {
      this.onStep();
    }, 1000);
  },
};

export default {
  name: "verification-code",

  props: {
    email: String,
  },

  data() {
    return {
      fetchCodeButtonText: "获取验证码(x)",
      fetchCodeButtonBusy: false,
      vcodeText: "",
      vcodeDTO: {},
    };
  },

  mounted() {
    this.reset();
  },

  methods: {
    checkEmailAddress(addr) {
      let text = "" + addr;
      let i = text.indexOf("@");
      let len = text.length;
      return 0 < i && i + 1 < len;
    },

    fetchVeriCode(email) {
      let url = "/api/v1/email-verification/";
      let data = {
        verification: {
          email,
        },
      };
      let p = {
        method: "POST",
        url,
        data,
      };
      let timer = new RetryTimer(this);
      this.fetchCodeButtonText = "正在获取...";
      this.fetchCodeButtonBusy = true;
      this.$store
        .dispatch("axios/execute", p)
        .then((resp) => {
          this.showAlert("验证码已发送至邮箱：" + email);
          timer.startWithTimeout(60);
          this.vcodeDTO = resp.data.verification;
        })
        .catch(() => {
          this.showAlert("操作失败！");
          timer.startWithTimeout(10);
        });
    },

    getVerificationDTO() {
      let dto = this.vcodeDTO;
      let vcode = this.vcodeText;
      let j = JSON.stringify(dto);
      dto = JSON.parse(j);
      dto.code = vcode;
      return dto;
    },

    handleClickFetchVeriCode() {
      let email = this.email;
      if (!this.checkEmailAddress(email)) {
        this.showAlert("请输入有效的Email地址！");
        return;
      }
      this.fetchVeriCode(email);
    },

    reset() {
      this.fetchCodeButtonBusy = false;
      this.fetchCodeButtonText = "获取验证码";
    },

    setFetchButtonText(text) {
      this.fetchCodeButtonText = text;
    },

    showAlert(msg) {
      this.$alert(msg);
    },

    verify() {
      let verification = this.getVerificationDTO();
      verification.code = this.vcodeText;
      let data = {
        verification,
      };
      let p = {
        method: "PUT",
        url: "/api/v1/email-verification/",
        data,
      };
      let result = this.$store.dispatch("axios/execute", p).catch((err) => {
        this.showAlert("验证失败！");
        return Promise.reject(err);
      });
      return result;
    },
  },
};
</script>

<style scoped>
.vcode {
  display: flex;
  width: 100%;
  flex-direction: row;
}

.btn-fetch-code {
  margin-left: 10px;
  min-width: 110px;
}

.input-vcode {
  flex-grow: 1;
  flex-basis: 0px;
}
</style>