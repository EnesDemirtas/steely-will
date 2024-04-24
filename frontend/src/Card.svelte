<script>
    import { onMount } from 'svelte';

    export let habit;

    let lastResetDate;
    let elapsedTime;

    onMount(() => {
        lastResetDate = new Date(habit.lastResetDate)

        const intervalId = setInterval(() => {
            elapsedTime = Math.floor((Date.now() - lastResetDate.getTime()) / 1000);
        }, 1000);

        return () => clearInterval(intervalId);
    });

    function formatElapsedTime(seconds) {
        const hours = Math.floor(seconds / 3600);
        const minutes = Math.floor((seconds % 3600) / 60);
        const secondsRemaining = seconds % 60;

        const formattedHours = hours.toString().padStart(2, '0');
        const formattedMinutes = minutes.toString().padStart(2, '0');
        const formattedSeconds = secondsRemaining.toString().padStart(2, '0');

        return `${formattedHours}:${formattedMinutes}:${formattedSeconds}`;
    }
</script>

<div class="card">
    <h3>{habit.name}</h3>
    <h4>{habit.failStack}</h4>
    <span>{elapsedTime ? formatElapsedTime(elapsedTime) : '00:00:00'}</span>
</div>