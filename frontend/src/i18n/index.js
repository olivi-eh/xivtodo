import { createI18n, Translation } from "vue-i18n";
import langEN from "./lang-en.json";
import langFR from "./lang-fr.json";
import langDE from "./lang-de.json";
import langJA from "./lang-ja.json";

const messages = {
  en: langEN,
  fr: langFR,
  de: langDE,
  ja: langJA,
};

const i18n = createI18n({
  legacy: false,
  locale: "en",
  fallbackLocale: "en",
  messages,
});

// No component defines a local i18n scope, so default <i18n-t> to the global one
const GlobalTranslation = {
  ...Translation,
  props: { ...Translation.props, scope: { ...Translation.props.scope, default: "global" } },
};

export default {
  install(app) {
    app.use(i18n, { globalInstall: false });
    app.component("i18n-t", GlobalTranslation);
  },
};
