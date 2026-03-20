<script lang="ts" setup>
import { codeToHtml } from "shiki";
import { onMounted, ref } from "vue";

const props = defineProps({
  code: {
    type: String,
    required: true,
  },
  lang: {
    type: String,
    default: "javascript",
  },
  title: {
    type: String,
    default: "window",
  },
});

const highlightedCode = ref("");

onMounted(async () => {
  highlightedCode.value = await codeToHtml(props.code, {
    lang: props.lang,
    theme: "dracula",
    rootStyle: "line-height: 1.5rem; font-size: 14px;",
  });
});
</script>

<template>
  <div class="window">
    <div class="winhead">
      <div class="button">
        <div class="close"></div>
        <div class="min"></div>
        <div class="max"></div>
      </div>
      <div class="title">{{ props.title }}</div>
    </div>
    <div class="code-container" v-html="highlightedCode"></div>
  </div>
</template>

<style scoped>
.window {
  margin: 10px auto;
  width: 90%;
  display: flex;
  flex-direction: column;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.4);
  border-radius: 10px;
  overflow: hidden;
}
.winhead {
  height: 35px;
  background: #111316;
  border-bottom: 0.5px solid rgba(162, 163, 164, 0.4);
  border-top-left-radius: 5px;
  border-top-right-radius: 5px;
  display: flex;
  align-items: center;
  position: relative;
}
.winhead .button {
  display: flex;
  align-items: center;
  z-index: 9;
}
.button div {
  width: 12px;
  height: 12px;
  margin-left: 6px;
  background: rgb(199, 199, 199);
  border: 0.5px solid rgb(169, 169, 169);
  border-radius: 50%;
}
.button .close {
  background: rgb(231, 69, 70);
  border: 0.5px solid rgb(208, 67, 67);
}
.button .min {
  background: rgb(246, 180, 39);
  border: 0.5px solid rgb(212, 154, 58);
}
.button .max {
  background: rgb(86, 202, 53);
}
.winhead .title {
  position: absolute;
  width: 100%;
  text-align: center;
}

.code-container {
  overflow: hidden;
  font-size: 14px;
  padding: 16px;
  background-color: #282a36;
}
</style>
