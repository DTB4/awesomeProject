<template>
  <div class="home">
    <example></example>
    <img alt="Vue logo" src="../assets/logo.png"/>
    <img alt="OmNom Logo" class="omnomlogo" src="../assets/big_logo1.png"/>
    <HelloWorld msg="Welcome to Your Vue.js App"/>
    <cart></cart>
  </div>
</template>

<script>
// @ is an alias to /src
import HelloWorld from "@/components/HelloWorld.vue";
import Example from "@/components/Example";
import Cart from "@/components/Cart";

export default {
  name: "Home",
  components: {
    Cart,
    Example,
    HelloWorld,
  },
};
</script>

<style>
.omnomlogo {
  height: 20vh;
}
</style>
