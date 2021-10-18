<template>
  <div class="products_section">
    <div v-if="showLoading">Loading...</div>
    <product
        v-for="(product, id) in products_array"
        :key="id"
        :product_ent="product"
    ></product>
  </div>
</template>

<script>
export default {
  name: "Products_section",
  data() {
    return {
      products_array: [],
      showLoading: true,
    };
  },
  props: {
    productType: {
      type: String,
      required: false,
      default: "",
    }
  },
  watch: {
    productType(newValue) {
      this.products_array = []
      this.showLoading = true
      this.getProductsByType(newValue)
    }
  },
  methods: {
    async getProductsByType(type) {
      if (type === "") {
        return
      }
      let input = "http://localhost:8081/productsbytype?_product_type=" + type
      let resp = await fetch(input, {
        method: "GET",
      });
      this.showLoading = false
      this.products_array = await resp.json()
    },
  },
  created() {

  },
  async mounted() {
    this.products_array = await this.getProductsByType(this.productType);
    this.showLoading = false
  },
};
</script>

<style scoped>
.products_section {
  overflow: auto;
  width: auto;
  height: 100%;
  border: solid black;
  background: antiquewhite;
  display: flex;
  flex-wrap: wrap;
  box-sizing: border-box;
  position: relative;
  justify-content: space-evenly;
}


.products_section > * {

}
</style>
