<template>
  <ProfileLayout>
    <div class="max-w-3xl mx-auto px-6 pt-32 pb-24">
      <!-- Loading -->
      <div v-if="loading" class="py-16 flex items-center justify-center">
        <p class="text-gray-400 font-mono animate-pulse">Loading post…</p>
      </div>

      <!-- 404 / not found -->
      <div
        v-else-if="notFound"
        class="py-24 flex flex-col items-center gap-4 text-center"
      >
        <h1
          class="text-3xl md:text-5xl font-roboto font-bold text-gray-700"
        >
          Post not found
        </h1>
        <p class="text-gray-500 font-mono text-sm">
          This post may have been unpublished or removed.
        </p>
        <router-link
          to="/blog"
          class="mt-2 bg-gray-900 text-white px-5 py-2 rounded-md hover:bg-gray-800 font-mono"
        >
          Back to blog
        </router-link>
      </div>

      <!-- Error -->
      <div
        v-else-if="error"
        class="py-24 flex flex-col items-center gap-4 text-center"
      >
        <p class="text-red-500 font-mono text-sm">Failed to load post.</p>
        <button
          class="text-gray-800 font-mono text-sm hover:underline"
          @click="getPost"
        >
          Try again
        </button>
      </div>

      <template v-else>
        <router-link
          to="/blog"
          class="inline-block font-mono text-sm text-gray-500 hover:text-gray-900 transition-colors mb-10"
        >
          &larr; back to notes
        </router-link>

        <article>
          <header class="mb-10">
            <h1
              class="text-3xl md:text-4xl font-roboto font-bold text-gray-900 mb-4 leading-tight"
            >
              {{ post.title }}
            </h1>
            <div class="font-mono text-xs md:text-sm text-gray-400">
              {{ formatFullDate(post.created_at) }}
              <span v-if="readingTime"> · {{ readingTime }} min read</span>
            </div>
          </header>

          <!-- cover image -->
          <img
            v-if="post.cover_image"
            :src="resolveImageUrl(post.cover_image)"
            alt=""
            class="mt-8 w-full rounded-2xl"
          />

          <!-- body -->
          <div class="mt-8">
            <MdPreview
              :modelValue="post.content"
              previewTheme="github"
              :sanitize="resolveImageUrls"
              class="text-gray-700 leading-relaxed"
            />
          </div>
        </article>
      </template>
    </div>
  </ProfileLayout>
</template>

<script setup>
import { computed, onMounted, ref } from "vue";
import { useRoute } from "vue-router";
import ProfileLayout from "../layouts/ProfileLayout.vue";
import { MdPreview } from "md-editor-v3";
import "md-editor-v3/lib/preview.css";
import { resolveImageUrls } from "../utils/markdownImages.js";

const route = useRoute();
const VITE_API_URL = import.meta.env.VITE_API_URL || "";

const get_blog_endpoint =
  import.meta.env.VITE_API_URL + import.meta.env.VITE_GET_BLOG_ENDPOINT;

const post = ref(null);
const loading = ref(true);
const error = ref(false);
const notFound = ref(false);

const readingTime = computed(() => {
  if (!post.value || !post.value.content) return 0;
  const words = post.value.content.trim().split(/\s+/).filter(Boolean).length;
  return Math.ceil(words / 200);
});

function resolveImageUrl(image) {
  if (!image || !image.url) return "";
  return VITE_API_URL + image.url;
}

// e.g. "Aug 12, 2026"
function formatFullDate(dateStr) {
  return new Intl.DateTimeFormat("en-US", {
    month: "short",
    day: "numeric",
    year: "numeric",
  }).format(new Date(dateStr));
}

async function getPost() {
  loading.value = true;
  error.value = false;
  notFound.value = false;

  try {
    const res = await fetch(get_blog_endpoint + "/" + route.params.id);

    if (res.status === 404) {
      notFound.value = true;
      loading.value = false;
      return;
    }

    const resJson = await res.json();

    if (!res.ok) {
      throw new Error(resJson.message);
    }

    post.value = resJson.data;
  } catch (e) {
    console.error(e);
    error.value = true;
  } finally {
    loading.value = false;
  }
}

onMounted(() => {
  getPost();
});
</script>
