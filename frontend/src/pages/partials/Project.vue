<template>
  <div class="lg:max-w-4xl mx-auto mt-52">
    <!-- title -->
    <div class="w-full h-5 border-b-2 border-gray-400 text-2xl mb-16">
      <span
        class="bg-white pr-4 font-bold text-gray-700 font-roboto text-2xl md:text-4xl"
        >project I've built</span
      >
    </div>
    <!-- projects list -->
    <ul
      id="projects"
      class="flex flex-col items-center relative gap-y-28 lg:gap-y-40"
    >
      <ProjectBlock
        v-for="(projectItem, index) in projects"
        :key="projectItem.id"
        :isEven="index % 2 != 0"
        :project_name="projectItem.name"
        :description="projectItem.description"
        :tech_stack="projectItem.tech_stack"
        :source_url="projectItem.source_url"
        :project_url="projectItem.source_url[0]"
        :image_url="'http://' + projectItem.images[0].image_url"
      />
    </ul>
  </div>
</template>
<script setup>
import { ref, onMounted } from "vue";
import ProjectBlock from "../../components/ProjectBlock.vue";

const image_path = import.meta.env.VITE_PUBLIC_IMAGES;
const api_endpoint =
  import.meta.env.VITE_API_URL + import.meta.env.VITE_GET_PROJECTS_ENDPOINT;
const projects = ref([]);

// fetch projects data from API
onMounted(async () => {
  try {
    const response = await fetch(api_endpoint);

    if (!response.ok) {
      throw new Error(
        `Failed to fetch projects data. Status: ${response.status}`
      );
    }

    const jsonResponse = await response.json();
    projects.value = jsonResponse.data;

    console.log(jsonResponse.data);
  } catch (e) {
    console.log(e);
  }
});
</script>
