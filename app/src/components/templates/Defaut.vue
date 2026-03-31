<script setup lang="ts">
import Footer from "../organisms/Footer.vue";
import HeaderOrganism from "../organisms/Header.vue";
</script>

<template>
  <div class="content">
    <HeaderOrganism />
    <div class="content-body">
      <slot></slot>
    </div>
    <!-- <div class="bottom-space"></div> -->
  </div>
  <Footer
    :inspiredBy="{
      name: 'GitHub Blog',
      url: 'https://www.figma.com/community/file/1138814951106121051',
      figma: 'https://www.figma.com/@millenamartins',
      user_figma: 'Millena Martins',
    }"
  />
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
