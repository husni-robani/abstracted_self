import { createRouter, createWebHistory } from "vue-router";
import Profile from "../pages/Profile.vue";
import Login from "../pages/Login.vue";

// routes
const routes = [
  { path: "/", name: "Welcome", component: Profile },
  { path: "/unlock", name: "Login", component: Login },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
