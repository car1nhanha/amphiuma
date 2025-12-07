<script setup lang="ts"></script>

<template>
  <section>
    <slot />
  </section>
</template>

<style scoped>
section {
  display: flex;
  align-items: center;

  padding: 32px;
  border-radius: 10px;

  background-color: var(--profile);
  box-shadow: 0 2px 28px 0 rgba(0, 0, 0, 0.2);
}
@media screen and (max-width: 900px) {
  section {
    max-width: 100%;
  }
}
</style>
