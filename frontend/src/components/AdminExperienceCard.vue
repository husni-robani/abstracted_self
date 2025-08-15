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
      <button
        @click="$emit('delete')"
        class="text-red-500 hover:text-red-700 font-mono text-sm hover:cursor-pointer"
      >
        Delete
      </button>
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
  </div>
</template>

<script setup>
import { defineEmits } from "vue";

defineProps({
  experience: {
    type: Object,
    required: true,
  },
});

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
</script>
