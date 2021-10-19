<template>
  <div class="product">
    <a class="product_link" rel="">
      <img :src="product_ent.img_url" alt="" class="product_img" width="100%"/>
      <h2 class="product_name">{{ product_ent.name }}</h2>
    </a>
    <h4 class="product_type">{{ product_ent.type }}</h4>
    <h4 class="product_price">{{ product_ent.price }} $</h4>
    <input
        class="add_to_cart"
        title="add(+)"
        type="button"
        value="Add to Cart"
        @click="addProduct(product_ent)"
    />
  </div>
</template>

<script>
import {mapActions} from "vuex";

export default {
  name: "Product",
  props: {
    product_ent: Object,
  },
  methods: {
    ...mapActions("cart", ["addProduct"]),
    // handler(product_ent){
    //   this.$store.dispatch('addProduct', product_ent)
    // }
  },
};
</script>

<style scoped>
.product {
  margin: 0.5vh 0.5vh;
  background: aliceblue;
  width: 20vh;
  height: auto;
  border: solid black 0.1pt;
}
</style>
