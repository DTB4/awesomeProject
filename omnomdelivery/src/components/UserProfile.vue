<template>
  <div id="user_profile_container_id" class="user_profile_container">
    <h2>Profile</h2>
    {{ first_name }}
    {{ last_name }}
    {{ email }}
  </div>
</template>

<script>
import {mapActions} from "vuex";
import refresh_tokens from "../mixins/refresh_tokens";

export default {
  mixins: [refresh_tokens],
  name: "Profile",
  data() {
    return {
      first_name: "",
      last_name: "",
      email: "",
    };
  },
  methods: {
    ...mapActions("tokens", ["removeTokens"]),
    async getUserProfile() {
      const response = await fetch("http://localhost:8081/profile", {
        method: "GET",
        mode: "cors",
        headers: {
          Accept: "*/*",
          Authorization: "Bearer " + localStorage.getItem("access_token"),
        },
      });
      if (response.ok) {
        let parsedResponse = await response.json();
        this.email = parsedResponse.email;
        this.first_name = parsedResponse.first_name;
        this.second_name = parsedResponse.last_name;
      } else if (response.status === 401) {
        //TODO: try to catch 401 error without "error" in console
        let ok = await this.refreshTokens()
        if (ok) {
          console.log("try again getUserProfile")
          await this.getUserProfile()
        } else {
          console.log("not ok response", response);
        }
      }
    },
  },
  async created() {
    await this.getUserProfile();
  },
};
</script>

<style scoped></style>
