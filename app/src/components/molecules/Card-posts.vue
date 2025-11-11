<script setup lang="ts">
import { useRoute, useRouter } from "vue-router";

const route = useRoute();
const router = useRouter();

export type IApiResponse = {
  Total: number;
  Items: IFiles[];
  User: IUser;
};

type IFiles = {
  Name: string;
  Path: string;
  Url: string;
};

type IUser = {
  Login: string;
  Name: string;
  Avatar_url: string;
  Url: string;
  Public_repos: number;
  Followers: number;
  Following: number;
};

const props = defineProps<{ post: IFiles }>();

const user = route.params.user;

const goToPost = (path: string) => {
  router.push(`/${user}/${path}`);
};
</script>

<template>
  <div class="content" @click="goToPost(props.post.Path)">
    <div class="title-and-date">
      <h4>{{ props.post.Name }}</h4>
      <!-- <span>{{ formatDistanceToNow(new Date(props.post.date)) }}</span> -->
    </div>
    <p>{{ props.post.Path.slice(0, 182) }}</p>
  </div>
</template>

<style scoped>
div.content {
  background-color: var(--post);
  border-radius: 10px;
  padding: 32px;

  display: flex;
  flex-direction: column;
  gap: 20px;

  cursor: pointer;
  transition: 0.3s;
}
div.content:hover {
  box-shadow: 0px 0px 50px 5px var(--post);
  transform: scale(1.02);
}

div.title-and-date {
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: flex-start;
  gap: 20px;
}

h4 {
  font-weight: bold;
  font-size: 20px;
  line-height: 160%;
  color: var(--title);
}

span {
  font-weight: 400;
  font-size: 14px;
  line-height: 160%;
  color: var(--span);
  min-width: 20%;
  text-align: right;
}

p {
  font-weight: 400;
  font-size: 16px;
  line-height: 160%;
  color: var(--text);
}
</style>
