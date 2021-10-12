<template>
  <div class="login_container" id="login_container_id">
    <h2>Login</h2>
    <form>
      <input
          @submit.prevent
          id="login_input_email"
          class="input"
          type="email"
          @input="email = $event.target.value"
          v-bind:value="email"
          placeholder="email"
          autocomplete="email"
      />
      <input
          @submit.prevent
          id="login_input_password"
          class="input"
          type="password"
          @input="password = $event.target.value"
          v-bind:value="password"
          placeholder="password"
          autocomplete="current-password"
      />
    </form>

    <input
        @submit.prevent
        id="login_submit"
        class="button"
        type="button"
        value="Login"
        @click="loginUser()"
    />
  </div>
</template>

<script>
import {mapActions} from "vuex";

export default {
  name: "Login",
  data() {
    return {
      email: "",
      password: "",
      show: {
        type: Boolean,
        default: false,
      },
    };
  },
  methods: {
    ...mapActions("tokens", ["addTokens"]),
    async loginUser() {
      const response = await fetch("http://localhost:8081/login", {
        method: "POST",
        mode: "cors",
        credentials: "include",
        headers: {
          Accept: "*/*",
        },
        body: JSON.stringify({email: this.email, password: this.password}),
      });
      if (response.status === 200) {
        //TODO: make redirect to main page
        let parsedResponse = await response.json();
        this.addTokens([
          parsedResponse.access_token,
          parsedResponse.refresh_token,
        ]);
        this.$emit("userLogin");
        this.$emit("hideDialogWindow")
        console.log("response from server", parsedResponse);
        this.email = "";
        this.password = "";

      } else {
        //TODO: make message for error in response
        console.log("not OK response", response);
      }
    },
  },
};
</script>

<style scoped></style>
