<template>
  <div
      class="cart"
      id="cart_id"
  >
    <div class="cart_product"
         v-for="(product, id) in cart_products_array" :key="id"
    >
      <div @click="decreaseProduct(id)">"---"</div>
      <div @click="increaseProduct(id)">"+++"</div>
      <div>{{product[0].name}}</div>
      <div>{{product[0].price}}</div>
      <div>"x{{product[1]}}"</div>
      <div @click="removeProduct(id)">..X..</div>
    </div>
    <h2>Total {{totalPrice}}</h2>
    <h1 @click="$forceUpdate">...O...</h1>
  </div>
</template>

<script>
import {mapActions} from "vuex";

export default {
  name: "Cart",
  data(){
    return{
      cart_products_array: [],
      totalPrice: 0,
    }
  },
  methods:{
    ...mapActions('cart', ['addProduct', 'removeProduct', 'increaseProduct', 'decreaseProduct']),
    calcTotalPrice(){
      if (this.cart_products_array.length==0){return 0}
      let total =0
      for (let i=0; i<this.cart_products_array.length; i++){
        total+=this.cart_products_array[i][0].price*this.cart_products_array[i][1]
      }
      this.totalPrice=total
    }
  },
  created() {
    this.cart_products_array=this.$store.getters["cart/getProducts"]
    this.calcTotalPrice()
    },
  beforeUpdate() {
    this.cart_products_array=this.$store.getters["cart/getProducts"]
    this.calcTotalPrice()
  }
}
</script>

<style scoped>

</style>