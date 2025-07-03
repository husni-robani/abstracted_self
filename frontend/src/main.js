import { createApp } from "vue";
import "./style.css";
import Profile from "./pages/Profile.vue";

createApp(Profile).mount("#app");

var typed = new Typed("#typed", {
  strings: [
    "Developer-in-Training",
    "Craft digital experiences",
    "Junior Developer",
    "Backend engineer",
  ],
  typeSpeed: 50,
  backSpeed: 50,
  backDelay: 2000,
  cursorChar: "_",
  shuffle: false,
  smartBackspace: false,
  loop: true,
});
