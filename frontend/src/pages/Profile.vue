<template>
  <ProfileLayout>
    <!-- Hero Section -->
    <Hero></Hero>
    <!-- Bio Section -->
    <Bio></Bio>
  </ProfileLayout>
</template>

<script setup>
import ProfileLayout from "../layouts/ProfileLayout.vue";
import Hero from "./partials/Hero.vue";
import Bio from "./partials/Bio.vue";
</script>
