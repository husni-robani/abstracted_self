import { createRouter, createWebHistory } from "vue-router";
import Profile from "../pages/Profile.vue";
import Login from "../pages/Login.vue";
import ProfileManagement from "../pages/ProfileManagement.vue";
import ProjectManagement from "../pages/ProjectManagement.vue";

// routes
const routes = [
  { path: "/", name: "Welcome", component: Profile },
  { path: "/unlock", name: "Login", component: Login },
  {
    path: "/admin/profile",
    name: "Profile Admin",
    component: ProfileManagement,
    meta: { title: "Profile", requiresAuth: true },
  },
  {
    path: "/admin/dashboard",
    name: "Dashboard",
    meta: { title: "Dashboard", requiresAuth: true },
  },
  {
    path: "/admin/projects",
    name: "Projects",
    meta: { title: "Projects", requiresAuth: true },
    component: ProjectManagement,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

function isTokenExpiringSoon(token) {
  try {
    const payload = JSON.parse(atob(token.split(".")[1]));
    const exp = payload.exp * 1000; // token expiration time in seconds
    const now = Date.now();
    return exp - now < 2 * 60 * 1000; // less than 2 minutes
  } catch {
    return true;
  }
}

router.beforeEach(async (to, from, next) => {
  const token = localStorage.getItem("token");
  const requiresAuth = to.meta.requiresAuth;

  if (!requiresAuth) {
    return next();
  }

  if (!token) {
    return next({ name: "Login" });
  }

  const checkTokenEnd =
    import.meta.env.VITE_API_URL + import.meta.env.VITE_CHECK_TOKEN_ENDPOINT;
  const renewTokenEnd =
    import.meta.env.VITE_API_URL + import.meta.env.VITE_RENEW_TOKEN_ENDPOINT;

  try {
    // renew token if condition true
    if (isTokenExpiringSoon(token)) {
      const renewResponse = await fetch(renewTokenEnd, {
        method: "GET",
        headers: {
          Authorization: `Bearer ${token}`,
          "Content-Type": "application/json",
        },
      });

      if (!renewResponse.ok) throw new Error("Token renewal failed");

      const renewData = await renewResponse.json();
      localStorage.setItem("token", renewData.data.token);
    }

    // Check token validity
    const checkResponse = await fetch(checkTokenEnd, {
      method: "POST",
      headers: {
        Authorization: `Bearer ${localStorage.getItem("token")}`,
      },
    });

    if (!checkResponse.ok) throw new Error("Token check failed");

    next(); // ✅ All good
  } catch (err) {
    console.error("Auth error:", err);
    localStorage.removeItem("token");
    next({ name: "Login" });
  }
});

export default router;
