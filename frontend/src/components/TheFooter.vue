<template>
  <footer
    class="footer d-flex justify-content-between mt-auto py-2 px-3 bg-light user-select-none"
    data-bs-theme="night"
  >
    <!-- Left-side info -->
    <div class="footer-copyrights text-muted my-auto d-none d-xl-block col-4">
      {{ $t("footer.trademark") }}<br />
      {{ $t("footer.copyright") }}<br />
      All artworks were drawn by
      <a href="https://bsky.app/profile/churubirb.bsky.social" target="_blank" rel="noopener noreferrer" style="text-decoration: underline;">Lumaere</a>
      for XIV ToDo. All rights reserved.
    </div>

    <!-- Center branding and links -->
    <div class="flex-fill text-center text-muted">
      <img src="@/assets/img/brand.png" alt="XIV ToDo branding" />
      <br />
      <router-link to="/help/">
        <small>{{ $t("page.help") }}</small>
      </router-link>
      ·
      <router-link to="/privacy/">
        <small>{{ $t("page.privacy") }}</small>
      </router-link>
      ·
      <a href="https://discord.gg/zfzhKhG3zj" target="_blank" rel="noopener noreferrer">
        <small>{{ $t("page.discord") }}</small>
      </a>
      ·
      <a href="https://github.com/olivi-eh/xivtodo" target="_blank" rel="noopener noreferrer">
        <small>{{ $t("page.github") }}</small>
      </a>
      ·
      <a href="https://bsky.app/profile/xivtodo.com" target="_blank" rel="noopener noreferrer">
        <small>{{ $t("page.bluesky") }}</small>
      </a>
      <br />
      <a href="https://ko-fi.com/olivi_eh" target="_blank" rel="noopener noreferrer" class="fw-bolder">
        <small>{{ $t("page.donate") }} <i class="fa-fw fas fa-heart"></i></small>
      </a>
    </div>

    <!-- Right-side info -->
    <div class="footer-copyrights text-muted text-end my-auto d-none d-xl-block col-4">
      {{ $t("footer.buildVersion", { buildNum: this.$store.state.env.VUE_APP_BUILD_NUM }) }}<br />
      {{ $t("footer.buildDate", { buildDate: this.$store.state.env.VUE_APP_BUILD_DATE }) }}
    </div>
  </footer>
</template>

<style lang="scss">
.footer {
  background-color: var(--bs-dark) !important;
  border-top: 1px solid var(--bs-gray-800);

  a,
  a:focus {
    color: var(--bs-secondary-color);
  }

  a:hover {
    color: var(--bs-body-color);
  }
}

.footer-copyrights {
  font-size: 0.6rem;
}

footer a {
  text-decoration: none;
}

footer img {
  height: 22px;
  margin-bottom: 5px;
  margin-top: 5px;
  filter: brightness(70%) grayscale(100%);
}

footer img:hover {
  filter: brightness(100%) grayscale(0%);
}
</style>

<script>
export default {
  Name: "TheFooter",
};
</script>
