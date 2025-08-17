<template>
  <button>{{ title }}</button>
</template>

<script>
export default {
  name: "SubmitButton",
  props: ["title"],
};
</script>

<style scoped>
button {
  padding: 10px 20px;
  background-color: black;
  color: white;
  border: none;
  cursor: pointer;
  width: 100%;
  border-radius: 4px;
  max-width: 400px;
  margin-top: 12px;
}

button:hover {
  background-color: #555;
}
</style>
