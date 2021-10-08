<template>
  <div class="user_profile_container" id="user_profile_container_id">
    <h2>Profile</h2>
    {{ first_name }}
    {{ last_name }}
    {{ email }}
  </div>
</template>

<script>
export default {
  name: "Profile",
  data() {
    return {
      first_name: "",
      last_name: "",
      email: "",
      data_access_token: "",
      data_refresh_token: "",
    };
  },
  methods: {
    async getUserProfile() {
      const response = await fetch("http://localhost:8081/profile", {
        method: "GET",
        mode: "cors",
        // credentials: 'include',
        headers: {
          Accept: "*/*",
          Authorization: "Bearer " + this.data_access_token,
        },
      });
      if (response.ok) {
        //TODO: make redirect to Login
        let parsedResponce = await response.json();
        console.log("responce from server", parsedResponce);
        this.email = parsedResponce.email;
        this.first_name = parsedResponce.first_name;
        this.second_name = parsedResponce.last_name;
        alert("OK");
      } else if (response.status === 401) {
        //TODO: try to catch 401 error without "error" in console
        let ok = await this.refreshTokens();
        if (ok) {
          await this.getUserProfile();
        }
      } else {
        console.log("something else happens", response);
      }
    },
    async refreshTokens() {
      console.log("refresh function is started");
      const response = await fetch("http://localhost:8081/refresh", {
        method: "GET",
        mode: "cors",
        // credentials: 'include',
        headers: {
          Accept: "*/*",
          Authorization: "Bearer " + this.data_refresh_token,
        },
      });
      if (response.status === 200) {
        let parsedResponce = await response.json();
        console.log("responce from server", parsedResponce);
        localStorage.setItem("access_token", parsedResponce.access_token);
        localStorage.setItem("refresh_token", parsedResponce.refresh_token);
        this.data_refresh_token = localStorage.getItem("refresh_token");
        this.data_access_token = localStorage.getItem("access_token");
        return true;
      } else {
        //TODO: make redirect to login
        alert("both tokens are expired. Login needed");
      }
    },
  },
  async mounted() {
    let get_access_token = localStorage.getItem("access_token");
    if (get_access_token !== undefined) {
      this.data_access_token = get_access_token;
      this.data_refresh_token = localStorage.getItem("refresh_token");
      await this.getUserProfile();
    } else {
      alert("You need to login for access to profile");
    }
  },
};
</script>

<style scoped></style>
