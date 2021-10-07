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
        state.products[id][1]++
    },
    decreaseProduct(state, id) {
        if (state.products[id][1]>1){
        state.products[id][1]--
        }else{
            state.products[id][1]=1
        }
    }
}


const actions = {
    addProduct(context, product) {
        context.commit('addProduct', product)
    },
    removeProduct(context, id) {
        context.commit('removeProduct', id)
    },
    increaseProduct(context, id) {
        context.commit('increaseProduct', id)
    },
    decreaseProduct(context, id) {
        context.commit('decreaseProduct', id)
    }
}

const getters = {
    getProducts: (state) => {
        return state.products
    },
}

export default ({
    namespaced: true,
    state,
    mutations,
    actions,
    getters
});
