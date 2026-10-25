<template>
  <PageHeader title="page.checklist" subtitle="pageHeader.forCharacter">
    <button
      v-if="this.appStore.hasCharacter"
      type="button"
      class="btn float-end"
      :class="{ 'btn-outline-primary': !showHidden, 'btn-primary': showHidden }"
      @click="showHidden = !showHidden"
    >
      <i class="fa-regular fa-pen-to-square"></i>
      {{ $t("shared.customize") }}
    </button>
  </PageHeader>

  <h2>{{ $t("checklist.thingsToDo") }}</h2>
  <div class="row">
    <div class="col-md">
      <h3>{{ $t("checklist.weeklies") }}</h3>
      <span class="text-muted">
        {{ $t("checklist.untilReset", { time: weeklyReset }) }}
      </span>

      <div v-if="showHidden" class="input-group mt-3 mb-2">
        <input v-model="customWeekly" type="text" class="form-control" :placeholder="$t('checklist.customWeekly')" />
        <button
          class="btn btn-outline-primary"
          :class="{ disabled: !customWeekly }"
          type="button"
          id="button-addon2"
          @click="addCustomWeekly"
        >
          {{ $t("checklist.addCustomWeekly") }}
        </button>
      </div>

      <ul class="list-group list-group-flush">
        <ChecklistItem
          v-for="item of [...this.appStore.checklistWeeklies]"
          :key="item.name"
          :item="item"
          type="weekly"
          :showHidden="showHidden"
          :draggable="showHidden"
          @dragstart="startDrag($event, item)"
          @dragend="endDrag()"
          @dragleave="endDrag()"
          @drop="onDrop($event, item)"
          @dragover="onDragOver($event, item)"
          @dragover.prevent
          @dragenter.prevent
          :dragHovered="item.name == this.draggedOverName"
        />
      </ul>
      <br />
    </div>

    <div class="col-md order-first order-md-last">
      <h3>{{ $t("checklist.dailies") }}</h3>
      <span class="text-muted">
        {{ $t("checklist.untilReset", { time: dailyReset }) }}
      </span>

      <div v-if="showHidden" class="input-group mt-3 mb-2">
        <input v-model="customDaily" type="text" class="form-control" :placeholder="$t('checklist.customDaily')" />
        <button
          class="btn btn-outline-primary"
          :class="{ disabled: !customDaily }"
          type="button"
          id="button-addon2"
          @click="addCustomDaily"
        >
          {{ $t("checklist.addCustomDaily") }}
        </button>
      </div>

      <ul class="list-group list-group-flush">
        <ChecklistItem
          v-for="item of [...this.appStore.checklistDailies]"
          :key="item.ID"
          :item="item"
          type="daily"
          :showHidden="showHidden"
          :draggable="showHidden"
          @dragstart="startDrag($event, item)"
          @dragend="endDrag()"
          @dragleave="endDrag()"
          @drop="onDrop($event, item)"
          @dragover="onDragOver($event, item)"
          @dragover.prevent
          @dragenter.prevent
          :dragHovered="item.name == this.draggedOverName"
        />
      </ul>
      <br />
      <span v-if="this.appStore.checklistLenHiddens > 0" class="d-none d-md-block">
        <span class="text-muted float-end">
          {{ $t("checklist.hiddenTasks", this.appStore.checklistLenHiddens) }}<br /><br />
        </span>
      </span>
    </div>
  </div>
  <div v-if="showHidden || this.appStore.checklistAdhocs.length > 0" class="row">
    <div class="col">
      <h3>{{ $t("checklist.scratchpad") }}</h3>
      <span class="text-muted">{{ $t("checklist.scratchpadDesc") }}</span>

      <div v-if="showHidden" class="input-group mt-3 mb-2">
        <input v-model="customAdhoc" type="text" class="form-control" :placeholder="$t('checklist.customItem')" />
        <button
          class="btn btn-outline-primary"
          :class="{ disabled: !customAdhoc }"
          type="button"
          id="button-addon2"
          @click="addCustomAdhoc"
        >
          {{ $t("checklist.addCustomItem") }}
        </button>
      </div>

      <ul class="list-group list-group-flush">
        <ChecklistItem
          v-for="item of [...this.appStore.checklistAdhocs]"
          :key="item.ID"
          :item="item"
          type="adhoc"
          :showHidden="showHidden"
          :draggable="showHidden"
          @dragstart="startDrag($event, item)"
          @dragend="endDrag()"
          @dragleave="endDrag()"
          @drop="onDrop($event, item)"
          @dragover="onDragOver($event, item)"
          @dragover.prevent
          @dragenter.prevent
          :dragHovered="item.name == this.draggedOverName"
        />
        <li v-if="this.appStore.checklistAdhocs.length == 0" class="list-group-item">
          {{ $t("checklist.scratchpadHelp") }}
        </li>
      </ul>
      <br />
    </div>
  </div>
</template>

<script>
import PageHeader from "@/components/PageHeader.vue";
import ChecklistItem from "@/components/ChecklistItem.vue";
import { updateChecklist } from "@/utilities/backend.js";
import { swapChecklistItems } from "@/utilities/checklist.js";
import { useAppStore } from "@/store/app";

export default {
  name: "ChecklistView",
  setup() {
    return { appStore: useAppStore() };
  },
  data() {
    return {
      intervalFunction: null,
      showHidden: false,
      weeklyReset: this.formatTimeDiff(this.weeklyResetTime(), true),
      dailyReset: this.formatTimeDiff(this.dailyResetTime(), false),
      customWeekly: "",
      customDaily: "",
      customAdhoc: "",
      draggedOverName: "",
    };
  },
  components: {
    PageHeader,
    ChecklistItem,
  },
  mounted() {
    this.$nextTick(function () {
      this.resetDailliesWeeklies();
    });
    this.intervalFunction = setInterval(() => {
      this.resetDailliesWeeklies();
    }, 1000);
  },
  unmounted() {
    clearInterval(this.intervalFunction);
  },
  methods: {
    weeklyResetTime() {
      let now = new Date();
      let then = new Date();

      then.setUTCHours(8, 0, 0);
      then.setUTCDate(now.getUTCDate() + (7 + 2 - now.getUTCDay()));

      let diff = then - now;
      let delta = 1000 * 60 * 60 * 24 * 7;

      return (diff % delta) + now.getTime();
    },
    dailyResetTime() {
      let now = new Date();
      let then = new Date();
      then.setUTCHours(15, 0, 0);
      then.setUTCDate(then.getUTCDate() + 1);

      let diff = then - now;
      let delta = 1000 * 60 * 60 * 24;

      return (diff % delta) + now.getTime();
    },
    formatTimeDiff(then, showDays) {
      let now = new Date();
      let diff = then - now;

      let dday = 1000 * 60 * 60 * 24;
      let dhour = 1000 * 60 * 60;
      let dminute = 1000 * 60;
      let dsecond = 1000;

      let days = Math.floor(diff / dday);
      let hours = Math.floor(diff / dhour) - 24 * Math.floor(diff / dday);
      let minutes = Math.floor(diff / dminute) - 60 * Math.floor(diff / dhour);
      let seconds = Math.floor(diff / dsecond) - 60 * Math.floor(diff / dminute);

      if (days > 0 && showDays) {
        return (
          this.$t("shared.time.daysShort", { days: days }) + " " + this.$t("shared.time.hoursShort", { hours: hours })
        );
      } else if (hours > 0) {
        return (
          this.$t("shared.time.hoursShort", { hours: hours }) +
          " " +
          this.$t("shared.time.minutesShort", { minutes: minutes })
        );
      } else {
        return (
          this.$t("shared.time.minutesShort", { minutes: minutes }) +
          " " +
          this.$t("shared.time.secondsShort", { seconds: seconds })
        );
      }
    },
    resetDailliesWeeklies() {
      this.weeklyReset = this.formatTimeDiff(this.weeklyResetTime(), true);
      this.dailyReset = this.formatTimeDiff(this.dailyResetTime(), false);
      let resetHappened = false;

      // Skip this if no active character is set.
      if (!this.appStore.hasCharacter) return;

      // Clear checked if past weekly reset time.
      if (this.appStore.checklistNextWeeklyReset < Date.now()) {
        resetHappened = true;
        this.appStore.setChecklistNextWeeklyReset(this.weeklyResetTime());
        let checklistWeeklies = this.appStore.checklistWeeklies;
        for (let i = 0; i < checklistWeeklies.length; i++) {
          checklistWeeklies[i].checked = false;
        }
      }

      // Clear checked if past daily reset time.
      if (this.appStore.checklistNextDailyReset < Date.now()) {
        resetHappened = true;
        this.appStore.setChecklistNextDailyReset(this.dailyResetTime());
        let checklistDailies = this.appStore.checklistDailies;
        for (let i = 0; i < checklistDailies.length; i++) {
          checklistDailies[i].checked = false;
        }
      }

      if (resetHappened) {
        let characterID = this.appStore.lodestoneData.Character.ID;
        updateChecklist(characterID, this.appStore.checklistData);
      }
    },
    addCustomWeekly() {
      let newItem = {
        name: this.customWeekly,
        custom: true,
        checked: false,
        hidden: false,
      };

      let weeklyChecklist = [...this.appStore.checklistWeeklies];
      weeklyChecklist.push(newItem);
      this.appStore.setChecklistWeeklies(weeklyChecklist);

      this.customWeekly = "";
      let characterID = this.appStore.lodestoneData.Character.ID;
      updateChecklist(characterID, this.appStore.checklistData);
    },
    addCustomDaily() {
      let newItem = {
        name: this.customDaily,
        custom: true,
        checked: false,
        hidden: false,
      };

      let dailyChecklist = [...this.appStore.checklistDailies];
      dailyChecklist.push(newItem);
      this.appStore.setChecklistDailies(dailyChecklist);

      this.customDaily = "";
      let characterID = this.appStore.lodestoneData.Character.ID;
      updateChecklist(characterID, this.appStore.checklistData);
    },
    addCustomAdhoc() {
      let newItem = {
        name: this.customAdhoc,
        custom: true,
        checked: false,
        hidden: false,
      };

      let adhocChecklist = [...this.appStore.checklistAdhocs];
      adhocChecklist.push(newItem);
      this.appStore.setChecklistAdhocs(adhocChecklist);

      this.customAdhoc = "";
      let characterID = this.appStore.lodestoneData.Character.ID;
      updateChecklist(characterID, this.appStore.checklistData);
    },
    startDrag(evt, item) {
      evt.dataTransfer.dropEffect = "move";
      evt.dataTransfer.effectAllowed = "move";
      evt.dataTransfer.setData("itemName", item.name);
    },
    endDrag() {
      this.draggedOverName = "";
    },
    onDrop(evt, item) {
      let fromName = evt.dataTransfer.getData("itemName");
      let toName = item.name;

      if (toName == fromName) return;

      swapChecklistItems(fromName, toName);

      let characterID = this.appStore.lodestoneData.Character.ID;
      updateChecklist(characterID, this.appStore.checklistData);
    },
    onDragOver(evt, item) {
      this.draggedOverName = item.name;
    },
  },
};
</script>
