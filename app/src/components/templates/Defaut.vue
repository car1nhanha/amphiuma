<script setup lang="ts">
import HeaderOrganism from "../organisms/Header.vue";
</script>

<template>
  <div class="content">
    <HeaderOrganism />
    <div class="content-body">
      <slot></slot>
    </div>
    <div class="bottom-space"></div>
  </div>
</template>

<style scoped>
div.content {
  height: 100%;
  max-width: 100%;
  background-color: var(--background);
  display: flex;
  flex-direction: column;
  align-items: center;
}

div.content-body {
  display: flex;
  flex-direction: column;
  width: 864px;
  margin-top: -88px;
  z-index: 1;
}
/*  */
.bottom-space {
  height: 200px;
}

@media screen and (max-width: 900px) {
  div.content-body {
    width: calc(96%);
    margin: -88px auto 0 auto;
  }
}
</style>
