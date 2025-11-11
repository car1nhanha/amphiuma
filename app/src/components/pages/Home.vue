<script setup lang="ts">
import { Icon } from "@iconify/vue";
import { ref } from "vue";
import { useRoute } from "vue-router";
import InputText from "../atoms/input-text.vue";
import CardHeader from "../molecules/Card-header.vue";
import CardPost, { type IApiResponse } from "../molecules/Card-posts.vue";
import TemplateDefault from "../templates/Defaut.vue";

const route = useRoute();
const userParams = route.params.user;

const apiListPosts = ref({} as IApiResponse);

fetch(`http://localhost:8080/v1/${userParams}`)
  .then((response) => response.json())
  .then((response: IApiResponse) => {
    apiListPosts.value = response;
  });

const user = {
  image: "https://avatars.githubusercontent.com/u/34972401",
  name: "Lucas Carinhanha",
  description:
    "Crux sacra sit mihi lux. Non draco sit mihi dux. Vade retro satana! Nunquam suade mihi vana. Sunt mala quae libas. Ipse venena bibas.",
  slug: "car1nhanha",
  // company: "nenhuma",
  followers: 4,
};

// const publications = {
//   quantity: 6,
// };

// const posts: Post[] = [
//   {
//     id: 1,
//     title: "Explorando Vue 3",
//     description:
//       "Programming languages all have built-in data structures, but these often differ from one language to another. This article attempts to list the built-in data structures available in JavaScript and what properties they have. These can be used to build other data structures. Wherever possible, comparisons with other languages are drawn. Dynamic typing JavaScript is a loosely typed and dynamic language. Variables in JavaScript are not directly associated with any particular value type, and any variable can be assigned (and re-assigned) values of all types:",
//     date: "2025-10-29",
//   },
//   {
//     id: 2,
//     title: "Melhores práticas com TypeScript",
//     description: "Aprenda como usar TypeScript em projetos Vue para obter segurança e produtividade.",
//     date: "2025-10-15",
//   },
//   {
//     id: 3,
//     title: "Otimizando performance no front-end",
//     description: "Dicas e técnicas para deixar sua aplicação web mais leve e rápida.",
//     date: "2025-09-30",
//   },
// ];
</script>

<template>
  <TemplateDefault>
    <CardHeader>
      <div class="card-content">
        <img :src="user.image" alt="User avatar" />
        <div class="card-body">
          <div class="name-line">
            <h2>{{ user.name }}</h2>
            <a :href="`https://github.com/${user.slug}`">GITHUB <Icon icon="solar:square-arrow-right-up-broken" /></a>
          </div>
          <p>{{ user.description }}</p>
          <ul>
            <li><Icon icon="octicon:mark-github-24" /> {{ user.slug }}</li>
            <li><Icon icon="flowbite:building-solid" /> {{ apiListPosts.Total }} posts</li>
            <li><Icon icon="ic:round-star" /> {{ user.followers }} followers</li>
          </ul>
        </div>
      </div>
    </CardHeader>

    <div class="find-container">
      <div class="find-container-header">
        <h3>Publicações</h3>
        <span>{{ apiListPosts.Total }} publicações</span>
      </div>
      <InputText placeholder="buscar conteúdo" />
    </div>

    <div class="posts">
      <CardPost v-for="(value, index) in apiListPosts.Items" :post="value" :key="index" />
    </div>
  </TemplateDefault>
</template>

<style scoped>
div.card-content {
  display: flex;
  flex-direction: row;
  gap: 32px;
}

img {
  height: 148px;
  width: 148px;
  border-radius: 8px;
  object-fit: cover;
}

div.card-body {
  display: flex;
  flex-direction: column;
  justify-content: center;
}

div.name-line {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  margin-bottom: 8px;
}
h2 {
  color: var(--title);
  font-weight: bold;
  font-size: 24px;
  line-height: 130%;
}
div.name-line a {
  text-decoration: none;
  color: var(--blue);
  font-weight: bold;
  font-size: 12px;
  line-height: 160%;

  display: flex;
  align-items: center;
  gap: 8px;
}

p {
  color: var(--text);
  font-size: 16px;
  font-weight: 400;
  line-height: 160%;
  margin-bottom: 24px;
  text-align: justify;
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

/*  */

div.find-container {
  margin-top: 72px;
  display: flex;
  flex-direction: column;
  gap: 12px;
  margin-bottom: 48px;
}
div.find-container-header {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center;
  gap: 12px;
}

div.find-container h3 {
  font-weight: bold;
  font-size: 18px;
  line-height: 160%;
}

div.find-container span {
  color: var(--span);
  font-weight: 400;
  font-size: 14px;
  line-height: 160%;
}

/*  */

.posts {
  display: grid;
  grid-template-columns: repeat(2, 1fr);
  gap: 32px;
}
</style>
