<script setup lang="ts">
import { ref } from "vue";

const user = ref("");
const projectName = import.meta.env.VITE_PROJECT_NAME;

const goToProfile = () => {
  if (user.value.trim() !== "") {
    window.location.href = `/${user.value}`;
  }
};
</script>

<template>
  <header>
    <h1>{{ projectName }}</h1>
    <p class="subtitle">Descubra e compartilhe</p>
    <p class="description">Transforme suas publicações do GitHub em conteúdo acessível, de forma simples e rápida.</p>
    <div class="find-container-header">
      <input placeholder="Digite o usuário" v-model="user" />
      <button @click="goToProfile">Explorar</button>
    </div>
  </header>
</template>

<style scoped>
header {
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 36px;
  height: 100vh;
  background-color: var(--profile);
}

h1 {
  background: var(--border);
  mix-blend-mode: normal;
  border-radius: 8px;
  width: min-content;
  padding: 10px;

  font-weight: 400;
  font-size: 20px;
  line-height: 24px;

  color: #ffffff;
  box-shadow: 0 2px 28px 0 rgba(0, 0, 0, 0.2);
}

.subtitle {
  font-weight: 400;
  font-size: 62px;
  line-height: 62px;
  color: #ffffff;
}

.description {
  font-weight: 300;
  font-size: 16px;
  line-height: 24px;

  display: flex;
  align-items: center;

  color: #ffffff;
}

/* input */
.find-container-header {
  display: flex;
  flex-direction: row;
  gap: 10px;
}

input {
  all: unset;
  width: 427px;
  padding: 16px;
  background: var(--input);
  border: 1px solid var(--border);
  border-radius: 17px;
  font-weight: 400;
  font-size: 18px;
  line-height: 27px;
  color: rgba(255, 255, 255, 0.5);
}

button {
  all: unset;
  background: var(--blue);
  border-radius: 17px;
  padding: 16px;
  font-style: normal;
  font-weight: 600;
  font-size: 17px;
  line-height: 27px;
  color: var(--title);
}
</style>
