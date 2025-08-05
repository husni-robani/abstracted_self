<template>
  <div class="bg-white rounded-lg shadow p-4 flex flex-col justify-between">
    <div>
      <h2 class="text-xl font-semibold font-mono text-gray-900 mb-2">
        {{ project.name }}
      </h2>
      <p class="text-sm text-gray-600 mb-2">
        {{ project.description }}
      </p>
      <div class="mb-2">
        <p class="text-xs text-gray-500 font-semibold">Tech Stack:</p>
        <ul class="flex flex-wrap gap-1">
          <li
            v-for="tech in project.tech_stack"
            :key="tech"
            class="bg-gray-200 text-xs px-2 py-1 rounded"
          >
            {{ tech }}
          </li>
        </ul>
      </div>
      <div class="mb-2">
        <p class="text-xs text-gray-500 font-semibold">Period:</p>
        <p class="text-sm text-gray-700">
          {{ formatDate(project.start_date) }} -
          {{ formatDate(project.end_date) }}
        </p>
      </div>
      <div class="mb-2">
        <p class="text-xs text-gray-500 font-semibold">Images:</p>
        <div
          class="overflow-x-auto max-w-full border-gray-200 rounded-md bg-white"
        >
          <div class="flex gap-4">
            <img
              v-for="(image, index) in project.images"
              :key="image.id"
              :src="asset_endpoint + '/' + image.file_name"
              class="w-48 h-32 object-cover rounded-md"
              alt="Preview"
            />
          </div>
        </div>
      </div>
    </div>

    <div class="flex justify-between items-center mt-4">
      <div class="flex gap-2">
        <a
          v-if="project.project_url"
          :href="project.project_url"
          target="_blank"
          class="text-xs text-indigo-600 hover:underline"
        >
          <ArrowTopRightOnSquareIcon class="w-6 h-6 text-gray-900" />
        </a>
        <a
          v-for="(url, index) in project.source_url"
          :key="index"
          :href="url"
          target="_blank"
          class="text-xs text-blue-600 hover:underline"
        >
          <GithubIcon class="w-6 h-6" />
        </a>
      </div>
      <router-link
        :to="`/`"
        class="text-sm font-mono text-gray-800 hover:underline"
      >
        Edit
      </router-link>
    </div>
  </div>
</template>

<script setup>
import { ArrowTopRightOnSquareIcon } from "@heroicons/vue/24/outline";
import GithubIcon from "../assets/GithubIcon.vue";

const asset_endpoint =
  import.meta.env.VITE_API_URL + import.meta.env.VITE_ASSET_IMAGES_ENDPOINT;

const props = defineProps({
  project: Object,
});

function formatDate(dateStr) {
  const options = { year: "numeric", month: "short", day: "numeric" };
  return new Date(dateStr).toLocaleDateString("en-US", options);
}
</script>
