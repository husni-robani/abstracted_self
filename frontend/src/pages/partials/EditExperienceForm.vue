<template>
  <form @submit.prevent="updateExperience" class="flex flex-col gap-5">
    <!-- job title & company -->
    <div class="flex flex-col md:flex-row gap-4">
      <InputField label="Job Title" v-model="form.job_title" class="w-full" />
      <InputField
        label="Company Name"
        v-model="form.company_name"
        class="w-full"
      />
    </div>

    <!-- workplace -->
    <InputField label="Work Place" v-model="form.work_place" class="w-full" />

    <!-- description -->
    <TextAreaField label="Description" v-model="form.description" :rows="4" />

    <!-- start & end date -->
    <div class="flex flex-col md:flex-row gap-4">
      <InputField
        label="Start Date"
        type="date"
        v-model="form.start_date"
        class="w-full"
      />
      <div class="w-full">
        <label class="block text-sm mb-1">End Date</label>
        <div class="flex gap-2">
          <input
            type="date"
            v-model="form.end_date"
            :disabled="isPresent"
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

    <!-- tech & accomplishments -->
    <div class="flex flex-col md:flex-row w-full gap-4">
      <TagInput
        class="w-full h-fit"
        label="Tech Stack"
        v-model="form.tech_stack"
        placeholder="type new tech..."
      />
      <TagInput
        class="w-full h-fit"
        label="Accomplishments"
        v-model="form.accomplishments"
        placeholder="type accomplishment..."
      />
    </div>

    <!-- save -->
    <div class="mt-4 text-right">
      <button
        type="submit"
        class="bg-gray-900 text-white px-5 py-2 rounded-md hover:bg-gray-800 hover:cursor-pointer"
      >
        Save
      </button>
    </div>
  </form>
</template>

<script setup>
import { reactive, watch, ref } from "vue";
import InputField from "../../components/InputField.vue";
import TagInput from "../../components/TagInput.vue";
import TextAreaField from "../../components/TextAreaField.vue";

const props = defineProps({
  experience: { type: Object, required: true },
});
const emits = defineEmits(["saved"]);

// helper → normalize backend data into YYYY-MM-DD for input[type=date]
function formatDateForInput(date) {
  if (!date) return "";

  // case 1: string (start_date)
  if (typeof date === "string") {
    if (date.trim() === "") return "";
    return new Intl.DateTimeFormat("en-CA").format(new Date(date));
  }

  // case 2: object with {Time, Valid} (end_date)
  if (typeof date === "object" && "Valid" in date) {
    if (!date.Valid || !date.Time) return "";
    return new Intl.DateTimeFormat("en-CA").format(new Date(date.Time));
  }

  return "";
}

// initial form state
const form = reactive({
  job_title: props.experience.job_title,
  company_name: props.experience.company_name,
  work_place: props.experience.work_place,
  start_date: formatDateForInput(props.experience.start_date),
  end_date: formatDateForInput(props.experience.end_date),
  description: props.experience.description,
  accomplishments: props.experience.accomplishments,
  tech_stack: props.experience.tech_stack,
});

const isPresent = ref(!props.experience.end_date?.Valid);

// keep form reactive if props change
watch(
  () => props.experience,
  (newVal) => {
    Object.assign(form, {
      ...newVal,
      start_date: formatDateForInput(newVal.start_date),
      end_date: formatDateForInput(newVal.end_date),
    });
    isPresent.value = !newVal.end_date?.Valid;
  },
  { immediate: true }
);

watch(isPresent, (newVal) => {
  if (newVal) form.end_date = null;
});

async function updateExperience() {
  emits("saved", form);
}
</script>
