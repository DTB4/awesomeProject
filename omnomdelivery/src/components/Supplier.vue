<template>
  <div class="supplier">
    <a class="supplier_link" rel="">
      <img
          class="supplier_img"
          width="100%"
          :src="supplier_ent.img_url"
          alt=""
      />
      <h2 class="supplier_name">{{ supplier_ent.name }}</h2>
    </a>
    <h4 class="supplier_type">{{ supplier_ent.type }}</h4>
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
  margin: 0.5vh 0.5vh;
  background: aliceblue;
  width: 20vh;
  height: auto;
  border: solid black 0.1pt;
}
</style>
