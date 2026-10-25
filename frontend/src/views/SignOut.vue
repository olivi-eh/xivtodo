<template>
  <div>Signing out...</div>
</template>

<script>
import { useAppStore } from "@/store/app";

export default {
  name: "SignOut",
  setup() {
    return { appStore: useAppStore() };
  },
  mounted() {
    this.appStore.deleteUserData();
    this.$router.push("/");
  },
};
</script>
