<script setup>
import { ref, onMounted } from "vue"
import autoAnimate from "@formkit/auto-animate";
const dropdown = ref()
const show = ref(false)
onMounted(() => {
  autoAnimate(dropdown.value)
})
</script>

<template>
  <div ref="dropdown" class="dropdown" position="relative">
    <strong class="dropdown-label" @click="show = !show"
    >Click me for read more about me</strong>
    <p class="dropdown-content" v-if="show">
      Sea of Tranquility a mote of dust suspended in a sunbeam hundreds of
      thousands concept of the number one realm of the galaxies radio telescope.
      As a patch of light descended from astronomers two ghostly white figures
      in coveralls and helmets are softly dancing emerged into consciousness
      Orion's sword encyclopaedia galactica. Another world bits of moving fluff
      network of wormholes muse about network of wormholes with pretty stories
      for which there's little good evidence and billions upon billions upon
      billions upon billions upon billions upon billions upon billions.
    </p>
  </div>
</template>

<style scoped>
.dropdown {
  border: 2px solid var(--gray-l);
  border-radius: 0.5em;
  background-color: white;
  margin-bottom: 2em;
}
[data-dark-mode="true"] .dropdown {
  border-color: var(--purple-md);
  background-color: var(--purple-d);
}
.dropdown strong {
  display: block;
  padding: 1em;
  cursor: pointer;
}
.dropdown p {
  margin: 1em;
}
</style>