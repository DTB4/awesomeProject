<template>
  <div v-if="show" class="logout" id="logout_id" @click="logout">Logout</div>
</template>

<script>
import { mapActions } from "vuex";

export default {
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
        // credentials: 'include',
        headers: {
          Accept: "*/*",
          Authorization: "Bearer " + localStorage.getItem("access_token"),
        },
      });
      if (response.ok) {
        this.removeTokens();
        this.$emit("userLogout");
        alert("Successful logout");
      } else {
        console.log("not ok response", response);
      }
    },
  },
};
</script>

<style scoped></style>
