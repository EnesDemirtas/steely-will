<script>
    import { GetAllHabits } from "../wailsjs/go/main/App.js";
    import { CreateHabit } from "../wailsjs/go/main/App.js";

    import Card from "./Card.svelte";

    let habits = [];
    let newHabitName;
    let newHabitGoalDays;

    function init() {
        getAllHabits();
    }

    init();

    function getAllHabits() {
        GetAllHabits().then((result) => (habits = result.Habits));
    }

    function createHabit() {
        let daysInt = parseInt(newHabitGoalDays);
        CreateHabit(newHabitName, daysInt);
    }
</script>

<h3>Steely Will</h3>
<div class="card-container">
    {#each habits as habit}
        <Card {habit}></Card>
    {/each}
</div>
<div>
    <input bind:value={newHabitName} placeholder="Habit name" />
    <input bind:value={newHabitGoalDays} placeholder="Goal as days" />
    <button class="btn" on:click={createHabit}> Create a new habit! </button>
</div>
