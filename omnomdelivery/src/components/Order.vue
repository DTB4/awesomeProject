<template>
  <div class="order_container" @click="getOrder(order.id), changeVisibility()">
    <div>
      <div>Order # {{ order.id }}</div>
      <div>Status: {{ order.status }}</div>
      <div>Date: {{ order.created }}</div>
      <div>Total: {{}}</div>
    </div>
    <div>
      <orders-product v-for="(product, id) in products"
                      :key="id"
                      :product="product"
                      :show-order="showOrder"></orders-product>
      <div v-if="showOrder">Total: {{ total }}</div>
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
      total: 0
    }
  },
  props: {
    order: Object,
  },
  methods: {
    getTotal() {
      let total = 0
      for (let i = 0; i < this.products.length; i++) {
        total = total + this.products[i].price
      }
      this.total = total.toFixed(2)
    },
    changeVisibility() {
      this.showOrder = this.showOrder !== true;
    },
    async getOrder(id) {
      if (this.products.length !== 0) {
        return
      }
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
        this.getTotal()
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

.order_container > * {
  margin-right: 10px;
}

</style>