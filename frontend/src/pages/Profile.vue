<template>
  <ProfileLayout>
    <!-- Hero Section -->
    <Hero></Hero>
    <!-- Bio Section -->
    <Bio></Bio>
    <Contact></Contact>
  </ProfileLayout>
</template>

<script setup>
import ProfileLayout from "../layouts/ProfileLayout.vue";
import Hero from "./partials/Hero.vue";
import Bio from "./partials/Bio.vue";
import Contact from "./partials/Contact.vue";
</script>
