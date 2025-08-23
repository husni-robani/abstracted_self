<template>
  <AdminLayout>
    <div class="max-w-7xl mx-auto space-y-6">
      <!-- Main Form -->
      <h1 class="text-3xl font-semibold font-mono mb-6">Your Profile</h1>
      <div
        class="mx-auto flex flex-col justify-around gap-5 bg-white shadow rounded-lg p-6"
      >
        <div class="flex flex-col w-full gap-4">
          <!-- name -->
          <InputField label="Name" v-model="profile.name" />
          <!-- summary -->
          <TextAreaField label="Summary" v-model="profile.summary" :rows="3" />
          <!-- bio -->
          <TextAreaField label="Bio" v-model="profile.bio" :rows="7" />
        </div>
        <!-- resume upload and Taglines Form -->
        <div class="flex flex-col md:flex-row gap-6">
          <!-- Taglines -->
          <div class="w-full h-fit">
            <TagInput
              label="Taglines"
              v-model="profile.taglines"
              placeholder="type new tagLine ..."
            />
          </div>
          <!-- resume upload & view -->
          <div class="flex items-center gap-4 w-full h-fit">
            <FileUpload
              label="Resume"
              v-model:model-value="profile.resume_file"
              class="w-full"
            />
            <a
              :href="completeResumeURL"
              target="_blank"
              class="flex items-center gap-1 px-2 py-1 border-b text-sm font-medium font-mono"
            >
              <ArrowTopRightOnSquareIcon class="w-4 h-4" />
              View
            </a>
          </div>
        </div>
        <div class="mt-6 text-right">
          <button
            @click="saveProfile"
            class="bg-gray-900 text-white px-5 py-2 rounded-md hover:bg-gray-800 font-mono hover:cursor-pointer"
          >
            Save
          </button>
        </div>
      </div>
      <!-- Profile Skillset Management -->
      <div class="space-y-6">
        <h2 class="text-xl font-mono font-bold">Profile Skill</h2>
        <!-- Filter -->
        <div class="flex space-x-2">
          <select
            v-model="selectedType"
            class="border rounded-lg px-3 py-2 w-full max-w-sm"
          >
            <option
              v-for="type in skillSet"
              :key="type.type_name"
              :value="type.type_name"
            >
              {{ type.type_name }}
            </option>
          </select>
          <button
            @click="showAddTypeModal = !showAddTypeModal"
            class="flex items-center px-3 py-2 rounded-2xl bg-gray-200 hover:bg-gray-300 space-x-1 hover:cursor-pointer"
          >
            <PlusIcon class="w-4 h-4 color text-gray-900" />
            <span class="text-gray-900 text-sm">Type</span>
          </button>
        </div>
        <!-- Skills -->
        <SkillManagement
          v-for="type in skillSet"
          :key="type.type_name"
          v-show="selectedType === type.type_name"
          :type_name="type.type_name"
          :skill_items="type.skill_items"
          :refresh_skillset="getSkillType"
        />
      </div>
    </div>
  </AdminLayout>
  <!-- add type modal -->
  <Modal :show="showAddTypeModal" @close="closeAddTypeModal">
    <form @submit.prevent="addType" class="space-y-4">
      <InputField
        :is-required="true"
        placeholder="Type the type name ..."
        v-model="newType"
      />
      <button
        type="submit"
        class="bg-gray-900 text-white px-5 py-2 rounded-md hover:bg-gray-800 font-mono hover:cursor-pointer"
      >
        Save
      </button>
    </form>
  </Modal>
</template>

<script setup>
import AdminLayout from "../layouts/AdminLayout.vue";
import InputField from "../components/InputField.vue";
import TextAreaField from "../components/TextAreaField.vue";
import TagInput from "../components/TagInput.vue";
import FileUpload from "../components/FileUpload.vue";
import Modal from "../components/Modal.vue";
import SkillManagement from "./partials/SkillManagement.vue";
import { ArrowTopRightOnSquareIcon, PlusIcon } from "@heroicons/vue/24/outline";
import { onMounted, reactive, ref } from "vue";

const profile = reactive({
  name: "",
  summary: "",
  bio: "",
  resume_file: null,
  resume_file_name: "",
  taglines: [],
});

const original_profile = reactive({});
const skillSet = reactive([]);

const selectedType = ref();
const newType = ref("");

const showAddTypeModal = ref(false);

const completeResumeURL = ref("");
const profile_endpoint =
  import.meta.env.VITE_API_URL + import.meta.env.VITE_GET_PROFILEDATA_ENDPOINT;
const get_asset_endpoint =
  import.meta.env.VITE_API_URL + import.meta.env.VITE_ASSET_DOCUMENTS_ENDPOINT;
const update_profile_endpoint =
  import.meta.env.VITE_API_URL + import.meta.env.VITE_UPDATE_PROFILE_ENDPOINT;
const add_skill_type_endpoint =
  import.meta.env.VITE_API_URL + import.meta.env.VITE_ADD_SKILL_TYPE_ENDPOINT;

async function getProfileData() {
  try {
    const response = await fetch(
      profile_endpoint +
        "?name=true&summary=true&bio=true&taglines=true&resume=true"
    );

    const fetchedData = await response.json();

    Object.assign(profile, fetchedData.data);
    Object.assign(original_profile, fetchedData.data);
    completeResumeURL.value =
      get_asset_endpoint + "/" + profile.resume_file_name;
  } catch (e) {
    console.error("Failed to fetch profile:", e);
  }
}

async function saveProfile() {
  const changedFields = {};

  for (const key in profile) {
    // handle file differently
    if (key == "resumeFile") continue;

    const originalValue = original_profile[key];
    const currentValue = profile[key];

    const isDifferent = Array.isArray(currentValue)
      ? JSON.stringify(currentValue) !== JSON.stringify(originalValue)
      : currentValue !== originalValue;

    if (isDifferent) {
      changedFields[key] = currentValue;
    }
  }

  // handle file upload if selected
  if (profile.resume_file) {
    changedFields.resume_file = profile.resume_file;
  }

  const formData = new FormData();

  for (const key in changedFields) {
    if (key === "taglines") {
      for (const i in changedFields[key]) {
        formData.append(key, changedFields[key][i]);
      }
      continue;
    }
    formData.append(key, changedFields[key]);
  }

  try {
    const token = localStorage.getItem("token");
    const res = await fetch(update_profile_endpoint, {
      method: "PUT",
      body: formData,
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });

    if (!res.ok) {
      throw new Error(`Update failed: ${res.status}`);
    }

    alert("Update successful!");
    Object.assign(original_profile, JSON.parse(JSON.stringify(profile)));
    profile.resume_file = null;
  } catch (e) {
    console.error(e);
    alert("Update failed");
  }
}

async function getSkillType() {
  try {
    const res = await fetch(profile_endpoint + "?skill_set=true");

    if (!res.ok) {
      throw new Error(`failed to get skill data`);
    }

    const resJson = await res.json();

    Object.assign(skillSet, resJson.data.skill_set);
  } catch (e) {
    console.error(e);
  }
}

function closeAddTypeModal() {
  showAddTypeModal.value = false;
  newType.value = "";
}

async function addType() {
  try {
    const token = localStorage.getItem("token");

    const res = await fetch(add_skill_type_endpoint, {
      method: "POST",
      body: JSON.stringify({ type_name: newType.value }),
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });

    if (!res.ok) {
      throw new Error(`failed to add new type`);
    }

    await getProfileData();
    showAddTypeModal.value = false;
    newType.value = "";
  } catch (e) {
    console.error(e);
  }
}

onMounted(async () => {
  await getProfileData();
  await getSkillType();
  selectedType.value = skillSet[0].type_name;
});
</script>
