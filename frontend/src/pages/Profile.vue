<template>
  <ProfileLayout>
    <!-- Hero Section -->
    <Hero></Hero>
    <!-- Bio Section -->
    <Bio id="about"></Bio>
    <!-- Project Section -->
    <Project id="projects"></Project>
    <!-- Contact Section -->
    <Contact id="contact"></Contact>
  </ProfileLayout>
</template>

<script setup>
import ProfileLayout from "../layouts/ProfileLayout.vue";
import Hero from "./partials/Hero.vue";
import Bio from "./partials/Bio.vue";
import Contact from "./partials/Contact.vue";
import Project from "./partials/Project.vue";
</script>
