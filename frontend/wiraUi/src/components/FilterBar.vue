<template>
    <div class="filter-bar">
      <button
        v-for="r in rarities"
        :key="r"
        @click="$emit('set-filter', r)"
      >
        {{ r }}
      </button>
    </div>
  </template>
  
  <script>
  export default {
    name: 'FilterBar',
    props: ['selected'],
    data() {
      return {
        rarities: ['All', 'Common', 'Rare', 'Legendary']
      }
    }
  }
  </script>
  
  <style scoped>
  .filter-bar {
    margin: 20px;
  }
  button {
    margin: 5px;
    padding: 8px 16px;
    background-color: #ddd;
    border: none;
    cursor: pointer;
  }
  button:hover {
    background-color: #bbb;
  }
  </style>
  