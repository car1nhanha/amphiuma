<script lang="ts" setup></script>

<template>
  <footer>
    <p>
      Inspired by
      <a href="https://www.figma.com/community/file/1398116799760144340/ai-dark-landing-landing-page-free-figma-file"
        >AI Dark Landing Landing Page</a
      >
      by <a href="https://www.figma.com/@robitaille">Chris Robitaille</a>.
    </p>
    <p>
      Developed by
      <a href="https://www.github.com/car1nhanha">car1nhanha</a>
    </p>
  </footer>
</template>

<style scoped>
footer {
  border-top: 1px solid var(--border);
  width: 864px;
  margin: 0 auto;
  margin-top: 100px;
  padding: 40px 0;
  color: #fff;

  display: flex;
  flex-direction: row;
  align-items: center;
  justify-content: space-between;
}

@media screen and (max-width: 900px) {
  footer {
    width: calc(96%);
    margin: 100px auto 0 auto;
    flex-direction: column;
    gap: 12px;
  }
}

footer a {
  color: #fff;
  text-decoration: underline;
}
</style>
