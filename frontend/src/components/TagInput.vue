<template>
  <div>
    <label class="block text-sm font-medium text-gray-700 font-mono mb-1">{{
      label
    }}</label>
    <input
      type="text"
      v-model="newTag"
      :placeholder="placeholder"
      @keydown.enter="handleEnter"
      class="w-full px-4 py-2 border border-gray-300 rounded-md shadow-sm font-mono mb-2"
    />
    <div class="flex flex-wrap gap-2">
      <span
        v-for="(tag, index) in modelValue"
        :key="index"
        class="flex items-center gap-2 bg-gray-100 text-gray-800 px-3 py-1 rounded-sm font-mono text-sm"
      >
        {{ tag }}
        <button
          type="button"
          @click="removeTag(index)"
          class="text-red-600 hover:text-red-800 hover:cursor-pointer text-sm"
        >
          ×
        </button>
      </span>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";

const props = defineProps({
  label: String,
  placeholder: String,
  modelValue: Array,
});
const emit = defineEmits(["update:modelValue"]);

const newTag = ref("");

function handleEnter(e) {
  e.preventDefault();
  e.stopPropagation();
  addTag();
}

function addTag() {
  if (newTag.value.trim() && !props.modelValue.includes(newTag.value.trim())) {
    emit("update:modelValue", [...props.modelValue, newTag.value.trim()]);
  }
  newTag.value = "";
}

function removeTag(index) {
  const updated = [...props.modelValue];
  updated.splice(index, 1);
  emit("update:modelValue", updated);
}
</script>
