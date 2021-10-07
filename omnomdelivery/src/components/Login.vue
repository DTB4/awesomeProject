<template>
  <div
      class="login_container"
      id="login_container_id"
  >
    <h2>Login</h2>
    <input
        @submit.prevent
        id="login_input_email"
        class="input"
        type="email"
        @input="email=$event.target.value"
        v-bind:value="email"
        placeholder="email"
    >
    <input
        @submit.prevent
        id="login_input_password"
        class="input"
        type="password"
        @input="password=$event.target.value"
        v-bind:value="password"
        placeholder="password"
    >
    <input
        @submit.prevent
        id="login_submit"
        class="button"
        type="button"
        value="Login"
        @click="loginUser()"
    >
    <logout></logout>
  </div>
</template>

<script>
import {mapActions} from "vuex";

export default {
  name: "Login",
  data() {
    return {
      email: '',
      password: '',
    }
  },
  methods: {
    ...mapActions('tokens', ['addTokens']),
    async loginUser() {
      const response = await fetch('http://localhost:8081/login', {
        method: 'POST',
        mode: 'cors',
        credentials: 'include',
        headers: {
          'Accept': '*/*',
        },
        body: JSON.stringify({email: this.email, password: this.password}),
      });
      if (response.status == 200) {
        //TODO: make redirect to main page
        let parsedResponce = await response.json();
        this.addTokens([parsedResponce.access_token, parsedResponce.refresh_token])
        console.log("responce from server", parsedResponce)
        this.email = ''
        this.password = ''
        alert("OK")

      } else {
        //TODO: make message for error in responce
        console.log('not OK response', response)
      }
    }
  }
}
</script>

<style scoped>

</style>