<template>
  <PageHeader title="page.encounters" subtitle="pageHeader.clearedByCharacter">
    <div v-if="this.$store.getters.hasCharacter" class="d-flex gap-2 float-end">
      <button
        type="button"
        class="btn"
        :class="{ 'btn-outline-primary': !filters.mounts, 'btn-primary': filters.mounts }"
        :aria-pressed="filters.mounts"
        @click="filters.mounts = !filters.mounts"
      >
        <i class="fa-fw fas fa-horse-saddle me-1"></i>
        {{ $t("page.mounts") }}
      </button>
      <button
        type="button"
        class="btn"
        :class="{ 'btn-outline-primary': !filters.minions, 'btn-primary': filters.minions }"
        :aria-pressed="filters.minions"
        @click="filters.minions = !filters.minions"
      >
        <i class="fa-fw fas fa-people-simple me-1"></i>
        {{ $t("page.minions") }}
      </button>
    </div>
  </PageHeader>

  <div class="row">
    <h2>{{ $t("shared.expansion.arr") }}</h2>

    <div class="col-12 col-lg-4">
      <DutyList :title="$t('encounters.dungeons')" :duties="db.arrDungeons" type="dngs" :showTotal="true" />
    </div>
    <div class="col-12 col-lg-4">
      <DutyList :title="$t('encounters.trials')" :duties="db.arrTrials" type="trials" :showTotal="true" />
      <DutyList :title="$t('encounters.heTrials')" :duties="db.arrTrialsEx" type="high" :showTotal="true" />
    </div>
    <div class="col-12 col-lg-4">
      <DutyList :title="$t('encounters.raids')" :duties="db.arrRaids" type="raids" :showTotal="true" />
      <DutyList :title="$t('encounters.heRaids')" :duties="db.arrRaidsSavUlt" type="high" :showTotal="true" />
      <DutyList :title="$t('encounters.guildhests')" :duties="db.arrGuildhests" type="guild" :showTotal="true" />
    </div>
  </div>
  <hr />
  <div class="row">
    <h2>{{ $t("shared.expansion.hw") }}</h2>

    <div class="col-12 col-lg-4">
      <DutyList :title="$t('encounters.dungeons')" :duties="db.hwDungeons" type="dngs" :showTotal="true" />
    </div>
    <div class="col-12 col-lg-4">
      <DutyList :title="$t('encounters.trials')" :duties="db.hwTrials" type="trials" :showTotal="true" />
      <DutyList :title="$t('encounters.heTrials')" :duties="db.hwTrialsEx" type="high" :showTotal="true" />
    </div>
    <div class="col-12 col-lg-4">
      <DutyList :title="$t('encounters.raids')" :duties="db.hwRaids" type="raids" :showTotal="true" />
      <DutyList :title="$t('encounters.heRaids')" :duties="db.hwRaidsSavUlt" type="high" :showTotal="true" />
    </div>
  </div>
  <hr />
  <div class="row">
    <h2>{{ $t("shared.expansion.sb") }}</h2>

    <div class="col-12 col-lg-4">
      <DutyList :title="$t('encounters.dungeons')" :duties="db.sbDungeons" type="dngs" :showTotal="true" />
    </div>
    <div class="col-12 col-lg-4">
      <DutyList :title="$t('encounters.trials')" :duties="db.sbTrials" type="trials" :showTotal="true" />
      <DutyList :title="$t('encounters.heTrials')" :duties="db.sbTrialsEx" type="high" :showTotal="true" />
    </div>
    <div class="col-12 col-lg-4">
      <DutyList :title="$t('encounters.raids')" :duties="db.sbRaids" type="raids" :showTotal="true" />
      <DutyList :title="$t('encounters.heRaids')" :duties="db.sbRaidsSavUlt" type="high" :showTotal="true" />
      <DutyList :title="$t('encounters.exploratory')" :duties="db.sbExploratory" type="expl" :showTotal="true" />
    </div>
  </div>
  <hr />
  <div class="row">
    <h2>{{ $t("shared.expansion.shb") }}</h2>

    <div class="col-12 col-lg-4">
      <DutyList :title="$t('encounters.dungeons')" :duties="db.shbDungeons" type="dngs" :showTotal="true" />
    </div>
    <div class="col-12 col-lg-4">
      <DutyList :title="$t('encounters.trials')" :duties="db.shbTrials" type="trials" :showTotal="true" />
      <DutyList :title="$t('encounters.heTrials')" :duties="db.shbTrialsEx" type="high" :showTotal="true" />
    </div>
    <div class="col-12 col-lg-4">
      <DutyList :title="$t('encounters.raids')" :duties="db.shbRaids" type="raids" :showTotal="true" />
      <DutyList :title="$t('encounters.heRaids')" :duties="db.shbRaidsSavUlt" type="high" :showTotal="true" />
      <DutyList :title="$t('encounters.exploratory')" :duties="db.shbExploratory" type="expl" :showTotal="true" />
    </div>
  </div>
  <hr />
  <div class="row">
    <h2>{{ $t("shared.expansion.ew") }}</h2>

    <div class="col-12 col-lg-4">
      <DutyList :title="$t('encounters.dungeons')" :duties="db.ewDungeons" type="dngs" :showTotal="true" />
      <DutyList :title="$t('encounters.vcDungeons')" :duties="db.ewVariantsCriterions" type="vc" :showTotal="true" />
    </div>
    <div class="col-12 col-lg-4">
      <DutyList :title="$t('encounters.trials')" :duties="db.ewTrials" type="trials" :showTotal="true" />
      <DutyList :title="$t('encounters.heTrials')" :duties="db.ewTrialsEx" type="high" :showTotal="true" />
    </div>
    <div class="col-12 col-lg-4">
      <DutyList :title="$t('encounters.raids')" :duties="db.ewRaids" type="raids" :showTotal="true" />
      <DutyList :title="$t('encounters.heRaids')" :duties="db.ewRaidsSavUlt" type="high" :showTotal="true" />
    </div>
  </div>
  <hr />
  <div class="row">
    <h2>{{ $t("shared.expansion.dt") }}</h2>

    <div class="col-12 col-lg-4">
      <DutyList :title="$t('encounters.dungeons')" :duties="db.dtDungeons" type="dngs" :showTotal="true" />
      <DutyList :title="$t('encounters.vcDungeons')" :duties="db.dtVariantsCriterions" type="vc" :showTotal="true" />
    </div>
    <div class="col-12 col-lg-4">
      <DutyList :title="$t('encounters.trials')" :duties="db.dtTrials" type="trials" :showTotal="true" />
      <DutyList :title="$t('encounters.heTrials')" :duties="db.dtTrialsEx" type="high" :showTotal="true" />
    </div>
    <div class="col-12 col-lg-4">
      <DutyList :title="$t('encounters.raids')" :duties="db.dtRaids" type="raids" :showTotal="true" />
      <DutyList :title="$t('encounters.heRaids')" :duties="db.dtRaidsSavUlt" type="high" :showTotal="true" />
      <DutyList :title="$t('encounters.exploratory')" :duties="db.dtExploratory" type="expl" :showTotal="true" />
    </div>
  </div>
</template>

<script>
import PageHeader from "@/components/PageHeader.vue";
import DutyList from "@/components/DutyList.vue";

import dbJson from "@/assets/db.json";

export default {
  name: "EncountersView",
  data() {
    return {
      db: dbJson,
      filters: {
        mounts: false,
        minions: false,
      },
    };
  },
  components: {
    PageHeader,
    DutyList,
  },
};
</script>
