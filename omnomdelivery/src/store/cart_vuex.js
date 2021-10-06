const state = {
    products: [],
}

const mutations = {
    addProduct(state, product) {
        state.products.push([product, 1])
    },
    removeProduct(state, id) {
        state.products = state.products.filter(i => i.id !== id)
    },
    increaseProduct(state, id) {
        state.product[id][1]++
    },
    decreaseProducts(state, id) {
        state.product[id][1]--
    }
}


const actions = {
    addProduct(context, product) {
        context.commit('addProduct', product)
        context.commit('increaseProduct', product.id)
    },
    removeProduct(context, id) {
        context.commit('removeProduct', id)
        context.commit('decreaseProduct', id)
    }
}

const getters = {
    gotProducts: (state) => {
        return state.products
    },
    gotProductsQty: (state) => {
        return state.quantity
    }
}

export default ({
    namespaced: true,
    state,
    mutations,
    actions,
    getters
});
