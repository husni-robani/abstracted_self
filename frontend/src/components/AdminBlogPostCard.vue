<template>
  <div
    class="bg-white border border-gray-200 rounded-lg shadow-sm p-4 flex flex-col justify-between"
  >
    <div>
      <!-- top row: title + status badge -->
      <div class="flex items-start justify-between gap-4">
        <h2 class="text-xl font-semibold font-mono text-gray-900">
          {{ post.title }}
        </h2>
        <!-- status badge -->
        <span
          v-if="post.published"
          class="inline-flex items-center gap-1.5 px-2.5 py-0.5 rounded-full text-xs font-mono bg-green-100 text-green-800"
        >
          <span class="w-1.5 h-1.5 rounded-full bg-green-500"></span>
          Published
        </span>
        <span
          v-else
          class="inline-flex items-center gap-1.5 px-2.5 py-0.5 rounded-full text-xs font-mono bg-amber-100 text-amber-800"
        >
          <span class="w-1.5 h-1.5 rounded-full bg-amber-500"></span>
          Draft
        </span>
      </div>

      <!-- summary -->
      <p v-if="post.blog_snippet" class="text-sm text-gray-600 mt-2 line-clamp-2">
        {{ post.blog_snippet }}
      </p>

      <!-- dates -->
      <p class="text-xs font-mono text-gray-400 mt-2">
        Created {{ formatDate(post.created_at) }}
        <span v-if="post.updated_at"> · Updated {{ formatDate(post.updated_at) }}</span>
      </p>

      <!-- optional cover thumbnail -->
      <img
        v-if="post.cover_image"
        :src="resolveCoverUrl(post.cover_image)"
        class="mt-3 w-40 h-24 object-cover border border-gray-200 rounded-md"
        alt="Cover"
      />
    </div>

    <!-- actions -->
    <div class="flex justify-between items-center mt-4">
      <div class="flex gap-2 items-center">
        <router-link
          :to="`/admin/blog/${post.id}/edit`"
          class="px-2 text-sm text-center font-mono text-gray-800 hover:underline"
        >
          Edit
        </router-link>
        <button
          v-if="post.published"
          class="px-2 text-center text-sm font-mono text-amber-700 hover:underline hover:cursor-pointer"
          @click="togglePublish"
        >
          Unpublish
        </button>
        <button
          v-else
          class="px-2 text-center text-sm font-mono text-green-700 hover:underline hover:cursor-pointer"
          @click="togglePublish"
        >
          Publish
        </button>
        <button
          class="px-2 text-center text-sm font-mono text-red-500 hover:underline hover:cursor-pointer"
          @click="deletePost"
        >
          Delete
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { defineEmits } from "vue";
import { extractImageIds } from "../utils/markdownImages.js";

const VITE_API_URL = import.meta.env.VITE_API_URL || "";

const get_blog_endpoint =
  import.meta.env.VITE_API_URL + import.meta.env.VITE_GET_BLOG_ENDPOINT;
const update_blog_endpoint =
  import.meta.env.VITE_API_URL + import.meta.env.VITE_UPDATE_BLOG_ENDPOINT;
const delete_blog_endpoint =
  import.meta.env.VITE_API_URL + import.meta.env.VITE_DELETE_BLOG_ENDPOINT;

const props = defineProps({
  post: Object,
});

const emit = defineEmits(["refetch"]);

function formatDate(dateStr) {
  const options = { year: "numeric", month: "short", day: "numeric" };
  return new Date(dateStr).toLocaleDateString("en-US", options);
}

// cover_image is an Image|null object with a relative `url` ("/images/{id}").
function resolveCoverUrl(image) {
  if (!image || !image.url) return "";
  return VITE_API_URL + image.url;
}

// Publish/Unpublish via PUT full-replace: fetch the full post (to recover content),
// flip `published`, and send back every field including the content image ids (AD-11).
async function togglePublish() {
  const token = localStorage.getItem("token");

  try {
    const fetchRes = await fetch(get_blog_endpoint + "/" + props.post.id, {
      method: "GET",
    });

    const resJson = await fetchRes.json();

    if (!fetchRes.ok) {
      throw new Error(resJson.message);
    }

    const post = resJson.data;
    const nextPublished = !post.published;

    const formData = new FormData();
    formData.append("title", post.title);
    formData.append("content", post.content);
    if (post.blog_snippet) formData.append("blog_snippet", post.blog_snippet);
    if (post.slug) formData.append("slug", post.slug);
    formData.append("published", String(nextPublished));

    const imageIds = extractImageIds(post.content);
    for (const id of imageIds) {
      formData.append("content_image_ids", id);
    }

    const updateRes = await fetch(update_blog_endpoint + "/" + props.post.id, {
      method: "PUT",
      body: formData,
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });

    const updateResJson = await updateRes.json();

    if (!updateRes.ok) {
      throw new Error(updateResJson.message);
    }

    emit("refetch");
    alert(nextPublished ? "Post published!" : "Post unpublished!");
  } catch (e) {
    console.error(e);
    alert("Something went wrong: " + e.message);
  }
}

async function deletePost() {
  const token = localStorage.getItem("token");

  if (!confirm("Are you sure ?")) {
    console.log("Delete Cancel");
    return;
  }

  try {
    const res = await fetch(delete_blog_endpoint + "/" + props.post.id, {
      method: "DELETE",
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });

    if (!res.ok) {
      const resJson = await res.json();
      throw new Error(resJson.message);
    }

    emit("refetch");
    alert("Post deleted!");
  } catch (e) {
    console.error(e);
    alert("Something went wrong: " + e.message);
  }
}
</script>
