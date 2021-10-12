<template>
  <div id="products_section_id" class="products_section">
    <product
        v-for="(product, id) in products_shown"
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
      products_shown: [],
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
      if (newValue === 'all') {
        this.products_shown = this.products_array
        return
      }
      this.products_shown = this.products_array.filter(product => product.type === newValue)
    }
  },
  methods: {
    async getAllProducts() {
      let resp = await fetch("http://localhost:8081/products", {
        method: "GET",
      });
      return resp.json();
    },
  },
  created() {

  },
  async mounted() {
    document.getElementById("products_section_id").innerText = "Loading";
    this.products_array = await this.getAllProducts();
    this.products_shown = this.products_array
    localStorage.setItem("Products", JSON.stringify(this.products_array))
    document.getElementById("products_section_id").innerText = "";
    this.$emit('products_loaded')
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
