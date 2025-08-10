<template>
  <div class="min-h-screen">
    <!-- //navbar -->
    <nav
      ref="navbar"
      :class="[
        'h-16 top-0 left-0 right-0 fixed flex items-center justify-between px-20 py-5 md:flex-row-reverse p-3 z-20 bg-white transition-all duration-300',
        isHidden ? '-translate-y-full' : 'translate-y-0',
        hasScrolled
          ? 'shadow-md border-b border-gray-200'
          : 'shadow-none border-b-white',
      ]"
    >
      <!-- Right: Resume -->
      <div class="flex flex-row items-center gap-8 mx-auto md:m-0">
        <div class="flex flex-col items-center">
          <div class="peer text-gray-500 font-medium hover:text-gray-800">
            <a
              :href="asset_document_endpoint + '/' + resume_file_name"
              target="_blank"
              class="py-1 px-6"
              >Resume</a
            >
          </div>
          <div
            class="transition-transform duration-500 peer-hover:scale-150 w-20 border-b-2 border-solid border-gray-500 peer-hover:border-gray-800"
          ></div>
        </div>
      </div>

      <!-- Left: Menu Links -->
      <div class="flex gap-8">
        <router-link
          to="/profile"
          class="text-gray-600 hover:text-gray-900 font-medium transition-colors duration-200"
        >
          Profile
        </router-link>
        <router-link
          to="/notes"
          class="text-gray-600 hover:text-gray-900 font-medium transition-colors duration-200"
        >
          My Notes
        </router-link>
      </div>
    </nav>
    <!-- //content -->
    <div class="relative lg:flex z-0 h-full">
      <!-- App -->
      <div class="w-11/12 mx-auto lg:w-10/12 p-5">
        <main><slot></slot></main>
      </div>
    </div>
    <!-- //footer -->
    <div class="flex flex-col max-w-4xl mx-auto m-5">
      <!-- link -->
      <div class="mx-auto flex flex-row lg:hidden">
        <!-- LinkedIn -->
        <a
          href="https://www.linkedin.com/in/husni-robani-706a78225/"
          target="_blank"
          class="mx-2"
        >
          <svg
            class="fill-gray-500 hover:fill-gray-800"
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 30 30"
            width="30px"
            height="30px"
          >
            <path
              d="M24,4H6C4.895,4,4,4.895,4,6v18c0,1.105,0.895,2,2,2h18c1.105,0,2-0.895,2-2V6C26,4.895,25.105,4,24,4z M10.954,22h-2.95 v-9.492h2.95V22z M9.449,11.151c-0.951,0-1.72-0.771-1.72-1.72c0-0.949,0.77-1.719,1.72-1.719c0.948,0,1.719,0.771,1.719,1.719 C11.168,10.38,10.397,11.151,9.449,11.151z M22.004,22h-2.948v-4.616c0-1.101-0.02-2.517-1.533-2.517 c-1.535,0-1.771,1.199-1.771,2.437V22h-2.948v-9.492h2.83v1.297h0.04c0.394-0.746,1.356-1.533,2.791-1.533 c2.987,0,3.539,1.966,3.539,4.522V22z"
            />
          </svg>
        </a>
        <!-- GitHub -->
        <a href="https://github.com/husni-robani" target="_blank" class="mx-2">
          <svg
            class="fill-gray-500 hover:fill-gray-800"
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 30 30"
            width="30px"
            height="30px"
          >
            <path
              d="M15,3C8.373,3,3,8.373,3,15c0,5.623,3.872,10.328,9.092,11.63C12.036,26.468,12,26.28,12,26.047v-2.051 c-0.487,0-1.303,0-1.508,0c-0.821,0-1.551-0.353-1.905-1.009c-0.393-0.729-0.461-1.844-1.435-2.526 c-0.289-0.227-0.069-0.486,0.264-0.451c0.615,0.174,1.125,0.596,1.605,1.222c0.478,0.627,0.703,0.769,1.596,0.769 c0.433,0,1.081-0.025,1.691-0.121c0.328-0.833,0.895-1.6,1.588-1.962c-3.996-0.411-5.903-2.399-5.903-5.098 c0-1.162,0.495-2.286,1.336-3.233C9.053,10.647,8.706,8.73,9.435,8c1.798,0,2.885,1.166,3.146,1.481C13.477,9.174,14.461,9,15.495,9 c1.036,0,2.024,0.174,2.922,0.483C18.675,9.17,19.763,8,21.565,8c0.732,0.731,0.381,2.656,0.102,3.594 c0.836,0.945,1.328,2.066,1.328,3.226c0,2.697-1.904,4.684-5.894,5.097C18.199,20.49,19,22.1,19,23.313v2.734 c0,0.104-0.023,0.179-0.035,0.268C23.641,24.676,27,20.236,27,15C27,8.373,21.627,3,15,3z"
            />
          </svg>
        </a>
        <!-- instagram -->
        <a
          href="https://www.instagram.com/husnirb/"
          target="_blank"
          class="mx-2"
        >
          <svg
            class="fill-gray-500 hover:fill-gray-800"
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 50 50"
            width="28px"
            height="28px"
          >
            <path
              d="M 16 3 C 8.83 3 3 8.83 3 16 L 3 34 C 3 41.17 8.83 47 16 47 L 34 47 C 41.17 47 47 41.17 47 34 L 47 16 C 47 8.83 41.17 3 34 3 L 16 3 z M 37 11 C 38.1 11 39 11.9 39 13 C 39 14.1 38.1 15 37 15 C 35.9 15 35 14.1 35 13 C 35 11.9 35.9 11 37 11 z M 25 14 C 31.07 14 36 18.93 36 25 C 36 31.07 31.07 36 25 36 C 18.93 36 14 31.07 14 25 C 14 18.93 18.93 14 25 14 z M 25 16 C 20.04 16 16 20.04 16 25 C 16 29.96 20.04 34 25 34 C 29.96 34 34 29.96 34 25 C 34 20.04 29.96 16 25 16 z"
            />
          </svg>
        </a>
      </div>
      <!-- cc -->
      <div class="mx-auto text-gray-500">
        <p>
          handcrafted by me inspired by
          <a
            href="https://brittanychiang.com"
            target="_blank"
            class="hover:text-gray-900"
            >Brittany Chiang</a
          >
          ©️2023
        </p>
      </div>
    </div>
    <!-- Side content 1 -->
    <div
      class="hidden lg:flex lg:flex-col justify-end fixed top-1/2 bottom-0 left-0 items-center w-1/12"
    >
      <div>
        <!-- GitHub -->
        <a href="https://github.com/husni-robani" target="_blank"
          ><svg
            class="transition ease-in-out hover:scale-110 duration-300 fill-gray-800 hover:fill-gray-500"
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 30 30"
            width="30px"
            height="30px"
          >
            <path
              d="M15,3C8.373,3,3,8.373,3,15c0,5.623,3.872,10.328,9.092,11.63C12.036,26.468,12,26.28,12,26.047v-2.051 c-0.487,0-1.303,0-1.508,0c-0.821,0-1.551-0.353-1.905-1.009c-0.393-0.729-0.461-1.844-1.435-2.526 c-0.289-0.227-0.069-0.486,0.264-0.451c0.615,0.174,1.125,0.596,1.605,1.222c0.478,0.627,0.703,0.769,1.596,0.769 c0.433,0,1.081-0.025,1.691-0.121c0.328-0.833,0.895-1.6,1.588-1.962c-3.996-0.411-5.903-2.399-5.903-5.098 c0-1.162,0.495-2.286,1.336-3.233C9.053,10.647,8.706,8.73,9.435,8c1.798,0,2.885,1.166,3.146,1.481C13.477,9.174,14.461,9,15.495,9 c1.036,0,2.024,0.174,2.922,0.483C18.675,9.17,19.763,8,21.565,8c0.732,0.731,0.381,2.656,0.102,3.594 c0.836,0.945,1.328,2.066,1.328,3.226c0,2.697-1.904,4.684-5.894,5.097C18.199,20.49,19,22.1,19,23.313v2.734 c0,0.104-0.023,0.179-0.035,0.268C23.641,24.676,27,20.236,27,15C27,8.373,21.627,3,15,3z"
            /></svg
        ></a>
        <!-- LinkedIn -->
        <a
          href="https://www.linkedin.com/in/husni-robani-706a78225/"
          target="_blank"
        >
          <svg
            class="transition ease-in-out hover:scale-110 duration-300 fill-gray-800 hover:fill-gray-500 my-2"
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 30 30"
            width="30px"
            height="30px"
          >
            <path
              d="M24,4H6C4.895,4,4,4.895,4,6v18c0,1.105,0.895,2,2,2h18c1.105,0,2-0.895,2-2V6C26,4.895,25.105,4,24,4z M10.954,22h-2.95 v-9.492h2.95V22z M9.449,11.151c-0.951,0-1.72-0.771-1.72-1.72c0-0.949,0.77-1.719,1.72-1.719c0.948,0,1.719,0.771,1.719,1.719 C11.168,10.38,10.397,11.151,9.449,11.151z M22.004,22h-2.948v-4.616c0-1.101-0.02-2.517-1.533-2.517 c-1.535,0-1.771,1.199-1.771,2.437V22h-2.948v-9.492h2.83v1.297h0.04c0.394-0.746,1.356-1.533,2.791-1.533 c2.987,0,3.539,1.966,3.539,4.522V22z"
            />
          </svg>
        </a>
        <a href="https://www.instagram.com/husnirb/" target="_blank">
          <svg
            class="transition ease-in-out hover:scale-110 duration-300 fill-gray-800 hover:fill-gray-500 my-2"
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 50 50"
            width="30px"
            height="30px"
          >
            <path
              d="M 16 3 C 8.83 3 3 8.83 3 16 L 3 34 C 3 41.17 8.83 47 16 47 L 34 47 C 41.17 47 47 41.17 47 34 L 47 16 C 47 8.83 41.17 3 34 3 L 16 3 z M 37 11 C 38.1 11 39 11.9 39 13 C 39 14.1 38.1 15 37 15 C 35.9 15 35 14.1 35 13 C 35 11.9 35.9 11 37 11 z M 25 14 C 31.07 14 36 18.93 36 25 C 36 31.07 31.07 36 25 36 C 18.93 36 14 31.07 14 25 C 14 18.93 18.93 14 25 14 z M 25 16 C 20.04 16 16 20.04 16 25 C 16 29.96 20.04 34 25 34 C 29.96 34 34 29.96 34 25 C 34 20.04 29.96 16 25 16 z"
            />
          </svg>
        </a>
      </div>
      <div class="h-24 border-l-2 border-solid border-gray-800 mt-7"></div>
    </div>
    <!-- side content 2 -->
    <div
      class="hidden lg:flex lg:flex-col justify-end fixed top-1/2 bottom-0 right-0 items-center w-1/12"
    >
      <div class="rotate-90 m-16">
        <a
          href="mailto:husnir2005@gmail.com"
          target="_blank"
          class="text-black hover:text-gray-500 font-semibold"
          >husnir2005@gmail.com</a
        >
      </div>
      <div class="h-24 border-l-2 border-solid border-gray-800 mt-7 mr-1"></div>
    </div>
  </div>
</template>
<script setup>
import { onMounted, onBeforeUnmount, ref } from "vue";

const asset_document_endpoint =
  import.meta.env.VITE_API_URL + import.meta.env.VITE_ASSET_DOCUMENTS_ENDPOINT;

const api_endpoint =
  import.meta.env.VITE_API_URL + import.meta.env.VITE_GET_PROFILEDATA_ENDPOINT;

const resume_file_name = ref("");
const isHidden = ref(false);
const hasScrolled = ref(false);
let lastScrollTop = 0;

onMounted(async () => {
  // Fetch resume
  try {
    const response = await fetch(api_endpoint + "?resume=true");
    if (!response.ok) throw new Error(`Failed: ${response.status}`);
    const jsonResponse = await response.json();
    resume_file_name.value = jsonResponse.data.resume_file_name;
  } catch (e) {
    console.log(e);
  }

  // Scroll listener
  window.addEventListener("scroll", handleScroll);
});

onBeforeUnmount(() => {
  window.removeEventListener("scroll", handleScroll);
});

function handleScroll() {
  const currentScroll =
    window.pageYOffset || document.documentElement.scrollTop;

  // Show shadow + border after scrolling down 1px
  hasScrolled.value = currentScroll > 0;

  // Hide navbar when scrolling down
  if (currentScroll > lastScrollTop && currentScroll > 50) {
    isHidden.value = true;
  } else {
    isHidden.value = false;
  }

  lastScrollTop = currentScroll <= 0 ? 0 : currentScroll;
}
</script>
