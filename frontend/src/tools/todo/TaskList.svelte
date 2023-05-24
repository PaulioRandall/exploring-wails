<script>
  import Task from "./Task.svelte";
  import NewTaskButton from "./NewTaskButton.svelte";

  let tasks = [
    {
      text: "Move tasks to svelte store",
    },
    {
      text: "Button to rearrange tasks (cursor drag)",
    },
    {
      text: "Modify tasks so they're objects (create ID generator)",
    },
    {
      text: "Setup SQLite database to read/write tasks",
    },
    {
      text: "Use textarea instead of text input for task edits",
    },
    {
      text: "Restyle input boxes during edits",
    },
  ];

  const newTask = () => {
    tasks.push("");
    tasks = tasks;
  };

  const deleteTask = (id) => {
    tasks.splice(id, 1);
    tasks = tasks;
  };

  const updateTask = (id, task) => {
    tasks[id] = task;
    tasks = tasks;
  };
</script>

<div class="task-list">
  {#each tasks as task, i}
    <Task
      id={i}
      {task}
      whenDeletePressed={deleteTask}
      whenEditComplete={updateTask}
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
