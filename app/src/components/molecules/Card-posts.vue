<script setup lang="ts">
import { useRoute } from "vue-router";

const route = useRoute();

export type IApiResponse = {
  Total: number;
  Items: IFiles[];
  User: IUser;
};

type IFiles = {
  Name: string;
  Path: string;
  Url: string;
  RepoInfo: {
    Name: string;
    FullName: string;
    Description: string;
  };
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

const goToPost = () => {
  const repo = props.post.RepoInfo.Name;
  const path = props.post.Path;
  const url = `/${user}/${repo}/${path}`;
  window.location.href = url;
};
</script>

<template>
  <div class="content" @click="goToPost()">
    <div class="title-and-date">
      <h4>{{ props.post.Name }}</h4>
      <span> {{ props.post.RepoInfo.FullName }} - {{ props.post.RepoInfo.Description }} </span>
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
  flex-direction: column;
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
  text-align: left;
}

p {
  font-weight: 400;
  font-size: 16px;
  line-height: 160%;
  color: var(--text);
}
</style>
