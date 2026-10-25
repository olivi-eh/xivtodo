import { defineStore } from "pinia";
import { injectWeeklyChecklist, injectDailyChecklist } from "@/utilities/checklist";

export const useAppStore = defineStore("app", {
  state: () => ({
    env: {},
    upstreamVersion: "",
    userData: null,
    activeCharacterID: 0,
    language: null,
    flags: {},
    signingIn: false,
  }),
  getters: {
    isDevEnv() {
      return this.env.DEV;
    },
    versionMatches() {
      return this.env.VITE_BUILD_NUM == this.upstreamVersion || this.upstreamVersion == null;
    },
    backendOffline() {
      return this.upstreamVersion == "OFFLINE";
    },
    flag() {
      return (flag) => this.flags[flag] || false;
    },
    isSignedIn() {
      return this.userData != null && this.userData != {};
    },
    discordAuthURI() {
      return (
        "https://discord.com/api/oauth2/authorize" +
        "?response_type=code" +
        "&client_id=" +
        this.env.VITE_DISCORD_CLIENT_ID +
        "&scope=identify" +
        "&redirect_uri=" +
        encodeURIComponent(this.env.VITE_DISCORD_REDIRECT_URI) +
        "&prompt=consent"
      );
    },
    discordUser() {
      return this.userData?.discordUser || {};
    },
    discordUsername() {
      let username = this.discordUser.username;
      if (this.discordUser.discriminator != 0) {
        username += "#" + this.discordUser.discriminator;
      }
      return username;
    },
    settings() {
      return this.userData?.settings || {};
    },
    theme() {
      const themeOption = this.settings.themeOption || 0;
      switch (themeOption) {
        case 0:
          return "night";
        case 1:
          return "day";
        default:
          return "night";
      }
    },
    characters() {
      return this.userData?.characters || [];
    },
    hasCharacter() {
      return this.characters?.length > 0 || false;
    },
    activeCharacter() {
      return this.characters[this.activeCharacterID];
    },
    lodestoneData() {
      return this.activeCharacter?.lodestoneData || {};
    },
    character() {
      return this.lodestoneData.Character || {};
    },
    profilePublic() {
      return this.lodestoneData.Character?.ActiveClassJob != null;
    },
    achievements() {
      let achievements = new Map();
      for (let achievement of this.lodestoneData.Achievements || []) {
        achievements.set(achievement.ID, achievement.Date);
      }
      return achievements;
    },
    achievementsPublic() {
      return this.lodestoneData.Achievements?.length > 0 || false;
    },
    mounts() {
      let mounts = {};
      for (let mount of this.lodestoneData.Mounts || []) {
        mounts[mount.ID] = true;
      }
      return mounts;
    },
    minions() {
      let minions = {};
      for (let minion of this.lodestoneData.Minions || []) {
        minions[minion.ID] = true;
      }
      return minions;
    },
    latestNewsSeen() {
      return this.settings.latestNewsSeen || 0;
    },
    latestCountdownSeen() {
      return this.settings.latestCountdownSeen || 0;
    },
    encounterIDs() {
      let encounterIDs = new Map();
      for (let encounterID of this.activeCharacter?.encounterIDs || []) {
        encounterIDs.set(encounterID, true);
      }
      return encounterIDs;
    },
    encountersVerified() {
      return this.activeCharacter?.encountersVerified || {};
    },
    checklistData() {
      return this.activeCharacter?.checklistData || {};
    },
    checklistWeeklies() {
      return injectWeeklyChecklist(this.checklistData.weeklyChecklist || []);
    },
    checklistDailies() {
      return injectDailyChecklist(this.checklistData.dailyChecklist || []);
    },
    checklistAdhocs() {
      return this.checklistData.adhocChecklist || [];
    },
    checklistLenHiddens() {
      let hiddens = 0;
      for (let item of this.checklistWeeklies) {
        if (item.hidden) {
          hiddens++;
        }
      }
      for (let item of this.checklistDailies) {
        if (item.hidden) {
          hiddens++;
        }
      }
      return hiddens;
    },
    checklistNextWeeklyReset() {
      return this.checklistData.nextWeeklyReset || 0;
    },
    checklistNextDailyReset() {
      return this.checklistData.nextDailyReset || 0;
    },
    activeCharacterOutOfDate() {
      return this.characterOutOfDate(this.activeCharacterID);
    },
    characterOutOfDate() {
      return (id) => {
        if (this.characters[id].invalid) {
          return false;
        }

        let now = new Date();
        let msBeforeUpdate = 1000 * 60 * 60 * 48; // 48 hours

        return now > this.characters[id].updated * 1000 + msBeforeUpdate;
      };
    },
  },
  actions: {
    setSigningIn(payload) {
      this.signingIn = payload;
    },

    setUpstreamVersion(payload) {
      this.upstreamVersion = payload;
    },
    setLanguage(payload) {
      this.language = payload;
    },
    setFlag(payload) {
      this.flags[payload.key] = payload.value;
    },
    setUserData(payload) {
      this.userData = { ...payload };
      if (this.userData.characters == null) {
        this.userData.characters = [];
      }
      this.clampActiveCharacter();

      this.signingIn = false;
    },
    deleteUserData() {
      this.userData = null;
    },
    setSettings(payload) {
      this.userData.settings = { ...payload };
    },
    addCharacter(payload) {
      for (let i = 0; i < this.userData.characters.length; i++) {
        if (this.userData.characters[i].lodestoneData.Character.ID == payload.lodestoneData.Character.ID) {
          this.userData.characters[i] = payload;
          return;
        }
      }

      this.userData.characters.push(payload);
      this.activeCharacterID = this.userData.characters.length - 1;
    },
    removeCharacter(payload) {
      this.userData.characters.splice(payload, 1);
      if (this.activeCharacterID == payload) {
        this.activeCharacterID = 0;
      } else if (this.activeCharacterID > payload) {
        // Keep the same character active now that it has shifted down by one
        this.activeCharacterID--;
      }
    },
    clampActiveCharacter() {
      if (this.activeCharacterID >= this.characters.length) {
        this.activeCharacterID = 0;
      }
    },
    invalidateCharacter(payload) {
      this.userData.characters[payload].invalid = true;
    },
    changeActiveCharacter(payload) {
      this.activeCharacterID = payload;
    },
    setEncounterIDs(payload) {
      this.userData.characters[this.activeCharacterID].encounterIDs = payload;
    },
    setChecklistNextWeeklyReset(payload) {
      this.userData.characters[this.activeCharacterID].checklistData.nextWeeklyReset = payload;
    },
    setChecklistNextDailyReset(payload) {
      this.userData.characters[this.activeCharacterID].checklistData.nextDailyReset = payload;
    },
    setChecklistWeeklies(payload) {
      this.userData.characters[this.activeCharacterID].checklistData.weeklyChecklist = payload;
    },
    setChecklistDailies(payload) {
      this.userData.characters[this.activeCharacterID].checklistData.dailyChecklist = payload;
    },
    setChecklistAdhocs(payload) {
      this.userData.characters[this.activeCharacterID].checklistData.adhocChecklist = payload;
    },
  },
});
