<template>
  <div>
    <label class="block text-sm font-medium text-gray-700 font-mono mb-1">{{
      label
    }}</label>
    <div
      class="flex items-center gap-4 px-4 py-3 border border-dashed border-gray-300 dark:border-gray-600 rounded-md cursor-pointer hover:border-gray-500 dark:hover:border-gray-400 transition"
      @click="fileInput.click()"
    >
      <ArrowUpTrayIcon class="w-5 h-5 text-gray-500" />
      <span class="text-sm text-gray-600 dark:text-gray-300">
        {{ fileName || placeholder }}
      </span>
    </div>
    <input type="file" class="hidden" ref="fileInput" @change="onFileChange" />
  </div>
</template>

<script setup>
import { ref, computed } from "vue";
import { ArrowUpTrayIcon } from "@heroicons/vue/24/solid";

// Props and emits
const props = defineProps({
  modelValue: File,
  label: {
    type: String,
    default: "Upload file",
  },
  placeholder: {
    type: String,
    default: "Click to upload",
  },
});
const emit = defineEmits(["update:modelValue"]);

// Refs
const fileInput = ref(null);

// Computed file name
const fileName = computed(() => {
  return typeof props.modelValue === "string"
    ? props.modelValue
    : props.modelValue?.name || "";
});

// File change handler
function onFileChange(event) {
  const file = event.target.files[0];
  if (file) {
    emit("update:modelValue", file);
  }
}
</script>
