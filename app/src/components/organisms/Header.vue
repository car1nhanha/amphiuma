<script setup lang="ts">
const projectName = import.meta.env.VITE_PROJECT_NAME;
</script>

<template>
  <header>
    <div class="content">
      <img src="/terminal-solid.svg" class="terminal-solid" />
      <h1>{{ projectName }}</h1>
    </div>

    <div class="low-light" />
    <div class="left-light" />
    <div class="right-light" />
  </header>
</template>

<style scoped>
header {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: center;
  height: 296px;
  width: 100%;
  background-color: var(--profile);
  position: relative;
  overflow: hidden;

  background-image: url("/header-effect-left.svg"), url("/header-effect-right.svg");
  background-position: left center, right center;
  background-repeat: no-repeat, no-repeat;
}

.content {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 20px;
}

img.terminal-solid {
  height: 40px;
  width: 45px;
}

h1 {
  font-size: 24;
  font-weight: 400;
  line-height: 160%;
  color: var(--blue);
  font-family: "Coda";
}

div.low-light {
  position: absolute;
  height: 0;
  width: 891px;
  margin: 0 auto;
  bottom: 9px;

  box-shadow: 0px -20px 200px 62px rgba(20, 88, 156, 0.2);
}

div.right-light {
  position: absolute;
  height: 0px;
  width: 0px;
  border-radius: 50%;
  top: 0px;
  right: 0;
  background-color: transparent;

  box-shadow: -50px 50px 200px 200px rgba(20, 88, 156, 0.2);
}

div.left-light {
  position: absolute;
  height: 0px;
  width: 0px;
  border-radius: 50%;
  top: 0;
  left: 0;
  background-color: transparent;

  box-shadow: -50px 50px 200px 200px rgba(20, 88, 156, 0.2);
}
</style>
