<script lang="ts">
  import { onDestroy, onMount } from 'svelte';
  import { GetIdleTime, Test } from '../wailsjs/go/main/App.js';
  import { seconds } from './lib/stores/seconds';

  let idleTime = 0
  let formattedTime = "00:00:00"
  let interval: NodeJS.Timer = null;

  function setFormattedTime(value: number) {
    let hours = Math.floor(value / 3600);
    let mins = Math.floor((value - (hours * 3600)) / 60);
    let secs = value % 60;

    formattedTime = `${hours.toString().padStart(2, '0')}:${mins.toString().padStart(2, '0')}:${secs.toString().padStart(2, '0')}`;
  }

  const unsubscribeSeconds = seconds.subscribe(setFormattedTime)


  function start() {
    if (interval) return

    interval = setInterval(seconds.increment, 500);
  }

  function stop() {
    clearInterval(interval);
    interval = null;
  }

  function reset() {
    stop();
    seconds.reset();
    formattedTime = "00:00:00";
  }

  onMount(() => {
    setInterval(() => {
      Test().then((res) => {
        res.
        console.log(res);
      })
    }, 1000)

    setInterval(() => {
      GetIdleTime().then((result) => {
        idleTime = result
      })
    }, 1000)
  })


  onDestroy(() => {
    unsubscribeSeconds();
  })
</script>

<main>
  <!-- <img alt="Wails logo" id="logo" src="{logo}"> -->
  <div>idle time : {idleTime}</div>
  <div>{formattedTime}</div>
  <button on:click={start}>start</button>
  <button on:click={stop}>stop</button>
  <!-- <div class="result" id="result">{resultText}</div> -->
  <!-- <div class="input-box" id="input">
    <input autocomplete="off" bind:value={name} class="input" id="name" type="text"/>
  </div> -->
</main>

<style>

  #logo {
    display: block;
    width: 50%;
    height: 50%;
    margin: auto;
    padding: 10% 0 0;
    background-position: center;
    background-repeat: no-repeat;
    background-size: 100% 100%;
    background-origin: content-box;
  }

  .result {
    height: 20px;
    line-height: 20px;
    margin: 1.5rem auto;
  }

  .input-box .btn {
    width: 60px;
    height: 30px;
    line-height: 30px;
    border-radius: 3px;
    border: none;
    margin: 0 0 0 20px;
    padding: 0 8px;
    cursor: pointer;
  }

  .input-box .btn:hover {
    background-image: linear-gradient(to top, #cfd9df 0%, #e2ebf0 100%);
    color: #333333;
  }

  .input-box .input {
    border: none;
    border-radius: 3px;
    outline: none;
    height: 30px;
    line-height: 30px;
    padding: 0 10px;
    background-color: rgba(240, 240, 240, 1);
    -webkit-font-smoothing: antialiased;
  }

  .input-box .input:hover {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

  .input-box .input:focus {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

</style>
