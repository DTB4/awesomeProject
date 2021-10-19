<template>
  <div class="logout" @click="logout">Logout</div>
</template>

<script>
import {mapActions, mapMutations} from "vuex";
import refresh_tokens from "../mixins/refresh_tokens";

export default {
  mixins: [refresh_tokens],
  name: "Logout",
  methods: {
    ...mapActions("tokens", ["removeTokens"]),
    ...mapMutations("tokens", ["setLogoutState"]),
    async logout() {
      const response = await fetch("http://localhost:8081/logout", {
        method: "GET",
        mode: "cors",
        headers: {
          Accept: "*/*",
          Authorization: "Bearer " + localStorage.getItem("access_token"),
        },
      });
      if (response.status === 200) {
        this.setLogoutState();
        this.removeTokens();
      } else if (response.status === 401) {
        //TODO: try to catch 401 error without "error" in console
        let ok = await this.refreshTokens();
        if (ok) {
          await this.logout();
        }
      } else {
        this.setLogoutState();
        this.removeTokens();
        console.log("not ok response", response);
      }
    },
  },
};
</script>

<style scoped></style>
