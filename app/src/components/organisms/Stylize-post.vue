<script setup lang="ts">
import { marked } from "marked";

const props = defineProps({
  postFile: String,
});
</script>

<template>
  <div v-html="marked.parse(props.postFile ?? '')" class="markdown-body" />
</template>

<style>
div.markdown-body {
  color: var(--text);
  line-height: 1.6;
  font-size: 16px;
  padding: 40px 32px;
}

/* ----------- HEADINGS ----------- */
div.markdown-body h1,
div.markdown-body h2,
div.markdown-body h3,
div.markdown-body h4,
div.markdown-body h5,
div.markdown-body h6 {
  font-weight: bold;
  color: var(--title);
  margin-top: 32px;
  margin-bottom: 16px;
  line-height: 1.3;
}

div.markdown-body h1 {
  font-size: 32px;
  border-bottom: 2px solid var(--border);
  padding-bottom: 8px;
}

div.markdown-body h2 {
  font-size: 26px;
  color: var(--subtitle);
  padding-bottom: 4px;
}

div.markdown-body h3 {
  font-size: 22px;
  color: var(--subtitle);
}

div.markdown-body h4 {
  font-size: 20px;
  color: var(--text);
}

div.markdown-body h5 {
  font-size: 18px;
}

div.markdown-body h6 {
  font-size: 16px;
  color: var(--span);
}

/* ----------- PARAGRAPHS ----------- */
div.markdown-body p {
  margin-bottom: 16px;
  color: var(--text);
  text-align: justify;
}

/* ----------- BOLD / ITALIC / LINKS ----------- */
div.markdown-body strong {
  color: var(--title);
  font-weight: bold;
}

div.markdown-body em {
  font-style: italic;
  color: var(--subtitle);
}

div.markdown-body a {
  color: var(--blue);
  text-decoration: none;
  transition: color 0.2s;
}

div.markdown-body a:hover {
  text-decoration: underline;
}

/* ----------- LISTS ----------- */
div.markdown-body ul,
div.markdown-body ol {
  margin: 16px 0 16px 32px;
}

div.markdown-body ul {
  list-style: disc;
}

div.markdown-body ol {
  list-style: decimal;
}

div.markdown-body li {
  margin-bottom: 8px;
}

/* ----------- BLOCKQUOTE ----------- */
div.markdown-body blockquote {
  border-left: 4px solid var(--blue);
  background-color: var(--post);
  color: var(--subtitle);
  padding: 12px 16px;
  margin: 16px 0;
  border-radius: 6px;
  font-style: italic;
}

/* ----------- CODE ----------- */
div.markdown-body code {
  background-color: var(--input);
  color: var(--blue);
  padding: 2px 6px;
  border-radius: 4px;
  font-family: "Fira Code", monospace;
  font-size: 14px;
}

div.markdown-body pre {
  background-color: var(--post);
  padding: 16px;
  border-radius: 8px;
  overflow-x: auto;
  margin: 20px 0;
  border: 1px solid var(--border);
}

div.markdown-body pre code {
  background: none;
  color: var(--subtitle);
  padding: 0;
  font-size: 14px;
  line-height: 1.5;
}

/* ----------- TABLES ----------- */
div.markdown-body table {
  width: 100%;
  border-collapse: collapse;
  margin: 24px 0;
}

div.markdown-body th,
div.markdown-body td {
  border: 1px solid var(--border);
  padding: 12px;
  text-align: left;
}

div.markdown-body th {
  background-color: var(--profile);
  color: var(--title);
}

div.markdown-body tr:nth-child(even) {
  background-color: var(--post);
}

/* ----------- IMAGES ----------- */
div.markdown-body img {
  max-width: 100%;
  border-radius: 8px;
  margin: 20px 0;
  display: inline-block;
}

/* ----------- HORIZONTAL RULE ----------- */
div.markdown-body hr {
  border: none;
  border-top: 1px solid var(--border);
  margin: 32px 0;
}

/* ----------- INLINE QUOTES ----------- */
div.markdown-body q {
  quotes: "“" "”";
  color: var(--subtitle);
  font-style: italic;
}

/* ----------- CODE SYNTAX HIGHLIGHT BASE ----------- */
div.markdown-body pre code[class*="language-"] {
  display: block;
  overflow-x: auto;
  padding: 1em;
  background: var(--post);
  border-radius: 6px;
}
</style>
