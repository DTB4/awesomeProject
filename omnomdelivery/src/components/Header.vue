<template>
  <header>
    <div class="drop_down_menu_div">
      <ul class="drop_down_menu">

      </ul>
    </div>
    <div class="header_links">
      <router-link to="/about">About</router-link>
      <router-link to="/">Home</router-link>
      <router-link to="/cooperation">Cooperation</router-link>
    </div>
  </header>
</template>

<script>
export default {
  name: "Header"
}
</script>

<style scoped>
header {
  height: 15vh;
  background: #2c3e50;

}
</style>