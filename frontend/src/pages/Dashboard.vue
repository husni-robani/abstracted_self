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
                  {{ (project.tech_stack || []).join(", ") }}
                </div>
              </div>
              <a
                :href="project.project_url"
                target="_blank"
                class="text-sm text-gray-500 hover:text-gray-900 font-mono"
              >
                View
              </a>
            </div>
          </li>
        </ul>
      </div>
    </div>
  </AdminLayout>
</template>

<script setup>
import AdminLayout from "../layouts/AdminLayout.vue";
import { ref, computed, onMounted } from "vue";

const projects_endpoint =
  import.meta.env.VITE_API_URL + import.meta.env.VITE_GET_PROJECTS_ENDPOINT;
const daily_visits_counts_endpoint =
  import.meta.env.VITE_API_URL +
  import.meta.env.VITE_GET_DAILY_VISITS_COUNTS_ENDPOINT;
const resume_status_endpoint =
  import.meta.env.VITE_API_URL + import.meta.env.VITE_RESUME_STATUS_ENDPOINT;
const profile_endpoint =
  import.meta.env.VITE_API_URL + import.meta.env.VITE_GET_PROFILEDATA_ENDPOINT;

const token = localStorage.getItem("token");

const totalProjects = ref(0);
const visitors = ref(0);
const skills = ref(0);
const resumeStatus = ref("Not Uploaded");
const recentProjects = ref([]);

const stats = computed(() => [
  { label: "Total Projects", value: totalProjects.value },
  { label: "Visitors", value: visitors.value },
  { label: "Skills", value: skills.value },
  { label: "Resume", value: resumeStatus.value },
]);

function formatDate(d) {
  const month = String(d.getMonth() + 1).padStart(2, "0");
  const day = String(d.getDate()).padStart(2, "0");
  return `${d.getFullYear()}-${month}-${day}`;
}

async function getTotalProjects() {
  try {
    const res = await fetch(projects_endpoint + "?limit=1");

    if (!res.ok) {
      throw new Error(`Failed to fetch total projects. Status: ${res.status}`);
    }

    const json = await res.json();
    totalProjects.value = json.data.pagination.total;
  } catch (e) {
    console.error(e);
  }
}

async function getRecentProjects() {
  try {
    const res = await fetch(projects_endpoint + "?limit=5");

    if (!res.ok) {
      throw new Error(`Failed to fetch recent projects. Status: ${res.status}`);
    }

    const json = await res.json();
    recentProjects.value = json.data.projects;
  } catch (e) {
    console.error(e);
  }
}

async function getVisitors() {
  try {
    const end = new Date();
    const start = new Date();
    start.setDate(start.getDate() - 30);

    const query = `?start_date=${formatDate(start)}&end_date=${formatDate(end)}`;
    const res = await fetch(daily_visits_counts_endpoint + query, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });

    if (!res.ok) {
      throw new Error(`Failed to fetch visitors. Status: ${res.status}`);
    }

    const json = await res.json();
    visitors.value = json.data.reduce((acc, d) => acc + d.count, 0);
  } catch (e) {
    console.error(e);
  }
}

async function getSkills() {
  try {
    const res = await fetch(profile_endpoint + "?skill_set=true");

    if (!res.ok) {
      throw new Error(`Failed to fetch skills. Status: ${res.status}`);
    }

    const json = await res.json();
    skills.value = json.data.skill_set.reduce(
      (acc, type) => acc + type.skill_items.length,
      0
    );
  } catch (e) {
    console.error(e);
  }
}

async function getResumeStatus() {
  try {
    const res = await fetch(resume_status_endpoint, {
      headers: {
        Authorization: `Bearer ${token}`,
      },
    });

    if (!res.ok) {
      throw new Error(`Failed to fetch resume status. Status: ${res.status}`);
    }

    const json = await res.json();
    resumeStatus.value = json.data.is_uploaded ? "Uploaded" : "Not Uploaded";
  } catch (e) {
    console.error(e);
  }
}

onMounted(() => {
  getTotalProjects();
  getRecentProjects();
  getVisitors();
  getSkills();
  getResumeStatus();
});
</script>
