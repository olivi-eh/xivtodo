# XIV ToDo

[![Website](https://img.shields.io/website?url=https%3A%2F%2Fxivtodo.com)](https://xivtodo.com/)
[![Discord](https://img.shields.io/discord/946996969354690600?label=discord)](https://discord.gg/zfzhKhG3zj)
[![ko-fi](https://img.shields.io/badge/buy%20me%20a%20coffee!-donate-success)](https://ko-fi.com/V7V569BFY)
[![fr translation](https://img.shields.io/badge/dynamic/json?color=blue&label=fr&style=flat&logo=crowdin&query=%24.progress.1.data.translationProgress&url=https%3A%2F%2Fbadges.awesome-crowdin.com%2Fstats-15746609-575411.json)](https://crowdin.com/project/xiv-todo)
[![de translation](https://img.shields.io/badge/dynamic/json?color=blue&label=de&style=flat&logo=crowdin&query=%24.progress.0.data.translationProgress&url=https%3A%2F%2Fbadges.awesome-crowdin.com%2Fstats-15746609-575411.json)](https://crowdin.com/project/xiv-todo)
[![ja translation](https://img.shields.io/badge/dynamic/json?color=blue&label=ja&style=flat&logo=crowdin&query=%24.progress.2.data.translationProgress&url=https%3A%2F%2Fbadges.awesome-crowdin.com%2Fstats-15746609-575411.json)](https://crowdin.com/project/xiv-todo)

[XIV ToDo](https://xivtodo.com) provides useful dashboards, tailored checklists and tools for Final Fantasy XIV. Whether you are a new player to the game or more experienced, you will find it useful! In its current form, XIV ToDo offers customizable daily and weekly checklists as well as automated completion trackers of encounters and questlines, with much more to come!

## Features

- Support for adding multiple characters.
- Profile containing publicly accessible Lodestone information, shareable through a public character page.
- List of all encounters up-to-date with automatic character completion status.
- List of all important questlines with automatic character completion status.
- Challenges tracker covering deep dungeons, exploration, field operations, Blue Mage, and more.
- Comprehensive list of dailies and weeklies with per-character completion and customization.
- Countdowns to upcoming maintenance, patches, and events, plus the latest news and updates.
- Available in English, French, German, and Japanese.

## Screenshots

<img src="https://user-images.githubusercontent.com/3271352/231013111-2b3e6dcb-7efa-4628-9998-39d869658d5b.png" width=400 /> <img src="https://user-images.githubusercontent.com/3271352/231013112-eaf34cc8-3b6a-4180-b90e-c6825eec787a.png" width=400 /> <img src="https://user-images.githubusercontent.com/3271352/231013106-2c46d733-ffeb-49b1-b574-e322083ba506.png" width=400 /> <img src="https://user-images.githubusercontent.com/3271352/231013109-e77d2faf-6bdd-4f0b-95d5-0f4c6aab4c16.png" width=400 />

## Development

### Frontend

1. Make sure you have the following requirements:
   - Node.js 24+ installed locally (the exact major version is pinned in `frontend/.nvmrc`)
   - npm installed locally

2. Clone the repository:
   ```sh
   git clone https://github.com/olivi-eh/xivtodo
   cd xivtodo/frontend/
   ```

3. If you use a Node version manager such as [nvm](https://github.com/nvm-sh/nvm) or [nvm.fish](https://github.com/jorgebucaran/nvm.fish), install and activate the pinned version:
   ```sh
   nvm install
   ```

4. Fetch dependencies:
   ```sh
   npm install
   ```

5. Run one of the following commands:
   ```sh
   npm run dev # serve the frontend locally on http://localhost:8080
   npm run build # build the dist directory
   npm run preview # serve the built dist directory on http://localhost:8080
   npm run lint # lint the frontend
   ```

The frontend uses a local backend (`localhost:8181`), which you need to run as described below. The production backend only accepts requests from `https://xivtodo.com`, so a local frontend cannot use it. To point the frontend at a backend elsewhere, create a `frontend/.env.local` file (ignored by git) that overrides `VITE_BACKEND_API_URI`.

### Backend

> [!WARNING]
> The backend always connects to the production Firestore project (`xiv-todo`, set in `backend/store/firestore.go`). Unless you intend to work with production data, run it against the local Firestore emulator as described below.

1. Make sure you have the following requirements:
   - Go 1.27+ installed locally (see `backend/go.mod`)
   - [Google Cloud CLI](https://cloud.google.com/sdk/docs/install) installed locally
   - A Java runtime installed locally, required by the Firestore emulator (e.g. `brew install openjdk`)
   - Discord application with an OAuth2 redirect to `http://localhost:8080/auth`

2. Clone the repository:
   ```sh
   git clone https://github.com/olivi-eh/xivtodo
   cd xivtodo/backend/
   ```

3. Fetch dependencies:
   ```sh
   go mod download
   ```

4. Install and start the Firestore emulator in a separate terminal. Its data is kept in memory and is discarded when it stops:
   ```sh
   gcloud components install cloud-firestore-emulator
   gcloud emulators firestore start --host-port=localhost:8282
   ```

5. Set the following environment variables. All three `DISCORD_*` variables are required, and the server exits on startup if any is missing. The backend also only accepts CORS requests from the origin of `DISCORD_REDIRECT_URI`, so it must match where the frontend is served:
   ```sh
   export FIRESTORE_EMULATOR_HOST=localhost:8282
   export DISCORD_CLIENT_ID=<discord_client_id>
   export DISCORD_CLIENT_SECRET=<discord_client_secret>
   export DISCORD_REDIRECT_URI=http://localhost:8080/auth
   ```

6. Run the backend server, which listens on `localhost:8181` by default (set `PORT` to change it):
   ```sh
   go run .
   ```

To work against the production Firestore instead, leave `FIRESTORE_EMULATOR_HOST` unset and authenticate with your own Google account, which needs read and write access to Firestore in the `xiv-todo` project:
```sh
gcloud auth application-default login
```

Prefer this over a downloaded service account key, as it avoids keeping a long-lived credential on disk. Never commit the Discord client secret or any Google credentials.

### Lodestone function

The `godestone/` directory holds the `lodestone-character` Cloud Run function (entry point `LodestoneCharacter`, region `asia-northeast1`), which scrapes character data from the Lodestone using [godestone](https://github.com/xivapi/godestone). It is deployed from source, so it has no build script or `go.sum`. The backend always calls the deployed function (see `backend/utils/gcf.go`), including when it runs locally.

### Data scripts

The `scripts/` directory holds one-off Python scripts that regenerate the static data in `frontend/src/assets/`. Run them only when that data needs updating. They require [uv](https://docs.astral.sh/uv/), which installs the pinned Python version and dependencies on first run:
```sh
cd scripts/
uv run fetch_collect_json.py # refresh the FFXIV Collect databases in frontend/src/assets/dbs/
./gendb.sh # rebuild frontend/src/assets/db.json from the *.csv exports placed in scripts/
uv run duty_names.py # write localized duty names from duties.csv to duty_names_<locale>.txt
```

## Contributing

Contributions are welcome and appreciated. Simply open up a new issue, or fork the repository to open up a new pull request against the upstream main branch.
