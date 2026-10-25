<template>
  <!-- <template v-if="this.$store.getters.isSignedIn && this.$store.getters.hasCharacter"> -->
  <template v-if="false">
    <PageHeader title="page.home" :hideNotSignedInAlert="true" :hideNoCharactersAlert="true">
      <router-link to="/settings/">
        <button type="button" class="btn btn-outline-success">
          {{ $t("home.addCharacter") }}
        </button>
      </router-link>
    </PageHeader>

    <div class="row">
      <div
        v-for="item of this.$store.getters.characters"
        :key="item.ID"
        class="container px-4 py-2 mb-3 border-bottom border-grey"
      >
        <h2 class="pb-2">
          <img
            class="avatar-home d-none d-sm-inline"
            :src="item.lodestoneData.Character.Avatar"
            :alt="$t('shared.portraitAlt')"
          />
          {{ item.lodestoneData.Character.Name }}
        </h2>
        <p>
          {{ $t("checklist.weeklies") }}
        </p>
        <div class="progress">
          <div
            class="progress-bar"
            role="progressbar"
            :style="{ width: taskCompletionPercentage(item.checklistData.weeklyChecklist) }"
          >
            {{ item.checklistData.weeklyChecklist.filter((x) => !x.hidden && x.checked).length }} /
            {{ item.checklistData.weeklyChecklist.filter((x) => !x.hidden).length }}
            ({{ taskCompletionPercentage(item.checklistData.weeklyChecklist) }})
          </div>
        </div>
        <br />
        <p>
          {{ $t("checklist.dailies") }}
        </p>
        <div class="progress">
          <div
            class="progress-bar"
            role="progressbar"
            :style="{ width: taskCompletionPercentage(item.checklistData.dailyChecklist) }"
          >
            {{ item.checklistData.dailyChecklist.filter((x) => !x.hidden && x.checked).length }} /
            {{ item.checklistData.dailyChecklist.filter((x) => !x.hidden).length }}
            ({{ taskCompletionPercentage(item.checklistData.dailyChecklist) }})
          </div>
        </div>
        <br />
      </div>
    </div>
  </template>

  <template v-else>
    <div class="px-4 py-4 my-4 text-center">
      <!-- Branding -->
      <img class="d-block mx-auto mb-4" :src="this.brandURI" alt="XIV ToDo branding" width="300" />
      <h1 class="visually-hidden">XIV ToDo</h1>

      <div class="container col-lg-8 mx-auto">
        <!-- Lead paragraph -->
        <p class="lead mb-4 fw-normal">
          {{ $t("home.intro") }}
        </p>

        <!-- Buttons -->
        <div class="d-grid gap-2 d-sm-flex mb-5 justify-content-center">
          <!-- First button -->
          <router-link v-if="this.$store.getters.hasCharacter" to="/checklist/">
            <button type="button" class="btn btn-outline-primary btn-lg px-4 gap-3">
              {{ $t("page.checklist") }}
            </button>
          </router-link>
          <router-link v-else-if="this.$store.getters.isSignedIn" to="/settings/">
            <button type="button" class="btn btn-primary btn-lg px-4 gap-3">
              {{ $t("home.addCharacter") }}
            </button>
          </router-link>
          <a v-else :href="this.$store.getters.discordAuthURI">
            <button type="button" class="btn btn-primary btn-lg px-4 gap-3">
              {{ $t("home.signInDiscord") }}
            </button>
          </a>

          <!-- Second button -->
          <router-link v-if="this.$store.getters.isSignedIn" to="/settings/">
            <button type="button" class="btn btn-outline-secondary btn-lg px-4 gap-3">
              {{ $t("page.settings") }}
            </button>
          </router-link>
          <button v-else type="button" class="btn btn-outline-secondary btn-lg px-4" @click="goto('features')">
            {{ $t("home.learnMore") }}
          </button>
        </div>
      </div>

      <!-- Screenshots -->
      <div
        class="container overflow-hidden border-bottom border-grey row justify-content-center"
        style="max-height: 20vh; margin: auto"
      >
        <div class="d-none d-md-block col-2 col-lg-1 px-0 pt-5">
          <img
            src="@/assets/screenshots/en-checklist.png"
            class="border rounded-3 mb-4"
            alt="Screenshot of checklist page"
            style="width: 400%"
          />
        </div>
        <div class="col-11 col-md-7 col-lg-5 px-0">
          <img
            :src="screenshotURI('encounters')"
            class="border rounded-3 mb-4"
            alt="Screenshot of encounters page"
            style="z-index: +1000; position: relative; width: 100%"
          />
        </div>
        <div class="d-none d-md-block col-2 col-lg-1 px-0 pt-5">
          <img
            src="@/assets/screenshots/en-questlines.png"
            class="border rounded-3 mb-4 float-end"
            alt="Screenshot of questlines page"
            style="width: 400%"
          />
        </div>
      </div>
    </div>

    <!-- Features -->
    <div id="features-grid" class="container px-4 py-2 mb-5 border-bottom border-grey">
      <h1 ref="features" class="pb-2 border-bottom">{{ $t("home.features") }}</h1>

      <div class="row row-cols-1 row-cols-sm-2 row-cols-md-3 row-cols-lg-3 g-4 py-5">
        <div class="col d-flex align-items-start">
          <i class="fa-fw fas fa-list flex-shrink-0 me-3"></i>
          <div>
            <h2 class="h4 fw-bold mb-0">{{ $t("home.feature.encountersQuestlines.title") }}</h2>
            <p>{{ $t("home.feature.encountersQuestlines.description") }}</p>
          </div>
        </div>
        <div class="col d-flex align-items-start">
          <i class="fa-fw fas fa-chart-pie flex-shrink-0 me-3"></i>
          <div>
            <h2 class="h4 fw-bold mb-0">{{ $t("home.feature.contentCompletion.title") }}</h2>
            <p>{{ $t("home.feature.contentCompletion.description") }}</p>
          </div>
        </div>
        <div class="col d-flex align-items-start">
          <i class="fa-fw fas fa-clipboard-list flex-shrink-0 me-3"></i>
          <div>
            <h2 class="h4 fw-bold mb-0">{{ $t("home.feature.dailiesWeeklies.title") }}</h2>
            <p>{{ $t("home.feature.dailiesWeeklies.description") }}</p>
          </div>
        </div>
        <div class="col d-flex align-items-start">
          <i class="fa-fw fas fa-user-friends flex-shrink-0 me-3"></i>
          <div>
            <h2 class="h4 fw-bold mb-0">{{ $t("home.feature.altFriendly.title") }}</h2>
            <p>{{ $t("home.feature.altFriendly.description") }}</p>
          </div>
        </div>
        <div class="col d-flex align-items-start">
          <i class="fa-fw fas fa-eye-slash flex-shrink-0 me-3"></i>
          <div>
            <h2 class="h4 fw-bold mb-0">{{ $t("home.feature.spoilerFree.title") }}</h2>
            <p>{{ $t("home.feature.spoilerFree.description") }}</p>
          </div>
        </div>
        <div class="col d-flex align-items-start">
          <i class="fa-fw fas fa-palette flex-shrink-0 me-3"></i>
          <div>
            <h2 class="h4 fw-bold mb-0">{{ $t("home.feature.customizable.title") }}</h2>
            <p>{{ $t("home.feature.customizable.description") }}</p>
          </div>
        </div>
      </div>
    </div>

    <!-- About -->
    <div class="container col-xxl-8 px-4 pb-5">
      <div class="row align-items-center g-5 pb-5">
        <div class="d-none d-md-block col-md-3 col-lg-3">
          <img
            loading="lazy"
            src="@/assets/img/viera.png"
            style="border-radius: 50%; width: 100%; border: 1px solid var(--bs-body-color)"
            class="d-block mx-lg-auto"
            alt="Drawing of a hamster"
          />
        </div>
        <div class="col-11 col-md-9">
          <h1 class="pb-2 border-bottom mb-3">{{ $t("home.about.header") }}</h1>
          <p class="lead text-justify fw-normal">
            <i18n-t keypath="home.about.text">
              <a href="https://olivi-eh.dev" target="_blank" rel="noopener noreferrer">Olivier</a>
              <a href="https://forms.gle/2t5nLB28xDyi3Tn6A" target="_blank" rel="noopener noreferrer">
                {{ $t("home.about.contact") }}
              </a>
            </i18n-t>
          </p>
          <div class="d-grid gap-2 d-md-flex justify-content-md-start">
            <a href="https://ko-fi.com/olivi_eh" target="_blank" rel="noopener noreferrer">
              <button type="button" class="btn btn-primary btn-lg px-4 me-md-2">
                <i class="fa-fw fal fa-heart"></i> {{ $t("page.donate") }}
              </button>
            </a>
          </div>
        </div>
      </div>
    </div>
  </template>
</template>

<style lang="scss">
#features-grid .fas {
  color: var(--bs-primary-color);
  font-size: 1.75em;
}

.avatar-home {
  border-radius: 50%;
  height: 35px;
  margin-top: -8px;
}

.progress {
  background-color: #292c2f; // @TODO: replace with var
}

.progress-bar {
  background-color: #3ca577; // @TODO: replace with var
  color: #ffffff; // @TODO: replace with var
  overflow: unset;
  padding-left: 5px;
  padding-right: 5px;
}
</style>

<script>
import PageHeader from "@/components/PageHeader.vue";
import brandImg from "@/assets/img/brand.png";
import brandDarkImg from "@/assets/img/brand_dark.png";

const screenshots = import.meta.glob("../assets/screenshots/*.png", { eager: true, import: "default" });

export default {
  name: "HomeView",
  components: {
    PageHeader,
  },
  computed: {
    brandURI() {
      const theme = this.$store.getters.theme;

      if (theme == "night") {
        return brandImg;
      } else {
        return brandDarkImg;
      }
    },
  },
  methods: {
    screenshotURI(page) {
      return screenshots[`../assets/screenshots/${this.$i18n.locale}-${page}.png`];
    },
    goto(refName) {
      var element = this.$refs[refName];
      var top = element.offsetTop;

      window.scrollTo(0, top - 70);
    },
    changeActiveCharacter(i) {
      this.$store.commit("changeActiveCharacter", i);
    },
    taskCompletionPercentage(tasks) {
      let total = tasks.filter((x) => !x.hidden).length;
      let checked = tasks.filter((x) => !x.hidden && x.checked).length;

      if (total == 0) {
        return "100%";
      }

      return Math.floor((checked / total) * 100) + "%";
    },
  },
};
</script>
