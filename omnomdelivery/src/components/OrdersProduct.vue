<template>
  <div>
    <div class="orders_product" v-if="showOrder"
    >
      <h4>Product ID : {{ product.product_id }}</h4>
      <h4>Quantity : {{ product.quantity }}</h4>
    </div>
  </div>
</template>

<script>
export default {
  name: "OrdersProduct",
  props: {
    showOrder: Boolean,
    product: Object,
  }
}
</script>

<style scoped>
.orders_product {
  display: flex;
  justify-content: space-around;
}
</style>