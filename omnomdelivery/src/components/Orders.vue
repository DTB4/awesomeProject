<template>
  <div class="orders_container">
    <Order v-for="(order, id) in orders"
           :key="id"
           :order="order"
    ></Order>
  </div>
</template>

<script>
import refresh_tokens from "../mixins/refresh_tokens";
import Order from "./Order";

export default {
  components: {Order},
  mixins: [refresh_tokens],
  name: "Orders",
  data() {
    return {
      orders: [],
    }
  },
  methods: {
    async getOrders() {
      const response = await fetch("http://localhost:8081/getmyorders", {
        method: "GET",
        mode: "cors",
        headers: {
          Accept: "*/*",
          Authorization: "Bearer " + localStorage.getItem("access_token"),
        },
      });
      if (response.ok) {
        this.orders = await response.json()

      } else if (response.status === 401) {
        //TODO: try to catch 401 error without "error" in console
        let ok = await this.refreshTokens();
        if (ok) {
          console.log("try again getOrders")
          await this.getOrders();
        }
      } else {
        console.log("not ok response", response);
      }
    },
  },
  created() {
    this.getOrders()
  }
}
</script>

<style scoped>
.orders_container {

}
</style>