import db from "@/assets/db.json";

const jobInitials = {
  1: "gla",
  2: "pgl",
  3: "mrd",
  4: "lnc",
  5: "arc",
  6: "cnj",
  7: "thm",
  8: "crp",
  9: "bsm",
  10: "arm",
  11: "gsm",
  12: "ltw",
  13: "wvr",
  14: "alc",
  15: "cul",
  16: "min",
  17: "btn",
  18: "fsh",
  19: "pld",
  20: "mnk",
  21: "war",
  22: "drg",
  23: "brd",
  24: "whm",
  25: "blm",
  26: "acn",
  27: "smn",
  28: "sch",
  29: "rog",
  30: "nin",
  31: "mch",
  32: "drk",
  33: "ast",
  34: "sam",
  35: "rdm",
  36: "blu",
  37: "gnb",
  38: "dnc",
  39: "rpr",
  40: "sge",
  41: "vpr",
  42: "pct",
  43: "bst", // Verify this
};

function hasMaxIDOneOf(achievements, encounter) {
  for (let id of encounter.MaxIDOneOf) {
    if (achievements.has(id)) return true;
  }
  return false;
}

function hasMaxIDAllOf(achievements, encounter) {
  for (let id of encounter.MaxIDAllOf) {
    if (!achievements.has(id)) return false;
  }
  return true;
}

// Checks if the character has cleared the specified encounter.
// Returns 1 if cleared, 0 if not, and -1 if unsure.
// The ordering of these conditionals is important.
function isEncounterCleared(achievements, characterData, encounter) {
  if (!characterData.lodestoneData.Achievements?.length > 0) return -1;
  if (encounter.ID && achievements.has(encounter.ID)) return 1;
  if (encounter.MaxIDOneOf && hasMaxIDOneOf(achievements, encounter)) return 1;
  if (encounter.MaxIDAllOf && hasMaxIDAllOf(achievements, encounter)) return 1;
  if (encounter.MaxIDAllOf && !hasMaxIDAllOf(achievements, encounter)) return 0;
  if (encounter.MinID && achievements.has(encounter.MinID)) return -1;
  if (encounter.MinID && !achievements.has(encounter.MinID)) return 0;
  if (!encounter.MinID && encounter.MaxIDOneOf) return 0;
  if (!encounter.ID) return -1;
  return 0;
}

function injectCharacterData(characterData) {
  let achievements = new Map();
  for (let achievement of characterData.lodestoneData.Achievements || []) {
    achievements.set(achievement.ID, achievement.Date);
  }

  // Set the "Playing Since" date based on achievement #789 ("Complete 10 unique quests.")
  // It's not a perfect solution, but that achievement is the first one that most player should get.
  if (achievements.has(789)) {
    characterData.lodestoneData.PlayingSince = achievements.get(789);
  }

  // Create a "Jobs" object with the job initials as keys.
  // This is a more convenient way to access job data ("ClassJobs" is an unsorted array).
  characterData.lodestoneData.Jobs = {};
  for (let classJob of characterData.lodestoneData.Character.ClassJobs) {
    characterData.lodestoneData.Jobs[jobInitials[classJob.UnlockedState.ID]] = classJob;
  }

  // Create an "EncountersMap" object with the encounter UUIDs as keys.
  characterData.encountersVerified = {};
  for (const category in db) {
    for (const encounter of db[category]) {
      if (!encounter.UUID) continue;
      characterData.encountersVerified[encounter.UUID] = isEncounterCleared(achievements, characterData, encounter);
    }
  }
}

export { injectCharacterData };
