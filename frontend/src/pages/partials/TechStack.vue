<template>
  <section id="tech-stack" class="mx-auto py-10">
    <h2 class="text-2xl font-semibold text-gray-800 mb-4">
      Technologies I've Worked With
    </h2>

    <!-- Tabs -->
    <div class="flex flex-wrap gap-2 border-b border-gray-200 mb-6">
      <button
        v-for="category in categories"
        :key="category"
        @click="activeTab = category"
        class="px-4 py-2 rounded-t-lg text-sm font-medium transition-colors hover:cursor-pointer"
        :class="[
          activeTab === category
            ? 'bg-gray-100 text-gray-900 border border-b-0 border-gray-200'
            : 'text-gray-500 hover:text-gray-800',
        ]"
      >
        {{ category }}
      </button>
    </div>

    <!-- Content -->
    <div class="bg-gray-50 shadow rounded-lg p-6 border border-gray-100">
      <div class="flex flex-wrap gap-6">
        <div
          v-for="skill in filteredSkills"
          :key="skill.name"
          class="relative group flex flex-col justify-center items-center w-20"
        >
          <!-- Icon -->
          <a :href="skill.url" target="_blank">
            <img
              :src="`${getIconEndpoint + '/' + skill.icon_filename}`"
              :alt="skill.name"
              class="w-12 h-12 object-contain group-hover:cursor-pointer"
            />
          </a>

          <!-- Tooltip above -->
          <div
            class="absolute -top-7 opacity-0 group-hover:opacity-100 group-hover:-translate-y-1 transition-all duration-300 ease-out bg-gray-400 text-white text-xs px-4 py-1 rounded-md whitespace-nowrap pointer-events-none"
          >
            {{ skill.name }}
          </div>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup>
import { ref, computed, onMounted } from "vue";

const getSkillSetEndpoint =
  import.meta.env.VITE_API_URL + import.meta.env.VITE_GET_PROFILEDATA_ENDPOINT;
const getIconEndpoint =
  import.meta.env.VITE_API_URL + import.meta.env.VITE_ASSET_ICONS_ENDPOINT;

const skillSet = ref([
  {
    type_name: String,
    skill_items: [
      {
        name: String,
        icon_filename: String,
        is_most_used: Boolean,
        url: String,
      },
    ],
  },
]);

// Tabs: "Most Used" + categories
const categories = computed(() => [
  "Most Used",
  ...skillSet.value.map((s) => s.type_name),
]);

const activeTab = ref("Most Used");

const filteredSkills = computed(() => {
  if (activeTab.value === "Most Used") {
    return skillSet.value.flatMap((s) =>
      s.skill_items.filter((item) => item.is_most_used)
    );
  }
  const category = skillSet.value.find((s) => s.type_name === activeTab.value);
  return category ? category.skill_items : [];
});

async function getSkillSet() {
  try {
    const res = await fetch(getSkillSetEndpoint + "?skill_set=true");
    const resJson = await res.json();

    Object.assign(skillSet.value, resJson.data.skill_set);
  } catch (e) {
    console.error("failed to get data: " + e);
  }
}

onMounted(() => {
  getSkillSet();
});
</script>
