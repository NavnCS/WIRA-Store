<template>
  <div class="nft-card">
    <!-- Model viewer canvas, only visible when model is being shown -->
    <div v-if="isModelVisible" class="model-viewer">
      <canvas ref="threeCanvas"></canvas>
    </div>

    <!-- Image display, clicking it triggers the model view -->
    <div v-else>
      <img :src="image" alt="NFT Image" @click="showModel" />
    </div>

    <!-- NFT Info -->
    <h3>{{ name }}</h3>
    <p>Rarity: {{ rarity }}</p>

    <!-- Favorite Button -->
    <button class="favorite-btn" @click="$emit('toggle-favorite', id)">
      {{ isFavorite ? '❤️' : '🤍' }}
    </button>
  </div>
</template>

<script>
// Import necessary Three.js functionality
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
      isModelVisible: false, // Flag to toggle between image and 3D model view
      scene: null, // Three.js scene
      camera: null, // Camera to view the scene
      renderer: null, // Renderer for the 3D model
      model: null // The loaded 3D model
    };
  },
  methods: {
    showModel() {
      this.isModelVisible = true;
      this.$nextTick(() => {
        this.initThreeJS(); // Initialize Three.js for rendering after DOM is updated
        this.loadModel(); // Load the 3D model when clicked
      });
    },
    initThreeJS() {
      const canvas = this.$refs.threeCanvas; // Canvas element to render the 3D model
      const width = canvas.clientWidth; // Get the width of the canvas after it's mounted
      const height = canvas.clientHeight; // Get the height of the canvas after it's mounted
      this.scene = new THREE.Scene();
      this.camera = new THREE.PerspectiveCamera(75, width / height, 0.1, 1000);
      this.renderer = new THREE.WebGLRenderer({ canvas });
      this.renderer.setSize(width, height); // Set the renderer size based on canvas dimensions
      this.camera.position.z = 5; // Set camera position
    },
    loadModel() {
      const loader = new GLTFLoader(); // Load 3D models in GLTF format
      loader.load(
        '/models/car.glb', // Path to your GLTF model (adjust path as needed)
        (gltf) => {
          this.model = gltf.scene; // Assign the loaded model
          this.scene.add(this.model); // Add the model to the scene
          this.animate(); // Start the animation loop
        },
        undefined,
        (error) => {
          console.error('An error occurred while loading the model:', error);
        }
      );
    },
    animate() {
      const animateScene = () => {
        requestAnimationFrame(animateScene); // Request next frame
        if (this.model) {
          this.model.rotation.x += 0.01; // Rotate the model for animation
          this.model.rotation.y += 0.01;
        }
        this.renderer.render(this.scene, this.camera); // Render the scene
      };
      animateScene();
    }
  },
  destroyed() {
    // Clean up Three.js resources when the component is destroyed
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
  max-width: 220px; /* Ensure each card has a maximum width */
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
  height: 200px; /* Adjust based on your design */
  background-color: #222;
  margin-bottom: 10px;
}

img {
  cursor: pointer;
  width: 100%; /* Ensure image is responsive */
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
