import { createRouter, createWebHistory } from "vue-router";
import Profile from "../pages/Profile.vue";
import Login from "../pages/Login.vue";
import ProfileManagement from "../pages/ProfileManagement.vue";
import ProjectManagement from "../pages/ProjectManagement.vue";
import UpdateProject from "../pages/UpdateProject.vue";
import Dashboard from "../pages/Dashboard.vue";
import ExperienceManagement from "../pages/ExperienceManagement.vue";
import BlogManagement from "../pages/BlogManagement.vue";
import Visitors from "../pages/Visitors.vue";
import Cookies from "js-cookie";
import { v4 as uuidv4 } from "uuid";

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
    component: Dashboard,
  },
  {
    path: "/admin/projects",
    name: "Projects",
    meta: { title: "Projects", requiresAuth: true },
    component: ProjectManagement,
  },
  {
    path: "/admin/projects/:id",
    name: "Update Project",
    meta: { title: "Projects", requiresAuth: true },
    component: UpdateProject,
  },
  {
    path: "/admin/experiences",
    name: "Experiences",
    meta: { title: "Experiences", requiresAuth: true },
    component: ExperienceManagement,
  },
  {
    path: "/admin/blog",
    name: "Blog",
    meta: { title: "Blog", requiresAuth: true },
    component: BlogManagement,
  },
  {
    path: "/admin/visitors",
    name: "Visitors",
    meta: { title: "Visitors", requiresAuth: true },
    component: Visitors,
  },
  {
    path: "/admin/blog/new",
    name: "New Post",
    meta: { title: "New Post", requiresAuth: true },
    component: () => import("../pages/BlogPostEditor.vue"),
  },
  {
    path: "/admin/blog/:id/edit",
    name: "Edit Post",
    meta: { title: "Edit Post", requiresAuth: true },
    component: () => import("../pages/BlogPostEditor.vue"),
  },
  {
    path: "/blog",
    name: "Blog List",
    component: () => import("../pages/BlogList.vue"),
  },
  {
    path: "/blog/:id",
    name: "Blog Post",
    component: () => import("../pages/BlogPost.vue"),
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
  scrollBehavior(to, from, savedPosition) {
    if (savedPosition) return savedPosition;
    if (to.hash) return { el: to.hash, behavior: "smooth" };
    return { top: 0 };
  },
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

  // visitor tracking — one visit per browser per calendar day.
  // The backend reads `visitor_identifier` from the COOKIE (not a JSON body)
  // and dedups server-side, so the cookie must be set BEFORE the request.
  if (to.path == "/") {
    const profileVisitEndpoint =
      import.meta.env.VITE_API_URL +
      import.meta.env.VITE_PROFILE_VISIT_ENDPOINT;

    if (!Cookies.get("visitor_identifier")) {
      const endOfDay = new Date();
      endOfDay.setHours(24, 0, 0, 0);
      Cookies.set("visitor_identifier", uuidv4(), { expires: endOfDay });

      try {
        const res = await fetch(profileVisitEndpoint, { method: "POST" });
        if (!res.ok) {
          throw new Error("profile visit request failed");
        }
      } catch (e) {
        console.error("tracking failed", e);
      }
    }
  }

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
