<template>
    <div class="nft-card">
      <img :src="image" alt="NFT Image" />
      <h3>{{ name }}</h3>
      <p>Rarity: {{ rarity }}</p>
      <button class="favorite-btn" @click="$emit('toggle-favorite', id)">
        {{ isFavorite ? '❤️' : '🤍' }}
      </button>
    </div>
  </template>
  
  <script>
  export default {
    name: 'NFTCard',
    props: {
      id: Number,
      name: String,
      rarity: String,
      image: String,
      isFavorite: Boolean
    }
  }
  </script>
  
  <style scoped>
  .nft-card {
    border: 1px solid #ccc;
    padding: 10px;
    width: 200px;
    margin: 10px auto;
    border-radius: 8px;
    text-align: center;
  }
  .favorite-btn {
    background: none;
    border: none;
    font-size: 1.5rem;
    cursor: pointer;
  }
  </style>
  