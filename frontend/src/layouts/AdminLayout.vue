<template>
  <div class="flex min-h-screen bg-gray-100 text-gray-800">
    <!-- Mobile sidebar toggle -->
    <button
      class="absolute top-4 left-4 z-50 p-2 rounded-md md:hidden bg-gray-900 text-white hover:bg-gray-800"
      @click="sidebarOpen = !sidebarOpen"
    >
      <Bars2Icon class="w-5 h-5" />
    </button>

    <!-- Sidebar -->
    <aside
      :class="[
        'fixed md:static inset-y-0 left-0 w-64 bg-gray-900 text-gray-100 p-4 transform transition-transform z-40',
        sidebarOpen ? 'translate-x-0' : '-translate-x-full md:translate-x-0',
      ]"
    >
      <div class="text-xl font-mono font-bold mb-6">Admin Panel</div>
      <nav class="space-y-1">
        <SidebarLink icon="HomeIcon" text="Dashboard" to="/admin/dashboard" />
        <SidebarLink icon="UserIcon" text="Profile" to="/admin/profile" />
        <SidebarLink icon="FolderIcon" text="Projects" to="/admin/projects" />
      </nav>
    </aside>

    <!-- Page content -->
    <main class="flex-1 p-2 m-6">
      <slot />
    </main>
  </div>
</template>

<script setup>
import { computed, ref } from "vue";
import { useRoute } from "vue-router";
import SidebarLink from "../components/SidebarLink.vue";
import { Bars2Icon } from "@heroicons/vue/24/outline";

const route = useRoute();
const pageTitle = computed(() => route.meta.title || "Admin");

const sidebarOpen = ref(false);
</script>
