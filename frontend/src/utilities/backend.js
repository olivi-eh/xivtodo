import { useAppStore } from "../store/app";
import router from "../router";
import { injectCharacterData } from "./character";

const apiEndpoint = import.meta.env.VITE_BACKEND_API_URI;

// The authorization code was rotated or revoked (e.g. by signing in elsewhere), so the local session is stale
const signOutIfUnauthorized = (response) => {
  const store = useAppStore();
  if (response.status == 401 && store.isSignedIn) {
    // Not routed through /signout, since navigating to a non-Home page clears sign-in messages
    store.deleteUserData();
    store.setSigningIn({ sessionExpired: true });
    router.push("/");
  }
};

const authorization = () => {
  const store = useAppStore();
  return store.discordUser.id + ":" + store.settings.authorizationCode;
};

const getVersion = () =>
  new Promise((resolve, reject) => {
    fetch(apiEndpoint + "/version", {
      method: "GET",
      cache: "no-cache",
    })
      .then((response) => {
        if (response.ok) {
          return response.json();
        } else {
          reject();
        }
      })
      .then((version) => {
        resolve(version.version);
      })
      .catch((err) => {
        console.log(err);
        reject();
      });
  });

const authenticate = (code) =>
  new Promise((resolve, reject) => {
    fetch(apiEndpoint + "/auth?code=" + code, {
      method: "GET",
      cache: "no-cache",
    })
      .then((response) => {
        if (response.ok) {
          return response.json();
        } else {
          throw response;
        }
      })
      .then((userData) => {
        let characters = [];
        if (userData.characters != null) {
          for (let characterID in userData.characters) {
            injectCharacterData(userData.characters[characterID]);
            characters.push(userData.characters[characterID]);
          }
        }
        userData.characters = characters;
        resolve(userData);
      })
      .catch((err) => {
        reject(err);
      });
  });

const getUserData = () =>
  new Promise((resolve, reject) => {
    const timestamp = Date.now();
    fetch(`${apiEndpoint}/users?t=${timestamp}`, {
      headers: {
        Authorization: authorization(),
        "Cache-Control": "no-cache, no-store, must-revalidate",
      },
      method: "GET",
      cache: "no-store",
    })
      .then((response) => {
        if (response.ok) {
          return response.json();
        } else {
          signOutIfUnauthorized(response);
          throw response;
        }
      })
      .then((userData) => {
        let characters = [];
        if (userData.characters != null) {
          for (let characterID in userData.characters) {
            injectCharacterData(userData.characters[characterID]);
            characters.push(userData.characters[characterID]);
          }
        }
        userData.characters = characters;
        resolve(userData);
      })
      .catch((err) => {
        reject(err);
      });
  });

const updateSettings = (payload) =>
  new Promise((resolve, reject) => {
    fetch(apiEndpoint + "/settings", {
      headers: {
        Authorization: authorization(),
      },
      body: JSON.stringify(payload),
      method: "PUT",
      cache: "no-cache",
    })
      .then((response) => {
        if (response.ok) {
          return response.json();
        } else {
          signOutIfUnauthorized(response);
          throw response;
        }
      })
      .then((settings) => {
        resolve(settings);
      })
      .catch((err) => {
        reject(err);
      });
  });

const addCharacter = (id) =>
  new Promise((resolve, reject) => {
    fetch(apiEndpoint + "/characters?id=" + id, {
      headers: {
        Authorization: authorization(),
      },
      method: "GET",
      cache: "no-cache",
    })
      .then((response) => {
        if (response.ok) {
          return response.json();
        } else {
          signOutIfUnauthorized(response);
          throw response;
        }
      })
      .then((characterData) => {
        injectCharacterData(characterData);
        resolve(characterData);
      })
      .catch((err) => {
        reject(err);
      });
  });

const removeCharacter = (id) =>
  new Promise((resolve, reject) => {
    fetch(apiEndpoint + "/characters?id=" + id, {
      headers: {
        Authorization: authorization(),
      },
      method: "DELETE",
      cache: "no-cache",
    })
      .then((response) => {
        if (response.ok) {
          resolve(response);
        } else {
          signOutIfUnauthorized(response);
          throw response;
        }
      })
      .catch((err) => {
        reject(err);
      });
  });

const updateChecklist = (id, payload) =>
  new Promise((resolve, reject) => {
    fetch(apiEndpoint + "/checklist?id=" + id, {
      headers: {
        Authorization: authorization(),
      },
      body: JSON.stringify(payload),
      method: "PUT",
      cache: "no-cache",
    })
      .then((response) => {
        if (response.ok) {
          resolve(response);
        } else {
          signOutIfUnauthorized(response);
          throw response;
        }
      })
      .catch((err) => {
        reject(err);
      });
  });

const updateEncounterIDs = (id, payload) =>
  new Promise((resolve, reject) => {
    fetch(apiEndpoint + "/encounters?id=" + id, {
      headers: {
        Authorization: authorization(),
      },
      body: JSON.stringify(payload),
      method: "PUT",
      cache: "no-cache",
    })
      .then((response) => {
        if (response.ok) {
          resolve(response);
        } else {
          signOutIfUnauthorized(response);
          throw response;
        }
      })
      .catch((err) => {
        reject(err);
      });
  });

export {
  getVersion,
  authenticate,
  getUserData,
  updateSettings,
  addCharacter,
  removeCharacter,
  updateChecklist,
  updateEncounterIDs,
};
