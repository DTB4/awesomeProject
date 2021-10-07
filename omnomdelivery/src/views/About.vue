<template>
  <div class="about">
    <h1>This is an about page</h1>
    <div>{{ result }}</div>
    <div>{{ step }}</div>
    <div ref="test">TEST</div>
    <card :title="name" :value="value" :data="data" :clickable="clickable" @increment="handler"></card>
    <registration></registration>
    <login></login>
    <user_profile></user_profile>

    <suppliers>Suppliers section</suppliers>
    <products>Products section</products>
  </div>

</template>
<script>

export default {
  name: "About",
  components: {},
  data() {
    return {
      name: "About",
      data: {a: 1, b: 2},
      clickable: false,
      value: 1,
      step: 0,
      updated: 0,
    }
  },
  methods: {
    handler(val) {
      this.value = val
    }
  },
  computed: {
    result() {
      return {a: this.data.a + this.value, b: this.data.b + this.value}
    }
  },
  watch: {
    step(val) {
      this.clickable = val > 10
    },
    value(newVal, oldVal) {
      this.step += newVal - oldVal
    }
  },
  created() {
    // this.$ref.test.innerText="NOT CHANGED"
  },
  mounted() {
    this.$refs.test.innerText = "CNANGED"
  },
  updated() {
    console.log("Updated!")
  }
}
</script>