<template>
  <div id="login_container_id" class="login_container">
    <h2>Login</h2>
    <form>
      <input
          id="login_input_email"
          autocomplete="email"
          class="input"
          placeholder="email"
          type="email"
          v-bind:value="email"
          @input="email = $event.target.value"
          @submit.prevent
      />
      <input
          id="login_input_password"
          autocomplete="current-password"
          class="input"
          placeholder="password"
          type="password"
          v-bind:value="password"
          @input="password = $event.target.value"
          @submit.prevent
      />
    </form>

    <input
        id="login_submit"
        class="button"
        type="button"
        value="Login"
        @click="loginUser()"
        @submit.prevent
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
