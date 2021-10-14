<template>
  <div>
    <div class="main_header">
      <div v-if="!isLogin" @click="showLoginWindow()">Login</div>
      <div @click="cartWindowVisible=true" @userLogout="$emit('userLogout')">Cart</div>
      <div v-if="!isLogin" @click="showRegistrationWindow()">Register</div>
      <logout v-if="isLogin" @userLogout="$emit('userLogout')"></logout>
      <div v-if="isLogin" @click="showProfileWindow()">Profile</div>
      <div v-if="isLogin" @click="showOrdersWindow()">Orders</div>
    </div>
    <dialog_window
        :show="loginWindowVisible"
        @hideDialogWindow="hideLoginWindow"
    >
      <login @hideDialogWindow="hideLoginWindow" @userLogin="$emit('userLogin')"></login>
    </dialog_window>
    <dialog_window
        :show="ordersWindowVisible"
        @hideDialogWindow="hideOrdersWindow"
    >
      <orders>Orders</orders>
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
      <cart :is-login="isLogin" @hideDialogWindow="hideCartWindow()"></cart>
    </dialog_window>
    <dialog_window
        :show="profileWindowVisible"
        @hideDialogWindow="hideProfileWindow()"
    >
      <user_profile @userLogout="$emit('userLogout')"></user_profile>
    </dialog_window>

  </div>
</template>

<script>
import Orders from "./Orders";

export default {
  name: "Header",
  components: {Orders},
  data() {
    return {
      loginWindowVisible: false,
      registrationWindowVisible: false,
      cartWindowVisible: false,
      profileWindowVisible: false,
      ordersWindowVisible: false,
    }
  },
  props: {
    isLogin: Boolean
  },
  methods: {
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
    }

  }
};
</script>

<style scoped>
.main_header {
  display: flex;
  justify-content: space-around;
  height: 15vh;
  background: #2c3e50;
}

.main_header > * {
  color: white;
}
</style>
