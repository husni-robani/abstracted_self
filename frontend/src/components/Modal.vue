<template>
  <div v-if="show" class="fixed inset-0 flex items-center justify-center z-50">
    <!-- backdrop -->
    <div class="absolute inset-0 bg-gray-900/50" @click="$emit('close')"></div>

    <!-- modal content -->
    <div
      class="relative bg-white md:rounded-lg shadow-2xl p-8 w-full max-w-7xl z-10"
    >
      <!-- close button -->
      <button
        class="absolute top-4 right-4 text-gray-500 hover:text-gray-800 text-lg hover:cursor-pointer"
        @click="$emit('close')"
      >
        ✕
      </button>

      <!-- slot for form -->
      <slot></slot>
    </div>
  </div>
</template>

<script setup>
defineProps({
  show: Boolean,
});
defineEmits(["close"]);
</script>
