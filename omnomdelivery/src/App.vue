<template>
  <div class="app">
    <omnom_header></omnom_header>
    <div>
      <router-link to="/">Home</router-link>
      <router-link to="/goods">Products</router-link>
      <router-link to="/about">About</router-link>
    </div>

    <router-view/>

    <Footer></Footer>
  </div>
</template>

<script>


import {mapActions} from "vuex";
import Footer from "./components/Footer";

export default {
  name: "index",
  components: {Footer},
  methods: {
    ...mapActions("tokens", ["loadLocalTokens"]),
    ...mapActions("cart", ["loadLocalCart"]),
  },
  created() {
    this.loadLocalTokens()
    this.loadLocalCart()

  }
};
</script>

<style>
.app{


}
input {
  width: 100%;
  border: 1px solid coral;
  margin-top: 1pt;
  padding: 1pt 2pt;

}
button{
  border: 1px solid coral;
  margin: 5pt;
  padding: 5pt 7pt;
  background: none;
  color: coral;
}
button:hover{
  box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.5);
  transform: translateY(-2px);
}
button:active{
  box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2);
  transform: translateY(2px);
  opacity: 0.7;
}
*{
  padding: 0;
  margin: 0;
  box-sizing: border-box;
}
</style>
