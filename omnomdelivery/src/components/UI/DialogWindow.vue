<template>
  <div v-if="show" class="dialog_window" @click="hideDialogWindow">
    <div class="dialog_content" @hideThisWindow="hideDialogWindow" @click.stop>
      <slot></slot>
    </div>
  </div>
</template>

<script>
export default {
  name: "DialogWindow",
  props: {
    show: {
      type: Boolean,
      default: false,
    },
  },
  methods: {
    hideDialogWindow() {
      this.$emit("hideDialogWindow");
    },
  },
};
</script>

<style scoped>
.dialog_window {
  z-index: 10;
  top: 0;
  bottom: 0;
  right: 0;
  left: 0;
  background: rgba(0, 0, 0, 0.7);
  position: fixed;
  display: flex;
}

.dialog_content {
  margin: auto;
  background: white;
  border-radius: 15px;
  min-height: 50px;
  min-width: 400px;
  padding: 20px;
}
</style>
