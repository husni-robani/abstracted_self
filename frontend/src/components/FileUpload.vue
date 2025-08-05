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
    <input
      type="file"
      class="hidden"
      ref="fileInput"
      @change="onFileChange"
      :multiple="multiple"
    />
  </div>
</template>

<script setup>
import { ref, computed } from "vue";
import { ArrowUpTrayIcon } from "@heroicons/vue/24/solid";

// Props and emits
const props = defineProps({
  modelValue: {
    type: [File, Array],
    default: null,
  },
  label: {
    type: String,
    default: "Upload file",
  },
  placeholder: {
    type: String,
    default: "Click to upload",
  },
  multiple: {
    type: Boolean,
    default: false,
  },
});
const emit = defineEmits(["update:modelValue"]);

// Refs
const fileInput = ref(null);

// Computed file name
const fileName = computed(() => {
  if (props.multiple && Array.isArray(props.modelValue)) {
    return props.modelValue.map((f) => f.name).join(", ");
  }
  if (!props.multiple && props.modelValue instanceof File) {
    return props.modelValue.name;
  }

  return props.placeholder;
});

// File change handler
function onFileChange(event) {
  const files = event.target.files;
  if (props.multiple) {
    emit("update:modelValue", Array.from(files));
  } else {
    emit("update:modelValue", files[0] || null);
  }
}
</script>
