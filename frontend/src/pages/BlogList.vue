<template>
  <ProfileLayout>
    <div class="max-w-4xl mx-auto px-6 pt-32 pb-24">
      <router-link
        to="/"
        class="inline-block font-mono text-sm text-gray-500 hover:text-gray-900 transition-colors mb-10"
      >
        &larr; back to profile
      </router-link>

      <!-- signature title -->
      <header class="mb-14 md:mb-16">
        <div class="relative flex items-end mb-4">
          <span
            class="relative z-10 bg-white pr-4 font-roboto font-bold text-gray-700 text-2xl md:text-4xl leading-none"
          >
            all byte notes
          </span>
          <div
            class="flex-1 border-b-2 border-gray-400 mb-[0.25em] md:mb-[0.35em]"
          ></div>
        </div>
        <p class="font-mono text-sm text-gray-400">{{ posts.length }} posts</p>
      </header>

      <!-- Loading -->
      <div v-if="loading" class="py-16 flex items-center justify-center">
        <p class="text-gray-400 font-mono animate-pulse">Loading posts…</p>
      </div>

      <!-- Empty -->
      <div v-else-if="posts.length === 0">
        <p class="text-gray-400 font-mono text-sm">
          nothing here yet — writings are on the way.
        </p>
      </div>

      <!-- Posts: journal rows, reverse-chronological -->
      <div v-else class="flex flex-col border-t border-gray-200">
        <router-link
          v-for="post in posts"
          :key="post.id"
          :to="`/blog/${post.id}`"
          class="group border-b border-gray-200 py-6 md:py-7 flex items-center gap-5 md:gap-7 hover:border-gray-300 transition-colors"
        >
          <!-- cover thumb OR fallback -->
          <img
            v-if="post.cover_image"
            :src="resolveImageUrl(post.cover_image)"
            alt=""
            class="flex-none h-16 w-24 md:h-20 md:w-32 rounded-md object-cover"
          />
          <div
            v-else
            class="flex-none h-16 w-24 md:h-20 md:w-32 rounded-md bg-gray-100 flex items-center justify-center text-gray-300"
          >
            <BookOpenIcon class="w-6 h-6" />
          </div>

          <div class="flex-1 min-w-0">
            <p
              class="font-mono text-[11px] uppercase tracking-widest text-gray-400 mb-2"
            >
              {{ formatMonthYear(post.created_at) }}
            </p>
            <h2
              class="font-roboto font-semibold text-xl md:text-2xl text-gray-900 group-hover:text-gray-600 transition-colors"
            >
              {{ post.title }}
            </h2>
            <p
              v-if="post.blog_snippet"
              class="text-sm text-gray-500 mt-2 line-clamp-1 md:line-clamp-2"
            >
              {{ post.blog_snippet }}
            </p>
          </div>

          <!-- hover arrow (desktop only) -->
          <svg
            class="hidden md:block flex-none w-5 h-5 text-gray-300 group-hover:text-gray-900 group-hover:translate-x-1 transition-all"
            viewBox="0 0 24 24"
            fill="none"
            aria-hidden="true"
          >
            <path
              d="M5 12h14M13 6l6 6-6 6"
              stroke="currentColor"
              stroke-width="1.75"
              stroke-linecap="round"
              stroke-linejoin="round"
            />
          </svg>
        </router-link>
      </div>
    </div>
  </ProfileLayout>
</template>

<script setup>
import { onMounted, ref } from "vue";
import ProfileLayout from "../layouts/ProfileLayout.vue";
import { BookOpenIcon } from "@heroicons/vue/24/outline";

const VITE_API_URL = import.meta.env.VITE_API_URL || "";

const get_blogs_endpoint =
  import.meta.env.VITE_API_URL + import.meta.env.VITE_GET_BLOGS_ENDPOINT;

const posts = ref([]);
const loading = ref(true);

function resolveImageUrl(image) {
  if (!image || !image.url) return "";
  return VITE_API_URL + image.url;
}

// Journal micro-meta date, e.g. "AUG 2026".
function formatMonthYear(dateStr) {
  return new Intl.DateTimeFormat("en-US", {
    month: "short",
    year: "numeric",
  })
    .format(new Date(dateStr))
    .toUpperCase();
}

async function getPosts() {
  loading.value = true;

  try {
    const res = await fetch(get_blogs_endpoint + "?published=true");
    const resJson = await res.json();

    if (!res.ok) {
      throw new Error(resJson.message);
    }

    const data = resJson.data || [];
    // Newest first (created_at drives the journal ordering, AD-3).
    posts.value = [...data].sort(
      (a, b) => new Date(b.created_at) - new Date(a.created_at)
    );
  } catch (e) {
    // Public page: swallow errors quietly per the repo convention.
    console.error("Failed to fetch posts:", e);
  } finally {
    loading.value = false;
  }
}

onMounted(() => {
  getPosts();
});
</script>
