import {mapActions} from "vuex";

export default {
    data() {
        return {}
    },
    methods: {
        ...mapActions("tokens", ["addTokens"]),
        async refreshTokens() {
            console.log("refresh function is started");
            const response = await fetch("http://localhost:8081/refresh", {
                method: "GET",
                mode: "cors",
                // credentials: 'include',
                headers: {
                    Accept: "*/*",
                    Authorization: "Bearer " + localStorage.getItem('refresh_token'),
                },
            });
            if (response.status === 200) {
                let parsedResponse = await response.json();
                localStorage.setItem("access_token", parsedResponse.access_token);
                localStorage.setItem("refresh_token", parsedResponse.refresh_token);
                this.addTokens([parsedResponse.access_token, parsedResponse.refresh_token])
                console.log("refresh status 200")
                return true
            } else {
                console.log("fail to refresh tokens", response.text())
                return false
            }
        },
    }
}