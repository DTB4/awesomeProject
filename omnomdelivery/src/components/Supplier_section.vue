<template>
  <div class="supplier_section" id="supplier_section_id">
    <supplier
      v-for="(supplier, id) in suppliers_array"
      :key="id"
      :supplier_ent="supplier"
    ></supplier>
  </div>
</template>

<script>
export default {
  name: "Supplier_section",
  data() {
    return {
      suppliers_array: [],
    };
  },
  methods: {
    async getAllSuppliers() {
      let resp = await fetch("http://localhost:8081/suppliers", {
        method: "GET",
      });
      let suppliersMassive = await resp.json();
      return suppliersMassive;
    },
  },
  created() {},
  async mounted() {
    document.getElementById("supplier_section_id").innerText = "Loading";
    this.suppliers_array = await this.getAllSuppliers();
    document.getElementById("supplier_section_id").innerText = "";
  },
};
</script>

<style scoped>
.supplier_section {
  width: auto;
  height: auto;
  border: solid black;
  background: antiquewhite;
  display: flex;
  flex-wrap: wrap;
  box-sizing: border-box;
  position: relative;
  justify-content: space-evenly;
}

@keyframes fadeIn {
  from {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}

.supplier_section > * {
  box-sizing: border-box;
  flex: 0 1 15em;
  opacity: 0;
  animation: fadeIn ease-in 1;
  animation-fill-mode: forwards;
  animation-duration: 1s;
}
</style>
