<template>
  <div class="container mx-auto">
    <section class="flex flex-col sm:justify-between items-center sm:flex-row mt-12 sm:mt-10 mb-5 sm:mb-8">
      <div class="container">
        <div id="default">
          <div class="w-full md:w-1/2 text-left">
            <transition name="fade" appear>
              <h3 class="font-general-semibold text-3xl md:text-3xl xl:text-4xl text-center sm:text-left text-ternary-dark dark:text-primary-light uppercase">Hi! My name is Dmitriy.</h3>
            </transition>
            <transition name="fade" appear>
              <p style="transition-delay: .9s" class="font-general-medium mt-2 text-lg sm:text-xl xl:text-2xl text-center sm:text-left leading-none text-gray-400">A Back-End Developer.</p>
            </transition>
          </div>
        </div>
        <div class="align-right">
          <img src='@/assets/17850BED-E8E1-4013-8E13-E9B17FDB5B2B.png' width="250" height="250" />
        </div>
        <div id="drop">
          <transition name="fordropdown" appear style="transition-delay: .9s">
            <ActualDropdown /> 
          </transition>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
import ActualDropdown from "@/components/ActualDropdown.vue";

export default {
  name: 'HelloWorld',
  components: {ActualDropdown},
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

.container {
  width: 100%;
}

#default {
  border-radius: 50px;
  width: 40%;
  height: 60%;
  display: flex;
  justify-content: center;
  /*flex-direction: row;*/
  background-color: #5b6dcd;
  color: #fff;
  padding: 100px;
  float: left;
}

#drop{
  padding: 100px;
  position:absolute;
  width:100%;
  top: 300px;
}

.align-right {
  float: right;
}

.fordropdown-enter-from,
.fordropdown-leave-to {
  opacity: 0;
}

.fordropdown-enter-active,
.fordropdown-leave-active {
  transition: all 2s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
}

.fade-enter-active,
.fade-leave-active {
  transition: all 2s ease;
}
</style>
