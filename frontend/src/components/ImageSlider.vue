<template>
  <div class="h-full w-full">
    <!-- If there are no slides, render nothing (or you can render a placeholder) -->
    <div v-if="slides.length === 0" class="h-full w-full"></div>

    <!-- Slider -->
    <div
      v-else
      class="flex h-full w-full"
      :style="{
        transform: `translateX(-${currentIndex * 100}%)`,
        transition: transitionStyle,
      }"
      @transitionend="handleTransitionEnd"
    >
      <div
        v-for="(img, idx) in slides"
        :key="imgKey(img, idx)"
        class="min-w-full bg-center bg-cover"
        :style="{
          backgroundImage: img ? `url(${asset_endpoint}/${img.file_name})` : '',
        }"
      ></div>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted, onUnmounted, watch, nextTick } from "vue";

const asset_endpoint =
  import.meta.env.VITE_API_URL + import.meta.env.VITE_ASSET_IMAGES_ENDPOINT;

const props = defineProps({
  image_urls: {
    type: Array,
    required: true,
  },
  interval: {
    type: Number,
    default: 5000, // ms between slides
  },
  transitionTime: {
    type: Number,
    default: 1000, // ms for slide animation
  },
});

// whether we should use clones (only when more than 1 image)
const hasMultiple = computed(
  () => Array.isArray(props.image_urls) && props.image_urls.length > 1
);

// slides: cloned version if multiple, otherwise original array (single slide)
const slides = computed(() => {
  const arr = props.image_urls || [];
  if (!hasMultiple.value) return arr.slice(); // don't clone
  return [arr[arr.length - 1], ...arr, arr[0]];
});

// start index: 1 when cloned (so first real image is visible), 0 when single
const currentIndex = ref(hasMultiple.value ? 1 : 0);

// transition control: disable when we want to jump (for seamless loop)
const isTransitioning = ref(hasMultiple.value);

const transitionStyle = computed(() =>
  isTransitioning.value
    ? `transform ${props.transitionTime}ms ease-in-out`
    : "none"
);

let intervalId = null;

function startAutoplay() {
  stopAutoplay();
  if (!hasMultiple.value) return; // do not start when single or none
  intervalId = setInterval(() => {
    isTransitioning.value = true;
    currentIndex.value += 1;
  }, props.interval);
}

function stopAutoplay() {
  if (intervalId) {
    clearInterval(intervalId);
    intervalId = null;
  }
}

// when moving into the cloned slides, jump to the real slide without transition
async function handleTransitionEnd() {
  if (!hasMultiple.value) return;

  const lastIndex = slides.value.length - 1;
  if (currentIndex.value === lastIndex) {
    // moved into cloned first (at the end) -> jump to real first (index 1)
    isTransitioning.value = false;
    currentIndex.value = 1;
    await nextTick(); // allow DOM update without transition
    isTransitioning.value = true;
  } else if (currentIndex.value === 0) {
    // moved into cloned last (at the start) -> jump to real last
    isTransitioning.value = false;
    currentIndex.value = lastIndex - 1;
    await nextTick();
    isTransitioning.value = true;
  }
}

onMounted(() => {
  startAutoplay();
});

onUnmounted(() => {
  stopAutoplay();
});

// if parent updates images, reset indices and autoplay accordingly
watch(
  () => props.image_urls,
  (newVal) => {
    const many = Array.isArray(newVal) && newVal.length > 1;
    isTransitioning.value = many;
    currentIndex.value = many ? 1 : 0;
    startAutoplay();
  },
  { deep: true, immediate: false }
);

// helper for v-for key (prefer file_name, fallback to index)
function imgKey(img, idx) {
  return (img && img.file_name) || idx;
}
</script>
