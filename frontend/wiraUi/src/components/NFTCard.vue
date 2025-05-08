<template>
  <div class="nft-card">
    <div v-if="isModelVisible" class="model-viewer">
      <canvas ref="threeCanvas"></canvas>
    </div>

    <div v-else>
      <img :src="image" alt="NFT Image" @click="showModel" />
    </div>

    <h3>{{ name }}</h3>
    <p>Rarity: {{ rarity }}</p>

    <button class="favorite-btn" @click="$emit('toggle-favorite', id)">
      {{ isFavorite ? '❤️' : '🤍' }}
    </button>
  </div>
</template>

<script>
import * as THREE from 'three';
import { GLTFLoader } from 'three/examples/jsm/loaders/GLTFLoader';

export default {
  name: 'NFTCard',
  props: {
    id: Number,
    name: String,
    rarity: String,
    image: String,
    isFavorite: Boolean
  },
  data() {
    return {
      isModelVisible: false, 
      scene: null, 
      camera: null,
      renderer: null, 
      model: null 
    };
  },
  methods: {
    showModel() {
      this.isModelVisible = true;
      this.$nextTick(() => {
        this.initThreeJS(); 
        this.loadModel();
      });
    },
    initThreeJS() {
      const canvas = this.$refs.threeCanvas; 
      const width = canvas.clientWidth;
      const height = canvas.clientHeight; 
      this.scene = new THREE.Scene();
      this.camera = new THREE.PerspectiveCamera(75, width / height, 0.1, 1000);
      this.renderer = new THREE.WebGLRenderer({ canvas });
      this.renderer.setSize(width, height); 
      this.camera.position.z = 5;
    },
    loadModel() {
      const loader = new GLTFLoader(); 
      loader.load(
        '/models/car.glb', 
        (gltf) => {
          this.model = gltf.scene; 
          this.scene.add(this.model); 
          this.animate();
        },
        undefined,
        (error) => {
          console.error('An error occurred while loading the model:', error);
        }
      );
    },
    animate() {
      const animateScene = () => {
        requestAnimationFrame(animateScene); 
        if (this.model) {
          this.model.rotation.x += 0.01; 
          this.model.rotation.y += 0.01;
        }
        this.renderer.render(this.scene, this.camera); 
      };
      animateScene();
    }
  },
  destroyed() {

    if (this.renderer) {
      this.renderer.dispose();
    }
  }
};
</script>

<style scoped>
.nft-card {
  border: 1px solid #ccc;
  padding: 10px;
  width: 100%;
  max-width: 220px; 
  margin: 10px auto;
  border-radius: 8px;
  text-align: center;
  background-color: #333;
  color: white;
  transition: transform 0.2s;
}

.nft-card:hover {
  transform: translateY(-5px);
  background-color: #444;
}

.model-viewer {
  width: 100%;
  height: 200px; 
  background-color: #222;
  margin-bottom: 10px;
}

img {
  cursor: pointer;
  width: 100%; 
  height: auto;
}

button {
  border-radius: 8px;
  border: 1px solid transparent;
  padding: 0.6em 1.2em;
  font-size: 1em;
  font-weight: 500;
  font-family: inherit;
  background-color: #1a1a1a;
  cursor: pointer;
  transition: border-color 0.25s;
  color: white;
}

button:hover {
  border-color: #646cff;
}
</style>
