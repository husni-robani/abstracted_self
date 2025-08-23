<template>
  <nav
    ref="navbar"
    :class="[
      'h-20 top-0 left-0 right-0 fixed flex items-center justify-between md:flex-row-reverse px-6 md:px-20 py-5 z-30 bg-white transition-all duration-300',
      isHidden ? '-translate-y-full' : 'translate-y-0',
      hasScrolled
        ? 'shadow-md border-b border-gray-200'
        : 'shadow-none border-b-white',
    ]"
  >
    <!-- Right: Resume (desktop only) -->
    <div class="hidden md:flex flex-row items-center gap-8 mx-auto md:m-0">
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

    <!-- Left: Menu Links (desktop only) -->
    <div class="hidden md:flex gap-8">
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
        href="#bytenotes"
        :class="[
          'font-medium transition-colors duration-200',
          activeSection === 'bytenotes'
            ? 'text-gray-500'
            : 'text-gray-300 hover:text-gray-900',
        ]"
      >
        ByteNotes
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

    <!-- Mobile: Hamburger Button -->
    <button
      @click="openSidebar"
      class="md:hidden text-gray-700 focus:outline-none"
    >
      <svg
        xmlns="http://www.w3.org/2000/svg"
        class="h-6 w-6"
        fill="none"
        viewBox="0 0 24 24"
        stroke="currentColor"
      >
        <path
          stroke-linecap="round"
          stroke-linejoin="round"
          stroke-width="2"
          d="M4 6h16M4 12h16M4 18h16"
        />
      </svg>
    </button>
  </nav>

  <!-- Sidebar (Mobile Drawer) -->
  <div
    class="fixed top-0 left-0 w-64 h-full bg-white shadow-lg transform transition-transform duration-300 z-40 md:hidden flex flex-col"
    :class="isSidebarOpen ? 'translate-x-0' : '-translate-x-full'"
  >
    <!-- Close Button -->
    <div class="flex justify-end p-4">
      <button @click="closeSidebar" class="text-gray-700">
        <svg
          xmlns="http://www.w3.org/2000/svg"
          class="h-6 w-6"
          fill="none"
          viewBox="0 0 24 24"
          stroke="currentColor"
        >
          <path
            stroke-linecap="round"
            stroke-linejoin="round"
            stroke-width="2"
            d="M6 18L18 6M6 6l12 12"
          />
        </svg>
      </button>
    </div>

    <!-- Sidebar Links -->
    <div class="flex flex-col gap-6 p-6">
      <a href="#about" @click="closeSidebar">About</a>
      <a href="#experiences" @click="closeSidebar">Experiences</a>
      <a href="#projects" @click="closeSidebar">Projects</a>
      <a href="#bytenotes" @click="closeSidebar">ByteNotes</a>
      <a href="#contact" @click="closeSidebar">Contact</a>

      <!-- Resume (mobile sidebar) -->
      <a
        :href="asset_document_endpoint + '/' + resume_file_name"
        target="_blank"
        class="text-gray-700 font-medium border-t pt-4 mt-4"
        @click="closeSidebar"
      >
        Resume
      </a>
    </div>
  </div>

  <!-- Overlay -->
  <div
    v-if="isSidebarOpen"
    @click="closeSidebar"
    class="fixed inset-0 bg-black/60 md:hidden z-30"
  ></div>
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
const isSidebarOpen = ref(false);
let lastScrollTop = 0;

function openSidebar() {
  isSidebarOpen.value = true;
  document.body.classList.add("overflow-hidden"); // disable scroll
}

function closeSidebar() {
  isSidebarOpen.value = false;
  document.body.classList.remove("overflow-hidden"); // re-enable scroll
}

function updateActiveSection() {
  const scrollPosition = window.scrollY + 100;
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
  try {
    const response = await fetch(api_endpoint + "?resume=true");
    if (!response.ok) throw new Error(`Failed: ${response.status}`);
    const jsonResponse = await response.json();
    resume_file_name.value = jsonResponse.data.resume_file_name;
  } catch (e) {
    console.log(e);
  }

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
