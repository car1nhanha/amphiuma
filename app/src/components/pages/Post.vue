<script setup lang="ts">
import { Icon } from "@iconify/vue";
import { formatDistanceToNow } from "date-fns";
import { ref } from "vue";
import { useRoute } from "vue-router";
import CardHeader from "../molecules/Card-header.vue";
import Stylize from "../organisms/Stylize-post.vue";
import TemplateDefault from "../templates/Defaut.vue";

const route = useRoute();

const user = route.params.user;
const path = route.params.path;

const file = ref("");
const header = ref({} as header);
const created_at = ref();

type Iresponse = {
  File: string;
  Header: header;
};

export type header = {
  Title: string;
  Date: string;
  Author: string;
  Description: string;
  File_origin: string;
};

const isTimestamp = (t: string) => /\d{13}/.test(t);

fetch(`http://localhost:8080/v1/${user}/${path}`)
  .then((response) => response.json())
  .then((response: Iresponse) => response)
  .then((content) => {
    file.value = content.File;
    header.value = content.Header;
    created_at.value = isTimestamp(content.Header.Date) ? parseInt(content.Header.Date) : content.Header.Date;
  })
  .catch(() => alert("deu erro"));
</script>

<template>
  <TemplateDefault>
    <CardHeader>
      <div class="card-content">
        <div class="line-actions">
          <a href="/"><Icon icon="solar:alt-arrow-left-line-duotone" /> return</a>
          <a :href="`${header.File_origin}`">read on github <Icon icon="solar:square-arrow-right-up-broken" /></a>
        </div>

        <h1>{{ header.Title }}</h1>

        <ul>
          <li><Icon icon="octicon:mark-github-24" /> {{ user }}</li>
          <li><Icon icon="solar:calendar-mark-bold" /> {{ formatDistanceToNow(new Date(created_at)) }}</li>
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
