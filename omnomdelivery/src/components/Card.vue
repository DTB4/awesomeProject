<template>
  <div>
    <div>{{ title }}</div>
    <div>{{ value }}</div>
    <div>{{ data }}</div>
    <div :title="'Jump to:'+ $router.getRoutes()[0].name"
         @click="$router.push({name: $router.getRoutes()[0].name})"
         :class="{clickable: clickable}"
         :style="{'font-size':'25px'}"

    >Go home
    </div>
    <div @click="increment" class="clickable">INC</div>
  </div>
</template>

<script>
export default {
  name: "Card",
  props: {
    clickable: {
      boolean: true,
      default: false,
    },
    title: {
      type: String,
      require: true
    },
    value: {
      type: Number,
      default: 1,
      validator(value) {
        return value > 0
      }

    },
    data: {
      type: Object
    }
  },
  methods: {
    increment() {
      this.$emit('increment', this.value + 1)
    }
  }
}
</script>

<style scoped>
.clickable {
  cursor: pointer;
}
</style>