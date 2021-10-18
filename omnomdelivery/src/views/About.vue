<template>
  <div class="about">
    <h1>This is an about page</h1>
    <little-round-button>+</little-round-button>

    <little-round-button>-</little-round-button>



  </div>
</template>
<script>
import LittleRoundButton from "../components/UI/LittleRoundButton";
export default {
  name: "About",
  components: {LittleRoundButton},
};
</script>
