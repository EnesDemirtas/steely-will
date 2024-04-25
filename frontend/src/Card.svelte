<script>
    import { onMount } from "svelte";

    export let habit;

    let lastResetDate;
    let elapsedTime;

    onMount(() => {
        lastResetDate = new Date(habit.lastResetDate);

        const intervalId = setInterval(() => {
            elapsedTime = Math.floor(
                (Date.now() - lastResetDate.getTime()) / 1000
            );
        }, 1000);

        return () => clearInterval(intervalId);
    });

    function formatElapsedTime(seconds) {
        const hours = Math.floor(seconds / 3600);
        const minutes = Math.floor((seconds % 3600) / 60);
        const secondsRemaining = seconds % 60;

        const formattedHours = hours.toString().padStart(2, "0");
        const formattedMinutes = minutes.toString().padStart(2, "0");
        const formattedSeconds = secondsRemaining.toString().padStart(2, "0");

        return `${formattedHours}:${formattedMinutes}:${formattedSeconds}`;
    }

    function getProgressPercentage() {
        const lastResetDate = new Date(habit.lastResetDate);
        const goal = new Date(
            lastResetDate.getTime() + habit.goalDays * 24 * 60 * 60 * 1000
        );
        const now = new Date();

        if (now > goal) {
            return 1;
        } else {
            const elapsedTime =
                (now.getTime() - lastResetDate.getTime()) / 1000;
            const goalDaysAsSeconds = habit.goalDays * 24 * 60 * 60;
            const percentage = (elapsedTime / goalDaysAsSeconds) * 100;
            return percentage.toFixed(2);
        }
    }
</script>

<div class="card">
    <h3>{habit.name}</h3>
    <h4><span>Fails: </span>{habit.failStack}</h4>
    <h4><span>Goal: </span>{habit.goalDays} days</h4>
    <h4>
        <span>{elapsedTime ? formatElapsedTime(elapsedTime) : "00:00:00"}</span>
    </h4>
    <h4><span>%{getProgressPercentage()} completed.</span></h4>
</div>
