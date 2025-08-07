<template>
  <AdminLayout>
    <div class="max-w-7xl mx-auto space-y-6">
      <h1 class="text-3xl font-semibold font-mono mb-6">Project Management</h1>
      <div
        class="mx-auto flex flex-col justify-around gap-5 bg-white shadow-lg border border-gray-200 rounded-lg p-6"
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
        <!-- Existing Images -->
        <div class="flex flex-col w-full">
          <label class="block text-sm font-medium text-gray-700 font-mono mb-1"
            >Images</label
          >
          <div
            class="flex items-center h-40 border border-gray-300 rounded-md shadow-sm px-4 py-2"
          >
            <div
              v-if="project.images.length"
              class="w-full overflow-x-auto whitespace-nowrap space-x-4"
            >
              <ImageDelete
                v-for="(img, index) in previewImages"
                :key="index"
                :src="get_image_endpoint + '/' + img.file_name"
                @delete="deleteImage(img.id)"
              />
            </div>
            <div
              v-else
              class="h-full flex items-center justify-center text-gray-500 italic"
            >
              Empty
            </div>
          </div>
        </div>

        <!-- New Images -->
        <div class="flex flex-col w-full">
          <FileUpload
            label="New Images"
            v-model="project.new_images"
            placeholder="Add images"
            class="w-full"
            :multiple="true"
          />
          <div
            class="mt-3 overflow-x-auto max-w-full border-gray-200 rounded-md bg-white"
          >
            <div class="flex gap-4">
              <img
                v-for="(src, index) in previewNewImages"
                :key="index"
                :src="src"
                class="w-52 h-32 border-gray-200 object-cover rounded-md"
                alt="Preview"
              />
            </div>
          </div>
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
        <div class="mt-6 text-right space-x-2">
          <button
            @click="cancel"
            class="bg-gray-800 text-white px-5 py-2 rounded-md hover:bg-gray-900 font-mono hover:cursor-pointer"
          >
            Cancel
          </button>
          <button
            @click="updateProject"
            class="bg-red-500 text-white px-5 py-2 rounded-md hover:bg-red-700 font-mono hover:cursor-pointer"
          >
            Update
          </button>
        </div>
      </div>
    </div>
  </AdminLayout>
</template>

<script setup>
import { onMounted, reactive, ref, watch } from "vue";
import AdminLayout from "../layouts/AdminLayout.vue";
import InputField from "../components/InputField.vue";
import TextAreaField from "../components/TextAreaField.vue";
import TagInput from "../components/TagInput.vue";
import FileUpload from "../components/FileUpload.vue";
import ImageDelete from "../components/ImageDelete.vue";
import { useRoute, useRouter } from "vue-router";

const router = useRouter();
const projectId = useRoute().params.id;

const get_project_id_endpoint =
  import.meta.env.VITE_API_URL +
  import.meta.env.VITE_GET_PROJECT_BY_ID_ENDPOINT;
const update_project_endpoint =
  import.meta.env.VITE_API_URL + import.meta.env.VITE_UPDATE_PROJECT_ENDPOINT;
const get_image_endpoint =
  import.meta.env.VITE_API_URL + import.meta.env.VITE_ASSET_IMAGES_ENDPOINT;

const project = reactive({
  name: "",
  description: "",
  tech_stack: [],
  source_url: [],
  project_url: "",
  start_date: "",
  end_date: "",
  images: [],
  new_images: [],
  images_to_delete: [],
});

const previewNewImages = ref([]);
const previewImages = ref([]);

watch(
  () => project.new_images,
  (newFiles) => {
    if (!newFiles || newFiles.length === 0) {
      previewNewImages.value = [];
      return;
    }

    previewNewImages.value = newFiles.map((file) => URL.createObjectURL(file));
  }
);

watch(
  () => project.images,
  (updatedImages) => {
    if (!updatedImages || updatedImages.length === 0) {
      previewImages = [];
      return;
    }

    previewImages.value = updatedImages.map((image) => image);
  }
);

async function getProject() {
  try {
    const res = await fetch(get_project_id_endpoint + "/" + projectId);
    const resJson = await res.json();

    if (!res.ok) {
      throw new Error(
        `Fetch project failed\nstatus: ${res.status}\nmessage: ${resJson.message}`
      );
    }

    Object.assign(project, resJson.data);

    // re-formated end_date and start_date if not ""
    if (project.end_date != "") {
      project.end_date = new Intl.DateTimeFormat("en-CA").format(
        new Date(resJson.data.end_date)
      );
    }
    if (project.start_date != "") {
      project.start_date = new Intl.DateTimeFormat("en-CA").format(
        new Date(resJson.data.start_date)
      );
    }
  } catch (e) {
    console.error(e);
  }
}

async function updateProject() {
  const formData = new FormData();

  for (const key in project) {
    if (
      key === "new_images" ||
      key === "images_to_delete" ||
      key === "tech_stack" ||
      key === "source_url"
    ) {
      for (const i in project[key]) {
        formData.append(key, project[key][i]);
      }

      continue;
    }

    if (key === "images" || key === "id") continue;

    formData.append(key, project[key]);
  }

  try {
    const token = localStorage.getItem("token");
    const res = await fetch(update_project_endpoint + "/" + projectId, {
      method: "PUT",
      body: formData,
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });

    const resJson = await res.json();

    if (!res.ok) {
      throw new Error(`Create project failed: ${resJson.message}`);
    }

    clearUpdate();
    alert("Update project successfull!");

    router.push("/admin/projects");
  } catch (e) {
    console.error(e);
    alert("Update project failed: " + e);
  }
}

async function clearUpdate() {
  await getProject();
  project.new_images = [];
  project.images_to_delete = [];
}

function cancel() {
  clearUpdate();

  router.push("/admin/projects");
}

onMounted(() => {
  getProject();
});

function deleteImage(id) {
  project.images_to_delete.push(id);
  project.images = project.images.filter(
    (imageObject) => imageObject.id !== id
  );
}
</script>
