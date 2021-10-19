<template>
  <div id="cart_id" class="cart">
    <div>
      <div v-if="cart_products_array.length === 0">Cart is empty</div>
      <div
          v-for="(product, id) in cart_products_array"
          :key="id"
          class="cart_product"
      >
        <button @click="decreaseProduct(id)">-</button>
        <button @click="increaseProduct(id)">+</button>
        <div>"x{{ product[1] }}"</div>
        <div>{{ product[0].name }}</div>
        <div>{{ product[0].price }} $</div>
        <button @click="removeProduct(id)">X</button>
      </div>
      <h2 style="justify-self: end">Total {{ totalPrice }}</h2>
      <button v-if="getLoginStatus && cart_products_array.length !== 0" @click="showConfirmationWindow=true">Make Order
      </button>
      <div v-if="!getLoginStatus">pls login to make order</div>
    </div>
    <div v-if="showConfirmationWindow">
      <input
          id="confirmation_input_address"
          class="input"
          placeholder="address"
          type="text"
          v-bind:value="address"
          @input="address = $event.target.value"
          @submit.prevent
      />
      <input
          id="confirmation_input_contact_number"
          class="input"
          placeholder="contact_number"
          type="text"
          v-bind:value="contactNumber"
          @input="contactNumber = $event.target.value"
          @submit.prevent
      />
      <button @click="createOrder">Confirm Order</button>
    </div>
  </div>
</template>

<script>
import {mapActions} from "vuex";
import refresh_tokens from "../mixins/refresh_tokens";
import {mapGetters} from "vuex/dist/vuex.mjs";

export default {
  mixins: [refresh_tokens],
  name: "Cart",
  data() {
    return {
      cart_products_array: [],
      showConfirmationWindow: false,
      address: "",
      contactNumber: "",
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
            this.cart_products_array[i][0].price * this.cart_products_array[i][1];
      }
      return total.toFixed(2);
    },
  },
  methods: {
    ...mapGetters("tokens", ["getLoginStatus"]),
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
      let productsBody = {};
      productsBody.address = this.address
      productsBody.contact_number = this.contactNumber
      productsBody.products = []
      console.log(productsBody);
      for (let i = 0; i < this.cart_products_array.length; i++) {
        productsBody.products.push({
          order_id: 0,
          product_id: this.cart_products_array[i][0].id,
          quantity: this.cart_products_array[i][1],
          name: this.cart_products_array[i][0].name,
        })

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
        let resp = await response.json()
        alert(`order Created with ID:  ${resp.order_id} and ${resp.product_qty} products. Total: ${resp.total}`)
      } else if (response.status === 401) {
        //TODO: try to catch 401 error without "error" in console
        let ok = await this.refreshTokens();
        if (ok) {
          console.log("try again createOrder")
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
  text-justify: auto;
  border-bottom: coral dotted 1px;
}

.cart_product > * {
  padding: 5px 5px 5px 5px;
}
</style>
