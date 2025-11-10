import { createApp, h } from "vue";
import { createRouter, createWebHistory, RouterView } from "vue-router";
import Home from "./components/pages/Home.vue";
import Post from "./components/pages/Post.vue";
import "./reset.css";
import "./style.css";

const router = createRouter({
  history: createWebHistory(),
  routes: [
    { path: "/:user", component: Home },
    { path: "/:user/:path", component: Post },
  ],
});

createApp({
  render: () => h(RouterView),
})
  .use(router)
  .mount("#app");
