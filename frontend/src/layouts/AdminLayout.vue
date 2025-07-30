<template>
  <div class="min-h-screen bg-gray-100 text-gray-800">
    <!-- Sticky top nav -->
    <header
      class="sticky top-0 z-40 bg-gray-900 text-gray-100 px-4 py-3 flex items-center justify-between shadow"
    >
      <div class="text-xl font-bold font-mono">Husni's Console</div>

      <!-- Mobile toggle -->
      <button
        class="md:hidden p-2 rounded-md bg-gray-800 hover:bg-gray-700"
        @click="navOpen = !navOpen"
      >
        <Bars3Icon class="w-5 h-5 text-white" />
      </button>

      <!-- Desktop nav -->
      <nav class="hidden md:flex items-center gap-4">
        <SidebarLink icon="HomeIcon" text="Dashboard" to="/admin/dashboard" />
        <SidebarLink icon="UserIcon" text="Profile" to="/admin/profile" />
        <SidebarLink icon="FolderIcon" text="Projects" to="/admin/projects" />
      </nav>
    </header>

    <!-- Mobile nav menu -->
    <div
      v-if="navOpen"
      class="md:hidden bg-gray-900 text-gray-100 px-4 py-2 space-y-2"
    >
      <SidebarLink icon="HomeIcon" text="Dashboard" to="/admin/dashboard" />
      <SidebarLink icon="UserIcon" text="Profile" to="/admin/profile" />
      <SidebarLink icon="FolderIcon" text="Projects" to="/admin/projects" />
    </div>

    <!-- Page content -->
    <main class="p-6">
      <slot />
    </main>
  </div>
</template>

<script setup>
import { ref } from "vue";
import SidebarLink from "../components/SidebarLink.vue";
import { Bars3Icon } from "@heroicons/vue/24/outline";

const navOpen = ref(false);
</script>
