<template>
  <div id="products_section_id" class="products_section">
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
      this.products_array=[]
      this.showLoading=true
      this.getAllProducts(newValue)
    }
  },
  methods: {
    async getAllProducts(type) {

      let input = "http://localhost:8081/productsbytype?_product_type=" + type
      let resp = await fetch(input, {
        method: "GET",
      });
      this.showLoading=false
      this.products_array = await resp.json()
    },
  },
  created() {

  },
  async mounted() {
    this.products_array = await this.getAllProducts(this.productType);
    this.showLoading=false
  },
};
</script>

<style scoped>
.products_section {
  width: auto;
  height: auto;
  border: solid black;
  background: antiquewhite;
  display: flex;
  flex-wrap: wrap;
  box-sizing: border-box;
  position: relative;
  justify-content: space-evenly;
}

@keyframes fadeIn {
  from {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}

.products_section > * {
  box-sizing: border-box;
  flex: 0 1 15em;
  opacity: 0;
  animation: fadeIn ease-in 1;
  animation-fill-mode: forwards;
  animation-duration: 1s;
}
</style>
