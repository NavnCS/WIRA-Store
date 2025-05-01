<template>
  <div id="app">
    <header>
      <h1>WIRA NFT Store</h1>
    </header>

    <FilterBar :selected="selectedRarity" @set-filter="setRarity" />

    <main class="nft-gallery">
      <NFTCard
      v-for="nft in filteredNFTs"
      :key="nft.id"
      :id="nft.id"
      :name="nft.name"
      :rarity="nft.rarity"
      :image="nft.image"
      :is-favorite="nft.isFavorite"
      @toggle-favorite="toggleFavorite"
      />
    </main>
  </div>
</template>

<script>
import NFTCard from './components/NFTCard.vue'
import FilterBar from './components/FilterBar.vue'

export default {
  name: 'App',
  components: {
    NFTCard,
    FilterBar
  },
  data() {
  return {
    selectedRarity: 'All',
    favorites: [],
    nfts: [
      { id: 1, name: 'Crypto Dog', rarity: 'Common', image: '/nft-placeholder.png' },
      { id: 2, name: 'Pixel Alien', rarity: 'Rare', image: '/nft-placeholder.png' },
      { id: 3, name: 'Golden Robot', rarity: 'Legendary', image: '/nft-placeholder.png' },
      { id: 4, name: 'Moon Cat', rarity: 'Common', image: '/nft-placeholder.png' },
    ]
  }
},
computed: {
  filteredNFTs() {
    let list = this.nfts
    if (this.selectedRarity !== 'All') {
      list = list.filter(nft => nft.rarity === this.selectedRarity)
    }
    return list.map(nft => ({
      ...nft,
      isFavorite: this.favorites.includes(nft.id)
    }))
  }
},
methods: {
  setRarity(rarity) {
    this.selectedRarity = rarity
  },
  toggleFavorite(id) {
    const index = this.favorites.indexOf(id)
    if (index >= 0) {
      this.favorites.splice(index, 1)
    } else {
      this.favorites.push(id)
    }
  }
}
}
</script>
