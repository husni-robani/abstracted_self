<template>
  <div class="relative inline-block w-52 h-32">
    <img
      :src="src"
      alt="Image"
      class="h-full w-full border border-gray-200 object-cover rounded-md"
    />
    <button
      @click="$emit('delete')"
      class="absolute top-1 right-1 px-2 py-1 text-white text-xs rounded backdrop-blur-sm bg-red-400/50 hover:bg-red-600/70 hover:cursor-pointer transition shadow"
    >
      ✕
    </button>
  </div>
</template>

<script>
export default {
  props: {
    src: {
      type: String,
      required: true,
    },
  },
};
</script>
