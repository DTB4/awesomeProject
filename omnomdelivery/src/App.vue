<template>
  <div v-on:userLogout="stateToLogout" v-on:userLogin="stateToLogin"  id="app">
    <div  class="main_header">
      <div  v-if="!isLogin"   @click="showLoginWindow()">Login</div>
      <div @click="showCartWindow()">Cart</div>
      <div v-if="!isLogin" @click="showRegistrationWindow()">Register</div>
      <logout v-if="isLogin" ></logout>
    </div>
    <dialog_window
      :show="loginWindowVisible"
      @hideLoginWindow="hideLoginWindow()"
    >
      <login></login>
    </dialog_window>
    <dialog_window
      :show="registrationWindowVisible"
      @hideLoginWindow="hideRegistrationWindow()"
    >
      <registration></registration>
    </dialog_window>
    <dialog_window
      :show="cartWindowVisible"
      @hideLoginWindow="hideCartWindow()"
    >
      <cart></cart>
    </dialog_window>
    <suppliers>Suppliers section</suppliers>
    <products>Products section</products>

    <footer></footer>
  </div>
</template>

<script>

export default {
  name: "index",
  data() {
    return {
      isLogin: false,
      loginWindowVisible: false,
      registrationWindowVisible: false,
      cartWindowVisible: false,
    };
  },
  methods: {
    showLoginWindow() {
      this.loginWindowVisible = true;
    },
    hideLoginWindow() {
      this.loginWindowVisible = false;
    },
    showRegistrationWindow() {
      this.registrationWindowVisible = true;
    },
    hideRegistrationWindow() {
      this.registrationWindowVisible = false;
    },
    showCartWindow() {
      this.cartWindowVisible = true;
    },
    hideCartWindow() {
      this.cartWindowVisible = false;
    },
    stateToLogin(event) {
      if (event){this.isLogin = true;}
    },
    stateToLogout() {
      this.isLogin = false;
    },
  },
};
</script>

<style>
.main_header {
  display: flex;
  justify-content: space-around;
}
</style>
