<template>
  <div class="min-h-screen bg-black flex items-center justify-center px-4">
    <div
      class="w-full max-w-md bg-gray-100 border border-gray-300 rounded-2xl px-8 py-10 shadow-xl"
    >
      <div class="mb-8 text-center">
        <h1 class="text-3xl font-mono font-semibold text-gray-800">
          Husni's Console
        </h1>
        <p class="text-gray-500 text-sm mt-2">Access requires key</p>
      </div>

      <form @submit.prevent="login" class="space-y-6">
        <div>
          <label class="block text-sm text-gray-700 mb-1">Access Key</label>
          <input
            v-model="accessKey"
            type="password"
            placeholder="Your secret key"
            class="w-full px-4 py-2 rounded-lg bg-white text-gray-800 border border-gray-300 focus:outline-none focus:ring-2 focus:ring-gray-600"
          />
        </div>

        <button
          type="submit"
          class="w-full py-2 px-4 rounded-lg bg-gray-900 text-white font-semibold hover:bg-gray-800 transition duration-200"
        >
          Unlock
        </button>

        <p v-if="error" class="text-red-500 text-center text-sm mt-2">
          {{ error }}
        </p>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { useRouter } from "vue-router";

const accessKey = ref("");
const error = ref("");
const router = useRouter();

const authAPI =
  import.meta.env.VITE_API_URL + import.meta.env.VITE_LOGIN_ENDPOINT;

const login = async () => {
  try {
    const response = await fetch(authAPI, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify({
        access_key: accessKey.value,
      }),
    });

    let jsonResponse = {};

    try {
      jsonResponse = await response.json(); // attempt to parse response (even if it's an error)
    } catch (parseErr) {
      throw new Error("invalid-json");
    }

    // failed
    if (!response.ok) {
      if (response.status === 401) {
        throw new Error(jsonResponse.errors.access_key || "invalid credential"); // invalid password
      } else if (response.status === 500) {
        throw new Error(jsonResponse.message || "server error"); // internal server error
      } else {
        throw new Error(jsonResponse.message); // other errors
      }
    }

    // success
    const token = jsonResponse.data.token;
    localStorage.setItem("token", token);
    router.push("/testuyy");
  } catch (err) {
    error.value =
      err.message === "invalid-json"
        ? "Server response is invalid."
        : err.message;
  }
};
</script>
