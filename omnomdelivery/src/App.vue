<template>
  <div id="app">
    <omnom_header :is-login="isLogin" @userLogin="stateToLogin" @userLogout="stateToLogout"></omnom_header>
    <div>
      <router-link to="/">Home</router-link>
      <router-link to="/about">About</router-link>
    </div>

    <router-view/>

    <footer></footer>
  </div>
</template>

<script>


import {mapActions} from "vuex";

export default {
  name: "index",
  components: {},
  data() {
    return {
      isLogin: false,
    }
  },
  methods: {
    ...mapActions("tokens", ["addTokens"]),
    ...mapActions("cart", ["loadBackup"]),

    stateToLogout() {
      this.isLogin = false;
    },
    stateToLogin() {
      this.isLogin = true;
    },
  },
  created() {
    let get_access_token = localStorage.getItem('access_token')
    let get_refresh_token = localStorage.getItem('refresh_token')
    if (get_access_token !== null && get_access_token !== 'null') {
      this.addTokens([get_access_token, get_refresh_token])
      this.isLogin = true
    }
    let get_products = localStorage.getItem('cart')
    if (get_products !== null && get_products !== 'null') {
      this.loadBackup()
    }
  }
};
</script>

<style>

</style>
