<template>
  <header data-bs-theme="night">
    <nav class="navbar fixed-top navbar-expand-md navbar-dark">
      <div class="container-fluid user-select-none">
        <router-link to="/" class="navbar-brand mb-0 h1">
          <img style="height: 28px" src="@/assets/img/brand.png" alt="XIV ToDo branding" />
        </router-link>
        <button
          class="navbar-toggler"
          type="button"
          data-bs-toggle="collapse"
          data-bs-target="#navbarSupportedContent"
          aria-controls="navbarSupportedContent"
          aria-expanded="false"
          aria-label="Toggle navigation"
        >
          <span class="navbar-toggler-icon"></span>
        </button>

        <div class="collapse navbar-collapse" id="navbarSupportedContent">
          <ul class="navbar-nav navbar-nav-scroll ms-auto flex-grow-1 mb-2 mb-md-0">
            <!-- Profile -->
            <li v-if="this.appStore.hasCharacter" class="nav-item">
              <router-link to="/profile" class="nav-link" @click="collapseNav">
                <i class="fa-fw fas fa-user d-none d-md-inline-block d-xl-none"></i>
                <span class="d-md-none d-xl-inline">{{ $t("page.profile") }}</span>
              </router-link>
            </li>

            <!-- Encounters -->
            <li class="nav-item">
              <router-link to="/encounters" class="nav-link" @click="collapseNav">
                <i class="fa-fw fas fa-swords d-none d-md-inline-block d-xl-none"></i>
                <span class="d-md-none d-xl-inline">{{ $t("page.encounters") }}</span>
              </router-link>
            </li>

            <!-- Questlines -->
            <li class="nav-item">
              <router-link to="/questlines" class="nav-link" @click="collapseNav">
                <i class="fa-fw fas fa-map-marker-exclamation d-none d-md-inline-block d-xl-none"></i>
                <span class="d-md-none d-xl-inline">{{ $t("page.questlines") }}</span>
              </router-link>
            </li>

            <!-- Challenges -->
            <li class="nav-item">
              <router-link to="/challenges" class="nav-link" @click="collapseNav">
                <i class="fa-fw fas fa-medal d-none d-md-inline-block d-xl-none"></i>
                <span class="d-md-none d-xl-inline">{{ $t("page.challenges") }}</span>
              </router-link>
            </li>

            <!-- Collection -->
            <!-- <li class="nav-item dropdown">
              <a
                class="nav-link dropdown-toggle"
                href="#"
                id="collectionDropdown"
                role="button"
                data-bs-toggle="dropdown"
                aria-expanded="false"
              >
              {{ $t("page.collection") }}
              </a>
              <ul
                class="dropdown-menu dropdown-menu-end dropdown-menu-dark"
                aria-labelledby="collectionDropdown"
              >
                <li>
                  <router-link to="/achievements" class="dropdown-item" @click="collapseNav">
                    <i class="fa-fw fas fa-trophy"></i> {{ $t("page.achievements") }}
                  </router-link>
                  <router-link to="/titles" class="dropdown-item" @click="collapseNav">
                    <i class="fa-fw fas fa-scroll"></i> {{ $t("page.titles") }}
                  </router-link>
                  <router-link to="/mounts" class="dropdown-item" @click="collapseNav">
                    <i class="fa-fw fas fa-horse-saddle"></i> {{ $t("page.mounts") }}
                  </router-link>
                  <router-link to="/minions" class="dropdown-item" @click="collapseNav">
                    <i class="fa-fw fas fa-cat-space"></i> {{ $t("page.minions") }}
                  </router-link>
                </li>
              </ul>
            </li> -->

            <!-- Checklist -->
            <li class="nav-item">
              <router-link to="/checklist" class="nav-link" @click="collapseNav">
                <span class="fa-fw fas fa-tasks d-none d-md-inline-block d-xl-none"></span>
                <span class="d-md-none d-xl-inline">{{ $t("page.checklist") }}</span>
              </router-link>
            </li>

            <!-- Items above are left-aligned -->
            <li class="nav-item flex-grow-1"></li>
            <!-- Items below are right-aligned -->

            <!-- Sandbox -->
            <li v-if="this.appStore.isDevEnv" class="nav-item">
              <router-link to="/sandbox" class="nav-link" aria-label="Sandbox">
                <span class="fa-fw fas fa-vial"></span>
                <span class="d-inline d-md-none ms-2">{{ $t("page.sandbox") }}</span>
              </router-link>
            </li>

            <!-- Countdowns -->
            <li class="nav-item">
              <a
                class="nav-link"
                href="#"
                type="button"
                data-bs-toggle="offcanvas"
                data-bs-target="#offcanvasCountdowns"
                aria-controls="offcanvasCountdowns"
                aria-label="Countdowns"
                @click="seenLatestCountdown()"
              >
                <span class="fa-fw fas fa-hourglass-clock position-relative">
                  <span
                    v-if="
                      this.appStore.isSignedIn && this.appStore.latestCountdownSeen < news.latestCountdownID
                    "
                    class="position-absolute top-0 start-100 translate-middle p-1 bg-success border border-light rounded-circle"
                  >
                  </span>
                </span>
                <span class="d-inline d-md-none ms-2">{{ $t("page.countdowns") }}</span>
              </a>
            </li>

            <!-- News -->
            <li class="nav-item">
              <a
                class="nav-link"
                href="#"
                type="button"
                data-bs-toggle="offcanvas"
                data-bs-target="#offcanvasUpdates"
                aria-controls="offcanvasUpdates"
                aria-label="News"
                @click="seenLatestNews()"
              >
                <span class="fa-fw fas fa-newspaper position-relative">
                  <span
                    v-if="this.appStore.isSignedIn && this.appStore.latestNewsSeen < news.latestID"
                    class="position-absolute top-0 start-100 translate-middle p-1 bg-success border border-light rounded-circle"
                  >
                  </span>
                </span>
                <span class="d-inline d-md-none ms-2">{{ $t("page.newsUpdates") }}</span>
              </a>
            </li>

            <!-- Language -->
            <li class="nav-item dropdown">
              <a
                class="nav-link"
                href="#"
                id="collectionDropdown"
                role="button"
                data-bs-toggle="dropdown"
                aria-expanded="false"
                aria-label="Language"
                @click="this.appStore.setFlag({ key: 'languageClicked', value: true })"
              >
                <span class="fa-fw fas fa-language position-relative">
                  <span
                    v-if="!this.appStore.flag('languageClicked')"
                    class="position-absolute top-0 start-100 translate-middle p-1 bg-success border border-light rounded-circle"
                  >
                  </span>
                </span>
                <span class="d-inline d-md-none ms-2">{{ $t("page.language") }}</span>
              </a>
              <ul class="dropdown-menu dropdown-menu-end dropdown-menu-dark" aria-labelledby="collectionDropdown">
                <!-- Preferred language -->
                <li>
                  <h6 class="dropdown-header">{{ $t("navbar.preferredLanguage") }}</h6>
                </li>
                <li>
                  <span v-if="this.appStore.language == 'en'" class="dropdown-item active" @click="collapseNav">
                    English
                    <i class="fa-fw fal fa-check"></i>
                  </span>
                  <a v-else class="dropdown-item" href="#" @click="changeLanguage('en')">English</a>
                </li>
                <li>
                  <span v-if="this.appStore.language == 'fr'" class="dropdown-item active" @click="collapseNav">
                    Français
                    <i class="fa-fw fal fa-check"></i>
                  </span>
                  <a v-else class="dropdown-item" href="#" @click="changeLanguage('fr')">Français</a>
                </li>
                <li>
                  <span v-if="this.appStore.language == 'de'" class="dropdown-item active" @click="collapseNav">
                    Deutsch
                    <i class="fa-fw fal fa-check"></i>
                  </span>
                  <a v-else class="dropdown-item" href="#" @click="changeLanguage('de')">Deutsch</a>
                </li>
                <li>
                  <span v-if="this.appStore.language == 'ja'" class="dropdown-item active" @click="collapseNav">
                    日本語
                    <i class="fa-fw fal fa-check"></i>
                  </span>
                  <a v-else class="dropdown-item" href="#" @click="changeLanguage('ja')">日本語</a>
                </li>

                <li><hr class="dropdown-divider" /></li>

                <!-- Feedback -->
                <li>
                  <a
                    class="dropdown-item"
                    href="https://crowdin.com/project/xiv-todo/"
                    target="_blank"
                    rel="noopener noreferrer"
                  >
                    <i class="fa-fw fas fa-comments-alt"></i> {{ $t("navbar.feedback") }}
                  </a>
                </li>
              </ul>
            </li>

            <!-- Sign-in -->
            <li v-if="!this.appStore.isSignedIn" class="nav-item">
              <a class="nav-link" :href="this.appStore.discordAuthURI">
                <span class="d-inline-block d-md-none fa-fw fas fa-sign-in me-2"></span>
                <span>{{ $t("home.signInDiscord") }}</span>
              </a>
            </li>

            <!-- Settings & character -->
            <li v-else class="nav-item dropdown">
              <a
                class="nav-link"
                href="#"
                id="navbarDropdown"
                role="button"
                data-bs-toggle="dropdown"
                aria-expanded="false"
              >
                <span v-if="this.appStore.hasCharacter">
                  <span class="d-inline-block d-md-none fa-fw fas fa-cog me-2"></span>
                  <span>{{ this.appStore.lodestoneData.Character.Name }}</span>
                  <img
                    v-if="this.appStore.hasCharacter"
                    class="avatar-nav d-none d-md-inline"
                    :src="this.appStore.character.Avatar"
                    :alt="$t('shared.portraitAlt')"
                  />
                </span>
                <template v-else>
                  <span class="d-inline-block d-md-none fa-fw fas fa-cog me-2"></span>
                  <span>{{ $t("navbar.settingsAndCharacters") }}</span>
                </template>
              </a>
              <ul class="dropdown-menu dropdown-menu-end dropdown-menu-dark" aria-labelledby="navbarDropdown">
                <!-- Active character -->
                <li>
                  <h6 class="dropdown-header">{{ $t("navbar.changeActiveCharacter") }}</h6>
                </li>
                <li v-if="!this.appStore.hasCharacter">
                  <span class="dropdown-item disabled">{{ $t("navbar.noCharacters") }}</span>
                </li>
                <li v-for="(item, i) of this.appStore.characters" :key="item.ID">
                  <span v-if="i == this.appStore.activeCharacterID" class="dropdown-item active">
                    <b>{{ item.lodestoneData.Character.Name }}</b>
                    – {{ item.lodestoneData.Character.World }}
                    <i class="fa-fw fal fa-check"></i>
                  </span>
                  <a v-else class="dropdown-item" href="#" @click="changeActiveCharacter(i)">
                    <b>{{ item.lodestoneData.Character.Name }}</b>
                    – {{ item.lodestoneData.Character.World }}
                  </a>
                </li>

                <li><hr class="dropdown-divider" /></li>

                <!-- Discord username & sign-out -->
                <li>
                  <h6 class="dropdown-header">
                    {{ this.appStore.discordUsername }}
                  </h6>
                </li>
                <li>
                  <router-link to="/signout" class="dropdown-item" @click="collapseNav">
                    <i class="fa-fw fas fa-power-off"></i> {{ $t("page.signOut") }}
                  </router-link>
                </li>

                <li><hr class="dropdown-divider" /></li>

                <!-- Settings -->
                <li>
                  <router-link to="/settings" class="dropdown-item" @click="collapseNav">
                    <i class="fa-fw fas fa-cog"></i> {{ $t("page.settings") }}
                  </router-link>
                </li>
              </ul>
            </li>
          </ul>
        </div>
      </div>
    </nav>
  </header>

  <!-- Countdowns off-canvas -->
  <TheCountdownsCanvas />

  <!-- News off-canvas -->
  <TheNewsCanvas />
</template>

<style lang="scss">
nav {
  margin-bottom: 10px;
  background-color: var(--bs-dark);
  border-bottom: 1px solid var(--bs-gray-800);

  .nav-link.router-link-exact-active {
    color: var(--bs-primary-color) !important;
    text-decoration: underline;
    text-decoration-thickness: 2px;
    text-underline-offset: 7px;
  }
}

.navbar-dark .navbar-nav .nav-link {
  color: var(--bs-grey-300);
}

.navbar-dark .navbar-nav .nav-link:hover {
  color: var(--bs-white);
  text-decoration: underline;
  text-decoration-thickness: 2px;
  text-underline-offset: 7px;
}

.avatar-nav {
  border-radius: 50%;
  height: 35px;
  margin-left: 10px;
  margin-right: -5px;
  margin-top: -5px;
  float: right;
}

.dropdown-menu-dark .dropdown-item.active,
.dropdown-menu-dark .dropdown-item:focus {
  background-color: var(--bs-dark);
}

.offcanvas {
  background-color: var(--bs-dark);
  color: var(--bs-body-color);
}

.offcanvas-header .btn-close {
  background: transparent
    url("data:image/svg+xml,%3csvg xmlns='http://www.w3.org/2000/svg' viewBox='0 0 16 16' fill='%23ddd'%3e%3cpath d='M.293.293a1 1 0 011.414 0L8 6.586 14.293.293a1 1 0 111.414 1.414L9.414 8l6.293 6.293a1 1 0 01-1.414 1.414L8 9.414l-6.293 6.293a1 1 0 01-1.414-1.414L6.586 8 .293 1.707a1 1 0 010-1.414z'/%3e%3c/svg%3e")
    center/1em auto no-repeat;
}
</style>

<script>
import TheNewsCanvas from "@/components/TheNewsCanvas.vue";
import TheCountdownsCanvas from "@/components/TheCountdownsCanvas.vue";
import news from "@/assets/news.json";
import { updateSettings } from "@/utilities/backend.js";
import { useAppStore } from "@/store/app";

export default {
  setup() {
    return { appStore: useAppStore() };
  },
  Name: "TheNavbar",
  data() {
    return {
      news: news,
    };
  },
  components: {
    TheNewsCanvas,
    TheCountdownsCanvas,
  },
  methods: {
    changeLanguage(lang) {
      this.$i18n.locale = lang;
      this.appStore.setLanguage(lang);
      this.collapseNav();
    },
    collapseNav() {
      let navCollapse = document.getElementById("navbarSupportedContent");
      navCollapse.classList.remove("show");
    },
    changeActiveCharacter(i) {
      this.appStore.changeActiveCharacter(i);
      this.collapseNav();
    },
    seenLatestNews() {
      if (!this.appStore.isSignedIn) return;

      let settings = { ...this.appStore.settings };
      if (settings.latestNewsSeen == news.latestID) return;

      settings.latestNewsSeen = news.latestID;
      this.appStore.setSettings(settings);
      updateSettings(settings);
    },
    seenLatestCountdown() {
      if (!this.appStore.isSignedIn) return;

      let settings = { ...this.appStore.settings };
      if (settings.latestCountdownSeen == news.latestCountdownID) return;

      settings.latestCountdownSeen = news.latestCountdownID;
      this.appStore.setSettings(settings);
      updateSettings(settings);
    },
  },
};
</script>
