<script>
  import { GetAllHabits } from "../wailsjs/go/main/App.js";
  import { CreateHabit } from "../wailsjs/go/main/App.js";

  let habits = [];
  let newHabitName;
  let newHabitGoalDays;

  function init() {
    getAllHabits()
  }

  init();

  function getAllHabits() {
    GetAllHabits().then((result) => (habits = result))
  }

  function createHabit() {
    let daysInt = parseInt(newHabitGoalDays)
    CreateHabit(newHabitName, daysInt)
  }
</script>

<h3>Steely Will</h3>
<div>
  {#each habits as habit}
    <span>{habit.Name}</span>
  {/each}
</div>
<div>
  <input bind:value={newHabitName} />
  <input bind:value={newHabitGoalDays} />
  <button class="btn" on:click={createHabit}>
    Create a new habit!
  </button>
</div>
