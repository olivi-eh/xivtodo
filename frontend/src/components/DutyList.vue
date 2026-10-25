<template>
  <h3><img v-if="type" :src="'/icons/duty-' + type + '.png'" :alt="'Duty icon for type ' + type" />{{ title }}</h3>
  <span v-if="showTotal" class="text-muted">
    {{ $t("shared.countTotal", { count: duties.length }) }}
  </span>
  <div class="list-group list-group-flush">
    <DutyListItem
      v-for="item of injectDutyCompletion(duties)"
      :key="item.Name"
      :duty="item"
      :filters="this.$parent.$data.filters || {}"
    />
  </div>
  <br />
</template>

<script>
import DutyListItem from "@/components/DutyListItem.vue";
import { useAppStore } from "@/store/app";

export default {
  props: {
    title: String,
    duties: Array,
    type: String,
    showTotal: Boolean,
  },
  components: {
    DutyListItem,
  },
  setup() {
    return { appStore: useAppStore() };
  },
  methods: {
    injectDutyCompletion(duties) {
      const achievements = this.appStore.achievements;
      const spoilersOption = this.appStore.settings.spoilersOption || 0;
      const encounterIDs = this.appStore.encounterIDs;
      const verifiedEncounters = this.appStore.encountersVerified;

      for (let item of duties) {
        let cleared = verifiedEncounters[item.UUID];
        if (cleared == -1 && encounterIDs.has(~~item.UUID)) cleared = 2;
        item.cleared = cleared;

        let blur = false;
        if (item.Spoilers) {
          let showIfAllowed = false;
          if (item.SpoilersUntil && achievements.has(item.SpoilersUntil)) showIfAllowed = true;
          if (spoilersOption == "1" || (cleared == 0 && !showIfAllowed)) {
            blur = true;
          }
          if (!this.appStore.achievementsPublic && !showIfAllowed) blur = true;
        }
        if (item.Expansion == "2x" && !achievements.has(788)) blur = true;
        if (item.Expansion == "3x" && !achievements.has(1139)) blur = true;
        if (item.Expansion == "4x" && !achievements.has(1794)) blur = true;
        if (item.Expansion == "5x" && !achievements.has(2298)) blur = true;
        if (item.Expansion == "6x" && !achievements.has(2958)) blur = true;
        if (item.Expansion == "7x" && !achievements.has(3496)) blur = true;
        if (spoilersOption == "2") blur = false;
        item.blur = blur;
      }
      return duties;
    },
  },
};
</script>
