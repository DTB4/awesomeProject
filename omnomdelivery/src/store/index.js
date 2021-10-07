import Vue from "vue";
import Vuex from "vuex";
import cartStore from "@/store/cart_vuex";
import tokensStore from "@/store/tokens_vuex"

Vue.use(Vuex);

const modules = {
    cart: cartStore,
    tokens: tokensStore,
}

export default new Vuex.Store({
    modules
});
