# XIV ToDo

[![Website](https://img.shields.io/website?url=https%3A%2F%2Fxivtodo.com)](https://xivtodo.com/)
[![Discord](https://img.shields.io/discord/946996969354690600?label=discord)](https://discord.gg/zfzhKhG3zj)
[![ko-fi](https://img.shields.io/badge/buy%20me%20a%20coffee!-donate-success)](https://ko-fi.com/V7V569BFY)
[![fr translation](https://img.shields.io/badge/dynamic/json?color=blue&label=fr&style=flat&logo=crowdin&query=%24.progress.1.data.translationProgress&url=https%3A%2F%2Fbadges.awesome-crowdin.com%2Fstats-15746609-575411.json)](https://crowdin.com/project/xiv-todo)
[![de translation](https://img.shields.io/badge/dynamic/json?color=blue&label=de&style=flat&logo=crowdin&query=%24.progress.0.data.translationProgress&url=https%3A%2F%2Fbadges.awesome-crowdin.com%2Fstats-15746609-575411.json)](https://crowdin.com/project/xiv-todo)
[![ja translation](https://img.shields.io/badge/dynamic/json?color=blue&label=ja&style=flat&logo=crowdin&query=%24.progress.2.data.translationProgress&url=https%3A%2F%2Fbadges.awesome-crowdin.com%2Fstats-15746609-575411.json)](https://crowdin.com/project/xiv-todo)

[XIV ToDo](https://xivtodo.com) provides useful dashboards, tailored checklists and tools for Final Fantasy XIV. Whether you are a new player to the game or more experienced, you will find it useful! In its current form, XIV ToDo offers customizable daily and weekly checklists as well as automated completion trackers of emcounters and questlines, with much more to come!

## Features

- Support for adding multiple characters.
- Profile containing publicly accessible Lodestone information.
- List of all encounters up-to-date with automatic character completion status.
- List of all important questlines with automatic character completion status.
- Comprehensive list of dailies and weeklies with per-character completion and customization.

## Screenshots

<img src="https://user-images.githubusercontent.com/3271352/231013111-2b3e6dcb-7efa-4628-9998-39d869658d5b.png" width=400 /> <img src="https://user-images.githubusercontent.com/3271352/231013112-eaf34cc8-3b6a-4180-b90e-c6825eec787a.png" width=400 /> <img src="https://user-images.githubusercontent.com/3271352/231013106-2c46d733-ffeb-49b1-b574-e322083ba506.png" width=400 /> <img src="https://user-images.githubusercontent.com/3271352/231013109-e77d2faf-6bdd-4f0b-95d5-0f4c6aab4c16.png" width=400 />

## Development

### Frontend

1. Make sure you have the following requirements:
   - Node.js 20+ installed locally
   - npm installed locally
   - Vue's CLI installed locally (`npm install -g @vue/cli`)

2. Clone the repository:
   ```sh
   git clone https://github.com/olivi-eh/xivtodo
   cd xivtodo/frontend/
   ```

3. Fetch dependencies:
   ```sh
   npm install
   ```

4. Run one of the following commands:
   ```sh
   npm run serve # serve the frontend locally
   npm run build # build the dist directory
   npm run lint # lint the frontend
   vue ui # open vue's interactive ui
   ```

By default, the frontend will attempt to use a local backend (`localhost:8181`). To develop against the production backend, modify the following environment variable in the `.env` file before serving:
```sh
VUE_APP_BACKEND_API_URI=https://api.xivtodo.com
```

**Note**: With that setup, you will not be able to sign in with Discord due to security policies. It is recommended to run both the frontend and backend locally.

### Backend

1. Make sure you have the following requirements:
   - Google Cloud project with the Firestore API enabled
   - Service Account with RW permissions to Firestore
   - Discord application with an OAuth2 redirect to `http://localhost:8080/auth`
   - Go 1.22+ installed locally

2. Clone the repository:
   ```sh
   git clone https://github.com/olivi-eh/xivtodo
   cd xivtodo/backend/
   ```

3. Fetch dependencies:
   ```sh
   go mod
   ```

4. Set the following environment variables:
   ```
   GOOGLE_APPLICATION_CREDENTIALS=backend-sa.json # Google Cloud service account key
   DISCORD_CLIENT_ID=<discord_client_id>
   DISCORD_CLIENT_SECRET=<discord_client_secret>
   DISCORD_REDIRECT_URI=http://localhost:8080/auth
   ```

5. Run the backend server:
   ```sh
   go run main
   ```

## Contributing

Contributions are welcome and appreciated. Simply open up a new issue, or fork the repository to open up a new pull request against the upstream main branch.
