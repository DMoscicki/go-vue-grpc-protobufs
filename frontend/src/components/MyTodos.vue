<template>
  <div id="todoss">
    <section>
      <h3 class="title-text">TODO gRPC Client</h3>
      <div class="todo-create">
        <v-text-field rounded label="Write task" v-model="inputField" v-on:keyup.enter="addTodo" hide-details="auto" > </v-text-field>
        <div>
          <v-btn rounded="lg" color="primary" @click="addTodo">Add Todo</v-btn>
            <v-dialog v-model="dialog" width="auto">
              <v-card>
                <v-card-text>
                  Вы не ввели значение в строку ввода. Пожалуйста введите значение.
                </v-card-text>
                <v-card-actions>
                  <v-btn color="primary" block @click="dialog = false">Close dialog</v-btn>
                </v-card-actions>
              </v-card>
            </v-dialog>
        </div>
      </div>
    </section>
    <section>
      <div id="list-complete-demo" class="row">
        <div class="offset-md-3 col-md-6 mt-3">
          <ul v-auto-animate ref="list">
            <li v-for="todo in todos" :key="todo.id">
              <span>{{todo.task}}</span>
              <button @click="deleteTodo(todo)" aria-label="Remove">
                <IconRemove />
              </button>
            </li>
          </ul>
        </div>
      </div>
    </section>
  </div>
</template>

<script>

import { AddTodos, DeleteTodos, GetTodos } from '@/protos/answer_pb'
import {todoServiceClient} from '@/protos/answer_grpc_web_pb'
import IconRemove from "@/components/IconRemove.vue";

export default {
  name: 'MyTodoView',
  components: {IconRemove},
  data: function () {
    return {inputField: "", todos: [], dialog:false};
  },
  created: function () {
    this.client = new todoServiceClient("http://localhost:8080", null, null)
    this.getTodos();
  },
  methods: {
    getTodos: function () {
      let getRequest = new GetTodos();
      this.client.getTodos(getRequest, {}, (err, response) => {
        this.todos = response.toObject().todosList;
        console.log(this.todos);
      });
    },
    addTodo: function () {
      let request = new AddTodos();
      if (this.inputField === "") {
        this.dialog = true
      } else {
        request.setTask(this.inputField)
        this.client.addTodo(request, {}, () => {
          this.inputField = "";
          this.getTodos();
        });
      }
    },
    deleteTodo: function (todo) {
      let deleteRequest = new DeleteTodos();
      deleteRequest.setId(todo.id);
      this.client.deleteTodo(deleteRequest, {}, (err, response) => {
        if (response.getMessage() === "Success") {
          this.getTodos();
        }
      });
      console.log("todo -> ", todo.id)
    },
  }
}
</script>

<style scoped>

ul {
  list-style-type: none;
  padding: 0;
  max-width: 300px;
}
li {
  display: flex;
  justify-content: space-between;
  padding: 0.75em;
  background-color: antiquewhite;
  margin-bottom: 0.5em;
  border-radius: 0.5em;
  box-shadow: 0 0 0.5em rgba(0, 0, 0, 0.1);
  font-size: 0.875em;
}
[data-dark-mode="true"] li {
  background-color: var(--purple-md);
}
li::before {
  display: none;
}
li button {
  appearance: none;
  border: none;
  padding: none;
  margin: none;
  background-color: transparent;
  display: flex;
  align-items: center;
  cursor: pointer;
}
li button svg {
  width: 1.2em;
  fill: red;
}
[data-dark-mode="true"] li button svg {
  fill: rgb(244, 67, 67);
}

/*ul {*/
/*  margin-top: 24px;*/
/*  border-radius: 10px;*/
/*  position: relative;*/
/*}*/
/*span {*/
/*  flex: 1;*/
/*}*/

.todo-enter-active,
.todo-move {
  transition: 0.4s ease all;
}
.todo-enter-from,
.todo-leave-to {
  opacity: 0;
  transform: scale(0.6);
}
.todo-leave-active {
  transition: 0.4s ease all;
  position: absolute;
}
.todo-list-enter-active,
.todo-list-leave-active {
  transition: 0.4s ease all;
}
.todo-list-enter-from,
.todo-list-leave-to {
  transform: translateY(30px);
  opacity: 0;
}

</style>