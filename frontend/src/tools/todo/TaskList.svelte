<script>
  import Task from "./Task.svelte";
  import NewTaskButton from "./NewTaskButton.svelte";

  let tasks = [
    "Move tasks to svelte store",
    "Button to rearrange tasks (cursor drag)",
    "Modify tasks so they're objects (create ID generator)",
    "Setup SQLite database to read/write tasks",
    "Use textarea instead of text input for task edits",
    "Restyle input boxes during edits",
  ];

  const newTask = () => {
    tasks.push("");
    tasks = tasks;
  };

  const deleteTask = (id) => {
    tasks.splice(id, 1);
    tasks = tasks;
  };

  const updateTaskText = (id, text) => {
    tasks[id] = text;
    tasks = tasks;
  };
</script>

<div class="task-list">
  {#each tasks as text, i}
    <Task
      id={i}
      {text}
      whenDeletePressed={deleteTask}
      whenEditComplete={updateTaskText}
    />
  {/each}
  <NewTaskButton onClick={newTask} />
</div>

<style>
  .task-list {
    display: flex;
    flex-direction: column;
    row-gap: 1rem;

    overflow: scroll;
    max-height: 100%;
  }
</style>
