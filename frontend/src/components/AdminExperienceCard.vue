<template>
  <div class="bg-white border border-gray-200 rounded-lg shadow-sm p-4">
    <div class="flex justify-between items-start">
      <div>
        <h3 class="text-lg font-mono font-bold text-gray-900">
          {{ experience.job_title }}
        </h3>
        <p class="text-sm font-mono text-gray-500">
          {{ experience.company_name }} • {{ experience.work_place }}
        </p>
        <p class="text-xs font-mono text-gray-400">
          {{ formatStartDate(experience.start_date) }} -
          {{ formatEndDate(experience.end_date) }}
        </p>
      </div>
      <!-- delete & edit buttons -->
      <div>
        <button
          @click="openEditExperienceModal"
          class="px-2 text-sm text-center font-mono text-gray-800 hover:underline hover:cursor-pointer"
        >
          Edit
        </button>
        <button
          @click="$emit('delete')"
          class="px-2 text-center text-sm font-mono text-red-500 hover:underline hover:cursor-pointer"
        >
          Delete
        </button>
      </div>
    </div>
    <div class="mt-2 text-sm font-mono text-gray-700">
      {{ experience.description }}
    </div>
    <!-- Accomplishments -->
    <ul class="mt-2 list-disc list-inside text-gray-700 space-y-1">
      <li v-for="(acc, index) in experience.accomplishments" :key="index">
        {{ acc }}
      </li>
    </ul>
    <!-- Tech stacks -->
    <div class="mt-2 text-sm font-mono flex flex-wrap gap-2">
      <span
        v-for="(tech, index) in experience.tech_stack"
        :key="index"
        class="px-1 py-1 text-sm bg-gray-100 text-gray-700 rounded-sm border border-gray-300"
      >
        {{ tech }}
      </span>
    </div>
    <!-- edit experience form modal -->
    <Modal :show="showModal" @close="closeEditExperienceModal">
      <EditExperienceForm
        v-if="selectedExperience"
        :experience="experience"
        @saved="handleSaved"
      />
    </Modal>
  </div>
</template>

<script setup>
import { defineEmits, ref } from "vue";
import Modal from "../components/Modal.vue";
import EditExperienceForm from "../pages/partials/EditExperienceForm.vue";

defineProps({
  experience: {
    type: Object,
    required: true,
  },
});

const updateExperienceEndpoint =
  import.meta.env.VITE_API_URL +
  import.meta.env.VITE_UPDATE_EXPERIENCE_ENDPOINT;
const showModal = ref(false);
const selectedExperience = ref(null);

const emit = defineEmits(["refetch", "delete"]);

function formatStartDate(dateStr) {
  const options = { year: "numeric", month: "short", day: "numeric" };
  return new Date(dateStr).toLocaleDateString("en-US", options);
}

function formatEndDate(date) {
  if (date.Valid) {
    const options = { year: "numeric", month: "short", day: "numeric" };
    return new Date(date.Time).toLocaleDateString("en-US", options);
  }

  return "present";
}

function closeEditExperienceModal() {
  showModal.value = false;
  selectedExperience.value = null;
}

function openEditExperienceModal(exp) {
  selectedExperience.value = { ...exp }; // pass copy to form
  showModal.value = true;
}

function closeModal() {
  showModal.value = false;
  selectedExperience.value = null;
}

async function handleSaved(updated) {
  if (updated.end_date == null) {
    updated.end_date = "";
  }
  const token = localStorage.getItem("token");
  try {
    const res = await fetch(updateExperienceEndpoint + "/" + updated.id, {
      method: "PATCH",
      headers: {
        "Content-Type": "application/json",
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify(updated),
    });

    if (!res.ok) {
      const resJson = await res.json();
      throw new Error(`update failed: ${resJson.message}`);
    }

    alert("Update successfull!");
  } catch (e) {
    console.error(e);
    alert(e);
  }

  closeModal();
}
</script>
