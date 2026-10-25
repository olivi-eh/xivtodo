// import enAchievementsJson from "@/assets/dbs/en/achievements.json";
import enMountsJson from "@/assets/dbs/en/mounts.json";
import enMinionsJson from "@/assets/dbs/en/minions.json";
// import enTitlesJson from "@/assets/dbs/en/titles.json";

// import frAchievementsJson from "@/assets/dbs/fr/achievements.json";
import frMountsJson from "@/assets/dbs/fr/mounts.json";
import frMinionsJson from "@/assets/dbs/fr/minions.json";
// import frTitlesJson from "@/assets/dbs/fr/titles.json";

// import deAchievementsJson from "@/assets/dbs/de/achievements.json";
import deMountsJson from "@/assets/dbs/de/mounts.json";
import deMinionsJson from "@/assets/dbs/de/minions.json";
// import deTitlesJson from "@/assets/dbs/de/titles.json";

// import jaAchievementsJson from "@/assets/dbs/ja/achievements.json";
import jaMountsJson from "@/assets/dbs/ja/mounts.json";
import jaMinionsJson from "@/assets/dbs/ja/minions.json";
// import jaTitlesJson from "@/assets/dbs/ja/titles.json";

function reduceJson(json) {
  return json.results.reduce(function (map, obj) {
    map[obj.id] = obj;
    return map;
  }, {});
}

export default {
  en: {
    // achievements: reduceJson(enAchievementsJson),
    mounts: reduceJson(enMountsJson),
    minions: reduceJson(enMinionsJson),
    // titles: reduceJson(enTitlesJson),
  },
  fr: {
    // achievements: reduceJson(frAchievementsJson),
    mounts: reduceJson(frMountsJson),
    minions: reduceJson(frMinionsJson),
    // titles: reduceJson(frTitlesJson),
  },
  de: {
    // achievements: reduceJson(deAchievementsJson),
    mounts: reduceJson(deMountsJson),
    minions: reduceJson(deMinionsJson),
    // titles: reduceJson(deTitlesJson),
  },
  ja: {
    // achievements: reduceJson(jaAchievementsJson),
    mounts: reduceJson(jaMountsJson),
    minions: reduceJson(jaMinionsJson),
    // titles: reduceJson(jaTitlesJson),
  },
};
