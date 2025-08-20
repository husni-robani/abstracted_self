<template>
  <div class="lg:max-w-4xl mx-auto mt-52">
    <div class="w-full lg:max-w-md h-5 border-b-2 border-gray-400 mb-6">
      <span
        class="bg-white pr-4 font-bold text-gray-700 font-roboto text-2xl md:text-4xl"
        >My Bio</span
      >
    </div>
    <!-- my bio content -->
    <div class="lg:flex">
      <div
        v-html="bio"
        class="lg:mr-20 lg:w-1/2 text-lg text-justify md:text-left"
      ></div>
      <!-- bio image -->
      <div class="mx-auto mt-6 lg:mt-0">
        <div
          class="flex relative mx-auto w-60 h-60 sm:w-80 sm:h-80 border-2 border-black rounded-md"
        >
          <img
            src="/images/profile.jpg"
            alt="profile"
            class="object object-scale-down h-60 w-60 absolute transition-transform hover:-translate-x-1 hover:-translate-y-1 -top-2 -left-[6px] sm:-top-3 sm:-left-[10px] aspect-square sm:w-80 sm:h-80 rounded-md"
          />
        </div>
      </div>
    </div>
    <!-- skill set -->
    <TechStack />
  </div>
</template>

<script setup>
import { onMounted, ref } from "vue";
import TechStack from "./TechStack.vue";

const bio = ref("");

const api_endpoint =
  import.meta.env.VITE_API_URL + import.meta.env.VITE_GET_PROFILEDATA_ENDPOINT;

onMounted(async () => {
  try {
    const response = await fetch(api_endpoint + "?cache=true&bio=true");

    if (!response.ok) {
      throw new Error(
        `Failed to fetch bio profile. Status: ${response.status}`
      );
    }

    const jsonResponse = await response.json();
    bio.value = jsonResponse.data.bio;
  } catch (e) {
    console.log(e);
  }
});
</script>
