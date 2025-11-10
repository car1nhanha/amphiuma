<script setup lang="ts">
import { formatDistanceToNow } from "date-fns";
import { useRouter } from "vue-router";

const router = useRouter();

export interface Post {
  id: number;
  title: string;
  description: string;
  date: string;
  file_origin?: string;
}

const props = defineProps<{ post: Post }>();

const goToPost = (id: string) => {
  router.push(`/post/${id}`);
};
</script>

<template>
  <div class="content" @click="goToPost(props.post.id)">
    <div class="title-and-date">
      <h4>{{ props.post.title }}</h4>
      <span>{{ formatDistanceToNow(new Date(props.post.date)) }}</span>
    </div>
    <p>{{ props.post.description.slice(0, 182) }}</p>
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
