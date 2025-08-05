<template>
  <AdminLayout>
    <div class="max-w-7xl mx-auto space-y-6">
      <h1 class="text-3xl font-semibold font-mono mb-6">Project Management</h1>
      <div>
        <!-- FORM TOGGLE BUTTOn -->
        <button
          @click="showForm = !showForm"
          class="flex items-center gap-1 text-sm text-gray-900 px-1 py-1 border-b font-mono hover:cursor-pointer mb-2"
        >
          <span>New Project</span>
          <!-- Conditional icon -->
          <ChevronDownIcon
            v-if="showForm"
            class="w-4 h-4 text-gray-900 transition-transform duration-200"
          />
          <ChevronRightIcon
            v-else
            class="w-4 h-4 text-gray-900 transition-transform duration-200"
          />
        </button>
        <!-- NEW PROJECT FORM -->
        <div
          v-show="showForm"
          class="mx-auto flex flex-col justify-around gap-5 bg-white shadow-lg border border-gray-200 rounded-lg p-6 transition-all duration-300 ease-in-out transform opacity-0 scale-95"
          :class="{ 'opacity-100 scale-100': showForm }"
        >
          <!-- name & url -->
          <div class="flex flex-col md:flex-row w-full gap-4">
            <InputField
              class="w-full"
              label="Project Name"
              v-model="project.name"
            />
            <InputField
              class="w-full"
              label="Project URL"
              v-model="project.project_url"
            />
          </div>
          <!-- description -->
          <div class="flex flex-col w-full">
            <TextAreaField
              label="Description"
              v-model="project.description"
              :rows="4"
            />
          </div>
          <!-- images -->
          <div class="flex flex-col w-full">
            <FileUpload
              label="images"
              v-model="project.images"
              class="w-full"
              :multiple="true"
            />
            <div
              class="mt-3 overflow-x-auto max-w-full border-gray-200 rounded-md bg-white"
            >
              <div class="flex gap-4">
                <img
                  v-for="(src, index) in imagePreviews"
                  :key="index"
                  :src="src"
                  class="w-48 h-32 object-cover rounded-md"
                  alt="Preview"
                />
              </div>
            </div>
          </div>
          <!-- start & end date -->
          <div class="flex flex-col md:flex-row w-full gap-4">
            <InputField
              class="w-full"
              label="Start Date"
              v-model="project.start_date"
              type="date"
            />
            <InputField
              class="w-full"
              label="End Date"
              v-model="project.end_date"
              type="date"
            />
          </div>
          <!-- source urls & tech stack -->
          <div class="flex flex-col md:flex-row w-full gap-4">
            <TagInput
              class="w-full h-fit"
              label="Tech Stack"
              v-model="project.tech_stack"
              placeholder="type new tech..."
            />
            <TagInput
              class="w-full h-fit"
              label="Source URLs"
              v-model="project.source_url"
              placeholder="type url..."
            />
          </div>
          <!-- button save -->
          <div class="mt-6 text-right">
            <button
              @click="saveProject"
              class="bg-gray-900 text-white px-5 py-2 rounded-md hover:bg-gray-800 font-mono hover:cursor-pointer"
            >
              Save
            </button>
          </div>
        </div>
      </div>
      <!-- PROJECT LIST -->
      <div class="flex flex-col gap-5">
        <AdminProjectCard
          v-for="(projectData, index) in projects"
          :key="projectData.id"
          :project="projectData"
        />
      </div>
    </div>
  </AdminLayout>
</template>

<script setup>
import AdminLayout from "../layouts/AdminLayout.vue";
import InputField from "../components/InputField.vue";
import TextAreaField from "../components/TextAreaField.vue";
import TagInput from "../components/TagInput.vue";
import FileUpload from "../components/FileUpload.vue";
import AdminProjectCard from "../components/AdminProjectCard.vue";
import { reactive, ref, watch, onMounted } from "vue";
import { ChevronDownIcon, ChevronRightIcon } from "@heroicons/vue/24/outline";

// Project Form Logic
let showForm = ref(false);
const store_project_endpoint =
  import.meta.env.VITE_API_URL + import.meta.env.VITE_STORE_PROJECT_ENDPOINT;
const imagePreviews = ref([]);
const project = reactive({
  name: "",
  description: "",
  tech_stack: [],
  source_url: [],
  project_url: "",
  start_date: "",
  end_date: "",
  images: [],
});

watch(
  () => project.images,
  (newFiles) => {
    if (!newFiles || newFiles.length === 0) {
      imagePreviews.value = [];
      return;
    }

    imagePreviews.value = newFiles.map((file) => URL.createObjectURL(file));
  }
);

async function saveProject() {
  const formData = new FormData();

  for (const key in project) {
    if (key === "tech_stack" || key === "source_url" || key === "images") {
      for (const i in project[key]) {
        formData.append(key, project[key][i]);
      }
      continue;
    }
    formData.append(key, project[key]);
  }

  try {
    const token = localStorage.getItem("token");
    const res = await fetch(store_project_endpoint, {
      method: "POST",
      body: formData,
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });

    const resJson = await res.json();

    if (!res.ok) {
      throw new Error(`Create project failed: ${resJson.message}`);
    }

    project.name = "";
    project.description = "";
    project.tech_stack = [];
    project.source_url = [];
    project.project_url = "";
    project.start_date = "";
    project.end_date = "";
    project.images = [];
    showForm = false;
    alert("Create project successful!");
  } catch (e) {
    console.error(e);
    alert("Create project failed: " + e);
  }
}

// Project List Logic
const get_projects_endpoint =
  import.meta.env.VITE_API_URL + import.meta.env.VITE_GET_PROJECTS_ENDPOINT;
const projects = reactive([]);
onMounted(async () => {
  try {
    const response = await fetch(get_projects_endpoint);

    const fetchedData = await response.json();

    Object.assign(projects, fetchedData.data);
  } catch (e) {
    console.error("Failed to fetch projects:", e);
  }
});
</script>
