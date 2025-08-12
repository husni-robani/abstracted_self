<template>
  <nav
    ref="navbar"
    :class="[
      'h-20 top-0 left-0 right-0 fixed flex items-center justify-between md:flex-row-reverse px-20 py-5 z-20 bg-white transition-all duration-300',
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
      <a
        href="#about"
        :class="[
          'font-medium transition-colors duration-200',
          activeSection === 'about'
            ? 'text-gray-500'
            : 'text-gray-300 hover:text-gray-900',
        ]"
      >
        About
      </a>
      <a
        href="#experiences"
        :class="[
          'font-medium transition-colors duration-200',
          activeSection === 'experiences'
            ? 'text-gray-500'
            : 'text-gray-300 hover:text-gray-900',
        ]"
      >
        Experiences
      </a>
      <a
        href="#projects"
        :class="[
          'font-medium transition-colors duration-200',
          activeSection === 'projects'
            ? 'text-gray-500'
            : 'text-gray-300 hover:text-gray-900',
        ]"
      >
        Projects
      </a>
      <a
        href="#contact"
        :class="[
          'font-medium transition-colors duration-200',
          activeSection === 'contact'
            ? 'text-gray-500'
            : 'text-gray-300 hover:text-gray-900',
        ]"
      >
        Contact
      </a>
    </div>
  </nav>
</template>

<script setup>
import { onMounted, onBeforeUnmount, ref } from "vue";

const asset_document_endpoint =
  import.meta.env.VITE_API_URL + import.meta.env.VITE_ASSET_DOCUMENTS_ENDPOINT;
const api_endpoint =
  import.meta.env.VITE_API_URL + import.meta.env.VITE_GET_PROFILEDATA_ENDPOINT;

const sectionIds = ["about", "projects", "contact"];
const activeSection = ref("");
const resume_file_name = ref("");
const isHidden = ref(false);
const hasScrolled = ref(false);
let lastScrollTop = 0;

function updateActiveSection() {
  const scrollPosition = window.scrollY + 100; // offset for navbar
  for (let id of sectionIds) {
    const el = document.getElementById(id);
    if (el) {
      const sectionTop = el.offsetTop;
      const sectionBottom = sectionTop + el.offsetHeight;
      if (scrollPosition >= sectionTop && scrollPosition < sectionBottom) {
        activeSection.value = id;
        break;
      }
    }
  }
}

onMounted(() => {
  window.addEventListener("scroll", updateActiveSection);
  updateActiveSection();
});

onBeforeUnmount(() => {
  window.removeEventListener("scroll", updateActiveSection);
});

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
  hasScrolled.value = currentScroll > 0;

  if (currentScroll > lastScrollTop && currentScroll > 50) {
    isHidden.value = true;
  } else {
    isHidden.value = false;
  }

  lastScrollTop = currentScroll <= 0 ? 0 : currentScroll;
}
</script>
