<template>
  <PageHeader title="page.profile" subtitle="pageHeader.ofCharacter" />

  <div v-if="!this.$store.getters.profilePublic" class="row">
    <p>{{ $t('message.profileNotPublic') }}</p>
  </div>
  <div v-else class="row">
    <!-- Character portrait & bio -->
    <div class="col-lg d-none d-lg-block">
      <!-- Portrait -->
      <img id="character-portrait" :src="this.$store.getters.character.Portrait" :alt="$t('shared.portraitAlt')" />
      <br />

      <!-- Bio -->
      <blockquote v-if="this.hasBio" class="blockquote fs-6 text-break p-2">
        <p>“{{ this.$store.getters.character.Bio }}”</p>
      </blockquote>
    </div>

    <!-- Character info -->
    <div class="col-md-6 col-lg">
      <h2>{{ $t("profile.characterInfo") }}</h2>

      <!-- Info: World -->
      <div class="info-block">
        <b>{{ $t("profile.world") }}</b>
        <br />
        {{ this.$store.getters.character.World }} ({{ this.$store.getters.character.DC }})
      </div>

      <!-- Info: Race & clan -->
      <div class="info-block">
        <b>{{ $t("profile.raceClan") }}</b>
        <br />
        <!-- <span> -->
        {{ this.race }} / {{ this.clan }}
        <i class="fa-fw fal" :class="{ 'fa-mars': this.isMasculine, 'fa-venus': !this.isMasculine }"></i>
        <!-- </span> -->
      </div>

      <!-- Info: Nameday -->
      <div class="info-block">
        <b>{{ $t("profile.nameday") }}</b>
        <br />
        {{ this.$store.getters.character.Nameday }}
      </div>

      <!-- Info: Guardian -->
      <div class="info-block">
        <b>{{ $t("profile.guardian") }}</b>
        <br />
        <img :src="this.guardianImageURI" />
        {{ this.localeName(this.$store.getters.character.GuardianDeity) }}
      </div>

      <!-- Info: Free Company -->
      <div v-if="this.$store.getters.character.FreeCompanyName" class="info-block">
        <b>{{ $t("profile.freeCompany") }}</b>
        <br />
        {{ this.$store.getters.character.FreeCompanyName }}
      </div>

      <!-- Info: City-state -->
      <div class="info-block">
        <b>{{ $t("profile.cityState") }}</b>
        <br />
        <img :src="this.cityStateImageURI" />
        {{ this.localeName(this.$store.getters.character.Town) }}
      </div>

      <!-- Info: Grand Company -->
      <div v-if="this.$store.getters.character.GrandCompanyInfo?.GrandCompany" class="info-block">
        <b>{{ $t("profile.grandCompany") }}</b>
        <br />
        <img :src="this.grandCompanyImageURI" />
        {{ this.localeName(this.$store.getters.character.GrandCompanyInfo.GrandCompany) }}
      </div>

      <!-- Info: Playing since -->
      <div v-if="this.$store.getters.achievementsPublic" class="info-block">
        <b>{{ $t("profile.playingSince") }}</b>
        <br />
        {{ this.playingSince }}
      </div>

      <!-- TODO: Fix and add Occult Crescent info
      <!- - Info: Exploratory missions - ->
      <div class="info-block">
        <b>{{ $t("profile.exploratoryMissions") }}</b>
        <br />

        <!- - Eureka - ->
        {{ $t("profile.elementalLevel") }}:
        <span :class="{ 'fw-bolder': this.isMaxLevelEureka }">
          {{ this.levelEureka }}
        </span>
        <br />

        <!- - Bozja - ->
        {{ $t("profile.resistanceRank") }}:
        <span :class="{ 'fw-bolder': this.isMaxLevelBozja }">
          {{ this.levelBozja }}
        </span>
      </div>
      -->
    </div>

    <!-- Job levels -->
    <div class="col-md-6 col-lg">
      <h2>{{ $t("profile.jobLevels") }}</h2>

      <!-- Tanks -->
      <div class="info-block row">
        <b>{{ $t("profile.tanks") }}</b>
        <br />
        <div class="pt-1">
          <JobLevel initial="gla" :title="$t('profile.job.gla')" type="tank" />
          <JobLevel initial="pld" :title="$t('profile.job.pld')" type="tank" />
          <JobLevel initial="mrd" :title="$t('profile.job.mrd')" type="tank" />
          <JobLevel initial="war" :title="$t('profile.job.war')" type="tank" />
          <JobLevel initial="drk" :title="$t('profile.job.drk')" type="tank" />
          <JobLevel initial="gnb" :title="$t('profile.job.gnb')" type="tank" />
        </div>
      </div>

      <!-- Healers -->
      <div class="info-block row">
        <b>{{ $t("profile.healers") }}</b>
        <br />
        <div class="pt-1">
          <JobLevel initial="cnj" :title="$t('profile.job.cnj')" type="healer" />
          <JobLevel initial="whm" :title="$t('profile.job.whm')" type="healer" />
          <JobLevel initial="sch" :title="$t('profile.job.sch')" type="healer" />
          <JobLevel initial="ast" :title="$t('profile.job.ast')" type="healer" />
          <JobLevel initial="sge" :title="$t('profile.job.sge')" type="healer" />
        </div>
      </div>

      <!-- DPS -->
      <div class="info-block row">
        <b>{{ $t("profile.dps") }}</b>
        <br />
        <div class="pt-1">
          <div style="display: inline-block">
            <JobLevel initial="pgl" :title="$t('profile.job.pgl')" type="dps" />
            <JobLevel initial="mnk" :title="$t('profile.job.mnk')" type="dps" />
            <JobLevel initial="lnc" :title="$t('profile.job.lnc')" type="dps" />
            <JobLevel initial="drg" :title="$t('profile.job.drg')" type="dps" />
            <JobLevel initial="rog" :title="$t('profile.job.rog')" type="dps" />
            <JobLevel initial="nin" :title="$t('profile.job.nin')" type="dps" />
            <JobLevel initial="sam" :title="$t('profile.job.sam')" type="dps" />
            <JobLevel initial="rpr" :title="$t('profile.job.rpr')" type="dps" />
            <JobLevel initial="vpr" :title="$t('profile.job.vpr')" type="dps" />
            <JobLevel initial="arc" :title="$t('profile.job.arc')" type="dps" />
            <JobLevel initial="bst" :title="$t('profile.job.bst')" type="dps" />
          </div>
          <div style="display: inline-block">
            <JobLevel initial="brd" :title="$t('profile.job.brd')" type="dps" />
            <JobLevel initial="mch" :title="$t('profile.job.mch')" type="dps" />
            <JobLevel initial="dnc" :title="$t('profile.job.dnc')" type="dps" />
            <JobLevel initial="thm" :title="$t('profile.job.thm')" type="dps" />
            <JobLevel initial="blm" :title="$t('profile.job.blm')" type="dps" />
            <JobLevel initial="acn" :title="$t('profile.job.acn')" type="dps" />
            <JobLevel initial="smn" :title="$t('profile.job.smn')" type="dps" />
            <JobLevel initial="rdm" :title="$t('profile.job.rdm')" type="dps" />
            <JobLevel initial="pct" :title="$t('profile.job.pct')" type="dps" />
            <JobLevel initial="blu" :title="$t('profile.job.blu')" type="dps" />
          </div>
        </div>
      </div>

      <!-- Crafters -->
      <div class="info-block row">
        <b>{{ $t("profile.crafters") }}</b>
        <br />
        <div class="pt-1">
          <JobLevel initial="crp" :title="$t('profile.job.crp')" type="crafter" />
          <JobLevel initial="bsm" :title="$t('profile.job.bsm')" type="crafter" />
          <JobLevel initial="arm" :title="$t('profile.job.arm')" type="crafter" />
          <JobLevel initial="gsm" :title="$t('profile.job.gsm')" type="crafter" />
          <JobLevel initial="ltw" :title="$t('profile.job.ltw')" type="crafter" />
          <JobLevel initial="wvr" :title="$t('profile.job.wvr')" type="crafter" />
          <JobLevel initial="alc" :title="$t('profile.job.alc')" type="crafter" />
          <JobLevel initial="cul" :title="$t('profile.job.cul')" type="crafter" />
        </div>
      </div>

      <!-- Gatherers -->
      <div class="info-block row">
        <b>{{ $t("profile.gatherers") }}</b>
        <br />
        <div class="pt-1">
          <JobLevel initial="min" :title="$t('profile.job.min')" type="gatherer" />
          <JobLevel initial="btn" :title="$t('profile.job.btn')" type="gatherer" />
          <JobLevel initial="fsh" :title="$t('profile.job.fsh')" type="gatherer" />
        </div>
      </div>
    </div>
  </div>

  <!-- Encounters summaries -->
  <h1>
    {{ $t("page.encounters")}}
    <span class="fs-3 fw-lighter d-block d-lg-inline">
      {{ $t("pageHeader.clearedByCharacter", { characterName: this.$store.getters.character.Name }) }}
    </span>
  </h1>
  <hr />

  <div class="row">
    <div class="col-lg mb-4">
      <h2>{{ $t("shared.expansion.arr") }}</h2>
      <DutyProgress :title="$t('encounters.dungeons')" :duties="db.arrDungeons" type="dngs" />
      <DutyProgress :title="$t('encounters.trials')" :duties="db.arrTrials" type="trials" />
      <DutyProgress :title="$t('encounters.raids')" :duties="db.arrRaids" type="raids" />
      <DutyProgress :title="$t('encounters.guildhests')" :duties="db.arrGuildhests" type="guild" />
      <DutyProgress :title="$t('encounters.heTrials')" :duties="db.arrTrialsEx" type="high" />
      <DutyProgress :title="$t('encounters.heRaids')" :duties="db.arrRaidsSavUlt" type="high" />
    </div>
    <div class="col-lg mb-4">
      <h2>{{ $t("shared.expansion.hw") }}</h2>
      <DutyProgress :title="$t('encounters.dungeons')" :duties="db.hwDungeons" type="dngs" />
      <DutyProgress :title="$t('encounters.trials')" :duties="db.hwTrials" type="trials" />
      <DutyProgress :title="$t('encounters.raids')" :duties="db.hwRaids" type="raids" />
      <DutyProgress :title="$t('encounters.heTrials')" :duties="db.hwTrialsEx" type="high" />
      <DutyProgress :title="$t('encounters.heRaids')" :duties="db.hwRaidsSavUlt" type="high" />
    </div>
    <div class="col-lg mb-4">
      <h2>{{ $t("shared.expansion.sb") }}</h2>
      <DutyProgress :title="$t('encounters.dungeons')" :duties="db.sbDungeons" type="dngs" />
      <DutyProgress :title="$t('encounters.trials')" :duties="db.sbTrials" type="trials" />
      <DutyProgress :title="$t('encounters.raids')" :duties="db.sbRaids" type="raids" />
      <DutyProgress :title="$t('encounters.exploratory')" :duties="db.sbExploratory" type="expl" />
      <DutyProgress :title="$t('encounters.heTrials')" :duties="db.sbTrialsEx" type="high" />
      <DutyProgress :title="$t('encounters.heRaids')" :duties="db.sbRaidsSavUlt" type="high" />
    </div>
  </div>
  <div class="row">
    <div class="col-lg mb-4">
      <h2>{{ $t("shared.expansion.shb") }}</h2>
      <DutyProgress :title="$t('encounters.dungeons')" :duties="db.shbDungeons" type="dngs" />
      <DutyProgress :title="$t('encounters.trials')" :duties="db.shbTrials" type="trials" />
      <DutyProgress :title="$t('encounters.raids')" :duties="db.shbRaids" type="raids" />
      <DutyProgress :title="$t('encounters.exploratory')" :duties="db.shbExploratory" type="expl" />
      <DutyProgress :title="$t('encounters.heTrials')" :duties="db.shbTrialsEx" type="high" />
      <DutyProgress :title="$t('encounters.heRaids')" :duties="db.shbRaidsSavUlt" type="high" />
    </div>
    <div class="col-lg mb-4">
      <h2>{{ $t("shared.expansion.ew") }}</h2>
      <DutyProgress :title="$t('encounters.dungeons')" :duties="db.ewDungeons" type="dngs" />
      <DutyProgress :title="$t('encounters.trials')" :duties="db.ewTrials" type="trials" />
      <DutyProgress :title="$t('encounters.raids')" :duties="db.ewRaids" type="raids" />
      <DutyProgress :title="$t('encounters.vcDungeons')" :duties="db.ewVariantsCriterions" type="vc" />
      <DutyProgress :title="$t('encounters.heTrials')" :duties="db.ewTrialsEx" type="high" />
      <DutyProgress :title="$t('encounters.heRaids')" :duties="db.ewRaidsSavUlt" type="high" />
    </div>
    <div class="col-lg mb-4">
      <h2>{{ $t("shared.expansion.dt") }}</h2>
      <DutyProgress :title="$t('encounters.dungeons')" :duties="db.dtDungeons" type="dngs" />
      <DutyProgress :title="$t('encounters.trials')" :duties="db.dtTrials" type="trials" />
      <DutyProgress :title="$t('encounters.raids')" :duties="db.dtRaids" type="raids" />
      <DutyProgress :title="$t('encounters.heTrials')" :duties="db.dtTrialsEx" type="high" />
      <DutyProgress :title="$t('encounters.heRaids')" :duties="db.dtRaidsSavUlt" type="high" />
    </div>
  </div>

  <!-- Questlines summaries -->
  <h1>
    {{ $t("page.questlines")}}
    <span class="fs-3 fw-lighter d-block d-lg-inline">
      {{ $t("pageHeader.clearedByCharacter", { characterName: this.$store.getters.character.Name }) }}
    </span>
  </h1>
  <hr />

  <div class="row">
    <div class="col-lg mb-4">
      <h2>{{ $t("questlines.stories") }}</h2>
      <DutyProgress :title="$t('questlines.mainScenarioQuests')" :duties="db.mainStoryQuest" type="msq" />
      <DutyProgress :title="$t('questlines.sideStoryQuests')" :duties="db.sideStoryQuests" type="sidequest" />
      <DutyProgress :title="$t('questlines.encounterQuests')" :duties="db.dutiesQuests" type="encounterquest" />
      <DutyProgress :title="$t('questlines.tribalQuests')" :duties="db.beastTribeQuests" type="tribal" />
      <DutyProgress :title="$t('questlines.alliedTribalQuests')" :duties="db.alliedBeastTribeQuests" type="tribal" />
      <DutyProgress :title="$t('questlines.customDeliveryQuests')" :duties="db.customDeliveryQuests" type="deliveries" />
    </div>
    <div class="col-lg mb-4">
      <h2>{{ $t("questlines.jobsRolesRelics") }}</h2>
      <DutyProgress :title="$t('questlines.combatJobQuests')" :duties="db.combatJobQuests" type="combat" />
      <DutyProgress :title="$t('questlines.gathererJobQuests')" :duties="db.gathererJobQuests" type="gathering" />
      <DutyProgress :title="$t('questlines.crafterJobQuests')" :duties="db.crafterJobQuests" type="crafting" />
      <DutyProgress :title="$t('questlines.roleQuestsShB')" :duties="db.shbRoleQuests" type="combat" />
      <DutyProgress :title="$t('questlines.roleQuestsEW')" :duties="db.ewRoleQuests" type="combat" />
      <DutyProgress :title="$t('questlines.roleQuestsDT')" :duties="db.dtRoleQuests" type="combat" />
      <DutyProgress :title="$t('questlines.relicQuests')" :duties="db.relicQuests" type="relic" />
      <DutyProgress :title="$t('questlines.crystallineMeanQuests')" :duties="db.crystallineMeanQuests" type="leveling" />
      <DutyProgress :title="$t('questlines.studiumDeliveryQuests')" :duties="db.studiumDeliveryQuests" type="leveling" />
      <DutyProgress :title="$t('questlines.wachumeqimeqiDeliveryQuests')" :duties="db.wachumeqimeqiDeliveryQuests" type="leveling" />
    </div>
  </div>

  <!-- Challenges summaries -->
  <h1>
    {{ $t("page.challenges")}}
    <span class="fs-3 fw-lighter d-block d-lg-inline">
      {{ $t("pageHeader.clearedByCharacter", { characterName: this.$store.getters.character.Name }) }}
    </span>
  </h1>
  <hr />

  <div class="row">
    <div class="col-lg mb-4">
      <h2>{{ $t("challenges.exploration") }}</h2>
      <DutyProgress :title="$t('shared.expansion.arr')" :duties="db.arrExploration" type="magnifier" />
      <DutyProgress :title="$t('shared.expansion.hw')" :duties="db.hwExploration" type="magnifier" />
      <DutyProgress :title="$t('shared.expansion.sb')" :duties="db.sbExploration" type="magnifier" />
      <DutyProgress :title="$t('shared.expansion.shb')" :duties="db.shbExploration" type="magnifier" />
      <DutyProgress :title="$t('shared.expansion.ew')" :duties="db.ewExploration" type="magnifier" />
      <DutyProgress :title="$t('shared.expansion.dt')" :duties="db.dtExploration" type="magnifier" />
      <DutyProgress :title="$t('challenges.exploratoryMissions')" :duties="db.explExploration" type="magnifier" />
      <DutyProgress :title="$t('challenges.sightseeingLogs')" :duties="db.sightseeingLogs" type="magnifier" />
      <br />
      <h2>{{ $t("challenges.achievementFates") }}</h2>
      <DutyProgress :title="$t('shared.expansion.arr')" :duties="db.arrFATEs" type="fate" />
      <DutyProgress :title="$t('shared.expansion.hw')" :duties="db.hwFATEs" type="fate" />
      <DutyProgress :title="$t('shared.expansion.sb')" :duties="db.sbFATEs" type="fate" />
      <DutyProgress :title="$t('shared.expansion.shb')" :duties="db.shbFATEs" type="fate" />
      <DutyProgress :title="$t('shared.expansion.ew')" :duties="db.ewFATEs" type="fate" />
      <DutyProgress :title="$t('shared.expansion.dt')" :duties="db.dtFATEs" type="fate" />
    </div>
    <div class="col-lg mb-4">
      <h2>{{ $t("challenges.deepDungeons") }}</h2>
      <DutyProgress :title="$t('challenges.potd')" :duties="db.potd" type="deepdungeons" />
      <DutyProgress :title="$t('challenges.potdSolo')" :duties="db.potdSolo" type="deepdungeons" />
      <DutyProgress :title="$t('challenges.hoh')" :duties="db.hoh" type="deepdungeons" />
      <DutyProgress :title="$t('challenges.hohSolo')" :duties="db.hohSolo" type="deepdungeons" />
      <DutyProgress :title="$t('challenges.eo')" :duties="db.eo" type="deepdungeons" />
      <DutyProgress :title="$t('challenges.eoSolo')" :duties="db.eoSolo" type="deepdungeons" />
      <DutyProgress :title="$t('challenges.pt')" :duties="db.eo" type="deepdungeons" />
      <DutyProgress :title="$t('challenges.ptSolo')" :duties="db.eoSolo" type="deepdungeons" />
      <br />
      <h2>{{ $t("challenges.blueMage") }}</h2>
      <DutyProgress :title="$t('challenges.blueARR')" :duties="db.arrBlu" type="blu" />
      <DutyProgress :title="$t('challenges.blueHW')" :duties="db.hwBlu" type="blu" />
      <DutyProgress :title="$t('challenges.blueSB')" :duties="db.sbBlu" type="blu" />
      <DutyProgress :title="$t('challenges.blueShB')" :duties="db.shbBlu" type="blu" />
      <DutyProgress :title="$t('challenges.maskedCarnivale')" :duties="db.maskedCarnivale" type="blu" />
    </div>
    <div class="col-lg mb-4">
      <h2>{{ $t("challenges.miscellaneous") }}</h2>
      <DutyProgress :title="$t('challenges.sharedFateShB')" :duties="db.shbSharedFATE" type="fate" />
      <DutyProgress :title="$t('challenges.sharedFateEW')" :duties="db.ewSharedFATE" type="fate" />
      <DutyProgress :title="$t('challenges.sharedFateDT')" :duties="db.dtSharedFATE" type="fate" />
      <DutyProgress :title="$t('challenges.bozjanSouthernFront')" :duties="db.bozjanChallenges" type="field" />
      <DutyProgress :title="$t('challenges.zadnor')" :duties="db.zadnorChallenges" type="field" />
      <DutyProgress :title="$t('challenges.variantDungeons')" :duties="db.variantChallenges" type="vc" />
      <DutyProgress :title="$t('challenges.treasureHuntClears')" :duties="db.treasureHunts" type="treasure" />
      <DutyProgress :title="$t('challenges.trusts')" :duties="db.trusts" type="leveling" />
      <DutyProgress :title="$t('challenges.islandSanctuary')" :duties="db.islandSanctuary" type="island" />
      <br />
      <h2>{{ $t("challenges.mentorStatus") }}</h2>
      <DutyProgress :title="$t('challenges.mentorBattle')" :duties="db.mentorBattle" type="mentor" />
      <DutyProgress :title="$t('challenges.mentorTrade')" :duties="db.mentorTrade" type="mentor" />
    </div>
  </div>
</template>

<style lang="scss">
.info-block {
  line-height: 1.3;
  margin-bottom: 15px;

  b {
    color: var(--bs-primary-color);
  }

  img {
    height: 1.3em;
    margin-left: 5px;
  }
}

#character-portrait {
  width: 100%;
  border-radius: 20px;
  padding: 0.1rem;
  background-color: var(--bs-body-color);
}
</style>

<script>
import PageHeader from "@/components/PageHeader.vue";
import JobLevel from "@/components/JobLevel.vue";
import DutyProgress from "@/components/DutyProgress.vue";

import dbJson from "@/assets/db.json";

export default {
  name: "ProfileView",
  data() {
    return {
      maxLevelEureka: 60,
      maxLevelBozja: 25,
      db: dbJson,
    };
  },
  components: {
    PageHeader,
    JobLevel,
    DutyProgress,
  },
  computed: {
    hasBio() {
      return this.$store.getters.character.Bio != "-";
    },
    isMasculine() {
      return this.$store.getters.character.Gender == 1;
    },
    gender() {
      return this.isMasculine ? "Masculine" : "Feminine";
    },
    race() {
      return this.genderedLocaleName(this.$store.getters.character.Race, this.gender);
    },
    clan() {
      return this.genderedLocaleName(this.$store.getters.character.Tribe, this.gender);
    },
    guardianImageURI() {
      return "/icons/guardian-" + this.$store.getters.character.GuardianDeity.ID + ".png";
    },
    cityStateImageURI() {
      return "/icons/town-" + this.$store.getters.character.Town.ID + ".png";
    },
    grandCompanyImageURI() {
      const grandCompanyID = this.$store.getters.character.GrandCompanyInfo.GrandCompany.ID;
      const grandCompanyRankID = this.$store.getters.character.GrandCompanyInfo.RankID;
      return "/icons/gc-" + grandCompanyID + "-" + grandCompanyRankID + ".png";
    },
    playingSince() {
      const date = new Date(this.$store.getters.lodestoneData.PlayingSince);
      const locale = this.$i18n.locale;
      return date.toLocaleDateString(locale, { year: "numeric", month: "long", day: "numeric" });
    },
    levelEureka() {
      return this.$store.getters.character.ClassJobElemental?.Level || 0;
    },
    levelBozja() {
      return this.$store.getters.character.ClassJobBozjan?.Level || 0;
    },
    isMaxLevelEureka() {
      return this.levelEureka == this.maxLevelEureka;
    },
    isMaxLevelBozja() {
      return this.levelBozja == this.maxLevelBozja;
    },
  },
  methods: {
    localeName(payload) {
      const locale = this.$i18n.locale;
      return payload["Name" + locale.toUpperCase()];
    },
    genderedLocaleName(payload, gender) {
      const locale = this.$i18n.locale;
      return payload["Name" + gender + locale.toUpperCase()];
    },
  },
};
</script>
