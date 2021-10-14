<template>
  <div id="app">
    <omnom_header :is-login="isLogin" @userLogin="stateToLogin" @userLogout="stateToLogout"></omnom_header>

    <select v-model="selectedSupplierType" @change="changeSupplierType">
      <option disabled value="">Choose type from list</option>
      <option value="all">All</option>
      <option v-for="(type, id) in supplierTypes" :key="id" :value="type">{{ type }}</option>
    </select>
    <select v-model="selectedWorkingTime" @change="changeWorkingTime">
      <option v-for="(time, id) in workingTimes" :key="id" :value="time">{{ time }}</option>
    </select>
    <suppliers :supplierType="selectedSupplierType" :workingTime="selectedWorkingTime">Suppliers section
    </suppliers>

    <select v-model="selectedProductType" @change="changeProductType">
      <option disabled value="">Choose type from list</option>
      <option value="all">All</option>
      <option v-for="(type, id) in productTypes" :key="id" :value="type">{{ type }}</option>
    </select>


    <products :productType="selectedProductType">Products section</products>

    <footer></footer>
  </div>
</template>

<script>

import {mapActions} from "vuex";

export default {
  //TODO: make short specified requests instead of GetAll
  name: "index",
  components: {},
  data() {
    return {
      isLogin: false,
      supplierTypes: [],
      productTypes: [],
      workingTimes: ["0:00", "1:00", "2:00", "3:00", "4:00", "5:00", "6:00", "7:00", "8:00", "9:00", "10:00", "11:00", "12:00", "13:00", "14:00", "15:00", "16:00", "17:00", "18:00", "19:00", "20:00", "21:00", "22:00", "23:00"],
      selectedSupplierType: "all",
      selectedProductType: "",
      selectedWorkingTime: "",
    };
  },
  methods: {
    ...mapActions("tokens", ["addTokens"]),
    ...mapActions("cart", ["loadBackup"]),
    stateToLogin() {
      this.isLogin = true;
    },
    changeWorkingTime(event) {
      this.$emit('change:selectedWorkingTime', event)
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
    async getTypesOfSuppliers() {
      let resp = await fetch("http://localhost:8081/supplierstypes", {
        method: "GET",
      });
      this.supplierTypes = (await resp.json()).types

    },
    async getTypesOfProducts() {
      let resp = await fetch("http://localhost:8081/productstypes", {
        method: "GET",
      });
      this.productTypes = (await resp.json()).types
    },
  },

  created() {
    this.getTypesOfSuppliers()
    this.getTypesOfProducts()
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
    let date = new Date
    this.selectedWorkingTime = date.getHours().toString() + ":00"
  }
};
</script>

<style>

</style>
