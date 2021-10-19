<template>
  <div class="product" v-if="product_ent">

    <div class="img_container" :style="{ backgroundImage: `url(${product_ent.img_url})` }">
      <!--      <img :src="product_ent.img_url" alt="" class="product_img"/>-->

    </div>
    <div class="product_info"><h3 class="product_name">{{ product_ent.name }}</h3>
      <div class="type_and_price"><h4 class="product_type">{{ product_ent.type }}</h4>
        <h4 class="product_price">{{ product_ent.price }} $</h4></div>
      <button
          class="add_to_cart"
          @click="addProduct(product_ent)"
      >Add to Cart
      </button>
    </div>
    <div class="drop_down_content">{{ product_ent.ingredients.toString().replace("[", "").replace("]", "") }}</div>
  </div>
</template>

<script>
import {mapActions} from "vuex";

export default {
  name: "Product",
  props: {
    product_ent: Object,
  },
  data() {
    return {
      stylePart: ""
    }
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
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2);
  max-width: 25vw;
  height: 30vh;
  text-align: center;
  font-family: arial, serif;
  background: aliceblue;
  margin: 5pt;
  border-radius: 10px;
  box-sizing: border-box;
  flex: 0 1 15em;
  opacity: 0;
  animation: fadeIn ease-in 1;
  animation-fill-mode: forwards;
  animation-duration: 1s;
}

.product:hover .drop_down_content {
  display: block;
}

@keyframes fadeIn {
  from {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}

.product:hover {
  box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.5);
  transform: translateY(-2px);
}

.img_container {
  width: 100%;
  height: 60%;
  background-position: center;
  background-size: 100%;
  border-radius: 10px;
}

.type_and_price {
  padding: 5pt;
  display: flex;
  flex-direction: row;
  justify-content: space-between;
}

.product_price {
  opacity: 0.5;
}

.product_type {
  color: coral;
}

.product_info {
  align-items: center;
}

button {

}

.drop_down_content {
  display: none;
  position: absolute;
  top: 40%;
  background-color: aliceblue;
  width: 100%;
  box-shadow: 0 8px 16px 0 rgba(0, 0, 0, 0.2);
  padding: 12px 16px;
  z-index: 1;
  opacity: 0.8;
}

/*.product_img{*/
/*  max-width: 100%;*/
/*  height: auto;*/
/*}*/

</style>
