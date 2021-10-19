const state = {
    access_token: String,
    refresh_token: String,
    is_login: Boolean,
};

const mutations = {
    addAccessToken(state, accessToken) {
        localStorage.setItem("access_token", accessToken);
        state.access_token = accessToken;
    },

};

const actions = {
    addTokens(context, tokens) {
        context.commit("addAccessToken", tokens[0]);
        context.commit("addRefreshToken", tokens[1]);
        context.commit("setLoginState");
    },
};

const getters = {
    getLoginStatus: (state) => {
        return state.is_login;
    },

};

export default {
    namespaced: true,
    state,
    mutations,
    actions,
    getters,
};
