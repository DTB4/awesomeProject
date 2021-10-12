<template>
  <div>
    <div id="supplier_section_id" class="supplier_section">
      <div v-if="showLoading">Loading...</div>
      <supplier v-for="(supplier, id) in suppliers_array_shown"
                :key="id"
                :supplier_ent="supplier"
                @showProductsFor="showProductsFor"
      ></supplier>

    </div>
    <div class="supplier_products_section">
      <product
          v-for="(product, id) in productsShown"
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
      suppliers_array_shown: [],
      showLoading: true,
      productsForID: 0,
      products: [],
      productsShown: [],
    };
  },
  props: {
    supplierType: {
      type: String,
      required: false,
      default: "",
    }
  },
  watch: {
    supplierType(newValue) {
      if (newValue === 'all') {
        this.suppliers_array_shown = this.suppliers_array
        return
      }
      this.suppliers_array_shown = this.suppliers_array.filter(supplier => supplier.type === newValue)
      this.productsForID = 0
    },
    productsForID(newValue) {
      this.productsShown = this.products.filter(product => product.id_supplier === newValue)
    }
  },
  methods: {
    showProductsFor(event) {
      this.productsForID = event
    },
    async getAllSuppliers() {
      let resp = await fetch("http://localhost:8081/suppliers", {
        method: "GET",
      });
      return resp.json();
    },
  },
  async created() {
    this.suppliers_array = await this.getAllSuppliers();
    this.showLoading = false
    this.suppliers_array_shown = this.suppliers_array
    localStorage.setItem('Suppliers', JSON.stringify(this.suppliers_array))
    this.$emit('suppliers_loaded')
    this.products = JSON.parse(localStorage.getItem('Products'))
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
