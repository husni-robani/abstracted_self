<template>
  <RouterLink
    :to="to"
    class="group flex items-center gap-3 py-2 px-3 rounded-md transition font-medium"
    :class="isActive ? activeClasses : inactiveClasses"
  >
    <component :is="iconComponent" class="w-5 h-5" />
    <span class="font-mono">{{ text }}</span>
  </RouterLink>
</template>

<script setup>
import {
  HomeIcon,
  FolderIcon,
  TagIcon,
  UserIcon,
  Cog6ToothIcon,
} from "@heroicons/vue/24/outline";
import { computed } from "vue";
import { useRoute } from "vue-router";

const props = defineProps({
  text: String,
  to: String,
  icon: String,
  prefixPath: String,
});

const route = useRoute();

const isActive = computed(() => route.path.startsWith(props.prefixPath));

const activeClasses = "bg-gray-800 text-white";
const inactiveClasses = "text-gray-400 hover:bg-gray-800 hover:text-white";

// Map icon prop to component
const icons = {
  HomeIcon,
  FolderIcon,
  TagIcon,
  UserIcon,
  Cog6ToothIcon,
};

const iconComponent = computed(() => icons[props.icon] || HomeIcon);
</script>
