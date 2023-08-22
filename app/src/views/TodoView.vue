<template>
  <div id="mytodo">
    <MyTodos />
  </div>
</template>

<script>
import MyTodos from "@/components/MyTodos.vue";

export default {
  name: "MyTodoView",
  components: {
    MyTodos
  }
}
</script>