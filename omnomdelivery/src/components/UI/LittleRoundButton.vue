<template>
  <div class="little_round_button">
    <slot></slot>
  </div>
</template>

<script>
export default {
  name: "LittleRoundButton"
}
</script>

<style scoped>
.little_round_button {
  color: coral;
  display: flex;
  justify-content: center;
  align-items: center;
  border: coral solid 1px;
  width: 10pt;
  height: 10pt;
  border-radius: 50%;
  box-shadow: black;
}
</style>