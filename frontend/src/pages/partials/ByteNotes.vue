<template>
  <div class="flex flex-col lg:max-w-4xl mx-auto mt-52">
    <!-- Page title -->
    <div class="w-full h-5 border-b-2 border-gray-400 text-2xl mb-16">
      <span
        class="bg-white pr-4 font-bold text-gray-700 font-roboto text-2xl md:text-4xl"
      >
        My Byte Notes
      </span>
    </div>

    <!-- Blog cards -->
    <div class="overflow-x-auto">
      <div class="flex space-x-4 py-4">
        <a
          v-for="(blog, index) in blogs"
          :key="blog.id"
          :href="blog.url"
          target="_blank"
          class="flex-shrink-0 w-80 bg-white rounded-2xl shadow border border-gray-300 hover:shadow-lg transition-transform hover:-translate-y-1 duration-200 flex flex-col"
        >
          <!-- Image -->
          <img
            :src="imageAssetEndpoint + '/' + blog.image"
            alt="Blog image"
            class="w-full h-44 object-cover rounded-t-2xl"
          />

          <div class="p-4 flex flex-col flex-grow">
            <!-- Title -->
            <h2 class="text-lg font-semibold mb-2 line-clamp-2">
              {{ blog.title }}
            </h2>

            <!-- Description -->
            <p class="text-sm text-gray-500 mb-4 line-clamp-4 flex-grow">
              {{ blog.blog_snippet }}
            </p>

            <!-- Author -->
            <div class="flex items-center space-x-2 mt-auto">
              <img
                src="/images/author_image.jpg"
                alt="Author"
                class="w-8 h-8 rounded-full object-cover"
              />
              <div class="text-sm text-gray-600">
                <p class="font-medium">Husni Robani</p>
                <p class="text-xs text-gray-400">on Medium</p>
              </div>
            </div>
          </div>
        </a>
      </div>
    </div>
  </div>
</template>

<script setup>
import { onMounted, ref } from "vue";

const getBlogsEndpoint =
  import.meta.env.VITE_API_URL + import.meta.env.VITE_GET_BLOGS_ENDPOINT;
const imageAssetEndpoint =
  import.meta.env.VITE_API_URL + import.meta.env.VITE_ASSET_IMAGES_ENDPOINT;

const blogs = ref([]);

async function getBlogs() {
  try {
    const res = await fetch(getBlogsEndpoint, {
      method: "GET",
    });

    if (!res.ok) {
      throw new Error("failed to get blogs data");
    }

    const resJson = await res.json();

    blogs.value = resJson.data;
  } catch (e) {
    console.error(e);
  }
}

onMounted(async () => {
  await getBlogs();
});
</script>
