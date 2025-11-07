<script setup lang="ts">
import { Icon } from "@iconify/vue";
import { formatDistanceToNow } from "date-fns";
import { ref } from "vue";
import CardHeader from "../molecules/Card-header.vue";
import type { Post } from "../molecules/Card-posts.vue";
import Stylize from "../organisms/Stylize-post.vue";
import TemplateDefault from "../templates/Defaut.vue";

const user = {
  slug: "car1nhanha",
};

const file = ref("");

fetch("/test.md")
  .then((response) => response.text())
  .then((text) => {
    file.value = text;
  })
  .catch(() => alert("deu erro"));

const post: Post = {
  id: 1,
  title: "Explorando Vue 3",
  description: file.value,
  date: "2025-10-29",
  file_origin: "explorando_vue_3.md",
};
</script>

<template>
  <TemplateDefault>
    <CardHeader>
      <div class="card-content">
        <div class="line-actions">
          <a href="/"><Icon icon="solar:alt-arrow-left-line-duotone" /> return</a>
          <a :href="`github.com/car1nhanha/repo/${post.file_origin}`"
            >read on github <Icon icon="solar:square-arrow-right-up-broken"
          /></a>
        </div>

        <h1>{{ post.title }}</h1>

        <ul>
          <li><Icon icon="octicon:mark-github-24" /> {{ user.slug }}</li>
          <li><Icon icon="solar:calendar-mark-bold" /> {{ formatDistanceToNow(new Date(post.date)) }}</li>
        </ul>
      </div>
    </CardHeader>

    <Stylize :postFile="file" />
  </TemplateDefault>
</template>

<style scoped>
div.card-content {
  display: flex;
  flex-direction: column;
  flex: 1;
}

div.line-actions {
  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;
  margin-bottom: 20px;
}

div.line-actions a {
  display: flex;
  flex-direction: row;
  align-items: center;
  gap: 8px;
  color: var(--blue);
  text-decoration: none;
  font-size: 12px;
  font-weight: bold;
  line-height: 160%;
  text-transform: uppercase;
}

h1 {
  color: var(--title);
  line-height: 160%;
  font-size: 24px;
  line-height: 130%;
  font-weight: bold;
  margin-bottom: 8px;
}

ul {
  list-style: none;
  display: flex;
  flex-direction: row;
  align-items: center;
  gap: 24px;
}
li {
  color: var(--span);
  font-weight: 400;
  font-size: 16px;
  line-height: 160%;

  display: flex;
  flex-direction: row;
  align-items: center;
  gap: 8px;
}
</style>
