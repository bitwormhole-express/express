<template>
  <div>
    <div><h1>寄件</h1></div>

    <el-steps :active="step.current" align-center>
      <el-step title="地址" description="填写收件人地址"></el-step>
      <el-step title="运输方式" description="选择运输方式"></el-step>
      <el-step title="发件" description="发送邮件"></el-step>
      <el-step title="完成" description="成功"></el-step>
    </el-steps>

    <div class="steps">
      <div v-show="step.current == 0"><Step1 @nextstep="handleNextStep" /></div>
      <div v-show="step.current == 1"><Step2 @nextstep="handleNextStep" /></div>
      <div v-show="step.current == 2"><Step3 @nextstep="handleNextStep" /></div>
      <div v-show="step.current == 3"><Step4 @nextstep="handleNextStep" /></div>
    </div>
  </div>
</template>

<script>
import Step1 from "./step1_address.vue";
import Step2 from "./step2_transport.vue";
import Step3 from "./step3_upload.vue";
import Step4 from "./step4_done.vue";

export default {
  name: "SendIndex",
  components: { Step1, Step2, Step3, Step4 },

  data() {
    return {
      step: {
        current: 0,
      },
    };
  },

  computed: {
    pStep() {
      let step = this.$route.query["step"];
      if (step == null) {
        step = "0";
      }
      return Number(step);
    },
  },

  mounted() {
    this.init();
  },

  methods: {
    handleNextStep() {
      this.step.current++;
    },

    init() {
      this.step.current = this.pStep;
    },
  },
};
</script>

<style scoped>
.steps {
  margin-top: 30px;
}
</style>