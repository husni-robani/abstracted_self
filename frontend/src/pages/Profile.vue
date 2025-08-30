<template>
  <div>
    <!-- Preloader -->
    <transition name="fade-bg">
      <div
        v-if="loading"
        class="fixed inset-0 bg-white flex items-center justify-center z-50"
      >
        <transition name="fade-text">
          <!-- <h1 v-if="showText" class="loader text-4xl font-bold text-gray-500">
            Welcome.
          </h1> -->
          <div v-if="showText" class="loader"></div>
        </transition>
      </div>
    </transition>

    <!-- Main Profile Page -->
    <ProfileLayout v-show="!loading">
      <Hero />
      <Bio id="about" />
      <Experiences id="experiences" />
      <Project id="projects" />
      <ByteNotes id="bytenotes" />
      <Contact id="contact" />
    </ProfileLayout>
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";
import ProfileLayout from "../layouts/ProfileLayout.vue";
import Hero from "./partials/Hero.vue";
import Bio from "./partials/Bio.vue";
import Contact from "./partials/Contact.vue";
import Project from "./partials/Project.vue";
import Experiences from "./partials/Experiences.vue";
import ByteNotes from "./partials/ByteNotes.vue";

const loading = ref(true);
const showText = ref(true);

onMounted(() => {
  // Step 1: Fade out the pulsing text
  setTimeout(() => {
    showText.value = false;
  }, 1000);

  // Step 2: Fade out the white background
  setTimeout(() => {
    loading.value = false;
  }, 2000);
});
</script>

<style>
/* Text fade */
.fade-text-enter-active,
.fade-text-leave-active {
  transition: opacity 1s ease;
}
.fade-text-enter-from,
.fade-text-leave-to {
  opacity: 0;
}

/* Background fade */
.fade-bg-enter-active,
.fade-bg-leave-active {
  transition: opacity 0.8s ease;
}
.fade-bg-enter-from,
.fade-bg-leave-to {
  opacity: 0;
}

/* HTML: <div class="loader"></div> */
.loader {
  width: fit-content;
  font-weight: bold;
  font-family: monospace;
  white-space: pre;
  font-size: 30px;
  line-height: 1.2em;
  height: 1.2em;
  overflow: hidden;
}
.loader:before {
  content: "Welcome.\ASelamat datang.\A Bienvenido.\A Ahlan wa sahlan.\A Ciao\A Bienvenue.";
  white-space: pre;
  display: inline-block;
  animation: l38 1s infinite steps(6);
}

@keyframes l38 {
  100% {
    transform: translateY(-100%);
  }
}
</style>
