<template>
  <div class="Goods">
    <select v-model="selectedProductType" @change="changeProductType">
      <option disabled value="">Choose type from list</option>
      <option  selected value="all">All</option>
      <option v-for="(type, id) in productTypes" :key="id" :value="type">{{ type }}</option>
    </select>

    <products :productType="selectedProductType">Products section></products>
  </div>
</template>

<script>
export default {
  name: "Goods",
  data(){
    return{
      productTypes: [],
      selectedProductType: "",
    }
  },
  methods: {
    changeProductType(event) {
      this.$emit('change:selectedProductType', event)
    },
    async getTypesOfProducts() {
      let resp = await fetch("http://localhost:8081/productstypes", {
        method: "GET",
      });
      this.productTypes = (await resp.json()).types
    },
  },
  created() {
    this.getTypesOfProducts()
  },
};
</script>

<style scoped>
.Goods{
  height: 90%;
}
</style>
