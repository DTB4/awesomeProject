const state = {
    access_token: String,
    refresh_token: String,
    is_login: Boolean,
}

const mutations = {
   addAccessToken(state, accessToken){
       state.access_token=accessToken
       localStorage.setItem('access_token', accessToken)
   },
    addRefreshToken(state, refreshToken){
        state.refresh_token=refreshToken
        localStorage.setItem('refresh_token', refreshToken)
    },
   removeAccessToken(state){
       state.access_token=''
       localStorage.removeItem('access_token')
   },
    removeRefreshToken(state){
        state.refresh_token=''
        localStorage.removeItem('refresh_token')
    },
    setLoginState(state){
       state.is_login=true
    },
    setLogoutState(state){
       state.is_login=false
    }
}


const actions = {
    addTokens (context, tokens){
        context.commit('addAccessToken', tokens[0])
        context.commit('addRefreshToken', tokens[1])
        context.commit('setLoginState')

    },
    removeTokens (context){
        context.commit('removeAccessToken')
        context.commit('removeRefreshToken')
        context.commit('setLogoutState')
    }
}

const getters = {
    getLoginStatus:(state)=>{
          return state.is_login
    },
    getAccessToken: (state)=>{
        return state.access_token
    },
    getRefreshToken: (state)=>{
        return state.refresh_token
    }
}

export default ({
    namespaced: true,
    state,
    mutations,
    actions,
    getters
});