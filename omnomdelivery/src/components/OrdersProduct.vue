<template>
  <div>
    <div v-if="showOrder" class="orders_product"
    >
      <h4>Product : {{ productName }}</h4>
      <h4>Quantity : {{ product.quantity }}</h4>
      <h4>Price : {{ product.price }}</h4>
      <h4>Sum: {{ ((product.quantity * (product.price * 100)) / 100).toFixed(2) }}</h4>
    </div>
  </div>
</template>

<script>
export default {
  name: "OrdersProduct",
  data() {
    return {
      productName: ""
    }
  },
  props: {
    showOrder: Boolean,
    product: Object,
  },
  created() {
    let products = JSON.parse(localStorage.getItem("Products"))
    let product = products.filter(i => i.id === this.product.product_id)
    this.productName = product[0].name

  }
}
</script>

<style scoped>
.orders_product {
  display: flex;
  justify-content: space-around;
}
</style>