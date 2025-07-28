import { createRouter, createWebHistory } from "vue-router";
import Profile from "../pages/Profile.vue";
import Login from "../pages/Login.vue";
import ProfileManagement from "../pages/ProfileManagement.vue";

// routes
const routes = [
  { path: "/", name: "Welcome", component: Profile },
  { path: "/unlock", name: "Login", component: Login },
  {
    path: "/admin/profile",
    name: "Profile Admin",
    component: ProfileManagement,
    meta: { title: "Profile" },
  },
  {
    path: "/admin/dashboard",
    name: "Dashboard",
    meta: { title: "Dashboard" },
  },
  {
    path: "/admin/projects",
    name: "Projects",
    meta: { title: "Projects" },
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
