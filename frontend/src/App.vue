<template>
  <div id="root" class="d-flex flex-column min-vh-100" :data-bs-theme="this.$store.getters.theme">
    <TheNavbar />

    <main class="flex-shrink-0">
      <!-- Alert: Backend offline -->
      <div v-if="this.$store.getters.backendOffline" class="container">
        <AlertMsg type="danger" :msg="$t('message.backendUnavailable')" />
      </div>

      <!-- Alert: New version available -->
      <!-- <div v-else-if="!this.$store.getters.versionMatches" class="container">
        <AlertMsg
          type="success"
          msg="A new version is available! <a href='javascript:window.location.reload()' class='alert-link'>Reload and update the page</a>."
        />
      </div> -->

      <!-- Alert: Global error -->
      <div v-if="this.globalError" class="container">
        <AlertMsg type="danger" :msg="this.globalError" />
      </div>

      <!-- Alert: Signing in... -->
      <div v-if="this.$store.state.signingIn == true" class="container">
        <AlertMsg type="normal" :msg="$t('message.signingDiscord')" />
      </div>

      <!-- Alert: Lodestone maintenance -->
      <div v-if="this.now < this.news.maintenance.start" class="container">
        <AlertMsg type="info" :msg="this.lodestoneMaintenanceSoonMessage" />
      </div>
      <div v-else-if="this.now > this.news.maintenance.start && this.now < this.news.maintenance.end" class="container">
        <AlertMsg type="warning" :msg="this.lodestoneMaintenanceMessage" />
      </div>
      <div v-else-if="this.now > this.news.maintenance.end && !this.news.maintenance.updated" class="container">
        <AlertMsg type="info" :msg="$t('message.lodestoneMaintenanceNotUpdated')" />
      </div>

      <!-- Page content -->
      <div class="container">
        <router-view />
      </div>
    </main>

    <TheFooter />
  </div>
</template>

<style lang="scss">
///
// Variables
///

[data-bs-theme="day"] {
  --bs-body-color: #222222;
  --bs-body-bg: #f6f8fc;
  --bs-dark: #212529;
  --bs-transparent: rgba(255, 255, 255, 0);

  --bs-primary: #148653;
  --bs-primary-color: var(--bs-primary);
  --bs-primary-hover: #39a273;
  --bs-primary-active: #2f8c63;
  --bs-primary-text-emphasis: var(--bs-white);

  --bs-secondary: #4b5a66;
  --bs-secondary-color: var(--bs-secondary);

  --bs-success: #008045;

  --bs-warning: #967100;

  --bs-info: #0d7c93;

  --xt-tank-color: #0971cd;
  --xt-healer-color: #167b01;
  --xt-dps-color: #d02323;
  --xt-crafter-color: #85670d;
  --xt-gatherer-color: #6741eb;

  .form-check-input:checked[type="radio"] {
    --bs-form-check-bg-image: url("data:image/svg+xml,%3csvg xmlns='http://www.w3.org/2000/svg' viewBox='-4 -4 8 8'%3e%3ccircle r='2' fill='%23222'/%3e%3c/svg%3e");
  }

  background-color: var(--bs-body-bg);
  color: var(--bs-body-color);
}

[data-bs-theme="night"] {
  --bs-body-color: #dddddd;
  --bs-body-bg: #1c2024;
  --bs-dark: #212529;
  --bs-transparent: rgba(255, 255, 255, 0);

  --bs-primary: #41b883;
  --bs-primary-color: var(--bs-primary);
  --bs-primary-hover: #39a273;
  --bs-primary-active: #2f8c63;
  --bs-primary-text-emphasis: #0d3724;

  --bs-secondary: #838d96;
  --bs-secondary-color: var(--bs-secondary);

  --bs-success: #53b462;
  --bs-success-bg-subtle: var(--bs-success);
  --bs-success-text-emphasis: #05310d;
  --bs-success-border-subtle: var(--bs-success-text-emphasis);

  --bs-warning: #f6e275;
  --bs-warning-bg-subtle: var(--bs-warning);
  --bs-warning-text-emphasis: #332d19;
  --bs-warning-border-subtle: var(--bs-warning-text-emphasis);

  --bs-danger: #b45353;
  --bs-danger-bg-subtle: var(--bs-danger);
  --bs-danger-text-emphasis: #310505;
  --bs-danger-border-subtle: var(--bs-danger-text-emphasis);

  --bs-info: #9dd1ff;
  --bs-info-bg-subtle: var(--bs-info);
  --bs-info-text-emphasis: #071e2d;
  --bs-info-border-subtle: var(--bs-info-text-emphasis);

  --xt-tank-color: #9dd1ff;
  --xt-healer-color: #a6ff93;
  --xt-dps-color: #ff8f8f;
  --xt-crafter-color: #ffd99d;
  --xt-gatherer-color: #bdaaff;

  .form-check-input:checked[type="radio"] {
    --bs-form-check-bg-image: url("data:image/svg+xml,%3csvg xmlns='http://www.w3.org/2000/svg' viewBox='-4 -4 8 8'%3e%3ccircle r='2' fill='%23ddd'/%3e%3c/svg%3e");
  }

  background-color: var(--bs-body-bg);
  color: var(--bs-body-color);
}

.navbar-nav {
  --bs-nav-link-font-weight: 600;
}

.offcanvas,
.offcanvas-lg,
.offcanvas-md,
.offcanvas-sm,
.offcanvas-xl,
.offcanvas-xxl {
  --bs-offcanvas-transition: transform 0.1s ease-in-out;
}

///
// Body
///

#root {
  font-family:
    "Montserrat",
    "Noto Sans JP",
    system-ui,
    -apple-system,
    "Segoe UI",
    Roboto,
    "Helvetica Neue",
    "Noto Sans",
    "Liberation Sans",
    Arial,
    sans-serif,
    "Apple Color Emoji",
    "Segoe UI Emoji",
    "Segoe UI Symbol",
    "Noto Color Emoji";
  font-size: 18px;
  font-optical-sizing: auto;
  font-kerning: auto;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}

main {
  padding-top: 75px;
}

///
// Scrollbars
///

* {
  scrollbar-width: thin;
  scrollbar-color: #41b883 var(--bs-gray-800);
}

*::-webkit-scrollbar {
  width: 12px;
}

*::-webkit-scrollbar-track {
  background: var(--bs-gray-800);
}

*::-webkit-scrollbar-thumb {
  background-color: #41b883;
  border-radius: 20px;
  border: 3px solid var(--bs-gray-800);
}

///
// Tooltips
///

.tt {
  position: relative;
}

.tt .tt-text {
  visibility: hidden;
  width: auto;
  white-space: nowrap;
  background-color: var(--bs-gray-800);
  color: var(--bs-gray-200);
  text-align: center;
  border-radius: 10px;
  padding: 8px 10px;
  position: absolute;
  z-index: 10;
  bottom: 130%;
  left: 50%;
  transform: translateX(-50%);
  pointer-events: none;
  -webkit-user-select: none !important;
  -moz-user-select: none !important;
  user-select: none !important;
}

.tt .tt-text::after {
  content: "";
  position: absolute;
  top: 100%;
  left: 50%;
  margin-left: -5px;
  border-width: 5px;
  border-style: solid;
  border-color: var(--bs-gray-800) transparent transparent transparent;
}

@media only screen and (max-width: 800px) {
  .tt .tt-text {
    display: none;
  }
}

.tt:hover .tt-text {
  visibility: visible;
}

///
// Typography
///

a,
abbr {
  text-underline-offset: 0.25em !important;
}

p a,
p a:focus {
  color: var(--bs-primary-color);
}

p a:hover {
  filter: brightness(150%);
}

a:hover .fa-fw {
  filter: brightness(150%);
}

h1 {
  font-weight: 700;
  margin-bottom: 0px !important;

  .border-bottom {
    border-color: var(--bs-primary) !important;
  }
}

h3 {
  font-weight: 700;

  img {
    height: 0.8em;
    margin-right: 10px;
  }
}

hr {
  height: 0px;
  border: none;
  border-top: 1px solid var(--bs-primary);
  opacity: 1;
  background-color: initial;
  margin-top: 12px;
  margin-bottom: 12px;
}

.text-justify {
  text-align: justify;
  text-justify: inter-word;
}

.text-decoration-line-through {
  text-decoration-thickness: 2px !important;
}

.text-primary {
  color: var(--bs-primary) !important;
}

.text-secondary {
  color: var(--bs-secondary) !important;
}

.text-success {
  color: var(--bs-success) !important;
}

.text-danger {
  color: var(--bs-danger) !important;
}

.text-warning {
  color: var(--bs-warning) !important;
}

.text-info {
  color: var(--bs-info) !important;
}

///
// Buttons
///

// Primary
.btn-primary,
.btn-outline-primary {
  --bs-btn-hover-color: var(--bs-white);
  --bs-btn-active-color: var(--bs-white);

  --bs-btn-border-color: var(--bs-primary);
  --bs-btn-hover-border-color: var(--bs-primary-hover);
  --bs-btn-active-border-color: var(--bs-primary-active);
  --bs-btn-disabled-border-color: var(--bs-primary);

  --bs-btn-focus-shadow-rgb: 0, 0, 0;
  --bs-btn-active-shadow: inset 0 3px 5px rgba(0, 0, 0, 0.125);
}

.btn-primary {
  --bs-btn-color: var(--bs-primary-text-emphasis);
  --bs-btn-disabled-color: var(--bs-primary-text-emphasis);

  --bs-btn-bg: var(--bs-primary);
  --bs-btn-hover-bg: var(--bs-primary-hover);
  --bs-btn-active-bg: var(--bs-primary-active);
  --bs-btn-disabled-bg: var(--bs-primary);
}

.btn-outline-primary {
  --bs-btn-color: var(--bs-primary);
  --bs-btn-disabled-color: var(--bs-primary);

  --bs-btn-bg: var(--bs-transparent);
  --bs-btn-hover-bg: var(--bs-primary);
  --bs-btn-active-bg: var(--bs-primary);
  --bs-btn-disabled-bg: var(--bs-transparent);
}

// Success outline
.btn-outline-success {
  color: var(--bs-success);
  border-color: var(--bs-success);
}

// Secondary outline
.btn-outline-secondary {
  color: var(--bs-secondary);
  border-color: var(--bs-secondary);
}

///
// Lists
///

.list-group-item-action:focus,
.list-group-item-action:hover {
  color: unset;
  background-color: var(--bs-body-bg);
  filter: brightness(95%);
}

.list-group-item-action:active {
  background-color: var(--bs-body-bg);
  filter: brightness(90%);
}

.list-group-item {
  color: var(--bs-body-color);
  background-color: var(--bs-body-bg);
  border-color: #484848;
  padding-bottom: 0.4rem;
}

///
// Tables
///

table {
  color: var(--bs-body-color) !important;
}

.table {
  --bs-table-color: unset;
  --bs-table-bg: unset;
}

tbody,
td,
tfoot,
th,
thead,
tr {
  border-color: #484848;
}

///
// Forms
///

.form-check-input {
  background-color: var(--bs-transparent);
  border: 1px solid var(--bs-body-color);
}

.form-check-input:checked {
  background-color: var(--bs-transparent);
  border: 1px solid var(--bs-body-color);
}

.form-control {
  border-color: var(--bs-body-color);
}

.form-control::placeholder {
  color: rgba(97, 108, 120, 0.6);
}

///
// Progress
///

.progress {
  background-color: var(--bs-dark);
}

.progress-bar {
  background-color: var(--bs-primary-color);
  overflow: unset;
  padding-left: 5px;
  padding-right: 5px;
}

///
// Misc
///

.cursor-grab {
  cursor: grab;
}

.cursor-pointer {
  cursor: pointer;
}

.filter-bright-50 {
  filter: brightness(50%);
}

.border {
  border-color: var(--bs-gray-800) !important;
}

.border-bottom {
  border-color: var(--bs-primary) !important;
}

.border-grey {
  border-color: var(--bs-gray-800) !important;
}

.indent-paragraphs p {
  margin-left: 20px;
}
</style>

<script>
import TheNavbar from "@/components/TheNavbar.vue";
import TheFooter from "@/components/TheFooter.vue";
import AlertMsg from "@/components/AlertMsg.vue";
import { getVersion, getUserData, addCharacter } from "@/utilities/backend.js";
import news from "@/assets/news.json";

export default {
  name: "App",
  components: {
    TheNavbar,
    TheFooter,
    AlertMsg,
  },
  data() {
    return {
      intervalFunction: null,
      tabFocused: true,
      news: news,
      now: new Date() / 1000,
    };
  },
  mounted() {
    this.$nextTick(function () {
      this.setLanguage();
      // this.checkUpstreamVersion();
      this.updateUserData();
      setTimeout(() => this.updateCharactersData(), 1000 * 10); // 10 seconds
    });
    this.intervalFunction = setInterval(() => {
      if (this.tabFocused) {
        // this.checkUpstreamVersion();
        this.updateCharactersData();
        this.now = new Date() / 1000;
      }
    }, 1000 * 120); // 2 minutes
  },
  unmounted() {
    clearInterval(this.intervalFunction);
  },
  watch: {
    // Erase any previous sign-in message when navigating away from Home. This must not live in a
    // computed property: any store mutation re-evaluates it and would erase messages set mid-navigation.
    $route: {
      handler(to) {
        if (to.name !== undefined && to.name !== "Home") {
          this.$store.commit("signingIn", false);
        }
      },
      immediate: true,
    },
  },
  computed: {
    windowTitle() {
      const routeName = this.$route.name;

      if (routeName === undefined || routeName === "Home") {
        return `XIV ToDo: ${this.$t("page.homeTitle")}`;
      } else {
        const i18nRouteName = this.$t(`page.${routeName.toLowerCase()}`);
        return `XIV ToDo - ${i18nRouteName}`;
      }
    },
    canonicalURL() {
      return "https://xivtodo.com" + this.$route.path;
    },
    globalError() {
      if (this.$store.state.signingIn?.sessionExpired) {
        return this.$t("message.sessionExpired", { url: this.$store.getters.discordAuthURI });
      } else if (this.$store.state.signingIn?.status == 401) {
        return this.$t("message.signingUnauthorizedError");
      } else if (this.$store.state.signingIn?.status == 500) {
        return this.$t("message.signingUnknownError");
      } else {
        return "";
      }
    },
    lodestoneMaintenanceSoonMessage() {
      return this.getMaintenanceMessage("message.lodestoneMaintenanceSoon", this.news.maintenance.start);
    },
    lodestoneMaintenanceMessage() {
      return this.getMaintenanceMessage("message.lodestoneMaintenance", this.news.maintenance.end);
    },
  },
  head() {
    return {
      title: this.windowTitle,
      link: [
        {
          rel: "canonical",
          href: this.canonicalURL,
        },
      ],
    };
  },
  created() {
    this.detectFocusOut();
  },
  methods: {
    getCountdownToTimestamp(targetTimestamp) {
      const secondsUntilTarget = targetTimestamp - this.now;
      const totalMinutesUntilTarget = Math.max(0, Math.ceil(secondsUntilTarget / 60));

      return {
        hours: Math.floor(totalMinutesUntilTarget / 60),
        minutes: totalMinutesUntilTarget % 60,
      };
    },
    getTimeRemainingText(targetTimestamp) {
      const countdown = this.getCountdownToTimestamp(targetTimestamp);
      const hoursRemaining = this.$t("shared.time.hoursShort", {
        hours: countdown.hours,
      });
      const minutesRemaining = this.$t("shared.time.minutesShort", {
        minutes: countdown.minutes,
      });

      return `${hoursRemaining} ${minutesRemaining}`;
    },
    getMaintenanceMessage(messageKey, targetTimestamp) {
      return this.$t(messageKey, {
        timeRemaining: this.getTimeRemainingText(targetTimestamp),
      });
    },
    detectFocusOut() {
      let inView = false;

      const onWindowFocusChange = (e) => {
        if ({ focus: 1, pageshow: 1 }[e.type]) {
          if (inView) return;
          this.tabFocused = true;
          inView = true;
        } else if (inView) {
          this.tabFocused = !this.tabFocused;
          inView = false;
        }
      };

      window.addEventListener("focus", onWindowFocusChange);
      window.addEventListener("blur", onWindowFocusChange);
      window.addEventListener("pageshow", onWindowFocusChange);
      window.addEventListener("pagehide", onWindowFocusChange);
    },
    setLanguage() {
      if (this.$store.getters.language === null) {
        let locale = navigator.language.split("-")[0];
        if (locale === "fr" || locale === "de" || locale === "ja") {
          this.$store.commit("setLanguage", locale);
        } else {
          this.$store.commit("setLanguage", "en");
        }
      }
      this.$i18n.locale = this.$store.getters.language;
    },
    checkUpstreamVersion() {
      getVersion()
        .then((version) => {
          this.$store.commit("setUpstreamVersion", version);
        })
        .catch(() => {
          this.$store.commit("setUpstreamVersion", "OFFLINE");
        });
    },
    updateUserData() {
      if (this.$store.getters.isSignedIn) {
        getUserData()
          .then((userData) => {
            this.$store.commit("setUserData", userData);
          })
          .catch((err) => {
            console.log(err);
          });
      }
    },
    updateCharactersData() {
      let i = 0;
      for (let character of this.$store.getters.characters) {
        if (this.$store.getters.characterOutOfDate(i)) {
          let localID = i;
          console.log("Updating data for " + character.lodestoneData.Character.Name + "...");

          let characterID = character.lodestoneData.Character.ID;
          addCharacter(characterID)
            .then((characterData) => {
              this.$store.commit("addCharacter", characterData);
              console.log("Data for " + characterData.lodestoneData.Character.Name + " updated.");
            })
            .catch((err) => {
              console.log(err);
              this.error = err;
              if (err.status == 404) {
                this.$store.commit("invalidateCharacter", localID);
              }
            });
        }
        i++;
      }
    },
  },
};
</script>
