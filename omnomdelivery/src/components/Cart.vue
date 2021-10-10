<template>
  <div class="cart" id="cart_id">
    <div v-if="cart_products_array.length === 0">Cart is empty</div>
    <div
        class="cart_product"
        v-for="(product, id) in cart_products_array"
        :key="id"
    >
      <button @click="decreaseProduct(id)">-</button>
      <button @click="increaseProduct(id)">+</button>
      <div>"x{{ product[1] }}"</div>
      <div>{{ product[0].name }}</div>
      <div>{{ product[0].price }} $</div>
      <button @click="removeProduct(id)">X</button>
    </div>
    <h2>Total {{ totalPrice }}</h2>
    <div @click="createOrder" v-if="isLogin">Create Order</div>
    <div v-if="!isLogin">pls login to make order</div>
  </div>
</template>

<script>
import {mapActions} from "vuex";
import refresh_tokens from "../mixins/refresh_tokens";
import {mapGetters} from "vuex/dist/vuex.mjs";

export default {
  mixins: [refresh_tokens],
  name: "Cart",
  props: {
    isLogin: {
      type: Boolean,
      required: true,
    }
  },
  data() {
    return {
      cart_products_array: [],
    };
  },
  computed: {
    totalPrice() {
      if (this.cart_products_array.length === 0) {
        return 0;
      }
      let total = 0;
      for (let i = 0; i < this.cart_products_array.length; i++) {
        total +=
            this.cart_products_array[i][0].price *
            100 *
            this.cart_products_array[i][1];
      }
      return total / 100;
    },
  },
  methods: {
    ...mapActions("tokens", ["removeTokens"]),
    ...mapGetters("cart", ["getProducts"]),
    ...mapActions("cart", [
      "addProduct",
      "removeProduct",
      "increaseProduct",
      "decreaseProduct",
      "getTotal",
      "clearCart",
    ]),
    async createOrder() {
      console.log('orderCreate Button pressed')
      let productsBody = []
      for (let i = 0; i < this.cart_products_array.length; i++) {
        productsBody.push({
          order_id: 0,
          product_id: this.cart_products_array[i][0].id,
          quantity: this.cart_products_array[i][1],
        })
        console.log(productsBody)
      }
      console.log(productsBody)
      const response = await fetch("http://localhost:8081/createorder", {
        method: "POST",
        mode: "cors",
        headers: {
          Accept: "*/*",
          Authorization: "Bearer " + localStorage.getItem("access_token"),
        },
        body: JSON.stringify(productsBody),
      });
      if (response.ok) {
        this.clearCart()
        this.$emit("hideDialogWindow")
        alert("order Created")
      } else if (response.status === 401) {
        //TODO: try to catch 401 error without "error" in console
        let ok = await this.refreshTokens();
        if (ok) {
          await this.createOrder();
        }
      } else {
        console.log("not ok response", response, response.text());
        alert("fail to create an order");
      }
    },
  },
  created() {
    this.cart_products_array = this.getProducts()
  },
};
</script>

<style scoped>
.cart_product {
  display: flex;
  justify-content: flex-start;
  padding: 5px;
}

.cart_product > * {
  padding: 0 5px 0 5px;
}
</style>
