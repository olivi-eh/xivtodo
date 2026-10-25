<template>
  <PageHeader title="page.help" :hideNotSignedInAlert="true" :hideNoCharactersAlert="true" />

  <div class="row">
    <div class="col-lg-10 indent-paragraphs">
      <!-- Discord sign-in -->
      <h2>{{ $t("help.discordSignIn.header") }}</h2>
      <p>{{ $t("help.discordSignIn.text") }}</p>
      <br />

      <!-- Adding characters -->
      <h2>{{ $t("help.addingCharacters.header") }}</h2>
      <p>
        <i18n-t keypath="help.addingCharacters.text1">
          <a :href="getLodestoneURL()" target="_blank" rel="noopener noreferrer">
            {{ $t("shared.lodestone") }}
          </a>
        </i18n-t>
      </p>
      <p>
        <span class="font-monospace"> {{ getLodestoneURL() }}character/12345/</span>
      </p>
      <p>
        <i18n-t keypath="help.addingCharacters.text2">
          <router-link to="/settings">{{ $t("page.settings") }}</router-link>
          <b>{{ $t("settings.addCharacter") }}</b>
        </i18n-t>
      </p>
      <br />

      <!-- Free trial characters -->
      <h2>{{ $t("help.freeTrialCharacters.header") }}</h2>
      <p>{{ $t("help.freeTrialCharacters.text") }}</p>
      <br />

      <!-- Enabling public achievements -->
      <h2>{{ $t("help.enablingPublicAchievements.header") }}</h2>
      <p>
        <i18n-t keypath="help.enablingPublicAchievements.text1">
          <a :href="getLodestoneURL()" target="_blank" rel="noopener noreferrer">
            {{ $t("shared.lodestone") }}
          </a>
          <a :href="getLodestoneURL() + 'my/setting/account/'" target="_blank" rel="noopener noreferrer">
            {{ $t("help.enablingPublicAchievements.lodestoneAccountManagement") }}
          </a>
          <b>{{ $t("page.achievements") }}</b>
          <b>{{ $t("help.enablingPublicAchievements.lodestonePublic") }}</b>
        </i18n-t>
      </p>
      <p>
        <i18n-t keypath="help.enablingPublicAchievements.text2">
          <router-link to="/settings">{{ $t("page.settings") }}</router-link>
        </i18n-t>
      </p>
      <br />

      <!-- Customizing character checklists -->
      <h2>{{ $t("help.customizingCharacterChecklists.header") }}</h2>
      <p>
        <i18n-t keypath="help.customizingCharacterChecklists.text">
          <b>{{ $t("shared.customize") }}</b>
          <router-link to="/checklist">{{ $t("page.checklist") }}</router-link>
          <b>{{ $t("shared.customize") }}</b>
        </i18n-t>
      </p>
      <br />

      <!-- Encounter and questline completion -->
      <h2>{{ $t("help.encounterQuestlineCompletion.header") }}</h2>
      <p>{{ $t("help.encounterQuestlineCompletion.text1") }}</p>
      <p>
        <i18n-t keypath="help.encounterQuestlineCompletion.text2">
          <router-link to="/settings">{{ $t("page.settings") }}</router-link>
        </i18n-t>
      </p>
      <p>{{ $t("help.encounterQuestlineCompletion.textState") }}</p>
      <p>
        <span class="text-success">
          <i class="fa-fw fal fa-badge-check text-success"></i>
          {{ db.arrDungeons[0]["Name" + $i18n.locale.toUpperCase()] }}
        </span>
        <br />
        &nbsp;&nbsp;{{ $t("help.encounterQuestlineCompletion.stateClearedVerified") }}<br />
        <span class="text-success">
          <i class="fa-fw fal fa-check-circle text-success"></i>
          {{ db.arrDungeons[0]["Name" + $i18n.locale.toUpperCase()] }}
        </span>
        <br />
        &nbsp;&nbsp;{{ $t("help.encounterQuestlineCompletion.stateCleared") }}<br />
        <span class="text-secondary">
          <i class="fa-fw fal fa-question-circle"></i>
          {{ db.arrDungeons[0]["Name" + $i18n.locale.toUpperCase()] }}
        </span>
        <br />
        &nbsp;&nbsp;{{ $t("help.encounterQuestlineCompletion.stateUnknown") }}<br />
        <span>
          <i class="fa-fw fal fa-circle"></i>
          {{ db.arrDungeons[0]["Name" + $i18n.locale.toUpperCase()] }}
        </span>
        <br />
        &nbsp;&nbsp;{{ $t("help.encounterQuestlineCompletion.stateNotCleared") }}
      </p>
      <br />

      <!-- Other questions -->
      <h2>{{ $t("help.otherQuestions.header") }}</h2>
      <p>
        <i18n-t keypath="help.otherQuestions.text">
          <a href="https://forms.gle/2t5nLB28xDyi3Tn6A" target="_blank" rel="noopener noreferrer">
            {{ $t("help.otherQuestions.textHere") }}
          </a>
          <a href="https://discord.gg/zfzhKhG3zj" target="_blank" rel="noopener noreferrer">
            {{ $t("page.discord") }}
          </a>
          <a href="https://bsky.app/profile/xivtodo.com" target="_blank" rel="noopener noreferrer">
            {{ $t("page.bluesky") }}
          </a>
        </i18n-t>
      </p>
      <br />
    </div>
  </div>
</template>

<script>
import PageHeader from "@/components/PageHeader.vue";
import dbJson from "@/assets/db.json";
import { getLodestoneURL } from "@/utilities/shared.js";

export default {
  name: "HelpView",
  data() {
    return {
      db: dbJson,
    };
  },
  components: {
    PageHeader,
  },
  methods: {
    getLodestoneURL: getLodestoneURL,
  },
};
</script>
