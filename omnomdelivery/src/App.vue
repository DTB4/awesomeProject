<template>
  <div id="app">
    <omnom_header :is-login="isLogin" @userLogin="stateToLogin" @userLogout="stateToLogout"></omnom_header>

    <suppliers>Suppliers section</suppliers>
    <products>Products section</products>

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
    };
  },
  methods: {
    ...mapActions("tokens", ["addTokens"]),
    ...mapActions("cart", ["loadBackup"]),
    stateToLogin() {
      this.isLogin = true;
    },
    stateToLogout() {
      this.isLogin = false;
    },
  },

  created() {
    console.log("on APP created")
    let get_access_token = localStorage.getItem('access_token')
    let get_refresh_token = localStorage.getItem('refresh_token')
    console.log("refresh_token from local storage", get_refresh_token)
    if (get_access_token !== null && get_access_token !== 'null') {
      console.log("tokens in local storage exist")
      this.addTokens([get_access_token, get_refresh_token])
      this.isLogin = true
    }
    let get_products = localStorage.getItem('cart')
    console.log("products from local storage", get_products)
    if (get_products !== null && get_products !== 'null') {
      console.log("products in local storage exist")
      this.loadBackup()
    }
  }
};
</script>

<style>
.main_header {
  display: flex;
  justify-content: space-around;
}
</style>
