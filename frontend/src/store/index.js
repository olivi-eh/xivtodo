import { createStore } from "vuex";
import { injectWeeklyChecklist, injectDailyChecklist } from "@/utilities/checklist";

const store = createStore({
  state: {
    env: {},
    upstreamVersion: "",
    userData: null,
    activeCharacterID: 0,
    language: null,
    flags: {},
  },
  getters: {
    isDevEnv(state) {
      return state.env.DEV;
    },
    versionMatches(state) {
      return state.env.VITE_BUILD_NUM == state.upstreamVersion || state.upstreamVersion == null;
    },
    backendOffline(state) {
      return state.upstreamVersion == "OFFLINE";
    },
    language(state) {
      return state.language;
    },
    flag: (state) => (flag) => {
      return state.flags[flag] || false;
    },
    userData(state) {
      return state.userData;
    },
    isSignedIn(state) {
      return state.userData != null && state.userData != {};
    },
    discordAuthURI(state) {
      return (
        "https://discord.com/api/oauth2/authorize" +
        "?response_type=code" +
        "&client_id=" +
        state.env.VITE_DISCORD_CLIENT_ID +
        "&scope=identify" +
        "&redirect_uri=" +
        encodeURIComponent(state.env.VITE_DISCORD_REDIRECT_URI) +
        "&prompt=consent"
      );
    },
    discordUser(state) {
      return state.userData?.discordUser || {};
    },
    discordUsername(state, getters) {
      let username = getters.discordUser.username;
      if (getters.discordUser.discriminator != 0) {
        username += "#" + getters.discordUser.discriminator;
      }
      return username;
    },
    settings(state) {
      return state.userData?.settings || {};
    },
    theme(state, getters) {
      const themeOption = getters.settings.themeOption || 0;
      switch (themeOption) {
        case 0:
          return "night";
        case 1:
          return "day";
        default:
          return "night";
      }
    },
    characters(state) {
      return state.userData?.characters || [];
    },
    hasCharacter(state, getters) {
      return getters.characters?.length > 0 || false;
    },
    activeCharacter(state, getters) {
      if (state.activeCharacterID >= getters.characters.length) {
        state.activeCharacterID = 0;
      }

      return getters.characters[state.activeCharacterID];
    },
    lodestoneData(state, getters) {
      return getters.activeCharacter?.lodestoneData || {};
    },
    character(state, getters) {
      return getters.lodestoneData.Character || {};
    },
    profilePublic(state, getters) {
      return getters.lodestoneData.Character?.ActiveClassJob != null;
    },
    achievements(state, getters) {
      let achievements = new Map();
      for (let achievement of getters.lodestoneData.Achievements || []) {
        achievements.set(achievement.ID, achievement.Date);
      }
      return achievements;
    },
    achievementsPublic(state, getters) {
      return getters.lodestoneData.Achievements?.length > 0 || false;
    },
    mounts(state, getters) {
      let mounts = {};
      for (let mount of getters.lodestoneData.Mounts || []) {
        mounts[mount.ID] = true;
      }
      return mounts;
    },
    minions(state, getters) {
      let minions = {};
      for (let minion of getters.lodestoneData.Minions || []) {
        minions[minion.ID] = true;
      }
      return minions;
    },
    latestNewsSeen(state, getters) {
      return getters.settings.latestNewsSeen || 0;
    },
    latestCountdownSeen(state, getters) {
      return getters.settings.latestCountdownSeen || 0;
    },
    encounterIDs(state, getters) {
      let encounterIDs = new Map();
      for (let encounterID of getters.activeCharacter?.encounterIDs || []) {
        encounterIDs.set(encounterID, true);
      }
      return encounterIDs;
    },
    encountersVerified(state, getters) {
      return getters.activeCharacter?.encountersVerified || {};
    },
    checklistData(state, getters) {
      return getters.activeCharacter?.checklistData || {};
    },
    checklistWeeklies(state, getters) {
      return injectWeeklyChecklist(getters.checklistData.weeklyChecklist || []);
    },
    checklistDailies(state, getters) {
      return injectDailyChecklist(getters.checklistData.dailyChecklist || []);
    },
    checklistAdhocs(state, getters) {
      return getters.checklistData.adhocChecklist || [];
    },
    checklistLenHiddens(state, getters) {
      let hiddens = 0;
      for (let item of getters.checklistWeeklies) {
        if (item.hidden) {
          hiddens++;
        }
      }
      for (let item of getters.checklistDailies) {
        if (item.hidden) {
          hiddens++;
        }
      }
      return hiddens;
    },
    checklistNextWeeklyReset(state, getters) {
      return getters.checklistData.nextWeeklyReset || 0;
    },
    checklistNextDailyReset(state, getters) {
      return getters.checklistData.nextDailyReset || 0;
    },
    activeCharacterOutOfDate(state, getters) {
      return getters.characterOutOfDate(state.activeCharacterID);
    },
    characterOutOfDate: (state, getters) => (id) => {
      if (getters.characters[id].invalid) {
        return false;
      }

      let now = new Date();
      let msBeforeUpdate = 1000 * 60 * 60 * 48; // 48 hours

      return now > getters.characters[id].updated * 1000 + msBeforeUpdate;
    },
  },
  mutations: {
    initialiseStore(state) {
      if (localStorage.getItem("store")) {
        this.replaceState(Object.assign(state, JSON.parse(localStorage.getItem("store"))));
      }

      state.env = { ...import.meta.env };
    },

    signingIn(state, payload) {
      state.signingIn = payload;
    },

    setUpstreamVersion(state, payload) {
      state.upstreamVersion = payload;
    },
    setLanguage(state, payload) {
      state.language = payload;
    },
    setFlag(state, payload) {
      state.flags[payload.key] = payload.value;
    },
    setUserData(state, payload) {
      state.userData = { ...payload };
      if (state.userData.characters == null) {
        state.userData.characters = [];
      }

      delete state.signingIn;
    },
    deleteUserData(state) {
      state.userData = null;
    },
    setSettings(state, payload) {
      state.userData.settings = { ...payload };
    },
    addCharacter(state, payload) {
      for (let i = 0; i < state.userData.characters.length; i++) {
        if (state.userData.characters[i].lodestoneData.Character.ID == payload.lodestoneData.Character.ID) {
          state.userData.characters[i] = payload;
          return;
        }
      }

      state.userData.characters.push(payload);
      state.activeCharacterID = state.userData.characters.length - 1;
    },
    removeCharacter(state, payload) {
      state.userData.characters.splice(payload, 1);
      if (state.activeCharacterID == payload) {
        state.activeCharacterID = 0;
      }
    },
    invalidateCharacter(state, payload) {
      state.userData.characters[payload].invalid = true;
    },
    changeActiveCharacter(state, payload) {
      state.activeCharacterID = payload;
    },
    setEncounterIDs(state, payload) {
      state.userData.characters[state.activeCharacterID].encounterIDs = payload;
    },
    checklistNextWeeklyReset(state, payload) {
      state.userData.characters[state.activeCharacterID].checklistData.nextWeeklyReset = payload;
    },
    checklistNextDailyReset(state, payload) {
      state.userData.characters[state.activeCharacterID].checklistData.nextDailyReset = payload;
    },
    setChecklistWeeklies(state, payload) {
      state.userData.characters[state.activeCharacterID].checklistData.weeklyChecklist = payload;
    },
    setChecklistDailies(state, payload) {
      state.userData.characters[state.activeCharacterID].checklistData.dailyChecklist = payload;
    },
    setChecklistAdhocs(state, payload) {
      state.userData.characters[state.activeCharacterID].checklistData.adhocChecklist = payload;
    },
  },
  actions: {},
  modules: {},
});

export default store;
