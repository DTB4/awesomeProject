import Vue from "vue";
import Vuex from "vuex";
import cartStore from "@/store/cart_vuex";

Vue.use(Vuex);

const modules = {
    cart: cartStore
}

export default new Vuex.Store({
    modules
});
