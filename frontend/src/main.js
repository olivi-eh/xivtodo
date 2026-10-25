import { createApp } from "vue";
import { VueHeadMixin, createHead } from "@unhead/vue/client";
import App from "./App.vue";
import router from "./router";
import { createPinia } from "pinia";
import { useAppStore } from "./store/app";
import i18n from "./i18n";
import { configure } from "vue-gtag";

configure({
  tagId: "G-9WXY5M43Z0",
  config: {
    cookie_domain: "xivtodo.com",
    cookie_flags: "SameSite=None;Secure",
  },
})

const head = createHead();
const app = createApp(App).use(createPinia());

const store = useAppStore();
if (localStorage.getItem("store")) {
  store.$patch((state) => Object.assign(state, JSON.parse(localStorage.getItem("store"))));
}
store.clampActiveCharacter();
store.env = { ...import.meta.env };
store.$subscribe(
  (_mutation, state) => {
    localStorage.setItem("store", JSON.stringify(state));
  },
  { detached: true, flush: "sync" },
);

app.use(router).use(i18n).use(head).mixin(VueHeadMixin).mount("#app");
