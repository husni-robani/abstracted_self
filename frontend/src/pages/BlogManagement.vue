<template>
  <AdminLayout>
    <div class="max-w-7xl mx-auto space-y-6">
      <!-- Header row -->
      <div class="flex items-center justify-between">
        <div>
          <h1 class="text-3xl font-semibold font-mono mb-1">Blog</h1>
          <p class="text-gray-500 font-mono">Manage your posts</p>
        </div>
        <router-link
          to="/admin/blog/new"
          class="bg-gray-900 text-white px-5 py-2 rounded-md hover:bg-gray-800 font-mono"
        >
          + New Post
        </router-link>
      </div>

      <!-- Loading -->
      <div v-if="loading" class="py-16 flex items-center justify-center">
        <p class="text-gray-400 font-mono animate-pulse">Loading posts…</p>
      </div>

      <!-- Error -->
      <div
        v-else-if="error"
        class="py-16 flex flex-col items-center gap-4 text-center"
      >
        <p class="text-red-500 font-mono text-sm">Failed to load posts.</p>
        <button
          class="text-gray-800 font-mono text-sm hover:underline"
          @click="getPosts"
        >
          Try again
        </button>
      </div>

      <!-- Empty -->
      <div
        v-else-if="posts.length === 0"
        class="bg-white border border-dashed border-gray-300 rounded-lg p-12 text-center"
      >
        <BookOpenIcon class="w-10 h-10 text-gray-300 mx-auto" />
        <p class="mt-3 text-gray-500 font-mono">No posts yet</p>
        <router-link
          to="/admin/blog/new"
          class="inline-block mt-4 bg-gray-900 text-white px-5 py-2 rounded-md hover:bg-gray-800 font-mono"
        >
          Write your first post
        </router-link>
      </div>

      <!-- Post list -->
      <div v-else class="flex flex-col gap-5">
        <AdminBlogPostCard
          v-for="postData in posts"
          :key="postData.id"
          :post="postData"
          @refetch="getPosts"
        />
      </div>
    </div>
  </AdminLayout>
</template>

<script setup>
import { onMounted, ref } from "vue";
import AdminLayout from "../layouts/AdminLayout.vue";
import AdminBlogPostCard from "../components/AdminBlogPostCard.vue";
import { BookOpenIcon } from "@heroicons/vue/24/outline";

const get_blogs_endpoint =
  import.meta.env.VITE_API_URL + import.meta.env.VITE_GET_BLOGS_ENDPOINT;

const posts = ref([]);
const loading = ref(true);
const error = ref(false);

async function getPosts() {
  loading.value = true;
  error.value = false;

  try {
    // No `published` param → backend returns all posts, including drafts.
    const res = await fetch(get_blogs_endpoint);
    const resJson = await res.json();

    if (!res.ok) {
      throw new Error(resJson.message);
    }

    const data = resJson.data || [];
    // Newest first, matching the public ordering (created_at drives the journal meta).
    posts.value = [...data].sort(
      (a, b) => new Date(b.created_at) - new Date(a.created_at)
    );
  } catch (e) {
    console.error("Failed to fetch posts:", e);
    error.value = true;
  } finally {
    loading.value = false;
  }
}

onMounted(() => {
  getPosts();
});
</script>
