import { createRouter, createWebHistory } from "vue-router";
import store from "../store";
import Home from "../views/Home.vue";
import { authenticate } from "../utilities/backend";

// Static routes must also be listed in staticRoutes in vite.config.js to be served directly
const routes = [
  {
    path: "/",
    name: "Home",
    component: Home,
    beforeEnter: (_to, _from, next) => {
      if (sessionStorage.getItem("redirect") !== null) {
        const redirect = sessionStorage.redirect;
        delete sessionStorage.redirect;
        next(redirect);
      } else {
        next();
      }
    },
  },
  {
    path: "/auth",
    component: Home,
    beforeEnter: (_to, _from, next) => {
      authenticate(_to.query.code)
        .then((userData) => {
          store.commit("setUserData", userData);
        })
        .catch((err) => {
          store.commit("signingIn", err);
          console.log(err);
        });

      store.commit("signingIn", true);
      next("/");
    },
  },
  {
    path: "/profile",
    name: "Profile",
    component: () => import("../views/Profile.vue"),
    beforeEnter: (_to, _from, next) => {
      if (!store.getters.hasCharacter) {
        next("/");
      }
      next();
    },
  },
  {
    path: "/signout",
    name: "SignOut",
    component: () => import("../views/SignOut.vue"),
  },
  {
    path: "/encounters",
    name: "Encounters",
    component: () => import("../views/Encounters.vue"),
  },
  {
    path: "/questlines",
    name: "Questlines",
    component: () => import("../views/Questlines.vue"),
  },
  {
    path: "/challenges",
    name: "Challenges",
    component: () => import("../views/Challenges.vue"),
  },
  // {
  //   path: "/achievements",
  //   name: "Achievements",
  //   component: () => import("../views/Achievements.vue"),
  // },
  // {
  //   path: "/titles",
  //   name: "Titles",
  //   component: () => import("../views/Titles.vue"),
  // },
  // {
  //   path: "/minions",
  //   name: "Minions",
  //   component: () => import("../views/Minions.vue"),
  // },
  // {
  //   path: "/mounts",
  //   name: "Mounts",
  //   component: () => import("../views/Mounts.vue"),
  // },
  {
    path: "/checklist",
    name: "Checklist",
    component: () => import("../views/Checklist.vue"),
  },
  {
    path: "/settings",
    name: "Settings",
    component: () => import("../views/Settings.vue"),
    beforeEnter: (_to, _from, next) => {
      if (!store.getters.userData) {
        next("/");
      }
      next();
    },
  },
  {
    path: "/character/:id",
    name: "Character",
    component: () => import("../views/Character.vue"),
  },
  {
    path: "/help",
    name: "Help",
    component: () => import("../views/Help.vue"),
  },
  {
    path: "/privacy",
    name: "Privacy",
    component: () => import("../views/Privacy.vue"),
  },
  {
    path: "/sandbox",
    name: "Sandbox",
    component: () => import("../views/Sandbox.vue"),
  },
  {
    path: "/:pathMatch(.*)",
    redirect: {
      name: "Home",
    },
  },
];

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
  // eslint-disable-next-line no-unused-vars
  scrollBehavior(to, from, savedPosition) {
    document.getElementById("app").scrollIntoView();
  },
});

export default router;
