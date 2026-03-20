<script lang="ts" setup>
import CodeBlock from "./CodeBlock.vue";

const amphiuma = `# Amphiuma

**Amphiuma** is a genus of aquatic
__salamanders__ from the southeastern
United States.`;

const frontMatter = `---
title: Amphiuma
date: 1762810343151
description: Open source is philosophy
---
# my content
`;
</script>

<template>
  <section class="section-container">
    <div class="card-body">
      <!-- <img src="/amphiuma.png" alt="Amphiuma" /> -->
      <CodeBlock :code="amphiuma" lang="markdown" title="README.md" />
      <div class="card-content">
        <h3 class="card-title">Create a .md file</h3>
        <p class="card-description">
          Create a .md file anywhere in your repository and write the content you want to share.
        </p>
      </div>
    </div>

    <div class="card-body">
      <div class="card-content">
        <h3 class="card-title">Use Front Matter</h3>
        <p class="card-description">Use Front Matter to add metadata to your .md file.</p>
      </div>
      <!-- <img src="/front-metter.png" alt="Front Matter" /> -->
      <CodeBlock :code="frontMatter" lang="markdown" title="front-matter.md" />
    </div>
  </section>
</template>

<style scoped>
.section-container {
  display: flex;
  flex-direction: row;
  gap: 32px;
}

@media screen and (max-width: 950px) {
  .section-container {
    flex-direction: column;
    width: 50%;
    margin: 0 auto;
  }
}

.card-body {
  flex: 1;
  background: #ffffff;
  border-radius: 27px;
  overflow: hidden;
}

.card-body img {
  display: block;
  width: 100%;
  height: auto;
  border-radius: 10px;
}

/*  */
.card-content {
  display: flex;
  flex-direction: column;
  gap: 16px;
  padding: 40px;
}
.card-title {
  font-weight: 500;
  font-size: 26px;
  line-height: 41px;
  color: #0d0e11;
}

.card-description {
  font-weight: 400;
  font-size: 20px;
  line-height: 34px;
  text-align: justify;
  color: #737782;
}
</style>
