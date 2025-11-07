# 🚀 Dominando o Vue 3 — Reatividade, Componentes e Performance

![Banner Vue 3](https://encrypted-tbn0.gstatic.com/images?q=tbn:ANd9GcSsteA1GaknD4QQmtH9fLEk95YHH8glkK360Q&s)

O **Vue 3** é uma das bibliotecas mais modernas e poderosas do ecossistema front-end.  
Com a **Composition API**, **reatividade refinada** e **melhor performance**, ele se tornou uma escolha sólida tanto para projetos pequenos quanto para aplicações de larga escala.

> "_O Vue 3 não é apenas uma atualização — é uma nova filosofia de arquitetura de componentes._"  
> — Evan You (criador do Vue)

---

## 🧱 1. A Nova Era da Composição

No Vue 2, a lógica de um componente era centralizada dentro de `data`, `computed` e `methods`.  
No Vue 3, com a **Composition API**, podemos **organizar a lógica por funcionalidades**, não por tipo de propriedade.

### Exemplo comparativo:

#### Antes (Vue 2 — Options API)

```js
export default {
  data() {
    return {
      count: 0,
    };
  },
  methods: {
    increment() {
      this.count++;
    },
  },
};
```

#### Agora (Vue 3 — Composition API)

```js
import { ref } from "vue";

export default {
  setup() {
    const count = ref(0);
    const increment = () => count.value++;
    return { count, increment };
  },
};
```

💡 O `ref()` é a base da **reatividade** — qualquer alteração em `count.value` re-renderiza automaticamente o DOM.

---

## ⚙️ 2. Reatividade Profunda

![Reactivity Graph](https://v1.vuejs.org/images/data.png)

A reatividade do Vue 3 é construída sobre **Proxies do ES6**, permitindo que objetos inteiros sejam observados.

```js
import { reactive, watch } from "vue";

const user = reactive({
  name: "Lucas",
  age: 25,
});

watch(
  () => user.age,
  (newVal, oldVal) => {
    console.log(`Idade mudou de ${oldVal} para ${newVal}`);
  }
);

user.age++; // dispara o watcher automaticamente
```

### 🧠 Internamente:

O Vue intercepta as operações `get` e `set` no objeto, e só re-renderiza os componentes que dependem das propriedades modificadas — **nada de renderização desnecessária**.

---

## 🧩 3. Componentização Inteligente

![Component Tree](https://vuejs.org/assets/prop-drilling.XJXa8UE-.png)

O Vue 3 incentiva a criação de componentes **menores e mais específicos**.

Um exemplo real de design de componentes pode ser visto na construção de um **Card de postagem**:

```vue
<script setup lang="ts">
defineProps({
  post: {
    type: Object,
    required: true,
  },
});
</script>

<template>
  <article class="post-card">
    <img :src="post.image" alt="Post image" />
    <h2>{{ post.title }}</h2>
    <p>{{ post.description }}</p>
    <a :href="`/post/${post.id}`">Ler mais →</a>
  </article>
</template>

<style scoped>
.post-card {
  background: var(--post);
  padding: 16px;
  border-radius: 8px;
  transition: all 0.3s;
}
.post-card:hover {
  transform: translateY(-4px);
}
</style>
```

---

## 🔄 4. Comunicação entre Componentes

O Vue 3 oferece múltiplas formas de comunicação:

| Método           | Direção     | Descrição                                        |
| ---------------- | ----------- | ------------------------------------------------ |
| `props`          | Pai → Filho | Passa dados diretamente para o componente filho  |
| `emit`           | Filho → Pai | Emite eventos personalizados                     |
| `provide/inject` | Global      | Compartilha dependências entre hierarquias       |
| Pinia / Vuex     | Global      | Armazena e gerencia estado de forma centralizada |

### Exemplo de `emit`

```vue
<!-- Botão.vue -->
<script setup>
const emit = defineEmits(["clicked"]);
</script>

<template>
  <button @click="emit('clicked')">Clique aqui</button>
</template>
```

```vue
<!-- App.vue -->
<template>
  <Botao @clicked="handleClick" />
</template>

<script setup>
const handleClick = () => alert("Botão clicado!");
</script>
```

---

## ⚡ 5. Performance e Lazy Loading

O Vue 3 introduz o **Tree Shaking nativo**, o que significa que apenas os recursos usados entram no build final.
Além disso, o **Lazy Loading de rotas** é incrivelmente simples com o `vue-router`:

```js
import { createRouter, createWebHistory } from "vue-router";

const routes = [
  {
    path: "/post/:id",
    component: () => import("../pages/Post.vue"), // só carrega quando necessário
  },
];

export const router = createRouter({
  history: createWebHistory(),
  routes,
});
```

> Dica: combine isso com `Suspense` e `defineAsyncComponent` para carregamentos assíncronos mais elegantes.

---

## 🧠 6. Hooks Personalizados

Com a Composition API, você pode criar **hooks reutilizáveis** (como no React):

```ts
// useMouse.ts
import { ref, onMounted, onUnmounted } from "vue";

export function useMouse() {
  const x = ref(0);
  const y = ref(0);

  const update = (e: MouseEvent) => {
    x.value = e.pageX;
    y.value = e.pageY;
  };

  onMounted(() => window.addEventListener("mousemove", update));
  onUnmounted(() => window.removeEventListener("mousemove", update));

  return { x, y };
}
```

E no componente:

```vue
<script setup lang="ts">
import { useMouse } from "../composables/useMouse";

const { x, y } = useMouse();
</script>

<template>
  <p>Mouse: {{ x }}, {{ y }}</p>
</template>
```

---

## 🧭 7. Deploy e SEO

![Vite + Vue](https://www.blog.vigil.global/assets/images/posts/vue-vite.png)

Vue 3 trabalha perfeitamente com **Vite**, garantindo builds rápidos e hot reload instantâneo.

```bash
pnpm create vite@latest my-vue-app --template vue-ts
pnpm dev
```

Para SEO, frameworks como **Nuxt 3** aproveitam a SSR (Server-Side Rendering) para gerar conteúdo indexável pelos mecanismos de busca.

---

## ✅ Conclusão

O Vue 3 traz uma abordagem moderna e minimalista, focada em **desempenho, reatividade e clareza**.
Se você está vindo de React, Angular ou Svelte, vai se surpreender com a **simplicidade e poder da Composition API**.

> "_A curva de aprendizado do Vue é suave, mas o potencial é imenso._"

---

### 📚 Recursos recomendados

- [Documentação oficial do Vue 3](https://vuejs.org/)
- [Curso gratuito de Composition API](https://vue-mastery.com)
- [Reatividade em Profundidade - Blog Vue](https://blog.vuejs.org)

---

### 🏁 Exemplo Final

```vue
<script setup>
import { ref } from "vue";

const name = ref("");
</script>

<template>
  <input v-model="name" placeholder="Digite seu nome" />
  <p>Olá, {{ name || "Visitante" }}!</p>
</template>
```

![Vue Logo](https://images.icon-icons.com/2699/PNG/512/vuejs_logo_icon_169247.png)

---

_Artigo escrito por **Lucas Carinhanha**, explorando o poder e a elegância do Vue 3 no desenvolvimento moderno de interfaces._
