<template>
  <AdminLayout>
    <div class="max-w-7xl mx-auto space-y-6">
      <!-- Top row: back + title -->
      <div class="flex items-center justify-between">
        <router-link
          to="/admin/blog"
          class="flex items-center gap-1 text-sm font-mono text-gray-800 hover:underline"
        >
          <ArrowLeftIcon class="w-4 h-4" /> Back to Blog
        </router-link>
        <h1 class="text-3xl font-semibold font-mono">
          {{ isEdit ? "Edit Post" : "New Post" }}
        </h1>
      </div>

      <!-- Edit-mode fetch loading -->
      <div v-if="isEdit && loading" class="py-16 flex items-center justify-center">
        <p class="text-gray-400 font-mono animate-pulse">Loading post…</p>
      </div>

      <!-- Edit-mode fetch error -->
      <div
        v-else-if="isEdit && error"
        class="py-16 flex flex-col items-center gap-4 text-center"
      >
        <p class="text-red-500 font-mono text-sm">Failed to load the post.</p>
        <button
          class="text-gray-800 font-mono text-sm hover:underline"
          @click="getPost"
        >
          Try again
        </button>
      </div>

      <template v-else>
        <!-- Header fields card -->
        <div class="bg-white shadow-lg border border-gray-200 rounded-lg p-6 space-y-5">
          <!-- title + status in one row -->
          <div class="flex flex-col md:flex-row gap-4">
            <InputField
              class="w-full"
              label="Title"
              v-model="post.title"
              :isRequired="true"
              placeholder="Required"
            />
            <div class="w-full md:w-48">
              <label class="block text-sm font-medium text-gray-700 font-mono mb-1"
                >Status</label
              >
              <select
                v-model="postStatus"
                class="w-full px-4 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring font-mono bg-white"
              >
                <option value="draft">Draft</option>
                <option value="published">Published</option>
              </select>
            </div>
          </div>

          <!-- summary / excerpt -->
          <TextAreaField label="Summary" v-model="post.blog_snippet" :rows="2" />

          <!-- cover image: FileUpload + preview thumb -->
          <div class="flex flex-col w-full">
            <FileUpload
              label="Cover Image"
              v-model="coverFile"
              :placeholder="isEdit ? 'Replace cover image (optional)' : 'Add cover image (required)'"
              class="w-full"
            />
            <p
              v-if="coverMissing"
              class="mt-2 text-sm font-mono text-amber-600"
            >
              A cover image is required to save this post.
            </p>
            <img
              v-if="coverPreview"
              :src="coverPreview"
              class="mt-3 w-52 h-32 object-cover border border-gray-200 rounded-md"
              alt="Cover preview"
            />
          </div>
        </div>

        <!-- Editor block -->
        <div class="bg-white border border-gray-200 rounded-lg shadow-sm overflow-hidden">
          <MdEditor
            v-model="post.content"
            theme="light"
            language="en-US"
            previewTheme="github"
            :sanitize="resolveImageUrls"
            :style="{ height: '600px' }"
            @onSave="handleEditorSave"
            @onUploadImg="onUploadImg"
          />
        </div>

        <!-- Action bar -->
        <div class="flex flex-col-reverse sm:flex-row justify-between gap-3">
          <router-link
            to="/admin/blog"
            class="bg-gray-800 text-white px-5 py-2 rounded-md hover:bg-gray-900 font-mono text-center"
          >
            Cancel
          </router-link>
          <div class="flex gap-3">
            <button
              v-if="!post.published"
              :disabled="saving || coverMissing"
              class="bg-gray-800 text-white px-5 py-2 rounded-md hover:bg-gray-900 font-mono disabled:opacity-50 disabled:cursor-not-allowed"
              @click="save(false)"
            >
              {{ saving ? "Saving…" : "Save Draft" }}
            </button>
            <button
              :disabled="saving || coverMissing"
              class="bg-gray-900 text-white px-5 py-2 rounded-md hover:bg-gray-800 font-mono disabled:opacity-50 disabled:cursor-not-allowed"
              @click="save(true)"
            >
              {{ saving ? "Saving…" : post.published ? "Update Post" : "Save & Publish" }}
            </button>
          </div>
        </div>
      </template>
    </div>
  </AdminLayout>
</template>

<script setup>
import { computed, reactive, ref, watch } from "vue";
import { useRoute, useRouter, onBeforeRouteLeave } from "vue-router";
import AdminLayout from "../layouts/AdminLayout.vue";
import InputField from "../components/InputField.vue";
import TextAreaField from "../components/TextAreaField.vue";
import FileUpload from "../components/FileUpload.vue";
import { ArrowLeftIcon } from "@heroicons/vue/24/outline";
import { MdEditor } from "md-editor-v3";
import "md-editor-v3/lib/style.css";

import { resolveImageUrls, extractImageIds } from "../utils/markdownImages.js";

const router = useRouter();
const route = useRoute();

const isEdit = computed(() => !!route.params.id);
const postId = computed(() => route.params.id);

// Endpoints
const store_blog_endpoint =
  import.meta.env.VITE_API_URL + import.meta.env.VITE_STORE_BLOG_ENDPOINT;
const update_blog_endpoint =
  import.meta.env.VITE_API_URL + import.meta.env.VITE_UPDATE_BLOG_ENDPOINT;
const get_blog_endpoint =
  import.meta.env.VITE_API_URL + import.meta.env.VITE_GET_BLOG_ENDPOINT;
const upload_image_endpoint =
  import.meta.env.VITE_API_URL + import.meta.env.VITE_UPLOAD_IMAGE_ENDPOINT;

const VITE_API_URL = import.meta.env.VITE_API_URL || "";

// State
const post = reactive({
  title: "",
  blog_snippet: "",
  content: "",
  published: false,
  slug: "",
});
const coverFile = ref(null); // selected File (new-mode cover or edit-mode replacement)
const existingCoverImage = ref(null); // Image object from GET (edit mode)
const loading = ref(false);
const error = ref(false);
const saving = ref(false);

// Status select mapping to boolean
const postStatus = computed({
  get: () => (post.published ? "published" : "draft"),
  set: (v) => {
    post.published = v === "published";
  },
});

const coverMissing = computed(() => isEdit.value ? false : !coverFile.value);

const coverPreview = computed(() => {
  if (coverFile.value) return URL.createObjectURL(coverFile.value);
  if (existingCoverImage.value && existingCoverImage.value.url) {
    return VITE_API_URL + existingCoverImage.value.url;
  }
  return "";
});

// Unsaved-changes guard via snapshot comparison
let originalSnapshot = "";

function makeSnapshot() {
  return JSON.stringify({
    title: post.title,
    blog_snippet: post.blog_snippet,
    content: post.content,
    published: post.published,
    cover: coverFile.value ? coverFile.value.name : "",
  });
}

function resetState() {
  post.title = "";
  post.blog_snippet = "";
  post.content = "";
  post.published = false;
  post.slug = "";
  coverFile.value = null;
  existingCoverImage.value = null;
}

onBeforeRouteLeave(() => {
  if (makeSnapshot() !== originalSnapshot) {
    return window.confirm("You have unsaved changes. Leave anyway?");
  }
  return true;
});

// Edit-mode full-load
async function getPost() {
  loading.value = true;
  error.value = false;

  try {
    const res = await fetch(get_blog_endpoint + "/" + postId.value);
    const resJson = await res.json();

    if (!res.ok) {
      throw new Error(resJson.message);
    }

    const data = resJson.data;
    post.title = data.title;
    post.blog_snippet = data.blog_snippet || "";
    post.content = data.content || "";
    post.published = !!data.published;
    post.slug = data.slug || "";
    existingCoverImage.value = data.cover_image || null;
    originalSnapshot = makeSnapshot();
  } catch (e) {
    console.error(e);
    error.value = true;
  } finally {
    loading.value = false;
  }
}

async function save(publish) {
  if (!post.title.trim()) {
    alert("Title is required.");
    return;
  }
  if (!post.content.trim()) {
    alert("Content is required.");
    return;
  }
  if (!isEdit.value && !coverFile.value) {
    alert("Cover image is required.");
    return;
  }

  saving.value = true;
  post.published = publish;

  try {
    const formData = new FormData();
    formData.append("title", post.title);
    formData.append("content", post.content);
    if (post.blog_snippet) formData.append("blog_snippet", post.blog_snippet);
    if (post.slug) formData.append("slug", post.slug);
    formData.append("published", String(post.published));

    // content_image_ids: full-replacement list parsed from the markdown (AD-11)
    const imageIds = extractImageIds(post.content);
    for (const id of imageIds) {
      formData.append("content_image_ids", id);
    }

    // Cover: required on create, optional replacement on update.
    if (coverFile.value) {
      formData.append("file", coverFile.value);
    }

    const token = localStorage.getItem("token");
    const endpoint = isEdit.value
      ? update_blog_endpoint + "/" + postId.value
      : store_blog_endpoint;
    const method = isEdit.value ? "PUT" : "POST";

    const res = await fetch(endpoint, {
      method,
      body: formData,
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });

    const resJson = await res.json();

    if (!res.ok) {
      throw new Error(resJson.message || "Save failed");
    }

    originalSnapshot = makeSnapshot();
    alert(publish ? "Post published!" : "Post saved!");
    router.push("/admin/blog");
  } catch (e) {
    console.error(e);
    alert("Save failed: " + e.message);
  } finally {
    saving.value = false;
  }
}

// Ctrl/Cmd+S from the editor → save with the current status
function handleEditorSave() {
  save(post.published);
}

// In-editor image upload: validate JPEG/PNG ≤ 300KB, POST /images, insert relative URL (AD-12, AD-5)
async function onUploadImg(files, callback) {
  const token = localStorage.getItem("token");
  const urls = [];

  for (const file of files) {
    if (!["image/jpeg", "image/png"].includes(file.type)) {
      alert("Only JPEG/PNG allowed");
      continue;
    }
    if (file.size > 300 * 1024) {
      alert("Max 300KB per image");
      continue;
    }

    const fd = new FormData();
    fd.append("file", file);

    try {
      const res = await fetch(upload_image_endpoint, {
        method: "POST",
        body: fd,
        headers: {
          Authorization: `Bearer ${token}`,
        },
      });

      if (!res.ok) {
        let message = "Upload failed";
        try {
          const j = await res.json();
          message = j.message || message;
        } catch (_) {
          /* ignore non-JSON body */
        }
        alert(message);
        continue;
      }

      const json = await res.json();
      urls.push({ url: json.data.url, alt: file.name, title: file.name });
    } catch (e) {
      console.error(e);
      alert("Upload failed: " + e.message);
    }
  }

  callback(urls);
}

// Re-initialize when toggling between /admin/blog/new and /admin/blog/:id/edit
// (the same dynamic component instance is reused by the router, so watch the param).
function init() {
  resetState();

  if (isEdit.value) {
    getPost();
  } else {
    originalSnapshot = makeSnapshot();
  }
}

watch(() => route.params.id, init, { immediate: true });
</script>
