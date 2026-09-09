<template>
  <div v-if="this.jobExists" class="job-level text-center user-select-none tt">
    <!-- Tooltip -->
    <span class="tt-text">{{ title }}</span>

    <!-- Icon -->
    <span
      :class="{
        'job-type-tank': type == 'tank',
        'job-type-healer': type == 'healer',
        'job-type-dps': type == 'dps',
        'job-type-gatherer': type == 'gatherer',
        'job-type-crafter': type == 'crafter',
      }"
      ><i :class="iconClass"></i>
    </span>

    <br />

    <!-- Level -->
    <span v-if="!this.jobUnlocked" class="text-secondary">-</span>
    <span v-else :class="{ 'fw-bolder': this.isMaxLevel }">
      {{ this.currentLevel }}
    </span>

    <!-- Progress -->
    <div v-if="this.jobUnlocked && !this.isMaxLevel" class="job-level-progress">
      <div class="job-level-progress-bar" :style="'width: ' + this.levelProgressPercentage + '%'"></div>
    </div>
  </div>
</template>

<style lang="scss">
.job-level {
  float: left;
  margin-right: 10px;
  margin-top: 5px;
  width: 30px;
}

.job-level-progress {
  height: 1px;
  border-right: 1px solid var(--bs-body-color);
}

.job-level-progress-bar {
  height: 1px;
  width: 100%;
  border-bottom: 1px solid var(--bs-body-color);
}

.job-type-tank {
  color: var(--xt-tank-color);
}

.job-type-healer {
  color: var(--xt-healer-color);
}

.job-type-dps {
  color: var(--xt-dps-color);
}

.job-type-gatherer {
  color: var(--xt-gatherer-color);
}

.job-type-crafter {
  color: var(--xt-crafter-color);
}
</style>

<script>
export default {
  data() {
    return {
      maxLevelCommon: 100,
      maxLevelBlue: 80,
      maxLevelBeast: 50,
      iconClass: "icon-job-" + this.initial,
    };
  },
  props: {
    initial: String,
    title: String,
    type: String,
  },
  computed: {
    jobExists() {
      return this.$store.getters.lodestoneData.Jobs[this.initial] != null;
    },
    currentLevel() {
      return this.$store.getters.lodestoneData.Jobs[this.initial].Level || 0;
    },
    jobUnlocked() {
      return this.currentLevel != 0;
    },
    maxLevel() {
      if (this.initial == "blu") {
        return this.maxLevelBlue;
      } else if (this.initial == "bst") {
        return this.maxLevelBeast;
      } else {
        return this.maxLevelCommon;
      }
    },
    isMaxLevel() {
      return this.currentLevel == this.maxLevel;
    },
    levelProgressPercentage() {
      const level = this.currentLevel;
      const maxLevel = this.maxLevel;

      // Using a logarithmic scale to make the progress bar look more natural
      return (Math.pow(1.03, level + 1) / Math.pow(1.03, maxLevel + 1)) * 100;
    },
  },
};
</script>
