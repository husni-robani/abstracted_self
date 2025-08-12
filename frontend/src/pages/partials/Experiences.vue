<template>
  <div class="lg:max-w-4xl mx-auto mt-52">
    <div class="w-full h-5 border-b-2 border-gray-400 text-2xl mb-16">
      <span
        class="bg-white pr-4 font-bold text-gray-700 font-roboto text-2xl md:text-4xl"
        >Experiences</span
      >
    </div>
    <div class="relative border-l border-gray-300">
      <div v-for="exp in experiences" :key="exp.id" class="mb-10 ml-6">
        <!-- Dot -->
        <div
          class="absolute w-3 h-3 bg-gray-400 rounded-full -left-1.5 border border-gray-300"
        ></div>

        <!-- Card -->
        <div class="p-6 bg-white rounded-lg shadow-sm border border-gray-200">
          <h3 class="text-lg font-semibold text-gray-900">
            {{ exp.job_title }} —
            <span class="text-gray-600">{{ exp.company_name }}</span>
          </h3>
          <p class="text-sm text-gray-500">
            {{ exp.work_place }} • {{ formatDate(exp.start_date) }} –
            {{ exp.end_date ? formatDate(exp.end_date) : "Present" }}
          </p>

          <p class="mt-3 text-gray-700 leading-relaxed">
            {{ exp.description }}
          </p>

          <!-- Accomplishments -->
          <ul class="mt-3 list-disc list-inside text-gray-700 space-y-1">
            <li v-for="(acc, index) in exp.accomplishments" :key="index">
              {{ acc }}
            </li>
          </ul>

          <!-- Tech stacks -->
          <div class="mt-4 flex flex-wrap gap-2">
            <span
              v-for="(tech, index) in exp.tech_stack"
              :key="index"
              class="px-3 py-1 text-sm bg-gray-100 text-gray-700 rounded-full border border-gray-300"
            >
              {{ tech }}
            </span>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { onMounted, reactive } from "vue";

const getExperiencesEndpoint =
  import.meta.env.VITE_API_URL + import.meta.env.VITE_GET_EXPERIENCES_ENDPOINT;

const experiences = reactive([]);

async function getAllExperiences() {
  try {
    const res = await fetch(getExperiencesEndpoint);
    const fetchData = await res.json();

    // Clear the reactive array first
    experiences.length = 0;

    // Add `tech_stack` to each experience
    await fetchData.data.forEach((exp) => {
      console.log("loop");
      experiences.push({
        ...exp,
        tech_stack: ["Laravel", "Vue.js", "TailwindCSS", "MySQL"], // Dummy tech_stack data
      });
    });
    // Object.assign(experiences, fetchData.data);
  } catch (err) {
    console.log(err);
  }
}

onMounted(() => {
  getAllExperiences();
});

function formatDate(date) {
  return new Date(date).toLocaleDateString("en-US", {
    month: "short",
    year: "numeric",
  });
}
</script>
