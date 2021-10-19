<template>
  <div>
    <div class="main_header">
      <div class="nav_bar">
        <div v-if="currentRouteName==='Home'" style="background-color: coral; opacity: 0.7;">{{this.currentRouteName}}</div>
        <router-link v-if="currentRouteName!=='Home'" to="/">Home</router-link>
        <div v-if="currentRouteName==='Goods'" style="background-color: coral; opacity: 0.7;">{{this.currentRouteName}}</div>
        <router-link v-if="currentRouteName!=='Goods'" to="/goods">Goods</router-link>
        <div v-if="currentRouteName==='About'" style="background-color: coral; opacity: 0.7;">{{this.currentRouteName}}</div>
        <router-link  v-if="currentRouteName!=='About'" to="/about">About</router-link>
      </div>
      <div class="buttons_side">
        <button v-if="!getLoginStatus()" @click="showLoginWindow()">Login</button>
        <button @click="cartWindowVisible=true">Cart</button>
        <button v-if="!getLoginStatus()" @click="showRegistrationWindow()">Register</button>
        <button v-if="getLoginStatus()">
          <logout></logout>
        </button>
        <button v-if="getLoginStatus()" @click="showOrdersWindow()">Orders</button>
        <div v-if="getLoginStatus()" @click="showProfileWindow()"><img class="profile_img"
                                                                       src="../assets/avatar_icon1.png" height="30pt">
        </div>
      </div>

    </div>
    <dialog_window
        :show="loginWindowVisible"
        @hideDialogWindow="hideLoginWindow"
    >
      <login @hideDialogWindow="hideLoginWindow"></login>
    </dialog_window>
    <dialog_window
        :show="ordersWindowVisible"
        @hideDialogWindow="hideOrdersWindow"
    >
      <orders @hideDialogWindow="hideOrdersWindow(), showLoginWindow()">Orders
      </orders>
    </dialog_window>
    <dialog_window
        :show="registrationWindowVisible"
        @hideDialogWindow="hideRegistrationWindow()"
    >
      <registration></registration>
    </dialog_window>
    <dialog_window
        :show="cartWindowVisible"
        @hideDialogWindow="hideCartWindow()"
    >
      <cart :is-login="getLoginStatus()" @hideDialogWindow="hideCartWindow()"></cart>
    </dialog_window>
    <dialog_window
        :show="profileWindowVisible"
        @hideDialogWindow="hideProfileWindow()"
    >
      <user_profile @hideDialogWindow="hideProfileWindow(), showLoginWindow()"></user_profile>
    </dialog_window>

  </div>
</template>

<script>
import Orders from "./Orders";
import {mapGetters, mapMutations} from "vuex";

export default {
  name: "Header",
  components: {Orders},
  props:{
    currentRouteName: String,
},
  data() {
    return {
      loginWindowVisible: false,
      registrationWindowVisible: false,
      cartWindowVisible: false,
      profileWindowVisible: false,
      ordersWindowVisible: false,
    }
  },
  methods: {
    ...mapMutations("tokens", ["setLoginState", "setLogoutState"]),
    ...mapGetters("tokens", ["getLoginStatus"]),

    showProfileWindow() {
      this.profileWindowVisible = true
    },
    hideProfileWindow() {
      this.profileWindowVisible = false
    },
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
    hideCartWindow() {
      this.cartWindowVisible = false;
    },
    showOrdersWindow() {
      this.ordersWindowVisible = true;
    },
    hideOrdersWindow() {
      this.ordersWindowVisible = false;
    },

  },

};
</script>

<style scoped>
.main_header {
  display: flex;
  width: 100%;
  align-items: center;
  /*height: 10%;*/
  background: #2c3e50;
  /*padding: 1pt;*/
}
.nav_bar{
  display: flex;
  width: 50%;
  justify-content: space-around;
  align-items: center;
  height: 10%;
  background: #2c3e50;
  padding: 1pt;
}
.nav_bar>*{
  text-decoration: white;
  color: white;
}
.buttons_side{
  width: 50%;
  display: flex;
  justify-content: space-around;
  align-items: center;
  height: 10%;
  background: #2c3e50;
  padding: 1pt;
}



</style>
