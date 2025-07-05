<template>
  <li
    :class="[
      'flex flex-col justicy items-center w-full lg:items-start lg:w-full',
      isEven ? 'lg:flex-row' : 'lg:flex-row-reverse',
    ]"
  >
    <!-- image -->
    <div class="w-full max-w-xl">
      <div
        class="bg-cover bg-center aspect-w-16 aspect-h-9 rounded-2xl border"
        :style="{ backgroundImage: `url(${props.image_url})` }"
      ></div>
    </div>
    <!-- content -->
    <div
      :class="[
        'w-full max-w-2xl mt-2 lg:mt-0 lg:w-1/2 lg:absolute bg-transparent',
        isEven ? 'right-0 lg:text-right' : 'left-0 lg:text-left',
      ]"
    >
      <!-- title -->
      <div
        :class="['flex flex-col', isEven ? 'lg:items-end' : 'lg:items-start']"
      >
        <p class="text-sm md:text-sm font-roboto font-extralight text-gray-400">
          featured project
        </p>
        <h1 class="text-lg md:text-xl font-bold lg:w-1/2">
          {{ props.project_name }}
        </h1>
      </div>
      <!-- description -->
      <div
        class="mt-2 lg:bg-gray-100 lg:p-2 lg:rounded-md lg:shadow-md lg:mt-4 text-base"
      >
        <p>
          {{ props.description }}
        </p>
      </div>
      <!-- stack -->
      <div
        :class="[
          'flex flex-row gap-2 bg-gray-100 p-2 rounded-md shadow-sm lg:bg-transparent lg:p-0 lg:shadow-none lg:rounded-none mt-2 md:mt-4',
          isEven ? 'lg:flex-row-reverse' : 'lg:flex-row',
        ]"
      >
        <span
          :class="[
            'flex flex-wrap flex-row gap-2 lg:w-2/3 text-sm font-roboto font-extralight text-gray-400',
            isEven ? 'justify-end' : 'justify-start',
          ]"
        >
          <span v-for="tech in props.tech_stack">{{ tech }}</span>
        </span>
      </div>
      <!-- link -->
      <div
        :class="[
          'flex  gap-2 mt-2 lg:mt-4',
          isEven ? 'lg:flex-row-reverse' : 'lg:flex-row',
        ]"
      >
        <a v-for="url in props.source_url" :href="url" target="_blank">
          <GithubIcon />
        </a>
      </div>
    </div>
  </li>
</template>

<script setup>
import GithubIcon from "../assets/GithubIcon.vue";
const props = defineProps({
  isEven: {
    type: Boolean,
    required: true,
  },
  project_name: {
    type: String,
    required: true,
  },
  description: {
    type: String,
    required: true,
  },
  tech_stack: {
    type: Array,
    required: true,
  },
  source_url: {
    type: Array,
    required: true,
  },
  project_url: {
    type: String,
    required: true,
  },
  image_url: {
    type: String,
    required: true,
  },
});
</script>
