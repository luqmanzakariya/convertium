<template>
  <button>{{ title }}</button>
</template>

<script>
export default {
  name: "NavbarButton",
  props: ["title"],
  methods: {
    handleClick() {
      this.$emit("click");
    },
  },
};
</script>

<style scoped>
button {
  background: grey;
  color: #ffffff;
  border: none;
  padding: 0.5rem 1rem;
  cursor: pointer;
  transition: background-color 0.3s, color 0.3s;
  border-radius: 8px;
  font-weight: 700;
}

button:hover {
  background: black;
  color: #ffd700;
}
</style>
