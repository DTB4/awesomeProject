<template>
  <div class="supplier" @click="$emit('showProductsFor', supplier_ent.id)">
    <div class="img_container" :style="{ backgroundImage: `url(${supplier_ent.img_url})` }"></div>
    <h4 class="supplier_type">{{ supplier_ent.type }}</h4>
    <h2 class="supplier_name">{{ supplier_ent.name }}</h2>
    <h4 class="supplier_work_hours">{{ supplier_ent.opening }} - {{ supplier_ent.closing }}</h4>
  </div>
</template>

<script>
export default {
  name: "Supplier",
  props: {
    supplier_ent: Object,
  },
};
</script>

<style scoped>
.supplier {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  box-shadow: 0 4px 8px 0 rgba(0, 0, 0, 0.2);
  max-width: 25vw;
  height: 30vh;
  font-family: arial, serif;
  background: aliceblue;
  margin: 5pt;
  border-radius: 10px;
  box-sizing: border-box;
  flex: 0 1 15em;
  opacity: 0;
  animation: fadeIn ease-in 1;
  animation-fill-mode: forwards;
  animation-duration: 1s;
}

@keyframes fadeIn {
  from {
    opacity: 0;
  }
  to {
    opacity: 1;
  }
}

.img_container {
  width: 100%;
  height: 60%;
  background-position: center;
  background-repeat: no-repeat;
  background-size: 100%;
  border-radius: 10px;
}

.supplier_type {
  margin-left: 5pt;
  text-align: left;
  color: coral;
}

.supplier_name {
  text-align: center;
}

.supplier_work_hours {
  text-align: center;
}

</style>
