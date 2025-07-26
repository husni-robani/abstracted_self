<template>
  <div class="lg:max-w-5xl mx-auto flex items-center h-screen">
    <div class="container">
      <p class="-mb-1 text-lg font-roboto text-gray-400">hi, my name is</p>
      <h1 class="mb-2 text-5xl md:text-8xl font-roboto text-gray-700">
        {{ profile.name }}
      </h1>
      <span
        id="typed"
        class="text-2xl md:text-6xl font-semibold text-gray-300"
      ></span>
      <p
        class="w-full md:w-2/3 lg:w-1/2 mt-7 font-normal text-gray-500 text-base md:text-lg"
      >
        {{ profile.summary }}
      </p>
    </div>
  </div>
</template>

<script setup>
import { onMounted, ref } from "vue";

const api_endpoint =
  import.meta.env.VITE_API_URL + import.meta.env.VITE_GET_PROFILEDATA_ENDPOINT;

const taglines = [];

const profile = ref({
  name: "",
  summary: "",
});

onMounted(async () => {
  try {
    const response = await fetch(api_endpoint + "?name=true&summary=true");

    if (!response.ok) {
      throw new Error(
        `Failed to fetch name and summary profile. Status: ${response.status}`
      );
    }

    const jsonResponse = await response.json();
    profile.value.name = jsonResponse.data.name;
    profile.value.summary = jsonResponse.data.summary;
  } catch (e) {
    console.log(e);
  }
});
</script>
