<script setup lang="ts">
defineProps<{ placeholder?: string }>();
</script>

<template>
  <input :placeholder="placeholder" />
</template>

<style scoped>
input {
  background-color: var(--input);
  padding: 12px 16px;
  border-radius: 6px;
  border: 1px solid var(--border);
  color: var(--label);
  font-size: 16px;
  font-weight: 400;
  line-height: 160%;
  transition: 0.2s;
}

input:focus {
  outline: none;
  box-shadow: 0px 0px 50px 5px var(--post);
}
</style>
