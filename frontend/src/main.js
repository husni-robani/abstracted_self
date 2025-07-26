import { createApp } from "vue";
import "./style.css";
import Profile from "./pages/Profile.vue";

createApp(Profile).mount("#app");

// Setu typed js
const api_endpoint =
  import.meta.env.VITE_API_URL + import.meta.env.VITE_GET_PROFILEDATA_ENDPOINT;

fetch(api_endpoint + "?taglines=true")
  .then((res) => res.json())
  .then((jsonResponse) => {
    const taglines = jsonResponse.data.taglines;

    var typed = new Typed("#typed", {
      strings: taglines,
      typeSpeed: 50,
      backSpeed: 50,
      backDelay: 2000,
      cursorChar: "_",
      shuffle: false,
      smartBackspace: false,
      loop: true,
    });
  });
