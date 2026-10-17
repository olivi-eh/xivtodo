<template>
  <div
    class="list-group-item align-items-center fs-6"
    :class="{
      'list-group-item-action':
        this.$store.getters.hasCharacter && isHovering && (duty.cleared == -1 || duty.cleared == 2),
      'cursor-pointer': this.$store.getters.hasCharacter && (duty.cleared == -1 || duty.cleared == 2),
    }"
    @click="check"
    @mouseover="isHovering = true"
    @mouseout="isHovering = false"
  >
    <div
      class="d-flex justify-content-between align-items-center"
      :class="{
        'text-secondary': duty.cleared == -1 && this.$store.getters.isSignedIn,
        'text-success': duty.cleared >= 1 && this.$store.getters.isSignedIn,
      }"
    >
      <span
        class="duty-list-item fw-medium"
        :class="{
          'duty-list-item-blur': duty.blur,
        }"
        :title="title"
      >
        <!-- Completion icon -->
        <i
          v-if="this.$store.getters.isSignedIn"
          class="me-2 fa-fw fal"
          :class="{
            'fa-question-circle': duty.cleared == -1 && !(duty.cleared == -1 && isHovering),
            'fa-badge-check': duty.cleared == 1,
            'fa-check-circle': duty.cleared == 2 || (duty.cleared == -1 && isHovering),
            'fa-circle': duty.cleared == 0,
          }"
        ></i>

        <!-- Duty name -->
        <span
          :class="{
            'blur-maybe': duty.blur && duty.cleared == -1 && this.$store.getters.isSignedIn,
            'blur-cleared': duty.blur && duty.cleared == 1,
            'blur-uncleared': duty.blur && duty.cleared == 0,
            'user-select-none': duty.blur && this.$store.getters.isSignedIn,
            'fw-bolder': duty.Bold,
          }"
        >
          {{ duty["Name" + $i18n.locale.toUpperCase()] || duty["NameEN"] }}
        </span>
      </span>

      <!-- MSQ icon -->
      <span
        v-if="duty.IsMSQ"
        class="icon-marker-msq"
        data-bs-toggle="tooltip"
        data-bs-placement="top"
        :title="$t('encounters.msqContent')"
      ></span>

      <!-- Patch number icon -->
      <span
        v-if="duty.Expansion && showPatchNums"
        :class="'icon-exp-' + duty.Expansion"
        data-bs-toggle="tooltip"
        data-bs-placement="top"
        :title="$t('encounters.unlockedInExp')"
      ></span>
    </div>

    <!-- Rewards -->
    <div
      v-if="
        !duty.blur &&
        (('Mounts' in duty && filters.mounts) || ('Minions' in duty && filters.minions))
      "
      class="mx-2 mt-2"
    >
      <!-- Mounts -->
      <template v-if="'Mounts' in duty && filters.mounts">
        <div
          v-for="mountID of duty.Mounts.split(' ')"
          :key="mountID"
          class="pb-2"
          :class="{ 'text-muted': this.$store.getters.mounts[mountID] }"
        >
          <img
            v-if="mountID"
            loading="lazy"
            width="40"
            height="40"
            :src="dbs.mounts[mountID].icon"
            :class="{ 'filter-bright-50': this.$store.getters.mounts[mountID] }"
          />
          {{ dbs.mounts[mountID].name }}
          <i
            v-if="this.$store.getters.mounts[mountID]"
            class="fa-fw fal fa-check"
            :title="$t('encounters.obtained')"
          ></i>
          <i
            v-else-if="dbs.mounts[mountID].tradeable"
            class="fa-fw fal fa-comments-dollar"
            :title="$t('encounters.tradeable')"
          ></i>
        </div>
      </template>

      <!-- Minions -->
      <template v-if="'Minions' in duty && filters.minions">
        <div
          v-for="minionID of duty.Minions.split(' ')"
          :key="minionID"
          class="pb-2"
          :class="{ 'text-muted': this.$store.getters.minions[minionID] }"
        >
          <template v-if="minionID in dbs.minions">
            <img
              v-if="minionID"
              loading="lazy"
              width="40"
              height="40"
              :src="dbs.minions[minionID].icon"
              :class="{ 'filter-bright-50': this.$store.getters.minions[minionID] }"
            />
            {{ dbs.minions[minionID].name }}
            <i
              v-if="this.$store.getters.minions[minionID]"
              class="fa-fw fal fa-check"
              :title="$t('encounters.obtained')"
            ></i>
            <i
              v-else-if="dbs.minions[minionID].tradeable"
              class="fa-fw fal fa-comments-dollar"
              :title="$t('encounters.tradeable')"
            ></i>
          </template>
          <span v-else>Cannot load minion #{{ minionID }}</span>
        </div>
      </template>
    </div>
  </div>
</template>

<style lang="scss">
.blur-maybe {
  color: transparent;
  text-shadow: 0 0 20px var(--bs-secondary);
}

.blur-cleared {
  color: transparent;
  text-shadow: 0 0 20px var(--bs-success);
  text-overflow: unset;
}

.blur-uncleared {
  color: transparent;
  text-shadow: 0 0 20px var(--bs-body-color);
  text-overflow: unset;
}

.duty-list-item {
  overflow: hidden;
  white-space: nowrap;
  text-overflow: "… ";
}

.duty-list-item-blur {
  text-overflow: " ";
}
</style>

<script>
import { getLodestoneURL } from "@/utilities/shared.js";
import { updateEncounterIDs } from "@/utilities/backend.js";

import dbs from "@/utilities/dbs.js";

export default {
  data() {
    return {
      isHovering: false,
    };
  },
  props: {
    duty: Object,
    filters: Object,
  },
  computed: {
    dbs() {
      return dbs[this.$store.getters.language];
    },
    title() {
      return this.duty.LodestoneID || this.duty.blur ? "" : this.duty.Name;
    },
    showPatchNums() {
      let patchNumsSetting = this.$store.getters.settings.patchNumsOption || 0;

      return (
        (patchNumsSetting == 0 && this.duty.blur) ||
        (patchNumsSetting == 1 && this.duty.cleared != 1 && this.duty.cleared != 2) ||
        patchNumsSetting == 2
      );
    },
  },
  methods: {
    getLodestoneURL: getLodestoneURL,
    check() {
      if (!this.$store.getters.hasCharacter) return;
      if (this.duty.cleared == 0 || this.duty.cleared == 1) return;

      let uuid = ~~this.duty.UUID;

      let encounterIDs = this.$store.getters.encounterIDs;
      if (encounterIDs.has(uuid)) {
        encounterIDs.set(uuid, false);
      } else {
        encounterIDs.set(uuid, true);
      }

      let encounterIDsAsArray = [];
      encounterIDs.forEach((v, k) => {
        if (v) {
          encounterIDsAsArray.push(k);
        }
      });

      this.$store.commit("setEncounterIDs", encounterIDsAsArray);

      let characterID = this.$store.getters.lodestoneData.Character.ID;
      updateEncounterIDs(characterID, encounterIDsAsArray);
    },
  },
};
</script>
