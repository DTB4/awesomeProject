<template>
  <div class="order_container" @click="getOrder(order.id)">
    <div>{{ order.id }}</div>
    <div>{{ order.status }}</div>
    <div>
      <orders-product v-for="(product, id) in products"
                      :key="id"
                      :show-order="showOrder"
                      :product="product"></orders-product>
    </div>

  </div>
</template>

<script>
import OrdersProduct from "./OrdersProduct";
import refresh_tokens from "../mixins/refresh_tokens";

export default {
  mixins: [refresh_tokens],
  name: "Order",
  components: {OrdersProduct},
  data() {
    return {
      products: [],
      showOrder: false,
    }
  },
  props: {
    order: Object,
  },
  methods: {
    async getOrder(id) {
      const response = await fetch("http://localhost:8081/getorder", {
        method: "POST",
        mode: "cors",
        headers: {
          Accept: "*/*",
          Authorization: "Bearer " + localStorage.getItem("access_token"),
        },
        body: JSON.stringify({order_id: id}),
      });
      if (response.ok) {
        this.products = await response.json()
        console.log(this.products)
        this.showOrder = true
      } else if (response.status === 401) {
        //TODO: try to catch 401 error without "error" in console
        let ok = await this.refreshTokens();
        if (ok) {
          await this.getOrder(id);
        }
      } else {
        console.log("not ok response", response);
      }
    },
  }

}
</script>

<style scoped>
.order_container {
  display: flex;
}
</style>