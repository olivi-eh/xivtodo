import { useAppStore } from "@/store/app";

function getLodestoneURL() {
  let url = "https://";
  switch (useAppStore().language) {
    case "ja":
      url += "jp";
      break;
    case "de":
      url += "de";
      break;
    case "fr":
      url += "fr";
      break;
    default:
      url += "na";
      break;
  }
  return url + ".finalfantasyxiv.com/lodestone/";
}

export { getLodestoneURL };
