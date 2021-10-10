const state = {
    products: [],
};

const mutations = {
    addProduct(state, product) {
        state.products.push([product, 1]);
    },
    removeProduct(state, id) {
        state.products.splice(id, 1);
    },
    increaseProduct(state, id) {
        state.products[id].splice(1, 1, ++state.products[id][1]);
    },
    decreaseProduct(state, id) {
        if (state.products[id][1] > 1) {
            state.products[id].splice(1, 1, --state.products[id][1]);
        } else {
            state.products[id][1] = 1;
        }
    },
    saveToLocalStorage(state) {
        localStorage.setItem('cart', JSON.stringify(state.products))
    },
    loadFromLocalStorage(state) {
        state.products = JSON.parse(localStorage.getItem('cart'))
    },
    removeAllProducts(state) {
        localStorage.removeItem('cart')
        state.products = []
    }
};

const actions = {
    addProduct(context, product) {
        context.commit("addProduct", product);
        context.commit("saveToLocalStorage");
    },
    removeProduct(context, id) {
        context.commit("removeProduct", id);
        context.commit("saveToLocalStorage");
    },
    increaseProduct(context, id) {
        context.commit("increaseProduct", id);
        context.commit("saveToLocalStorage");
    },
    decreaseProduct(context, id) {
        context.commit("decreaseProduct", id);
        context.commit("saveToLocalStorage");
    },
    loadBackup(context) {
        context.commit("loadFromLocalStorage");
    },
    clearCart(context) {
        context.commit("removeAllProducts")
    }
};

const getters = {
    getProducts: (state) => {
        return state.products;
    },
};

export default {
    namespaced: true,
    state,
    mutations,
    actions,
    getters,
};
