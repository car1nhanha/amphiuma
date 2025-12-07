<script setup lang="ts">
const props = defineProps<{ placeholder?: string; modelValue?: string }>();
const emit = defineEmits<{ (e: "update:modelValue", value: string): void; (e: "change", value: string): void }>();

const onInput = (event: Event) => {
  const value = (event.target as HTMLInputElement).value;
  emit("update:modelValue", value);
};

const onChange = (event: Event) => {
  const value = (event.target as HTMLInputElement).value;
  emit("change", value);
};
</script>

<template>
  <input :placeholder="props.placeholder" :value="props.modelValue" @input="onInput" @change="onChange" />
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
