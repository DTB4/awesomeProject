<template>
<div
  class="logout"
  id="logout_id"

>
  <h2
      @click="logout()"
  >Logout</h2>
</div>
</template>

<script>
import {mapActions} from "vuex";

export default {
  name: "Logout",
  methods:{
    ...mapActions('tokens', ['removeTokens', 'getLoginStatus']),
    async logout(){
      const response = await fetch('http://localhost:8081/logout', {
        method: 'POST',
        mode: 'cors',
        // credentials: 'include',
        headers: {
          'Accept': '*/*',
          'Authorization': 'Bearer ' + localStorage.getItem('access_token'),
        },
      });
      if (response.ok){
        this.removeTokens()
        alert("Successful logout")
      }else{
        console.log('not ok response', response)
      }
    }
  },
}
</script>

<style scoped>

</style>