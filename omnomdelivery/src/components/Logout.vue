<template>
  <div v-if="show" id="logout_id" class="logout" @click="logout">Logout</div>
</template>

<script>
import {mapActions} from "vuex";
import refresh_tokens from "../mixins/refresh_tokens";

export default {
  mixins: [refresh_tokens],
  name: "Logout",
  data() {
    return {
      isLogout: false,
      show: {
        type: Boolean,
        default: false,
      },
    };
  },
  methods: {
    ...mapActions("tokens", ["removeTokens"]),
    async logout() {
      const response = await fetch("http://localhost:8081/logout", {
        method: "POST",
        mode: "cors",
        headers: {
          Accept: "*/*",
          Authorization: "Bearer " + localStorage.getItem("access_token"),
        },
      });
      if (response.ok) {
        this.removeTokens();
        this.$emit("userLogout");

      } else if (response.status === 401) {
        //TODO: try to catch 401 error without "error" in console
        let ok = await this.refreshTokens();
        if (ok) {
          await this.logout();
        }
      } else {
        this.removeTokens();
        this.$emit("userLogout");
        console.log("not ok response", response);
      }
    },
  },
};
</script>

<style scoped></style>
