<template>
  <ProfileLayout>
    <!-- Hero Section -->
    <Hero></Hero>
    <!-- Bio Section -->
    <Bio id="about"></Bio>
    <!-- Experiences Section -->
    <Experiences id="experiences" />
    <!-- Project Section -->
    <Project id="projects"></Project>
    <!-- ByteNotes section -->
    <ByteNotes id="bytenotes" />
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
import Experiences from "./partials/Experiences.vue";
import ByteNotes from "./partials/ByteNotes.vue";
</script>
