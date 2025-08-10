<template>
  <li
    :class="[
      'flex flex-col justify items-center w-full lg:items-start lg:w-full',
      isEven ? 'lg:flex-row' : 'lg:flex-row-reverse',
    ]"
  >
    <!-- Image -->
    <div
      class="w-full max-w-xl relative overflow-hidden rounded-2xl border border-gray-200 aspect-[16/9]"
    >
      <ImageSlider
        :image_urls="props.image_urls"
        :interval="5000"
        :transitionTime="1000"
      />
    </div>

    <!-- content -->
    <div
      :class="[
        'w-full max-w-2xl mt-2 lg:mt-0 lg:w-1/2 lg:absolute bg-transparent',
        isEven ? 'right-0 lg:text-right' : 'left-0 lg:text-left',
      ]"
    >
      <div
        :class="['flex flex-col', isEven ? 'lg:items-end' : 'lg:items-start']"
      >
        <p class="text-sm font-roboto font-extralight text-gray-400">
          featured project
        </p>
        <h1 class="text-lg md:text-xl font-bold lg:w-1/2">
          {{ props.project_name }}
        </h1>
      </div>

      <div
        class="mt-2 lg:bg-gray-100 lg:p-2 lg:rounded-md lg:shadow-md lg:mt-4 text-base"
      >
        <p class="line-clamp-6">{{ props.description }}</p>
      </div>

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
          <span v-for="tech in props.tech_stack" :key="tech">{{ tech }}</span>
        </span>
      </div>

      <div
        :class="[
          'flex gap-2 mt-2 lg:mt-4',
          isEven ? 'lg:flex-row-reverse' : 'lg:flex-row',
        ]"
      >
        <a v-show="props.project_url !== ''">
          <ArrowTopRightOnSquareIcon
            class="text-gray-800 hover:text-gray-500 hover:cursor-pointer w-[30px] h-[30px]"
          />
        </a>
        <a
          v-for="(url, idx) in props.source_url"
          :key="idx"
          :href="url"
          target="_blank"
        >
          <GithubIcon />
        </a>
      </div>
    </div>
  </li>
</template>

<script setup>
import {
  ArrowTopRightOnSquareIcon,
  ComputerDesktopIcon,
} from "@heroicons/vue/24/outline";
import GithubIcon from "../assets/GithubIcon.vue";
import ImageSlider from "./ImageSlider.vue";

const props = defineProps({
  isEven: Boolean,
  project_name: String,
  description: String,
  tech_stack: Array,
  source_url: Array,
  project_url: String,
  image_urls: Array,
});
</script>
