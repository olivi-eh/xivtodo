<template>
  <PageHeader title="page.achievements" subtitle="pageHeader.obtainedByCharacter" />

  <div class="row my-4">
    <table class="table align-middle">
      <thead>
        <tr>
          <th scope="col"></th>
          <th scope="col">Name</th>
          <th scope="col" class="d-none d-md-block">Description</th>
          <th scope="col">Patch</th>
        </tr>
      </thead>
      <tbody>
        <tr
          v-for="item of db"
          :key="item.id"
          :class="{
            'row-obtained': achievements.has(item.id),
          }"
        >
          <th scope="row"><img :src="item.icon" /></th>
          <td>{{ item.name }}</td>
          <td class="d-none d-md-table-cell">
            {{ item.description }}
          </td>
          <td>{{ item.patch }}</td>
        </tr>
      </tbody>
    </table>
  </div>
</template>

<style lang="scss">
.row-obtained {
  color: var(--bs-secondary-color);
  img {
    filter: brightness(50%);
  }
}
</style>

<script>
import PageHeader from "@/components/PageHeader.vue";
import achievementsJson from "@/assets/dbs/en/achievements.json";
import { useAppStore } from "@/store/app";

export default {
  name: "AchievementsView",
  setup() {
    return { appStore: useAppStore() };
  },
  data() {
    return {
      db: achievementsJson.results.filter((e) => e.type.id == 1).sort((a, b) => a.id - b.id),
      achievements: this.appStore.achievements,
    };
  },
  components: {
    PageHeader,
  },
};
</script>
