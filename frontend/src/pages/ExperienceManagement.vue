<template>
  <AdminLayout>
    <div class="max-w-7xl mx-auto space-y-6">
      <!-- Page Header -->
      <div class="mb-6">
        <h1 class="text-3xl font-semibold font-mono">Experiences</h1>
        <p class="text-gray-500 font-mono">Manage your work history</p>
      </div>
      <!-- Form -->
      <div>
        <!-- FORM TOGGLE BUTTOn -->
        <button
          @click="showForm = !showForm"
          class="flex items-center gap-1 text-sm text-gray-900 px-1 py-1 border-b font-mono hover:cursor-pointer mb-2"
        >
          <span>Add Experience</span>
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
        <!-- New Experience Form -->
        <form
          v-show="showForm"
          @submit.prevent="saveExperience"
          class="mx-auto flex flex-col justify-around gap-5 bg-white shadow-lg border border-gray-200 rounded-lg p-6 transition-all duration-300 ease-in-out transform opacity-0 scale-95"
          :class="{ 'opacity-100 scale-100': showForm }"
        >
          <!-- job title & company -->
          <div class="flex flex-col md:flex-row w-full gap-4">
            <InputField
              class="w-full"
              label="Job Title"
              v-model="newExperience.job_title"
              :isRequired="true"
              placeholder="Required"
            />
            <InputField
              class="w-full"
              label="Company Name"
              v-model="newExperience.company_name"
              :isRequired="true"
              placeholder="Required"
            />
          </div>
          <!-- work place -->
          <div class="flex flex-col w-full">
            <InputField
              class="w-full"
              label="Work Place"
              v-model="newExperience.work_place"
              :isRequired="true"
              placeholder="Required"
            />
          </div>
          <!-- description -->
          <div class="flex flex-col w-full">
            <TextAreaField
              label="Description"
              v-model="newExperience.description"
              :rows="4"
            />
          </div>
          <!-- start & end date -->
          <div class="flex flex-col md:flex-row w-full gap-4 items-center">
            <InputField
              class="w-full"
              label="Start Date"
              v-model="newExperience.start_date"
              type="date"
              :isRequired="true"
              placeholder="Required"
            />
            <div class="w-full">
              <label
                class="block text-sm font-medium text-gray-700 font-mono mb-1"
              >
                End Date
              </label>
              <div class="flex flex-row w-full gap-x-4">
                <input
                  type="date"
                  placeholder="end date"
                  v-model="newExperience.end_date"
                  :disabled="isPresent"
                  @input="$emit('update:modelValue', $event.target.value)"
                  class="w-full px-4 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring font-mono disabled:bg-gray-100 disabled:text-gray-500"
                />
                <label class="flex items-center gap-1 text-sm text-gray-700">
                  <input
                    type="checkbox"
                    v-model="isPresent"
                    class="h-4 w-4 text-blue-600 border-gray-300 rounded hover:cursor-pointer"
                  />
                  Present
                </label>
              </div>
            </div>
          </div>
          <!-- tech stack & accomplishments -->
          <div class="flex flex-col md:flex-row w-full gap-4">
            <TagInput
              class="w-full h-fit"
              label="Tech Stack"
              v-model="newExperience.tech_stack"
              placeholder="type new tech..."
            />
            <TagInput
              class="w-full h-fit"
              label="Accomplishments"
              v-model="newExperience.accomplishments"
              placeholder="type accomplishment..."
            />
          </div>
          <!-- button save -->
          <div class="mt-6 text-right">
            <button
              type="submit"
              class="bg-gray-900 text-white px-5 py-2 rounded-md hover:bg-gray-800 font-mono hover:cursor-pointer"
            >
              Save
            </button>
          </div>
        </form>
      </div>
      <!-- Experience List -->
      <div class="space-y-4">
        <AdminExperienceCard
          v-for="exp in experiences"
          :key="exp.id"
          :experience="exp"
          @delete="deleteExperience(exp.id)"
          @refetch="getExperiences"
        />
      </div>
    </div>
  </AdminLayout>
</template>

<script setup>
import { onMounted, reactive, ref, watch } from "vue";
import AdminLayout from "../layouts/AdminLayout.vue";
import AdminExperienceCard from "../components/AdminExperienceCard.vue";
import InputField from "../components/InputField.vue";
import TagInput from "../components/TagInput.vue";
import TextAreaField from "../components/TextAreaField.vue";
import { ChevronDownIcon, ChevronRightIcon } from "@heroicons/vue/24/outline";

let showForm = ref(false);
const getExperiencesEndpoint =
  import.meta.env.VITE_API_URL + import.meta.env.VITE_GET_EXPERIENCES_ENDPOINT;
const storeExperienceEndpoint =
  import.meta.env.VITE_API_URL + import.meta.env.VITE_STORE_EXPERIENCE_ENDPOINT;
const deleteExperienceEndpoint =
  import.meta.env.VITE_API_URL +
  import.meta.env.VITE_DELETE_EXPERIENCE_ENDPOINT;

const newExperience = ref({
  job_title: "",
  company_name: "",
  work_place: "",
  start_date: "",
  end_date: "",
  description: "",
  accomplishments: [],
  tech_stack: [],
});

const isPresent = ref(false);

watch(
  () => isPresent,
  (newVal) => {
    if (newVal) experiences.endDate = "";
  }
);

const experiences = reactive([]);

const deleteExperience = async (id) => {
  if (!confirm("Are you sure ?")) {
    console.log("Delete Cancel");
    return;
  }

  try {
    const token = localStorage.getItem("token");

    const res = await fetch(deleteExperienceEndpoint + "/" + id, {
      method: "DELETE",
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });

    const resJson = await res.json();

    if (!res.ok) {
      throw new Error(resJson.message);
    }

    alert("Experience deleted!");
    getExperiences();
  } catch (e) {}
};

async function getExperiences() {
  try {
    const res = await fetch(getExperiencesEndpoint);

    const resJson = await res.json();

    Object.assign(experiences, resJson.data);
  } catch (e) {
    console.log("failed to get experiences: " + e);
  }
}

async function saveExperience() {
  try {
    const token = localStorage.getItem("token");
    const res = await fetch(storeExperienceEndpoint, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify(newExperience.value),
    });

    const jsonResponse = await res.json();

    if (!res.ok) {
      throw new Error(jsonResponse.message);
    }

    newExperience.value = [];
    showForm.value = false;
    getExperiences();
    alert("Create experience successful!");
  } catch (e) {
    console.error("Store experience failed: " + e);
    alert("Save experience failed!");
  }
}

onMounted(() => {
  getExperiences();
});
</script>
