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
  </div>
</template>

<script>
import { mapActions } from "vuex";

export default {
  name: "Cart",
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
    ...mapActions("cart", [
      "addProduct",
      "removeProduct",
      "increaseProduct",
      "decreaseProduct",
      "getTotal",
    ]),
  },
  created() {
    this.cart_products_array = this.$store.getters["cart/getProducts"];
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
