<template>
  <div>
    <label class="block text-sm font-medium text-gray-700 font-mono mb-1">{{
      label
    }}</label>
    <textarea
      :value="modelValue"
      @input="$emit('update:modelValue', $event.target.value)"
      :rows="rows"
      class="w-full px-4 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring font-mono"
    />
  </div>
</template>

<script setup>
defineProps({
  label: String,
  rows: {
    type: Number,
    default: 4,
  },
  modelValue: String,
});

defineEmits(["update:modelValue"]);
</script>
