<template>
  <div id="app">
    <h1>Todo App</h1>

    <div>
      <h2>Add Todo</h2>
      <input v-model="newTodo.title" placeholder="Title" />
      <input v-model="newTodo.description" placeholder="Description" />
      <button @click="addTodo">Add</button>
    </div>

    <h2>Todos</h2>
    <ul>
      <li v-for="todo in todos" :key="todo.id">
        <h3>{{ todo.title }}</h3>
        <p>{{ todo.description }}</p>
        <label>
          <input type="checkbox" v-model="todo.is_completed" @change="updateTodo(todo)" />
          Completed
        </label>
        <button @click="deleteTodo(todo.id)">Delete</button>
      </li>
    </ul>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      todos: [],
      newTodo: {
        title: '',
        description: '',
      },
    };
  },
  methods: {
    fetchTodos() {
      axios.get('http://localhost:3000/api/todos').then((response) => {
        this.todos = response.data;
      }).catch(err => console.log(err));
    },
    addTodo() {
      axios.post('http://localhost:3000/api/todos', this.newTodo).then(() => {
        this.fetchTodos();
        this.newTodo.title = '';
        this.newTodo.description = '';
      });
    },
    updateTodo(todo) {
      axios.put(`http://localhost:3000/api/todos/${todo.id}`, todo).then(() => {
        this.fetchTodos();
      });
    },
    deleteTodo(id) {
      axios.delete(`http://localhost:3000/api/todos/${id}`).then(() => {
        this.fetchTodos();
      });
    },
  },
  mounted() {
    this.fetchTodos();
  },
};
</script>
