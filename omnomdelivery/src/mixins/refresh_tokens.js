import {mapActions} from "vuex";

export default {
    data() {
        return {
            url: "http://localhost:8081"
        }
    },
    methods: {
        ...mapActions("tokens", ["addTokens", "removeTokens"]),
        async refreshTokens() {
            console.log("access token expired. try to refresh");
            const response = await fetch(`${this.url}/refresh`, {
                method: "GET",
                mode: "cors",
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
                console.log("refresh successful")
                return true
            } else if (response.status === 401) {
                console.log("both tokens are expired")
                this.$emit("userLogout")
                this.$emit("hideDialogWindow")
                this.removeTokens()
                return false

            } else {
                console.log("fail to refresh tokens", response.text())
                this.$emit("userLogout")
                this.$emit("hideDialogWindow")
                this.removeTokens()
                return false
            }
        },
    }
}