<template>
  <div id="app">
    <omnom_header :is-login="isLogin" @userLogin="stateToLogin" @userLogout="stateToLogout"></omnom_header>

    <select v-if="showSuppliersSelect" v-model="selectedSupplierType" @change="changeSupplierType">
      <option disabled value="">Choose type from list</option>
      <option value="all">All</option>
      <option v-for="(type, id) in supplierTypes" :key="id" :value="type">{{ type }}</option>
    </select>
    <suppliers :supplierType="selectedSupplierType" @suppliers_loaded="getTypesOfSuppliers">Suppliers section
    </suppliers>

    <select v-if="showProductsSelect" v-model="selectedProductType" @change="changeProductType">
      <option disabled value="">Choose type from list</option>
      <option value="all">All</option>
      <option v-for="(type, id) in productTypes" :key="id" :value="type">{{ type }}</option>
    </select>

    <products :productType="selectedProductType" @products_loaded="getTypesOfProducts">Products section</products>

    <footer></footer>
  </div>
</template>

<script>

import {mapActions} from "vuex";

export default {
  name: "index",
  components: {},
  data() {
    return {
      isLogin: false,
      supplierTypes: [],
      productTypes: [],
      showSuppliersSelect: false,
      showProductsSelect: false,
      selectedSupplierType: "",
      selectedProductType: "",
    };
  },
  methods: {
    ...mapActions("tokens", ["addTokens"]),
    ...mapActions("cart", ["loadBackup"]),
    stateToLogin() {
      this.isLogin = true;
    },
    changeSupplierType(event) {
      this.$emit('change:selectedSupplierType', event)
    },
    changeProductType(event) {
      this.$emit('change:selectedProductType', event)
    },
    stateToLogout() {
      this.isLogin = false;
    },
    getTypesOfSuppliers() {
      let types = []
      let suppliers = JSON.parse(localStorage.getItem("Suppliers"))
      for (let i = 0; i < suppliers.length; i++) {
        types.push(suppliers[i].type)
      }
      types = types.filter((x, i, a) => a.indexOf(x) === i)
      this.supplierTypes = types
      this.showSuppliersSelect = true
    },
    getTypesOfProducts() {
      let types = []
      let suppliers = JSON.parse(localStorage.getItem("Products"))
      for (let i = 0; i < suppliers.length; i++) {
        types.push(suppliers[i].type)
      }
      types = types.filter((x, i, a) => a.indexOf(x) === i)
      this.productTypes = types
      this.showProductsSelect = true
    },
  },

  created() {
    let get_access_token = localStorage.getItem('access_token')
    let get_refresh_token = localStorage.getItem('refresh_token')
    if (get_access_token !== null && get_access_token !== 'null') {
      this.addTokens([get_access_token, get_refresh_token])
      this.isLogin = true
    }
    let get_products = localStorage.getItem('cart')
    if (get_products !== null && get_products !== 'null') {
      this.loadBackup()
    }
  }
};
</script>

<style>

</style>
