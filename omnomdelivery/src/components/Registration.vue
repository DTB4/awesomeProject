<template>
  <div id="registration_container_id" class="registration_container">
    <h2>Registration</h2>
    <input
        id="registration_input_email"
        class="input"
        placeholder="email"
        type="email"
        v-bind:value="email"
        @input="(email = $event.target.value), checkEmail()"
        @submit.prevent
    />
    <input
        id="registration_input_password"
        class="input"
        placeholder="password"
        type="password"
        v-bind:value="password"
        @input="(password = $event.target.value), checkPassword()"
        @submit.prevent
    />
    <input
        id="registration_input_first_name"
        class="input"
        placeholder="first_name"
        type="text"
        v-bind:value="first_name"
        @input="first_name = $event.target.value"
        @submit.prevent
    />
    <input
        id="registration_input_second_name"
        class="input"
        placeholder="second_name"
        type="text"
        v-bind:value="second_name"
        @input="second_name = $event.target.value"
        @submit.prevent
    />
    <button
        @click="registerUser()"
        @submit.prevent
    >Register
    </button>
    <div id="registration_message_container" class="message_container">
      <h2 id="email_message" v-text="password_message"></h2>
      <h2 id="password_message" v-text="email_message"></h2>
    </div>
  </div>
</template>

<script>
export default {
  name: "Registration",
  data() {
    return {
      email_ok: false,
      password_ok: false,
      email: "",
      password: "",
      email_message: "",
      password_message: "",
      first_name: "",
      second_name: "",
    };
  },
  methods: {
    //TODO: learn how to work with inline message pop-up
    checkEmail() {
      switch (this.email) {
        case this.email === "":
          this.email_message = "Email is empty";
          return;
          // case (!this.email.contains('@')):this.email_message="Input is not an email";
          //   return;
        case this.email.length < 5:
          this.email_message = "Email is too short";
          return;
        default:
          this.email_ok = true;
      }
    },
    //TODO: learn how to work with inline message pop-up
    checkPassword() {
      switch (this.password) {
        case this.password === "":
          this.password_message = "Password is empty";
          return;
        case this.password.length < 6:
          this.password_message = "Password is too short";
          return;
        default:
          this.password_ok = true;
      }
    },
    async registerUser() {
      if (!this.email_ok || !this.password_ok) {
        alert("Check input errors");
      } else {
        const response = await fetch("http://localhost:8081/registration", {
          method: "POST",
          mode: "cors",
          credentials: "include",
          headers: {
            Accept: "*/*",
          },
          body: JSON.stringify({
            first_name: this.first_name,
            last_name: this.second_name,
            email: this.email,
            password_hash: this.password,
          }),
        });
        if (response.status === 200) {
          //TODO: make redirect to Login
          let parsedResponce = await response.json();
          console.log("responce from server", parsedResponce);
          this.email = "";
          this.password = "";
          this.first_name = "";
          this.second_name = "";
          alert("OK");
        } else {
          let parsedResponce = await response.json(); // parses JSON response into native JavaScript objects
          alert(parsedResponce);
        }
      }
    },
  },
};
</script>

<style scoped></style>
