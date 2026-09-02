<template>
  <div v-if="blogs.length > 0" class="flex flex-col lg:max-w-4xl mx-auto mt-52">
    <!-- signature section title -->
    <div class="w-full h-5 border-b-2 border-gray-400 text-2xl mb-16">
      <span
        class="bg-white pr-4 font-bold text-gray-700 font-roboto text-2xl md:text-4xl"
      >
        My Byte Notes
      </span>
    </div>

    <p class="font-mono text-[11px] uppercase tracking-widest text-gray-400 mb-2">
      journal rows — 3 newest notes
    </p>

    <!-- Journal rows: 3 newest posts -->
    <div class="flex flex-col border-t border-gray-200">
      <router-link
        v-for="blog in blogs"
        :key="blog.id"
        :to="`/blog/${blog.id}`"
        class="group border-b border-gray-200 py-6 md:py-7 flex items-center gap-5 md:gap-7 hover:border-gray-300 transition-colors"
      >
        <!-- cover thumb OR fallback -->
        <img
          v-if="blog.cover_image"
          :src="resolveCoverUrl(blog.cover_image)"
          alt=""
          class="flex-none h-16 w-24 md:h-20 md:w-32 rounded-md object-cover"
        />
        <div
          v-else
          class="flex-none h-16 w-24 md:h-20 md:w-32 rounded-md bg-gray-100 flex items-center justify-center text-gray-300"
        >
          <BookOpenIcon class="w-6 h-6" />
        </div>

        <div class="flex-1 min-w-0">
          <p
            class="font-mono text-[11px] uppercase tracking-widest text-gray-400 mb-2"
          >
            {{ formatMonthYear(blog.created_at) }}
          </p>
          <h3
            class="font-roboto font-semibold text-xl md:text-2xl text-gray-900 group-hover:text-gray-600 transition-colors"
          >
            {{ blog.title }}
          </h3>
          <p
            v-if="blog.blog_snippet"
            class="text-sm text-gray-500 mt-2 line-clamp-1 md:line-clamp-2"
          >
            {{ blog.blog_snippet }}
          </p>
        </div>

        <!-- hover arrow (desktop only) -->
        <svg
          class="hidden md:block flex-none w-5 h-5 text-gray-300 group-hover:text-gray-900 group-hover:translate-x-1 transition-all"
          viewBox="0 0 24 24"
          fill="none"
          aria-hidden="true"
        >
          <path
            d="M5 12h14M13 6l6 6-6 6"
            stroke="currentColor"
            stroke-width="1.75"
            stroke-linecap="round"
            stroke-linejoin="round"
          />
        </svg>
      </router-link>
    </div>

    <!-- Show More -->
    <div class="mt-10 md:mt-12 flex justify-center">
      <router-link
        to="/blog"
        class="group inline-flex items-center gap-2 border border-gray-900 text-gray-900 font-mono text-sm px-6 py-2 rounded-full hover:bg-gray-900 hover:text-white transition-colors"
      >
        Show More
        <svg
          width="16"
          height="16"
          viewBox="0 0 24 24"
          fill="none"
          class="transition-transform group-hover:translate-x-1"
        >
          <path
            d="M5 12h14M13 6l6 6-6 6"
            stroke="currentColor"
            stroke-width="1.75"
            stroke-linecap="round"
            stroke-linejoin="round"
          />
        </svg>
      </router-link>
    </div>
  </div>
</template>

<script setup>
import { onMounted, ref } from "vue";
import { BookOpenIcon } from "@heroicons/vue/24/outline";

const VITE_API_URL = import.meta.env.VITE_API_URL || "";

const getBlogsEndpoint =
  import.meta.env.VITE_API_URL + import.meta.env.VITE_GET_BLOGS_ENDPOINT;

const blogs = ref([]);

// cover_image is an Image|{ url } object (default palette fallback below).
function resolveCoverUrl(image) {
  if (!image || !image.url) return "";
  return VITE_API_URL + image.url;
}

// Journal micro-meta date, e.g. "AUG 2026".
function formatMonthYear(dateStr) {
  return new Intl.DateTimeFormat("en-US", {
    month: "short",
    year: "numeric",
  })
    .format(new Date(dateStr))
    .toUpperCase();
}

async function getBlogs() {
  try {
    const res = await fetch(getBlogsEndpoint + "?published=true");

    if (!res.ok) {
      throw new Error("failed to get blogs data");
    }

    const resJson = await res.json();

    const data = resJson.data || [];
    // Newest first by created_at (AD-3), then only the 3 newest show here.
    blogs.value = [...data]
      .sort((a, b) => new Date(b.created_at) - new Date(a.created_at))
      .slice(0, 3);
  } catch (e) {
    // Public section: swallow errors quietly per the repo convention.
    console.error(e);
  }
}

onMounted(async () => {
  await getBlogs();
});
</script>
