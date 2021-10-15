<template>
  <div>
    <div id="supplier_section_id" class="supplier_section">
      <div v-if="showLoading">Loading...</div>
      <supplier v-for="(supplier, id) in suppliers_array"
                :key="id"
                :supplier_ent="supplier"
                @showProductsFor="showProductsFor"
      ></supplier>

    </div>
    <div class="supplier_products_section">
      <product
          v-for="(product, id) in products"
          :key="id"
          :product_ent="product"
      ></product>
    </div>
  </div>
</template>

<script>
export default {
  name: "Supplier_section",
  data() {
    return {
      suppliers_array: [],
      showLoading: true,
      productsForID: 0,
      products: [],
    };
  },
  //TODO: make sorting by work time
  props: {
    workingTime: {
      type: String,
      required: true,
    },
    supplierType: {
      type: String,
      required: false,
      default: "",
    }
  },
  watch: {
    async workingTime(newValue) {
      await this.getSuppliers(this.supplierType, newValue)
      this.productsForID = 0
      this.products = []

    },
    async supplierType(newValue) {
      await this.getSuppliers(newValue, this.workingTime)
      this.productsForID = 0
      this.products = []

    },
    async productsForID(newValue) {
      await this.getProductsForSupplierByID(newValue)
    }
  },
  methods: {
    async getProductsForSupplierByID(id) {
      if (id === 0) {
        return
      }
      let input = "http://localhost:8081/productsbysupplier?" + "_supplier_id=" + id
      let resp = await fetch(input, {
        method: "GET",
      });
      this.products = await resp.json()

    },
    showProductsFor(event) {
      this.productsForID = event
    },
    async getSuppliers(type, time) {
      let input = "http://localhost:8081/supplierparam?" + "_type=" + type + "&_time=" + time
      let resp = await fetch(input, {
        method: "GET",
      });
      this.suppliers_array = await resp.json()

    },
  },
  async created() {
    await this.getSuppliers(this.supplierType, this.workingTime);
    this.showLoading = false
    console.log(this.suppliers_array)
  },
};
</script>

<style scoped>
.supplier_section,
.supplier_products_section {
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

.supplier_section, > * {
  box-sizing: border-box;
  flex: 0 1 15em;
  opacity: 0;
  animation: fadeIn ease-in 1;
  animation-fill-mode: forwards;
  animation-duration: 1s;
}
</style>
