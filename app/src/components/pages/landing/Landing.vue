<script setup lang="ts">
import { ref } from "vue";
import CardLanding from "../../molecules/Card-landing.vue";

const user = ref("");
const projectName = import.meta.env.VITE_PROJECT_NAME;

const goToProfile = () => {
  if (user.value.trim() !== "") {
    window.location.href = `/${user.value}`;
  }
};
</script>

<template>
  <header class="header">
    <h1 class="h1">{{ projectName }}</h1>
    <p class="subtitle">Discover and share</p>
    <p class="description">Turn your GitHub posts into accessible content in a simple and fast way.</p>
    <div class="find-container-header">
      <input class="input" placeholder="Enter username" v-model="user" />
      <button @click="goToProfile">Explore</button>
    </div>
  </header>

  <div class="body">
    <div class="body-container">
      <h2 class="body-title">How does it work?</h2>
      <CardLanding />
    </div>
  </div>
</template>

<style src="./landing.css"></style>
