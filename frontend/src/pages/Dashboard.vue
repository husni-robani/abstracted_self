<template>
  <AdminLayout>
    <div class="max-w-7xl mx-auto space-y-6">
      <!-- Page Header -->
      <div>
        <h1 class="text-3xl font-semibold font-mono mb-6">Dashboard</h1>
        <p class="text-gray-500 font-mono">Welcome back, Husni 👋</p>
      </div>

      <!-- Stats Section -->
      <div class="grid md:grid-cols-4 gap-4">
        <div
          v-for="stat in stats"
          :key="stat.label"
          class="bg-white border border-gray-200 rounded-lg p-4 shadow-sm"
        >
          <div class="text-sm font-mono text-gray-500">{{ stat.label }}</div>
          <div class="text-2xl font-mono font-bold text-gray-900">
            {{ stat.value }}
          </div>
        </div>
      </div>

      <!-- Two-column section -->
      <div class="grid md:grid-cols-2 gap-6">
        <!-- Recent Projects -->
        <div class="bg-white border border-gray-200 rounded-lg shadow-sm">
          <div class="px-4 py-3 border-b border-gray-200">
            <h2 class="text-lg font-mono font-bold text-gray-900">
              Recent Projects
            </h2>
          </div>
          <ul class="divide-y divide-gray-200">
            <li
              v-for="project in recentProjects"
              :key="project.id"
              class="px-4 py-3 hover:bg-gray-50"
            >
              <div class="flex justify-between items-center">
                <div>
                  <div class="font-mono font-semibold text-gray-900">
                    {{ project.name }}
                  </div>
                  <div class="text-sm text-gray-500">
                    {{ project.tech.join(", ") }}
                  </div>
                </div>
                <a
                  :href="project.url"
                  target="_blank"
                  class="text-sm text-gray-500 hover:text-gray-900 font-mono"
                >
                  View
                </a>
              </div>
            </li>
          </ul>
        </div>

        <!-- Recent Activities -->
        <div class="bg-white border border-gray-200 rounded-lg shadow-sm">
          <div class="px-4 py-3 border-b border-gray-200">
            <h2 class="text-lg font-mono font-bold text-gray-900">
              Recent Activities
            </h2>
          </div>
          <ul class="divide-y divide-gray-200">
            <li
              v-for="(activity, index) in activities"
              :key="index"
              class="px-4 py-3 text-sm font-mono text-gray-700"
            >
              {{ activity }}
            </li>
          </ul>
        </div>
      </div>
    </div>
  </AdminLayout>
</template>

<script setup>
import AdminLayout from "../layouts/AdminLayout.vue";

const stats = [
  { label: "Total Projects", value: 12 },
  { label: "Visitors", value: 18 },
  { label: "Skills", value: 10 },
  { label: "Resume", value: "Uploaded" },
];

const recentProjects = [
  { id: 1, name: "Portfolio Website", tech: ["Vue", "Tailwind"], url: "#" },
  { id: 2, name: "E-commerce App", tech: ["Nuxt", "Supabase"], url: "#" },
  { id: 3, name: "Admin Dashboard", tech: ["Vue", "Go"], url: "#" },
];

const activities = [
  "Updated profile bio",
  'Added a new project "Portfolio Website"',
  "Updated skills: Vue, Tailwind",
  "Uploaded a new resume",
];
</script>
