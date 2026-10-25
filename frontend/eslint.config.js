import js from "@eslint/js";
import pluginVue from "eslint-plugin-vue";
import prettier from "@vue/eslint-config-prettier";
import globals from "globals";

export default [
  { ignores: ["dist/**", "public/**"] },
  js.configs.recommended,
  ...pluginVue.configs["flat/essential"],
  prettier,
  {
    languageOptions: {
      ecmaVersion: "latest",
      sourceType: "module",
      globals: { ...globals.browser },
    },
    rules: {
      "max-len": ["warn", { code: 120 }],
    },
  },
];
